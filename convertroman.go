package convertroman

import "strings"

var (
	// lookup arrays used for converting from an int to a roman numeral extremely quickly.
	// method from here: https://rosettacode.org/wiki/Roman_numerals/Encode#Go
	r0 = [10]string{"", "i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix"}
	r1 = [10]string{"", "x", "xx", "xxx", "xl", "l", "lx", "lxx", "lxxx", "xc"}
	r2 = [10]string{"", "c", "cc", "ccc", "cd", "d", "dc", "dcc", "dccc", "cm"}
	r3 = [4]string{"", "m", "mm", "mmm"}
)

// FromInt converts an integer value to a roman numeral string.
// The function will return "0" if the input integer is less than 1.
// The function will return "OOB" if the input integer is greater than
// or equal to 4,000,000 (the roman numeral vinculum limit).
func FromInt(input int) string {
	if input < 1 {
		return "0"
	} else if input < 4000 {
		return baseRange(input)
	} else if input < 4000000 {
		thousands, remainder := separateThousands(input)
		return vinculumRange(thousands) + baseRange(remainder)
	} else {
		return "OOB"
	}
}

// FromIntCapital very inefficiently converts an integer value to a CAPITAL roman numeral string.
// Note that this package is optimized for lowercase roman numerals. As such, this function is a slow afterthought.
func FromIntCapital(input int) string {
	if input < 4000 {
		// if below vinculum range, return the translated result
		return strings.ToUpper(FromInt(input))
	} else {
		// if in vinculum range, remove the ANSI escape codes, capitalize the result, and reapply the ANSI escape codes
		noPrefix := strings.TrimPrefix(FromInt(input), "\033[53m")
		secondIndex := strings.Index(noPrefix, "\033[0m")
		capitalPreANSI := strings.ToUpper(noPrefix[:secondIndex] + noPrefix[secondIndex+4:])
		return "\033[53m" + capitalPreANSI[:secondIndex] + "\033[0m" + capitalPreANSI[secondIndex:]
	}
}

// separateThousands separates an integer into two parts: the thousands and the remainder.
func separateThousands(num int) (int, int) {
	return num / 1000, num % 1000
}

// vinculumRange returns a roman numeral string for an integer greater than 3999.
func vinculumRange(input int) string {
	return "\033[53m" + baseRange(input) + "\033[0m"
}

// baseRange returns a roman numeral string for an integer between 1 and 3999.
func baseRange(input int) string {
	return r3[input%1e4/1e3] + r2[input%1e3/1e2] + r1[input%100/10] + r0[input%10]
}

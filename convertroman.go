package convertroman

var (
	// lookup arrays used for converting from an int to a roman numeral extremely quickly.
	// method from here: https://rosettacode.org/wiki/Roman_numerals/Encode#Go
	r0 = [10]string{"", "i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix"}
	r1 = [10]string{"", "x", "xx", "xxx", "xl", "l", "lx", "lxx", "lxxx", "xc"}
	r2 = [10]string{"", "c", "cc", "ccc", "cd", "d", "dc", "dcc", "dccc", "cm"}
	r3 = [4]string{"", "m", "mm", "mmm"}
)

// FromInt converts an integer value to a roman numeral string.
func FromInt(input int) string {
	// check the range of the input integer
	if input < 1 {
		return "0"
	} else if input < 4000 {
		return baseRange(input)
	} else {
		thousands, remainder := separateThousands(input)
		return vinculumRange(thousands) + baseRange(remainder)
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

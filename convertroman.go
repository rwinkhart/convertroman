package convertroman

// Currently only supports Roman Numerals without viniculum (1-3999; "OOB" is returned for values outside this range)
// See here for details on viniculum:
// https://en.wikipedia.org/wiki/Roman_numerals#Large_numbers

var (
	// lookup arrays used for converting from an int to a roman numeral extremely quickly.
	// method from here: https://rosettacode.org/wiki/Roman_numerals/Encode#Go
	r0 = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	r1 = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	r2 = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	r3 = [4]string{"", "M", "MM", "MMM"}
)

// FromInt converts an integer value to a roman numeral string.
// "OOB" is returned if the integer is not between 1 and 3999.
func FromInt(input int) string {
	// ensure provided integer is within the valid range
	if input < 1 || input > 3999 {
		return "OOB"
	}

	// convert the integer to a roman numeral string and return it
	return r3[input%1e4/1e3] + r2[input%1e3/1e2] + r1[input%100/10] + r0[input%10]
}

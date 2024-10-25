package convertroman

import (
	"fmt"
	"testing"
)

var testCases = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12,
	"XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17,
	"XVIII": 18, "XIX": 19, "XX": 20, "XXXI": 31, "XXXII": 32,
	"XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37,
	"XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42,
	"XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47,
	"XLVIII": 48, "XLIX": 49, "L": 50, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95,
	"XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	"CI": 101, "CII": 102, "CIII": 103, "CIV": 104, "CV": 105,
	"CVI": 106, "CVII": 107, "CVIII": 108, "CIX": 109, "CXLIX": 149,
	"CCCXLIX": 349, "CDLVI": 456, "D": 500, "DCIV": 604, "DCCLXXXIX": 789,
	"DCCCXLIX": 849, "CMIV": 904, "M": 1000, "MVII": 1007, "MLXVI": 1066,
	"MCCXXXIV": 1234, "MDCCLXXVI": 1776, "MMXXI": 2021, "MMDCCCVI": 2806,
	"MMCMXCIX": 2999, "MMM": 3000, "MMMCMLXXIX": 3979, "MMMCMXCIX": 3999,
}

func TestFromInt(t *testing.T) {
	for expected, input := range testCases {
		out := FromInt(input)
		if out == "OOB" {
			t.Errorf("FromInt(%d) returned an OOB (out of bounds) error", input)
		} else if out != expected {
			t.Errorf("FromInt(%d) = %s; want %s", input, out, expected)
		}
	}
	out := FromInt(100000)
	if out != "OOB" {
		t.Errorf("FromInt(%d) expected an error", 100000)
	}
	out = FromInt(0)
	if out != "OOB" {
		t.Errorf("FromInt(%d) expected an error", 0)
	}
}

func BenchmarkFromInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = FromInt(3999)
	}
}

func ExampleFromInt() {
	roman := FromInt(4)
	if roman == "OOB" {
		panic("Input integer is out of bounds")
	}
	fmt.Println(roman == "IV") // True
}

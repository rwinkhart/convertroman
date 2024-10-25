package convertroman

import (
	"fmt"
	"testing"
)

var testCases = map[string]int{
	"i": 1, "ii": 2, "iii": 3, "iv": 4, "v": 5, "vi": 6,
	"vii": 7, "viii": 8, "ix": 9, "x": 10, "xi": 11, "xii": 12,
	"xiii": 13, "xiv": 14, "xv": 15, "xvi": 16, "xvii": 17,
	"xviii": 18, "xix": 19, "xx": 20, "xxxi": 31, "xxxii": 32,
	"xxxiii": 33, "xxxiv": 34, "xxxv": 35, "xxxvi": 36, "xxxvii": 37,
	"xxxviii": 38, "xxxix": 39, "xl": 40, "xli": 41, "xlii": 42,
	"xliii": 43, "xliv": 44, "xlv": 45, "xlvi": 46, "xlvii": 47,
	"xlviii": 48, "xlix": 49, "l": 50, "lxxxix": 89, "xc": 90,
	"xci": 91, "xcii": 92, "xciii": 93, "xciv": 94, "xcv": 95,
	"xcvi": 96, "xcvii": 97, "xcviii": 98, "xcix": 99, "c": 100,
	"ci": 101, "cii": 102, "ciii": 103, "civ": 104, "cv": 105,
	"cvi": 106, "cvii": 107, "cviii": 108, "cix": 109, "cxlix": 149,
	"cccxlix": 349, "cdlvi": 456, "d": 500, "dciv": 604, "dcclxxxix": 789,
	"dcccxlix": 849, "cmiv": 904, "m": 1000, "mvii": 1007, "mlxvi": 1066,
	"mccxxxiv": 1234, "mdcclxxvi": 1776, "mmxxi": 2021, "mmdcccvi": 2806,
	"mmcmxcix": 2999, "mmm": 3000, "mmmcmlxxix": 3979, "mmmcmxcix": 3999,
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

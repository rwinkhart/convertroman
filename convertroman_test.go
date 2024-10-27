package convertroman

import (
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
	"\033[53miv\033[0m": 4000, "\033[53mv\033[0m": 5000, "\033[53mvi\033[0m": 6000,
	"\033[53mvii\033[0m": 7000, "\033[53mviii\033[0m": 8000, "\033[53mix\033[0m": 9000, "\033[53mx\033[0m": 10000, "\033[53mxi\033[0m": 11000, "\033[53mxii\033[0m": 12000,
	"\033[53mxiii\033[0m": 13000, "\033[53mxiv\033[0m": 14000, "\033[53mxv\033[0mxcix": 15099, "\033[53mxvi\033[0m": 16000, "\033[53mxvii\033[0m": 17000,
	"\033[53mxviii\033[0mi": 18001, "\033[53mxix\033[0m": 19000, "\033[53mxx\033[0m": 20000, "\033[53mxxxi\033[0m": 31000, "\033[53mxxxii\033[0m": 32000,
	"\033[53mxxxiii\033[0m": 33000, "\033[53mxxxiv\033[0m": 34000, "\033[53mxxxv\033[0m": 35000, "\033[53mxxxvi\033[0m": 36000, "\033[53mxxxvii\033[0m": 37000,
	"\033[53mxxxviii\033[0m": 38000, "\033[53mxxxix\033[0m": 39000, "\033[53mxl\033[0m": 40000, "\033[53mxli\033[0m": 41000, "\033[53mxlii\033[0m": 42000,
	"\033[53mxliii\033[0m": 43000, "\033[53mxliv\033[0m": 44000, "\033[53mxlv\033[0mlvi": 45056, "\033[53mxlvi\033[0m": 46000, "\033[53mxlvii\033[0m": 47000,
	"\033[53mxlviii\033[0m": 48000, "\033[53mxlix\033[0m": 49000, "\033[53ml\033[0m": 50000, "\033[53mlxxxix\033[0mc": 89100, "\033[53mxc\033[0m": 90000,
	"\033[53mxci\033[0m": 91000, "\033[53mxcii\033[0m": 92000, "\033[53mxciii\033[0m": 93000, "\033[53mxciv\033[0m": 94000, "\033[53mxcv\033[0m": 95000,
	"\033[53mxcvi\033[0m": 96000, "\033[53mxcvii\033[0m": 97000, "\033[53mxcviii\033[0m": 98000, "\033[53mxcix\033[0m": 99000, "\033[53mc\033[0m": 100000,
	"\033[53mci\033[0m": 101000, "\033[53mcii\033[0mcxxvii": 102127, "\033[53mciii\033[0m": 103000, "\033[53mciv\033[0m": 104000, "\033[53mcv\033[0m": 105000,
	"\033[53mcvi\033[0m": 106000, "\033[53mcvii\033[0m": 107000, "\033[53mcviii\033[0m": 108000, "\033[53mcix\033[0m": 109000, "\033[53mcxlix\033[0m": 149000,
	"\033[53mcccxlix\033[0m": 349000, "\033[53mcdlvi\033[0m": 456000, "\033[53md\033[0m": 500000, "\033[53mdciv\033[0m": 604000, "\033[53mdcclxxxix\033[0m": 789000,
	"\033[53mdcccxlix\033[0m": 849000, "\033[53mcmiv\033[0m": 904000, "\033[53mm\033[0m": 1000000, "\033[53mmvii\033[0m": 1007000, "\033[53mmlxvi\033[0m": 1066000,
	"\033[53mmccxxxiv\033[0m": 1234000, "\033[53mmdcclxxvi\033[0m": 1776000, "\033[53mmmxxi\033[0m": 2021000, "\033[53mmmdcccvi\033[0mcdxlii": 2806442,
	"\033[53mmmcmxcix\033[0m": 2999000, "\033[53mmmm\033[0m": 3000000, "\033[53mmmmcmlxxix\033[0m": 3979000, "\033[53mmmmcmxcix\033[0m": 3999000,
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
	out := FromInt(4000000)
	if out != "OOB" {
		t.Errorf("FromInt(%d) expected an error", 4000000)
	}
	out = FromInt(0)
	if out != "0" {
		t.Errorf("FromInt(%d) something went wrong (0 should resolve to 0)", 0)
	}
}

func BenchmarkFromInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = FromInt(3999999)
	}
}

func BenchmarkFromIntCapital(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = FromIntCapital(3999999)
	}
}

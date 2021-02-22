package bitmap64_test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"github.com/danil/iso8583/bitmap64"
)

var testArray = [8]byte{0x46, 0x01, 0xa8, 0xe1, 0xa2, 0x0a, 0xf1, 0xf6}

var NewTestCases = []struct {
	input     []byte
	expected  bitmap64.Bitmap
	line      int
	benchmark bool
}{
	{
		input:     testArray[:],
		expected:  bitmap64.Bitmap(testArray),
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
}

func TestNew(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range NewTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := bitmap64.New(tc.input)
			if b != tc.expected {
				t.Errorf("[bitmap] expected: %#v, received: %#v - %s", tc.expected, b, linkToExample)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range NewTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bitmap64.New(tc.input)
			}
		})
	}
}

func getLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

var BitmapGetTestCases = []struct {
	input     [8]byte
	index     int
	expected  bool
	line      int
	benchmark bool
}{
	{testArray, 1, false, getLine(), true},
	{testArray, 2, true, getLine(), false},
	{testArray, 3, false, getLine(), false},
	{testArray, 4, false, getLine(), false},
	{testArray, 5, false, getLine(), false},
	{testArray, 6, true, getLine(), false},
	{testArray, 7, true, getLine(), false},
	{testArray, 8, false, getLine(), false},
	{testArray, 9, false, getLine(), false},
	{testArray, 10, false, getLine(), false},
	{testArray, 11, false, getLine(), false},
	{testArray, 12, false, getLine(), false},
	{testArray, 13, false, getLine(), false},
	{testArray, 14, false, getLine(), false},
	{testArray, 15, false, getLine(), false},
	{testArray, 16, true, getLine(), false},
	{testArray, 17, true, getLine(), false},
	{testArray, 18, false, getLine(), false},
	{testArray, 19, true, getLine(), false},
	{testArray, 20, false, getLine(), false},
	{testArray, 21, true, getLine(), false},
	{testArray, 22, false, getLine(), false},
	{testArray, 23, false, getLine(), false},
	{testArray, 24, false, getLine(), false},
	{testArray, 25, true, getLine(), false},
	{testArray, 26, true, getLine(), false},
	{testArray, 27, true, getLine(), false},
	{testArray, 28, false, getLine(), false},
	{testArray, 29, false, getLine(), false},
	{testArray, 30, false, getLine(), false},
	{testArray, 31, false, getLine(), false},
	{testArray, 32, true, getLine(), false},
	{testArray, 33, true, getLine(), false},
	{testArray, 34, false, getLine(), false},
	{testArray, 35, true, getLine(), false},
	{testArray, 36, false, getLine(), false},
	{testArray, 37, false, getLine(), false},
	{testArray, 38, false, getLine(), false},
	{testArray, 39, true, getLine(), false},
	{testArray, 40, false, getLine(), false},
	{testArray, 41, false, getLine(), false},
	{testArray, 42, false, getLine(), false},
	{testArray, 43, false, getLine(), false},
	{testArray, 44, false, getLine(), false},
	{testArray, 45, true, getLine(), false},
	{testArray, 46, false, getLine(), false},
	{testArray, 47, true, getLine(), false},
	{testArray, 48, false, getLine(), false},
	{testArray, 49, true, getLine(), false},
	{testArray, 50, true, getLine(), false},
	{testArray, 51, true, getLine(), false},
	{testArray, 52, true, getLine(), false},
	{testArray, 53, false, getLine(), false},
	{testArray, 54, false, getLine(), false},
	{testArray, 55, false, getLine(), false},
	{testArray, 56, true, getLine(), false},
	{testArray, 57, true, getLine(), false},
	{testArray, 58, true, getLine(), false},
	{testArray, 59, true, getLine(), false},
	{testArray, 60, true, getLine(), false},
	{testArray, 61, false, getLine(), false},
	{testArray, 62, true, getLine(), false},
	{testArray, 63, true, getLine(), false},
	{testArray, 64, false, getLine(), false},
}

func TestBitmapGet(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapGetTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%d %t %d", tc.index, tc.expected, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := bitmap64.Bitmap(tc.input)
			ok := b.Get(tc.index)
			if ok != tc.expected {
				t.Errorf("[bit %d] expected: %t, received: %t - %s", tc.index, tc.expected, ok, linkToExample)
			}
		})
	}
}

func BenchmarkBitmapGet(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BitmapGetTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bitmap64.Bitmap(tc.input)
				_ = b.Get(tc.index)
			}
		})
	}
}

var BitmapGetOutOfRangeTestCases = []struct {
	input     [8]byte
	index     int
	expected  interface{}
	line      int
	benchmark bool
}{
	{testArray, 0, "index out of range from 1 to 64", getLine(), true},
	{testArray, 1, nil, getLine(), true},
	{testArray, 64, nil, getLine(), false},
	{testArray, 65, "index out of range from 1 to 64", getLine(), false},
}

func TestBitmapGetOutOfRange(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapGetOutOfRangeTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%d %t %d", tc.index, tc.expected, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			defer func() {
				if p := recover(); p != tc.expected {
					t.Errorf("[panic] expected: %#v, received: %#v - %s", tc.expected, p, linkToExample)
				}
			}()
			b := bitmap64.Bitmap(tc.input)
			b.Get(tc.index)
		})
	}
}

var BitmapSetTestCases = []struct {
	input     [8]byte
	index     int
	line      int
	benchmark bool
}{
	{testArray, 1, getLine(), true},
	{testArray, 2, getLine(), false},
	{testArray, 3, getLine(), false},
	{testArray, 4, getLine(), false},
	{testArray, 5, getLine(), false},
	{testArray, 6, getLine(), false},
	{testArray, 7, getLine(), false},
	{testArray, 8, getLine(), false},
	{testArray, 9, getLine(), false},
	{testArray, 10, getLine(), true},
	{testArray, 11, getLine(), true},
	{testArray, 12, getLine(), false},
	{testArray, 13, getLine(), false},
	{testArray, 14, getLine(), false},
	{testArray, 15, getLine(), false},
	{testArray, 16, getLine(), false},
	{testArray, 17, getLine(), false},
	{testArray, 18, getLine(), false},
	{testArray, 19, getLine(), false},
	{testArray, 20, getLine(), true},
	{testArray, 21, getLine(), true},
	{testArray, 22, getLine(), false},
	{testArray, 23, getLine(), false},
	{testArray, 24, getLine(), false},
	{testArray, 25, getLine(), false},
	{testArray, 26, getLine(), false},
	{testArray, 27, getLine(), false},
	{testArray, 28, getLine(), false},
	{testArray, 29, getLine(), false},
	{testArray, 30, getLine(), true},
	{testArray, 31, getLine(), true},
	{testArray, 32, getLine(), false},
	{testArray, 33, getLine(), false},
	{testArray, 34, getLine(), false},
	{testArray, 35, getLine(), false},
	{testArray, 36, getLine(), false},
	{testArray, 37, getLine(), false},
	{testArray, 38, getLine(), false},
	{testArray, 39, getLine(), false},
	{testArray, 40, getLine(), true},
	{testArray, 41, getLine(), true},
	{testArray, 42, getLine(), false},
	{testArray, 43, getLine(), false},
	{testArray, 44, getLine(), false},
	{testArray, 45, getLine(), false},
	{testArray, 46, getLine(), false},
	{testArray, 47, getLine(), false},
	{testArray, 48, getLine(), false},
	{testArray, 49, getLine(), false},
	{testArray, 50, getLine(), true},
	{testArray, 51, getLine(), true},
	{testArray, 52, getLine(), false},
	{testArray, 53, getLine(), false},
	{testArray, 54, getLine(), false},
	{testArray, 55, getLine(), false},
	{testArray, 56, getLine(), false},
	{testArray, 57, getLine(), false},
	{testArray, 58, getLine(), false},
	{testArray, 59, getLine(), false},
	{testArray, 60, getLine(), false},
	{testArray, 61, getLine(), false},
	{testArray, 62, getLine(), false},
	{testArray, 63, getLine(), false},
	{testArray, 64, getLine(), false},
}

func TestBitmapSet(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapSetTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%d %d", tc.index, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := bitmap64.Bitmap([8]byte{})
			b.Set(tc.index)
			if !b.Get(tc.index) {
				t.Errorf("[bit %d] expected true, but it is false - %s", tc.index, linkToExample)
			}
		})
	}
}

func BenchmarkBitmapSet(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BitmapSetTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := bitmap64.Bitmap(tc.input)
				b.Set(tc.index)
			}
		})
	}
}

var BitmapSetOutOfRangeTestCases = []struct {
	input     [8]byte
	index     int
	expected  interface{}
	line      int
	benchmark bool
}{
	{testArray, 0, "index out of range from 1 to 64", getLine(), true},
	{testArray, 1, nil, getLine(), true},
	{testArray, 42, nil, getLine(), false},
	{testArray, 65, "index out of range from 1 to 64", getLine(), false},
}

func TestBitmapSetOutOfRange(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapSetOutOfRangeTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%d %t %d", tc.index, tc.expected, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			defer func() {
				if p := recover(); p != tc.expected {
					t.Errorf("[panic] expected: %#v, received: %#v - %s", tc.expected, p, linkToExample)
				}
			}()
			b := bitmap64.Bitmap(tc.input)
			b.Set(tc.index)
		})
	}
}

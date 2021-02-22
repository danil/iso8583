package tsysbitmap64r_test

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"testing"

	"github.com/danil/iso8583/tsysbitmap64r"
)

var NewTestCases = []struct {
	name      string
	input     []byte
	expected  map[int]bool
	line      int
	benchmark bool
}{
	{
		name:      "13th bit is set and this is a recurring payment",
		input:     []byte{0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52},
		expected:  map[int]bool{13: true},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		name:      "13th bit is not setted",
		input:     []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		expected:  map[int]bool{13: false},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: false,
	},
}

func TestNew(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range NewTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%s %s", tc.name, strconv.Itoa(tc.line)), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := tsysbitmap64r.New(tc.input)
			for k, v := range tc.expected {
				if b.Get(k) != v {
					t.Errorf("[bit %[1]d] expected: %[1]d=%[2]t, received: %[1]d=%[3]t set - %s", k, v, b.Get(k), linkToExample)
				}
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
				_ = tsysbitmap64r.New(tc.input)
			}
		})
	}
}

var NewStringTestCases = []struct {
	name      string
	input     string
	expected  map[int]bool
	line      int
	benchmark bool
}{
	{
		name:      "13th bit is set and this is a recurring payment",
		input:     "0000000000001000000000000000000000000000000000000000000000000000",
		expected:  map[int]bool{13: true},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		name:      "13th bit is not setted",
		input:     "0000000000000000000000000000000000000000000000000000000000000000",
		expected:  map[int]bool{13: false},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: false,
	},
}

func TestNewString(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range NewStringTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%s %s", tc.name, strconv.Itoa(tc.line)), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b, err := tsysbitmap64r.NewString(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			for k, v := range tc.expected {
				if b.Get(k) != v {
					t.Errorf("[bit %[1]d] expected: %[1]d=%[2]t, received: %[1]d=%[3]t set - %s", k, v, b.Get(k), linkToExample)
				}
			}
		})
	}
}

func BenchmarkNewString(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range NewStringTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := tsysbitmap64r.NewString(tc.input)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func getLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
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
			b := tsysbitmap64r.Bitmap(tc.input)
			b.Get(tc.index)
		})
	}
}

var testArray = [8]byte{0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

var BitmapGetTestCases = []struct {
	input     [8]byte
	index     int
	expected  bool
	line      int
	benchmark bool
}{
	{testArray, 1, false, getLine(), false},
	{testArray, 2, false, getLine(), false},
	{testArray, 3, false, getLine(), false},
	{testArray, 4, false, getLine(), false},
	{testArray, 5, false, getLine(), false},
	{testArray, 6, false, getLine(), false},
	{testArray, 7, false, getLine(), false},
	{testArray, 8, false, getLine(), false},
	{testArray, 9, false, getLine(), false},
	{testArray, 10, false, getLine(), false},
	{testArray, 11, false, getLine(), false},
	{testArray, 12, false, getLine(), false},
	{testArray, 13, true, getLine(), true},
	{testArray, 14, false, getLine(), false},
	{testArray, 15, false, getLine(), false},
	{testArray, 16, false, getLine(), false},
	{testArray, 17, false, getLine(), false},
	{testArray, 18, false, getLine(), false},
	{testArray, 19, false, getLine(), false},
	{testArray, 20, false, getLine(), false},
	{testArray, 21, false, getLine(), false},
	{testArray, 22, false, getLine(), false},
	{testArray, 23, false, getLine(), false},
	{testArray, 24, false, getLine(), false},
	{testArray, 25, false, getLine(), false},
	{testArray, 26, false, getLine(), false},
	{testArray, 27, false, getLine(), false},
	{testArray, 28, false, getLine(), false},
	{testArray, 29, false, getLine(), false},
	{testArray, 30, false, getLine(), false},
	{testArray, 31, false, getLine(), false},
	{testArray, 32, false, getLine(), false},
	{testArray, 33, false, getLine(), false},
	{testArray, 34, false, getLine(), false},
	{testArray, 35, false, getLine(), false},
	{testArray, 36, false, getLine(), false},
	{testArray, 37, false, getLine(), false},
	{testArray, 38, false, getLine(), false},
	{testArray, 39, false, getLine(), false},
	{testArray, 40, false, getLine(), false},
	{testArray, 41, false, getLine(), false},
	{testArray, 42, false, getLine(), false},
	{testArray, 43, false, getLine(), false},
	{testArray, 44, false, getLine(), false},
	{testArray, 45, false, getLine(), false},
	{testArray, 46, false, getLine(), false},
	{testArray, 47, false, getLine(), false},
	{testArray, 48, false, getLine(), false},
	{testArray, 49, false, getLine(), false},
	{testArray, 50, false, getLine(), false},
	{testArray, 51, false, getLine(), false},
	{testArray, 52, false, getLine(), false},
	{testArray, 53, false, getLine(), false},
	{testArray, 54, false, getLine(), false},
	{testArray, 55, false, getLine(), false},
	{testArray, 56, false, getLine(), false},
	{testArray, 57, false, getLine(), false},
	{testArray, 58, false, getLine(), false},
	{testArray, 59, false, getLine(), false},
	{testArray, 60, false, getLine(), false},
	{testArray, 61, false, getLine(), false},
	{testArray, 62, false, getLine(), false},
	{testArray, 63, false, getLine(), false},
	{testArray, 64, false, getLine(), false},
}

func TestBitmapGet(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapGetTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%d %t %d", tc.index, tc.expected, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := tsysbitmap64r.Bitmap(tc.input)
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
				b := tsysbitmap64r.Bitmap(tc.input)
				_ = b.Get(tc.index)
			}
		})
	}
}

var BitmapGetSetTestCases = []struct {
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
	{testArray, 10, getLine(), false},
	{testArray, 11, getLine(), false},
	{testArray, 12, getLine(), false},
	{testArray, 13, getLine(), false},
	{testArray, 14, getLine(), false},
	{testArray, 15, getLine(), false},
	{testArray, 16, getLine(), false},
	{testArray, 17, getLine(), false},
	{testArray, 18, getLine(), false},
	{testArray, 19, getLine(), false},
	{testArray, 20, getLine(), false},
	{testArray, 21, getLine(), false},
	{testArray, 22, getLine(), false},
	{testArray, 23, getLine(), false},
	{testArray, 24, getLine(), false},
	{testArray, 25, getLine(), false},
	{testArray, 26, getLine(), false},
	{testArray, 27, getLine(), false},
	{testArray, 28, getLine(), false},
	{testArray, 29, getLine(), false},
	{testArray, 30, getLine(), false},
	{testArray, 31, getLine(), false},
	{testArray, 32, getLine(), false},
	{testArray, 33, getLine(), false},
	{testArray, 34, getLine(), false},
	{testArray, 35, getLine(), false},
	{testArray, 36, getLine(), false},
	{testArray, 37, getLine(), false},
	{testArray, 38, getLine(), false},
	{testArray, 39, getLine(), false},
	{testArray, 40, getLine(), false},
	{testArray, 41, getLine(), false},
	{testArray, 42, getLine(), false},
	{testArray, 43, getLine(), false},
	{testArray, 44, getLine(), false},
	{testArray, 45, getLine(), false},
	{testArray, 46, getLine(), false},
	{testArray, 47, getLine(), false},
	{testArray, 48, getLine(), false},
	{testArray, 49, getLine(), false},
	{testArray, 50, getLine(), false},
	{testArray, 51, getLine(), false},
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

func TestBitmapGetSet(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapGetSetTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%d %d", tc.index, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := tsysbitmap64r.Bitmap([8]byte{})
			b.Set(tc.index)
			if !b.Get(tc.index) {
				t.Errorf("[bit %d] expected true, but it is false - %s", tc.index, linkToExample)
			}
		})
	}
}

func BenchmarkBitmapGetSet(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BitmapGetSetTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			_, testFile, _, _ := runtime.Caller(0)
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			for i := 0; i < b.N; i++ {
				b := tsysbitmap64r.Bitmap(tc.input)
				b.Set(tc.index)
				if !b.Get(tc.index) {
					log.Fatalf("[bit %d] expected true, but it is false - %s", tc.index, linkToExample)
				}
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
	{testArray, 64, nil, getLine(), false},
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
			b := tsysbitmap64r.Bitmap(tc.input)
			b.Set(tc.index)
		})
	}
}

var BitmapStringTestCases = []struct {
	name      string
	input     [8]byte
	expected  string
	line      int
	benchmark bool
}{
	{
		name:      "13th bit is set and this is a recurring payment",
		input:     [8]byte{0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		expected:  "0000000000001000000000000000000000000000000000000000000000000000",
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		name:      "13th bit is not setted",
		input:     [8]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		expected:  "0000000000000000000000000000000000000000000000000000000000000000",
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: false,
	},
}

func TestBitmapString(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapStringTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%s %s", tc.name, strconv.Itoa(tc.line)), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := tsysbitmap64r.Bitmap(tc.input)
			if b.String() != tc.expected {
				t.Errorf("[string] expected: %q, received: %q - %s", tc.expected, b.String(), linkToExample)
			}
		})
	}
}

func BenchmarkBitmapString(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BitmapStringTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := tsysbitmap64r.Bitmap(tc.input)
				_ = b.String()
			}
		})
	}
}

var BitmapMarshalISO8583TestCases = []struct {
	name      string
	input     [8]byte
	expected  []byte
	line      int
	benchmark bool
}{
	{
		name:      "13th bit is set and this is a recurring payment",
		input:     [8]byte{0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		expected:  []byte{0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		name:      "13th bit is not setted",
		input:     [8]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		expected:  []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: false,
	},
}

func TestBitmapMarshalISO8583(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range BitmapMarshalISO8583TestCases {
		tc := tc
		t.Run(fmt.Sprintf("%s %s", tc.name, strconv.Itoa(tc.line)), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := tsysbitmap64r.Bitmap(tc.input)
			p, err := b.MarshalISO8583()
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if !bytes.Equal(p, tc.expected) {
				t.Errorf("[marshal ISO 8583] expected: %#v, received: %#v - %s", tc.expected, p, linkToExample)
			}
		})
	}
}

func BenchmarkBitmapMarshalISO8583(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range BitmapMarshalISO8583TestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b := tsysbitmap64r.Bitmap(tc.input)
				_, err := b.MarshalISO8583()
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

package codec8583_test

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

func getLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

var testCases = []struct {
	fromEBCDIC string
	toEBCDIC   string
	fromASCII  string
	toASCII    string
	line       int
	benchmark  bool
}{
	{"\x00", "\x00", "\x00", "\x00", getLine(), false},
	{"\x01", "\x01", "\x01", "\x01", getLine(), false},
	{"\x02", "\x02", "\x02", "\x02", getLine(), false},
	{"\x03", "\x03", "\x03", "\x03", getLine(), false},
	{"\x04", "\x04", "\x1a", "\x1a", getLine(), false},
	{"\x05", "\x05", "\t", "\t", getLine(), false},
	{"\x06", "\x06", "\x1a", "\x1a", getLine(), false},
	{"\x07", "\x07", "\x7f", "\x7f", getLine(), false},
	{"\x08", "\x08", "\x1a", "\x1a", getLine(), false},
	{"\x09", "\x09", "\x1a", "\x1a", getLine(), false},
	{"\x0a", "\x0a", "\x1a", "\x1a", getLine(), false},
	{"\x0b", "\x0b", "\v", "\v", getLine(), false},
	{"\x0c", "\x0c", "\f", "\f", getLine(), false},
	{"\x0d", "\x0d", "\r", "\r", getLine(), false},
	{"\x0e", "\x0e", "\x0e", "\x0e", getLine(), false},
	{"\x0f", "\x0f", "\x0f", "\x0f", getLine(), false},
	{"\x10", "\x10", "\x10", "\x10", getLine(), false},
	{"\x11", "\x11", "\x11", "\x11", getLine(), false},
	{"\x12", "\x12", "\x12", "\x12", getLine(), false},
	{"\x13", "\x13", "\x13", "\x13", getLine(), false},
	{"\x14", "\x14", "\x1a", "\x1a", getLine(), false},
	{"\x15", "\x15", "\x85", "\x85", getLine(), false},
	{"\x16", "\x16", "\b", "\b", getLine(), false},
	{"\x17", "\x17", "\x1a", "\x1a", getLine(), false},
	{"\x18", "\x18", "\x18", "\x18", getLine(), false},
	{"\x19", "\x19", "\x19", "\x19", getLine(), false},
	{"\x1a", "\x1a", "\x1a", "\x1a", getLine(), false},
	{"\x1b", "\x1b", "\x1a", "\x1a", getLine(), false},
	{"\x1c", "\x1c", "\x1c", "\x1c", getLine(), false},
	{"\x1d", "\x1d", "\x1d", "\x1d", getLine(), false},
	{"\x1e", "\x1e", "\x1e", "\x1e", getLine(), false},
	{"\x1f", "\x1f", "\x1f", "\x1f", getLine(), false},
	{"\x20", "\x20", "\x1a", "\x1a", getLine(), false},
	{"\x21", "\x21", "\x1a", "\x1a", getLine(), false},
	{"\x22", "\x22", "\x1a", "\x1a", getLine(), false},
	{"\x23", "\x23", "\x1a", "\x1a", getLine(), false},
	{"\x24", "\x24", "\x1a", "\x1a", getLine(), false},
	{"\x25", "\x25", "\n", "\n", getLine(), false},
	{"\x26", "\x26", "\x17", "\x17", getLine(), false},
	{"\x27", "\x27", "\x1b", "\x1b", getLine(), false},
	{"\x28", "\x28", "\x1a", "\x1a", getLine(), false},
	{"\x29", "\x29", "\x1a", "\x1a", getLine(), false},
	{"\x2a", "\x2a", "\x1a", "\x1a", getLine(), false},
	{"\x2b", "\x2b", "\x1a", "\x1a", getLine(), false},
	{"\x2c", "\x2c", "\x1a", "\x1a", getLine(), false},
	{"\x2d", "\x2d", "\x05", "\x05", getLine(), false},
	{"\x2e", "\x2e", "\x06", "\x06", getLine(), false},
	{"\x2f", "\x2f", "\a", "\a", getLine(), false},
	{"\x30", "\x30", "\x1a", "\x1a", getLine(), false},
	{"\x31", "\x31", "\x1a", "\x1a", getLine(), false},
	{"\x32", "\x32", "\x16", "\x16", getLine(), false},
	{"\x33", "\x33", "\x1a", "\x1a", getLine(), false},
	{"\x34", "\x34", "\x1a", "\x1a", getLine(), false},
	{"\x35", "\x35", "\x1a", "\x1a", getLine(), false},
	{"\x36", "\x36", "\x1a", "\x1a", getLine(), false},
	{"\x37", "\x37", "\x04", "\x04", getLine(), false},
	{"\x38", "\x38", "\x1a", "\x1a", getLine(), false},
	{"\x39", "\x39", "\x1a", "\x1a", getLine(), false},
	{"\x3a", "\x3a", "\x1a", "\x1a", getLine(), false},
	{"\x3b", "\x3b", "\x1a", "\x1a", getLine(), false},
	{"\x3c", "\x3c", "\x14", "\x14", getLine(), false},
	{"\x3d", "\x3d", "\x15", "\x15", getLine(), false},
	{"\x3e", "\x3e", "\x1a", "\x1a", getLine(), false},
	{"\x3f", "\x3f", "\x1a", "\x1a", getLine(), false},
	{"\x40", "\x40", " ", " ", getLine(), false},
	{"\x41", "\x41", "\x1a", "\x1a", getLine(), false},
	{"\x42", "\x42", "\x1a", "\x1a", getLine(), false},
	{"\x43", "\x43", "\x1a", "\x1a", getLine(), false},
	{"\x44", "\x44", "\x1a", "\x1a", getLine(), false},
	{"\x45", "\x45", "\x1a", "\x1a", getLine(), false},
	{"\x46", "\x46", "\x1a", "\x1a", getLine(), false},
	{"\x47", "\x47", "\x1a", "\x1a", getLine(), false},
	{"\x48", "\x48", "\x1a", "\x1a", getLine(), false},
	{"\x49", "\x49", "\x1a", "\x1a", getLine(), false},
	{"\x4a", "\x4a", "\xa2", "\xa2", getLine(), false}, // ¢
	{"\x4b", "\x4b", ".", ".", getLine(), false},
	{"\x4c", "\x4c", "<", "<", getLine(), false},
	{"\x4d", "\x4d", "(", "(", getLine(), false},
	{"\x4e", "\x4e", "+", "+", getLine(), false},
	{"\x4f", "\x4f", "|", "|", getLine(), false},
	{"\x50", "\x50", "&", "&", getLine(), false},
	{"\x51", "\x51", "\x1a", "\x1a", getLine(), false},
	{"\x52", "\x52", "\x1a", "\x1a", getLine(), false},
	{"\x53", "\x53", "\x1a", "\x1a", getLine(), false},
	{"\x54", "\x54", "\x1a", "\x1a", getLine(), false},
	{"\x55", "\x55", "\x1a", "\x1a", getLine(), false},
	{"\x56", "\x56", "\x1a", "\x1a", getLine(), false},
	{"\x57", "\x57", "\x1a", "\x1a", getLine(), false},
	{"\x58", "\x58", "\x1a", "\x1a", getLine(), false},
	{"\x59", "\x59", "\x1a", "\x1a", getLine(), false},
	{"\x5a", "\x5a", "!", "!", getLine(), false},
	{"\x5b", "\x5b", "$", "$", getLine(), false},
	{"\x5c", "\x5c", "*", "*", getLine(), false},
	{"\x5d", "\x5d", ")", ")", getLine(), false},
	{"\x5e", "\x5e", ";", ";", getLine(), false},
	{"\x5f", "\x5f", "¬", "¬", getLine(), false},
	{"\x60", "\x60", "-", "-", getLine(), false},
	{"\x61", "\x61", "/", "/", getLine(), false},
	{"\x62", "\x62", "\x1a", "\x1a", getLine(), false},
	{"\x63", "\x63", "\x1a", "\x1a", getLine(), false},
	{"\x64", "\x64", "\x1a", "\x1a", getLine(), false},
	{"\x65", "\x65", "\x1a", "\x1a", getLine(), false},
	{"\x66", "\x66", "\x1a", "\x1a", getLine(), false},
	{"\x67", "\x67", "\x1a", "\x1a", getLine(), false},
	{"\x68", "\x68", "\x1a", "\x1a", getLine(), false},
	{"\x69", "\x69", "\x1a", "\x1a", getLine(), false},
	{"\x6a", "\x6a", "¦", "¦", getLine(), false},
	{"\x6b", "\x6b", ",", ",", getLine(), false},
	{"\x6c", "\x6c", "%", "%", getLine(), false},
	{"\x6d", "\x6d", "_", "_", getLine(), false},
	{"\x6e", "\x6e", ">", ">", getLine(), false},
	{"\x6f", "\x6f", "?", "?", getLine(), false},
	{"\x70", "\x70", "\x1a", "\x1a", getLine(), false},
	{"\x71", "\x71", "\x1a", "\x1a", getLine(), false},
	{"\x72", "\x72", "\x1a", "\x1a", getLine(), false},
	{"\x73", "\x73", "\x1a", "\x1a", getLine(), false},
	{"\x74", "\x74", "\x1a", "\x1a", getLine(), false},
	{"\x75", "\x75", "\x1a", "\x1a", getLine(), false},
	{"\x76", "\x76", "\x1a", "\x1a", getLine(), false},
	{"\x77", "\x77", "\x1a", "\x1a", getLine(), false},
	{"\x78", "\x78", "\x1a", "\x1a", getLine(), false},
	{"\x79", "\x79", "`", "`", getLine(), false},
	{"\x7a", "\x7a", ":", ":", getLine(), false},
	{"\x7b", "\x7b", "#", "#", getLine(), false},
	{"\x7c", "\x7c", "@", "@", getLine(), false},
	{"\x7d", "\x7d", "'", "'", getLine(), false},
	{"\x7e", "\x7e", "=", "=", getLine(), false},
	{"\x7f", "\x7f", `"`, `"`, getLine(), false},
	{"\x80", "\x80", "\x1a", "\x1a", getLine(), false},
	{"\x81", "\x81", "a", "a", getLine(), false},
	{"\x82", "\x82", "b", "b", getLine(), false},
	{"\x83", "\x83", "c", "c", getLine(), false},
	{"\x84", "\x84", "d", "d", getLine(), false},
	{"\x85", "\x85", "e", "e", getLine(), false},
	{"\x86", "\x86", "f", "f", getLine(), false},
	{"\x87", "\x87", "g", "g", getLine(), false},
	{"\x88", "\x88", "h", "h", getLine(), false},
	{"\x89", "\x89", "i", "i", getLine(), false},
	{"\x8a", "\x8a", "\x1a", "\x1a", getLine(), false},
	{"\x8b", "\x8b", "\x1a", "\x1a", getLine(), false},
	{"\x8c", "\x8c", "\x1a", "\x1a", getLine(), false},
	{"\x8d", "\x8d", "\x1a", "\x1a", getLine(), false},
	{"\x8e", "\x8e", "\x1a", "\x1a", getLine(), false},
	{"\x8f", "\x8f", "±", "±", getLine(), false},
	{"\x90", "\x90", "\x1a", "\x1a", getLine(), false},
	{"\x91", "\x91", "j", "j", getLine(), false},
	{"\x92", "\x92", "k", "k", getLine(), false},
	{"\x93", "\x93", "l", "l", getLine(), false},
	{"\x94", "\x94", "m", "m", getLine(), false},
	{"\x95", "\x95", "n", "n", getLine(), false},
	{"\x96", "\x96", "o", "o", getLine(), false},
	{"\x97", "\x97", "p", "p", getLine(), false},
	{"\x98", "\x98", "q", "q", getLine(), false},
	{"\x99", "\x99", "r", "r", getLine(), false},
	{"\x9a", "\x9a", "\x1a", "\x1a", getLine(), false},
	{"\x9b", "\x9b", "\x1a", "\x1a", getLine(), false},
	{"\x9c", "\x9c", "\x1a", "\x1a", getLine(), false},
	{"\x9d", "\x9d", "\x1a", "\x1a", getLine(), false},
	{"\x9e", "\x9e", "\x1a", "\x1a", getLine(), false},
	{"\x9f", "\x9f", "\x1a", "\x1a", getLine(), false},
	{"\xa0", "\xa0", "\x1a", "\x1a", getLine(), false},
	{"\xa1", "\xa1", "~", "~", getLine(), false},
	{"\xa2", "\xa2", "s", "s", getLine(), false},
	{"\xa3", "\xa3", "t", "t", getLine(), false},
	{"\xa4", "\xa4", "u", "u", getLine(), false},
	{"\xa5", "\xa5", "v", "v", getLine(), false},
	{"\xa6", "\xa6", "w", "w", getLine(), false},
	{"\xa7", "\xa7", "x", "x", getLine(), false},
	{"\xa8", "\xa8", "y", "y", getLine(), false},
	{"\xa9", "\xa9", "z", "z", getLine(), false},
	{"\xaa", "\xaa", "\x1a", "\x1a", getLine(), false},
	{"\xab", "\xab", "\x1a", "\x1a", getLine(), false},
	{"\xac", "\xac", "\x1a", "\x1a", getLine(), false},
	{"\xad", "\xad", "\x1a", "\x1a", getLine(), false},
	{"\xae", "\xae", "\x1a", "\x1a", getLine(), false},
	{"\xaf", "\xaf", "\x1a", "\x1a", getLine(), false},
	{"\xb0", "\xb0", "^", "^", getLine(), false},
	{"\xb1", "\xb1", "\x1a", "\x1a", getLine(), false},
	{"\xb2", "\xb2", "\x1a", "\x1a", getLine(), false},
	{"\xb3", "\xb3", "\x1a", "\x1a", getLine(), false},
	{"\xb4", "\xb4", "\x1a", "\x1a", getLine(), false},
	{"\xb5", "\xb5", "\x1a", "\x1a", getLine(), false},
	{"\xb6", "\xb6", "\x1a", "\x1a", getLine(), false},
	{"\xb7", "\xb7", "\x1a", "\x1a", getLine(), false},
	{"\xb8", "\xb8", "\x1a", "\x1a", getLine(), false},
	{"\xb9", "\xb9", "\x1a", "\x1a", getLine(), false},
	{"\xba", "\xba", "[", "[", getLine(), false},
	{"\xbb", "\xbb", "]", "]", getLine(), false},
	{"\xbc", "\xbc", "\x1a", "\x1a", getLine(), false},
	{"\xbd", "\xbd", "\x1a", "\x1a", getLine(), false},
	{"\xbe", "\xbe", "\x1a", "\x1a", getLine(), false},
	{"\xbf", "\xbf", "\x1a", "\x1a", getLine(), false},
	{"\xc0", "\xc0", "{", "{", getLine(), false},
	{"\xc1", "\xc1", "A", "A", getLine(), false},
	{"\xc2", "\xc2", "B", "B", getLine(), false},
	{"\xc3", "\xc3", "C", "C", getLine(), false},
	{"\xc4", "\xc4", "D", "D", getLine(), false},
	{"\xc5", "\xc5", "E", "E", getLine(), false},
	{"\xc6", "\xc6", "F", "F", getLine(), false},
	{"\xc7", "\xc7", "G", "G", getLine(), false},
	{"\xc8", "\xc8", "H", "H", getLine(), false},
	{"\xc9", "\xc9", "I", "I", getLine(), false},
	{"\xca", "\xca", "\x1a", "\x1a", getLine(), false},
	{"\xcb", "\xcb", "\x1a", "\x1a", getLine(), false},
	{"\xcc", "\xcc", "\x1a", "\x1a", getLine(), false},
	{"\xcd", "\xcd", "\x1a", "\x1a", getLine(), false},
	{"\xce", "\xce", "\x1a", "\x1a", getLine(), false},
	{"\xcf", "\xcf", "\x1a", "\x1a", getLine(), false},
	{"\xd0", "\xd0", "}", "}", getLine(), false},
	{"\xd1", "\xd1", "J", "J", getLine(), false},
	{"\xd2", "\xd2", "K", "K", getLine(), false},
	{"\xd3", "\xd3", "L", "L", getLine(), false},
	{"\xd4", "\xd4", "M", "M", getLine(), false},
	{"\xd5", "\xd5", "N", "N", getLine(), false},
	{"\xd6", "\xd6", "O", "O", getLine(), false},
	{"\xd7", "\xd7", "P", "P", getLine(), false},
	{"\xd8", "\xd8", "Q", "Q", getLine(), false},
	{"\xd9", "\xd9", "R", "R", getLine(), false},
	{"\xda", "\xda", "\x1a", "\x1a", getLine(), false},
	{"\xdb", "\xdb", "\x1a", "\x1a", getLine(), false},
	{"\xdc", "\xdc", "\x1a", "\x1a", getLine(), false},
	{"\xdd", "\xdd", "\x1a", "\x1a", getLine(), false},
	{"\xde", "\xde", "\x1a", "\x1a", getLine(), false},
	{"\xdf", "\xdf", "\x1a", "\x1a", getLine(), false},
	{"\xe0", "\xe0", "\\", "\\", getLine(), false},
	{"\xe1", "\xe1", "\x1a", "\x1a", getLine(), false},
	{"\xe2", "\xe2", "S", "S", getLine(), false},
	{"\xe3", "\xe3", "T", "T", getLine(), false},
	{"\xe4", "\xe4", "U", "U", getLine(), false},
	{"\xe5", "\xe5", "V", "V", getLine(), false},
	{"\xe6", "\xe6", "W", "W", getLine(), false},
	{"\xe7", "\xe7", "X", "X", getLine(), false},
	{"\xe8", "\xe8", "Y", "Y", getLine(), false},
	{"\xe9", "\xe9", "Z", "Z", getLine(), false},
	{"\xea", "\xea", "\x1a", "\x1a", getLine(), false},
	{"\xeb", "\xeb", "\x1a", "\x1a", getLine(), false},
	{"\xec", "\xec", "\x1a", "\x1a", getLine(), false},
	{"\xed", "\xed", "\x1a", "\x1a", getLine(), false},
	{"\xee", "\xee", "\x1a", "\x1a", getLine(), false},
	{"\xef", "\xef", "\x1a", "\x1a", getLine(), false},
	{"\xf0", "\xf0", "0", "0", getLine(), false},
	{"\xf1", "\xf1", "1", "1", getLine(), false},
	{"\xf2", "\xf2", "2", "2", getLine(), false},
	{"\xf3", "\xf3", "3", "3", getLine(), false},
	{"\xf4", "\xf4", "4", "4", getLine(), false},
	{"\xf5", "\xf5", "5", "5", getLine(), false},
	{"\xf6", "\xf6", "6", "6", getLine(), false},
	{"\xf7", "\xf7", "7", "7", getLine(), false},
	{"\xf8", "\xf8", "8", "8", getLine(), false},
	{"\xf9", "\xf9", "9", "9", getLine(), false},
	{"\xfa", "\xfa", "\x1a", "\x1a", getLine(), false},
	{"\xfb", "\xfb", "\x1a", "\x1a", getLine(), false},
	{"\xfc", "\xfc", "\x1a", "\x1a", getLine(), false},
	{"\xfd", "\xfd", "\x1a", "\x1a", getLine(), false},
	{"\xfe", "\xfe", "\x1a", "\x1a", getLine(), false},
	{"\xff", "\xff", "\x1a", "\x1a", getLine(), false},
	{"\xc8\x85\x93\x93\x96\x6b\x40\xe6\x96\x99\x93\x84\x5a", "\xc8\x85\x93\x93\x96\x6b\x40\xe6\x96\x99\x93\x84\x5a", "Hello, World!", "Hello, World!", getLine(), true},
}

func TestXTextEncodingCodePage037Decode(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range testCases {
		tc := tc
		switch tc.toASCII {
		case string(encoding.ASCIISub), "\x85", "\xad", "\xa0", "\xa2":
			continue
		}
		t.Run(fmt.Sprintf("%#v %#v %d", tc.fromEBCDIC, tc.toASCII, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			decoder := charmap.CodePage037.NewDecoder()
			dst := make([]byte, len(tc.toASCII))
			nDst, nSrc, err := decoder.Transform(dst, []byte(tc.fromEBCDIC), true)
			if err != nil {
				t.Fatalf("unexpected decoder transform error: %v - %s", err, linkToExample)
			}
			if nSrc != len(tc.fromEBCDIC) {
				t.Errorf("unexpected source bytes number: %d, expected %d %s",
					nSrc, len(tc.fromEBCDIC), linkToExample)
			}
			if nDst != len(tc.toASCII) {
				t.Errorf("unexpected destination bytes number: %d, expected %d %s",
					nDst, len(tc.toASCII), linkToExample)
			}
			if string(dst[:nDst]) != tc.toASCII {
				t.Errorf("unexpected UTF-8 string: %q, expected: %q %s", string(dst[:nDst]), tc.toASCII, linkToExample)
			}
		})
	}
}

func BenchmarkXTextEncodingCodePage037Decode(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range testCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("%#v %#v %d", tc.fromEBCDIC, tc.toASCII, tc.line), func(b *testing.B) {
			decoder := charmap.CodePage037.NewDecoder()
			for i := 0; i < b.N; i++ {
				dst := make([]byte, len(tc.toASCII))
				_, _, err := decoder.Transform(dst, []byte(tc.fromEBCDIC), true)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func TestXTextEncodingCodePage037Encode(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range testCases {
		tc := tc
		switch tc.fromASCII {
		case string(encoding.ASCIISub), "\x85", "\xad", "\xa0", "\xa2":
			continue
		}
		t.Run(fmt.Sprintf("%#v %#v %d", tc.fromASCII, tc.toEBCDIC, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			encoder := charmap.CodePage037.NewEncoder()
			dst := make([]byte, len(tc.toEBCDIC))
			nDst, nSrc, err := encoder.Transform(dst, []byte(tc.fromASCII), true)
			if err != nil {
				t.Fatalf("unexpected encoder transform error: %v - %s", err, linkToExample)
			}
			if nSrc != len(tc.fromASCII) {
				t.Errorf("unexpected source bytes number: %d, expected: %d %s",
					nSrc, len(tc.fromASCII), linkToExample)
			}
			if nDst != len(tc.toEBCDIC) {
				t.Errorf("unexpected destination bytes number: %d, expected: %d %s",
					nDst, len(tc.toEBCDIC), linkToExample)
			}
			if string(dst[:nDst]) != tc.toEBCDIC {
				t.Errorf("unexpected EBCDIC string: %#v, expected: %#v %s", string(dst[:nDst]), tc.toEBCDIC, linkToExample)
			}
		})
	}
}

func BenchmarkXTextEncodingCodePage037Encode(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range testCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("%#v %#v %d", tc.fromASCII, tc.toEBCDIC, tc.line), func(b *testing.B) {
			encoder := charmap.CodePage037.NewEncoder()
			for i := 0; i < b.N; i++ {
				dst := make([]byte, len(tc.toEBCDIC))
				_, _, err := encoder.Transform(dst, []byte(tc.fromASCII), true)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

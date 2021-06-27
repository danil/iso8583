package scan8583_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"github.com/danil/iso8583/codec8583"
	"github.com/danil/iso8583/scan8583"
	"github.com/protoscan/protoscan"
)

var ScannerScanTestCases = []struct {
	line      int
	split     protoscan.SplitFunc
	input     []byte
	expected  int
	benchmark bool
}{
	{
		line:      line(),
		split:     scan8583.ScanISO8583Indiscriminately,
		input:     payload8583(tsysMsgs[0], tsysMsgs[1], tsysMsgs[2], tsysMsgs[3]),
		expected:  4,
		benchmark: true,
	},
	{
		line:     line(),
		split:    scan8583.ScanISO8583Indiscriminately,
		input:    payload8583(tsysMsgs[3]),
		expected: 1,
	},
	{
		line:     line(),
		split:    scan8583.ScanISO8583Indiscriminately,
		input:    []byte{},
		expected: 0,
	},
}

func TestScannerScan(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range ScannerScanTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			r := bytes.NewReader(tc.input)
			s := protoscan.Protoscan{
				Reader: r,
				Split:  tc.split,
			}
			var (
				count   int
				payload []byte
			)
			for s.Scan() {
				count++
				payload = append(payload, s.Tokens[:s.Indexes[0]]...)
			}
			err := s.Err()
			if err != nil {
				t.Fatalf("unexpected scan error: %v %s", err, linkToExample)
			}
			if !bytes.Equal(payload, tc.input) {
				t.Errorf("unexpected payload, expected: %#v, received: %#v %s", tc.input, payload, linkToExample)
			}
			if len(payload) != len(tc.input) {
				t.Errorf("unexpected payload length, expected: %#v, received: %#v %s", len(tc.input), len(payload), linkToExample)
			}
			if count != tc.expected {
				t.Errorf("unexpected token count, expected: %#v, received: %#v %s", tc.expected, count, linkToExample)
			}
		})
	}
}

var tsysValidationTestCases = []struct {
	line      int
	input     []byte
	benchmark bool
}{
	{
		line:  line(),
		input: payload8583(tsysMsgs[0], tsysMsgs[1], tsysMsgs[2], tsysMsgs[3]),
	},
	{
		line:  line(),
		input: payload8583(tsysMsgs[3]),
	},
	{
		line:  line(),
		input: []byte{},
	},
}

func TestTSYSValidation(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range tsysValidationTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			r := bytes.NewReader(tc.input)
			s := protoscan.Protoscan{
				Reader: r,
				Split:  scan8583.ScanISO8583Indiscriminately,
			}
			var (
				tokens  [][]byte
				payload []byte
			)
			for s.Scan() {
				token := append([]byte{}, s.Tokens[:s.Indexes[0]]...)
				tokens = append(tokens, token)
				payload = append(payload, token...)
			}
			err := s.Err()
			if err != nil {
				t.Fatalf("unexpected scan error: %v %s", err, linkToExample)
			}
			if !bytes.Equal(payload, tc.input) {
				t.Errorf("unexpected payload, expected: %#v, received: %#v %s", tc.input, payload, linkToExample)
			}
			if len(payload) != len(tc.input) {
				t.Errorf("unexpected payload length, expected: %#v, received: %#v %s", len(tc.input), len(payload), linkToExample)
			}
			for i, token := range tokens {
				message := token[scan8583.ISO8583HeadSize:]
				tr := tsysMessage{}
				err := codec8583.TSYSUnmarshaler.Unmarshal(message, &tr)
				if err != nil {
					t.Errorf("unexpected error on unmarshaling tsys message %d: %v %s",
						i, err, linkToExample)
				}
			}
		})
	}
}

func BenchmarkScannerScan(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range ScannerScanTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := bytes.NewReader(tc.input)
				s := protoscan.Protoscan{
					Reader: r,
					Split:  tc.split,
				}
				for s.Scan() {
				}
			}
		})
	}
}

func line() int { _, _, l, _ := runtime.Caller(1); return l }

var tsysMsgs = [][]byte{
	[]byte{0x30, 0x31, 0x30, 0x30, 0xf6, 0x7e, 0x44, 0x91, 0x2c, 0xe0, 0xb0, 0x18, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x33, 0x34, 0x32, 0x39, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x31, 0x38, 0x31, 0x37, 0x30, 0x34, 0x34, 0x33, 0x36, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x34, 0x31, 0x37, 0x30, 0x34, 0x34, 0x33, 0x31, 0x32, 0x31, 0x38, 0x32, 0x32, 0x31, 0x32, 0x31, 0x32, 0x31, 0x38, 0x36, 0x30, 0x31, 0x31, 0x39, 0x30, 0x31, 0x30, 0x30, 0x30, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x36, 0x39, 0x39, 0x39, 0x39, 0x30, 0x31, 0x33, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x31, 0x34, 0x30, 0x36, 0x30, 0x31, 0x32, 0x38, 0x4d, 0x54, 0x46, 0x20, 0x54, 0x45, 0x53, 0x54, 0x41, 0x42, 0x43, 0x31, 0x32, 0x33, 0x54, 0x45, 0x53, 0x54, 0x4d, 0x54, 0x46, 0x31, 0x39, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x20, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x20, 0x20, 0x20, 0x4d, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x74, 0x65, 0x72, 0x20, 0x20, 0x20, 0x20, 0x46, 0x38, 0x34, 0x30, 0x38, 0x34, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x38, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4d, 0x41, 0x50, 0x49, 0x4d, 0x41, 0x53, 0x54, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x38, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x37, 0x38, 0x55, 0x44, 0x30, 0x37, 0x33, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x30, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x32, 0x30, 0x35, 0x30, 0x31, 0x30, 0x30, 0x36, 0x30, 0x31, 0x32, 0x30, 0x37, 0x30, 0x31, 0x32, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x43, 0x53, 0x30, 0x30, 0x34, 0x33, 0x30, 0x38, 0x31, 0x41, 0x44, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x35, 0x39, 0x37, 0x31, 0x31, 0x31},
	[]byte{0x30, 0x31, 0x30, 0x30, 0xf6, 0x7e, 0x44, 0x81, 0x2c, 0xe0, 0xa0, 0x18, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x33, 0x34, 0x32, 0x39, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x30, 0x30, 0x30, 0x31, 0x32, 0x31, 0x34, 0x31, 0x38, 0x30, 0x36, 0x34, 0x30, 0x36, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x30, 0x33, 0x31, 0x38, 0x30, 0x36, 0x34, 0x30, 0x31, 0x32, 0x31, 0x34, 0x32, 0x32, 0x31, 0x32, 0x31, 0x32, 0x31, 0x34, 0x35, 0x39, 0x39, 0x39, 0x39, 0x30, 0x33, 0x30, 0x30, 0x30, 0x30, 0x36, 0x39, 0x39, 0x39, 0x30, 0x33, 0x37, 0x33, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x36, 0x32, 0x30, 0x34, 0x32, 0x37, 0x30, 0x30, 0x4d, 0x54, 0x46, 0x20, 0x54, 0x45, 0x53, 0x54, 0x41, 0x42, 0x43, 0x31, 0x32, 0x33, 0x54, 0x45, 0x53, 0x54, 0x4d, 0x54, 0x46, 0x31, 0x39, 0x4d, 0x69, 0x73, 0x63, 0x20, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x56, 0x65, 0x67, 0x61, 0x73, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x4e, 0x56, 0x38, 0x34, 0x30, 0x38, 0x34, 0x30, 0x30, 0x38, 0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4d, 0x41, 0x50, 0x49, 0x4d, 0x41, 0x53, 0x54, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x38, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x37, 0x38, 0x55, 0x44, 0x30, 0x37, 0x33, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x30, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x32, 0x30, 0x35, 0x30, 0x31, 0x30, 0x30, 0x36, 0x30, 0x31, 0x30, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x43, 0x53, 0x30, 0x30, 0x34, 0x33, 0x30, 0x38, 0x31, 0x41, 0x44, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x35, 0x39, 0x36, 0x36, 0x33, 0x32},
	[]byte{0x30, 0x31, 0x30, 0x30, 0xf2, 0x3a, 0x44, 0x81, 0x2c, 0xe0, 0x93, 0x10, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x32, 0x35, 0x31, 0x34, 0x32, 0x35, 0x34, 0x35, 0x39, 0x32, 0x32, 0x32, 0x37, 0x34, 0x31, 0x37, 0x32, 0x37, 0x30, 0x39, 0x30, 0x35, 0x32, 0x35, 0x30, 0x35, 0x32, 0x35, 0x36, 0x30, 0x31, 0x31, 0x30, 0x35, 0x31, 0x30, 0x30, 0x32, 0x30, 0x36, 0x39, 0x39, 0x39, 0x39, 0x30, 0x35, 0x33, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x31, 0x34, 0x35, 0x31, 0x39, 0x39, 0x32, 0x32, 0x32, 0x37, 0x34, 0x30, 0x35, 0x34, 0x32, 0x35, 0x34, 0x30, 0x30, 0x30, 0x37, 0x37, 0x38, 0x31, 0x30, 0x30, 0x30, 0x30, 0x37, 0x37, 0x38, 0x31, 0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x41, 0x54, 0x4d, 0x20, 0x4e, 0x45, 0x4f, 0x50, 0x41, 0x4c, 0x49, 0x4d, 0x4f, 0x56, 0x53, 0x4b, 0x59, 0x20, 0x31, 0x30, 0x20, 0x20, 0x20, 0x20, 0x3e, 0x4d, 0x4f, 0x53, 0x43, 0x4f, 0x57, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x55, 0x36, 0x34, 0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x38, 0x38, 0x9f, 0x26, 0x8, 0x33, 0xe, 0xc5, 0x84, 0x6d, 0x83, 0xa4, 0xea, 0x9f, 0x10, 0x12, 0x1, 0x10, 0xa0, 0x0, 0x3, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x9f, 0x37, 0x4, 0x76, 0x40, 0x6, 0x37, 0x9f, 0x36, 0x2, 0x0, 0x5d, 0x95, 0x5, 0x80, 0x0, 0x4, 0x0, 0x0, 0x9a, 0x3, 0x17, 0x5, 0x25, 0x9c, 0x1, 0x1, 0x9f, 0x2, 0x6, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x5f, 0x2a, 0x2, 0x9, 0x78, 0x82, 0x2, 0x39, 0x0, 0x9f, 0x1a, 0x2, 0x6, 0x43, 0x9f, 0x33, 0x3, 0x60, 0x40, 0x20, 0x31, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x34, 0x32, 0x33, 0x30, 0x38, 0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4f, 0x43, 0x54, 0x49, 0x44, 0x48, 0x49, 0x31, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x36, 0x32, 0x55, 0x44, 0x30, 0x35, 0x37, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x32, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x32, 0x30, 0x35, 0x30, 0x31, 0x32, 0x30, 0x36, 0x30, 0x31, 0x32, 0x30, 0x37, 0x30, 0x31, 0x32, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x43, 0x53, 0x30, 0x30, 0x32, 0x34, 0x31},
	[]byte{0x30, 0x34, 0x32, 0x31, 0xf6, 0x38, 0x44, 0x81, 0xe, 0xe0, 0xa1, 0x18, 0x0, 0x0, 0x0, 0x42, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31, 0x33, 0x31, 0x34, 0x33, 0x36, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x34, 0x31, 0x33, 0x35, 0x32, 0x35, 0x33, 0x30, 0x35, 0x31, 0x36, 0x36, 0x30, 0x31, 0x30, 0x39, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x36, 0x39, 0x39, 0x39, 0x39, 0x30, 0x35, 0x37, 0x31, 0x33, 0x36, 0x31, 0x39, 0x39, 0x30, 0x39, 0x34, 0x34, 0x38, 0x30, 0x35, 0x38, 0x31, 0x33, 0x32, 0x30, 0x30, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x20, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3e, 0x4d, 0x6f, 0x73, 0x63, 0x6f, 0x77, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x55, 0x36, 0x34, 0x33, 0x36, 0x34, 0x33, 0x32, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x33, 0x39, 0x31, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x33, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x34, 0x30, 0x36, 0x31, 0x33, 0x31, 0x34, 0x33, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x39, 0x39, 0x39, 0x39, 0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4f, 0x43, 0x54, 0x49, 0x44, 0x48, 0x49, 0x31, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x36, 0x37, 0x55, 0x44, 0x30, 0x36, 0x32, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x30, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x30, 0x30, 0x35, 0x30, 0x31, 0x30, 0x30, 0x36, 0x30, 0x31, 0x30, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x49, 0x52, 0x30, 0x30, 0x30, 0x43, 0x53, 0x30, 0x30, 0x32, 0x34, 0x31},
}

func payload8583(messages ...[]byte) []byte {
	var s []byte
	for _, msg := range messages {
		msgLen := make([]byte, 2)
		binary.BigEndian.PutUint16(msgLen, uint16(len(msg)))
		s = append(s, append(msgLen, msg...)...)
	}
	return s
}

type tsysMessage struct {
	MTI                                string `iso8583:"MTI"`
	PrimaryAccountNumber               string `iso8583:"2"`
	ProcessingCode                     string `iso8583:"3"`
	AmountOrig                         string `iso8583:"4"`
	Amount                             string `iso8583:"6"`
	TransmissionDateTime               string `iso8583:"7"`
	BillingRate                        string `iso8583:"10"`
	TraceNumber                        string `iso8583:"11"`
	LocalTime                          string `iso8583:"12"`
	LocalDate                          string `iso8583:"13"`
	DateExpiration                     string `iso8583:"14"`
	DateSettlement                     string `iso8583:"15"`
	DateCapture                        string `iso8583:"17"`
	MerchantType                       string `iso8583:"18"`
	AcquiringInstitutionCountryCode    string `iso8583:"19"`
	POSDataCode                        string `iso8583:"22"`
	PointOfServiceConditionCode        string `iso8583:"25"`
	TransactionFee                     string `iso8583:"28"`
	ONLINEIssuerAuthorizationFeeAmount string `iso8583:"31"`
	AcquirerInstitutionID              string `iso8583:"32"`
	TrackData                          string `iso8583:"35"`
	RetrievalReference                 string `iso8583:"37"`
	AuthIDCode                         string `iso8583:"38"`
	RespCode                           string `iso8583:"39"`
	CardAccptrTermnlID                 string `iso8583:"41"`
	CardAccptrIDCode                   string `iso8583:"42"`
	CardAccptrNameLoc                  string `iso8583:"43"`
	AdditionalResponseData             string `iso8583:"44"`
	CurrencyOrig                       string `iso8583:"49"`
	Currency                           string `iso8583:"51"`
	PersonalIdentificationNumberData   string `iso8583:"52"`
	SecurityRelatedControlInformation  string `iso8583:"53"`
	AddtnlAmounts                      string `iso8583:"54"`
	ICCRelatedData                     string `iso8583:"55"`
	OriginalDataSerials                string `iso8583:"56"`
	AdditionalInformation              string `iso8583:"60"`
	OtherAmtTrans                      string `iso8583:"61"`
	NetworkManagementInformationCode   string `iso8583:"70"`
	BusinessDate                       string `iso8583:"73"`
	OrigDataElemts                     string `iso8583:"90"`
	NumberOfAccounts                   string `iso8583:"93"`
	QuerySequence                      string `iso8583:"94"`
	ReplacementAmount                  string `iso8583:"95"`
	MoreFlag                           string `iso8583:"99"`
	MessageOriginator                  string `iso8583:"100"`
	AccountFrom                        string `iso8583:"102"`
	AccountTo                          string `iso8583:"103"`
	PrivateData                        string `iso8583:"104"`
	AdditionalInformationPart2         string `iso8583:"116"`
	AdditionalAmountAccountTo          string `iso8583:"117"`
	AdditionalInformationPart1         string `iso8583:"120"`
	Transfercurrencies                 string `iso8583:"122"`
	CardholderUtilityAccount           string `iso8583:"125"`
	PrivateUseFields                   string `iso8583:"126"`
}

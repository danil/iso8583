package scan8583

import (
	"encoding/binary"
	"errors"
)

// iso8583Size is the number of bytes in which are encoded in Big-endian,
// unsigned number that determines the size of the ISO 8583 message.
const iso8583Size = 2

// ScanISO8583Indiscriminately is a split function for the ISO 8583 message format.
// Split function returns hint is a number of bytes hinted to read and
// returns advance is a needed number of bytes by which the carriage is to shift
// and returns a token and an error if occurs.
// Each token is an ISO 8583 message size plus message itself.
// The returned token may holds invalid or inconsistent
// or incomplete ISO 8583 message or holds not ISO 8583 at all
// because this split function do not performs any message validation.
func ScanISO8583Indiscriminately(data []byte, _atEOF bool) (int, [][]byte, int, []byte, error) {
	if len(data) == 0 {
		return iso8583Size, nil, 0, nil, nil
	}
	size := int(binary.BigEndian.Uint16(data[:iso8583Size]))
	if len(data) == iso8583Size {
		return size, nil, 0, nil, nil
	}
	if len(data) > size+iso8583Size {
		return 0, nil, 0, nil, errors.New("buffer exceeds hinted size of the ISO 8583 token")
	}
	if len(data) < size+iso8583Size {
		return size + iso8583Size - len(data), nil, 0, nil, nil
	}
	return 0, nil, len(data), data, nil
}

// ScanISO8583 is a split function for a Protoscan that returns each
// ISO 8583 message size plus message itself as a token.
// The returned token intends to be valid and consistent
// and complete ISO 8583 message
// but unfortunately this split function is not implemented yet)
func ScanISO8583(data []byte) (int, int, []byte, error) {
	panic("not implemented")
}

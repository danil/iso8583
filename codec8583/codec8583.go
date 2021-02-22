package codec8583

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"regexp"
	"strconv"
	"sync"
)

// Format maps codes to the ISO 8583 message MIT/bitmaps/each individual field.
type Format map[int]Codec

var (
	// MTIAsciiCodec encodes/decodes ISO 8583 MTI (Message Type Indicator)
	// the size of which is 4 bytes in ASCII character encoding.
	MTIAsciiCodec = FIX{4, ASCII, EncodeMTI, DecodeMTI}
	// MTIEbcdicCodec encodes/decodes ISO 8583 MTI (Message Type Indicator)
	// the size of which is 4 bytes in EBCDIC character encoding.
	MTIEbcdicCodec = FIX{4, EBCDIC, EncodeMTI, DecodeMTI}
	// BitmapCodec encodes/decodes ISO 8583 bitmap the size of which is 8 bytes (64 bit).
	BitmapCodec = FIX{8, NOPCharset, EncBitmap, DecBitmap}
)

type (
	// EncodeFunc encodes ISO 8583 MTI/bitmaps/fields.
	EncodeFunc func(Hasher, Codec, []byte) ([]byte, error)
	// DecodeFunc decodes ISO 8583 MTI/bitmaps/fields.
	DecodeFunc func(Hasher, Codec, []byte) ([]byte, error)
)

// EncodeMTI intends to encode the ISO 8583 MTI (Message Type Indicator).
func EncodeMTI(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncBitmap intends to encode the ISO 8583 bitmap.
func EncBitmap(_ Hasher, _ Codec, dec []byte) ([]byte, error) { return dec, nil }

// EncA intends to encode Alpha, including blanks <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncA(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncN intends to encode numeric values only <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncN(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncXN intends to encode numeric (amount) values,
// where the first byte is either 'C' to indicate a positive or Credit value,
// or 'D' to indicate a negative or Debit value,
// followed by the numeric value (using n digits)
// <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncXN(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncNS intends to encode numeric and special characters only <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncNS(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncAN intends to encode alphanumeric <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncAN(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncANS intends to encode alphabetic, numeric and special characters <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncANS(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncZ intends to encode tracks 2 and 3 code set as defined in ISO/IEC 7813 and ISO/IEC 4909 respectively
// <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>,
// <https://en.wikipedia.org/wiki/ISO/IEC_7813>.
func EncZ(_ Hasher, codec Codec, dec []byte) ([]byte, error) { return codec.Charset().Encode(dec) }

// EncB intends to encode binary data <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func EncB(_ Hasher, _ Codec, dec []byte) ([]byte, error) {
	enc := make([]byte, base64.StdEncoding.DecodedLen(len(dec)))
	n, err := base64.StdEncoding.Decode(enc, dec)
	return enc[:n], err
}

// DecodeMTI intends to decode the ISO 8583 MTI (Message Type Indicator).
func DecodeMTI(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecBitmap intends to decode the ISO 8583 bitmap.
func DecBitmap(_ Hasher, _ Codec, enc []byte) ([]byte, error) { return enc, nil }

// DecA intends to decode Alpha, including blanks <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecA(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecN intends to decode numeric values only <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecN(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecXN intends to decode numeric (amount) values,
// where the first byte is either 'C' to indicate a positive or Credit value,
// or 'D' to indicate a negative or Debit value,
// followed by the numeric value (using n digits)
// <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecXN(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecAN intends to decode alphanumeric <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecAN(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecNS intends to decode numeric and special characters only <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecNS(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecANS intends to decode alphabetic, numeric and special characters <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecANS(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecZ intends to decode tracks 2 and 3 code set as defined in ISO/IEC 7813 and ISO/IEC 4909 respectively
// <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>,
// <https://en.wikipedia.org/wiki/ISO/IEC_7813>.
func DecZ(_ Hasher, codec Codec, enc []byte) ([]byte, error) { return codec.Charset().Decode(enc) }

// DecB intends to decode binary data <https://en.wikipedia.org/wiki/ISO_8583#Data_elements>.
func DecB(_ Hasher, _ Codec, enc []byte) ([]byte, error) {
	dec := make([]byte, base64.StdEncoding.EncodedLen(len(enc)))
	base64.StdEncoding.Encode(dec, enc)
	return dec, nil
}

// DecNullify just skips the field.
func DecNullify(_ Hasher, _ Codec, _ []byte) ([]byte, error) { return []byte{}, nil }

// DecPANTruncate intends to decode and truncate (mask) PAN,
// 1234567890123456 -> 1234XXXXXXXX3456.
func DecPANTruncate(_ Hasher, codec Codec, enc []byte) ([]byte, error) {
	dec, err := codec.Charset().Decode(enc)
	if err != nil {
		return nil, err
	}
	for i := 4; i < len(dec)-4; i++ {
		dec[i] = 'X'
	}
	return dec, nil
}

// DecFirstPANTruncate intends to decode and truncate (mask) PAN located at the beginning of the byte slice,
// 1234567890123456D99122011969100000377 -> 1234XXXXXXXX3456D99122011969100000377
func DecFirstPANTruncate(_ Hasher, codec Codec, enc []byte) ([]byte, error) {
	dec, err := codec.Charset().Decode(enc)
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile("[^0-9]")
	if err != nil {
		return nil, err
	}
	loc := re.FindIndex(dec)
	l := len(dec)
	if loc != nil {
		l = loc[0]
	}
	for i := 4; i < l-4; i++ {
		dec[i] = 'X'
	}
	return dec, nil
}

// DecHash256 intends to decode and hash (obfuscate)
// like Central Bank of the Russian Federation obfuscates PAN,
// using the SHA 256 sum function.
func DecHash256(h Hasher, codec Codec, enc []byte) ([]byte, error) {
	dec, err := codec.Charset().Decode(enc)
	if err != nil {
		return nil, err
	}
	s256, err := h.Sum256(dec)
	if err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(s256)), nil
}

// Codec encodes/decodes ISO 8583 MTI/bitmaps/fields.
type Codec interface {
	Encode(Hasher, []byte) ([]byte, error)
	Decode(Hasher, []byte) ([]byte, error)
	Read(io.Reader) ([]byte, error)
	Write(io.Writer, []byte) error
	Len() int
	Charset() Charset
}

// FIX is an codec for the fixed length fields/MTI/bitmaps of the ISO 8583.
type FIX struct {
	Length int
	ChrSet Charset
	Enc    EncodeFunc
	Dec    DecodeFunc
}

func (c FIX) Len() int                                    { return c.Length }
func (c FIX) Charset() Charset                            { return c.ChrSet }
func (c FIX) Encode(h Hasher, dec []byte) ([]byte, error) { return c.Enc(h, c, dec) }
func (c FIX) Decode(h Hasher, enc []byte) ([]byte, error) { return c.Dec(h, c, enc) }

const (
	fieldMinLen = 1
	fieldMaxLen = 999
)

// readFix reads a field with fixed length.
func (c FIX) Read(r io.Reader) ([]byte, error) {
	if c.Length < fieldMinLen {
		return nil, fmt.Errorf("fixed field length too small to read: %d < %d", c.Length, fieldMinLen)
	}
	if c.Length > fieldMaxLen {
		return nil, fmt.Errorf("fixed field length too big to read: %d > %d", c.Length, fieldMaxLen)
	}
	data := make([]byte, c.Length)
	_, err := r.Read(data)
	return data, err
}

// Write writes a field with fixed length.
func (c FIX) Write(w io.Writer, data []byte) error {
	if c.Length < fieldMinLen {
		return fmt.Errorf("fixed field length too small to write: %d < %d", c.Length, fieldMinLen)
	}
	if c.Length > fieldMaxLen {
		return fmt.Errorf("fixed field length too big to write: %d > %d", c.Length, fieldMaxLen)
	}
	_, err := w.Write(data)
	return err
}

// LVAR is an codec for the variable length fields of the ISO 8583,
// the length coded in 1 byte.
type LVAR struct {
	MaxLen int
	ChrSet Charset
	Enc    EncodeFunc
	Dec    DecodeFunc
}

func (c LVAR) Len() int                                    { return c.MaxLen }
func (c LVAR) Charset() Charset                            { return c.ChrSet }
func (c LVAR) Encode(h Hasher, dec []byte) ([]byte, error) { return c.Enc(h, c, dec) }
func (c LVAR) Decode(h Hasher, enc []byte) ([]byte, error) { return c.Dec(h, c, enc) }

// Read reads a field with variable length (VAR).
func (c LVAR) Read(r io.Reader) ([]byte, error) {
	return varRead(r, 1, c.MaxLen, c.ChrSet.Dec)
}

// Write writes a field with variable length (VAR).
func (c LVAR) Write(w io.Writer, data []byte) error {
	return varWrite(w, data, 1, c.MaxLen, c.ChrSet.Enc)
}

// LLVAR is an codec for the variable length fields of the ISO 8583,
// the length coded in 2 bytes.
type LLVAR struct {
	MaxLen int
	ChrSet Charset
	Enc    EncodeFunc
	Dec    DecodeFunc
}

func (c LLVAR) Len() int                                    { return c.MaxLen }
func (c LLVAR) Charset() Charset                            { return c.ChrSet }
func (c LLVAR) Encode(h Hasher, dec []byte) ([]byte, error) { return c.Enc(h, c, dec) }
func (c LLVAR) Decode(h Hasher, enc []byte) ([]byte, error) { return c.Dec(h, c, enc) }

// Read reads a field with variable length (VAR).
func (c LLVAR) Read(r io.Reader) ([]byte, error) {
	return varRead(r, 2, c.MaxLen, c.ChrSet.Dec)
}

// Write writes a field with variable length (VAR).
func (c LLVAR) Write(w io.Writer, data []byte) error {
	return varWrite(w, data, 2, c.MaxLen, c.ChrSet.Enc)
}

// LLLVAR is an codec for the variable length fields of the ISO 8583,
// the length coded in 3 bytes.
type LLLVAR struct {
	MaxLen int
	ChrSet Charset
	Enc    EncodeFunc
	Dec    DecodeFunc
}

func (c LLLVAR) Len() int                                    { return c.MaxLen }
func (c LLLVAR) Charset() Charset                            { return c.ChrSet }
func (c LLLVAR) Encode(h Hasher, dec []byte) ([]byte, error) { return c.Enc(h, c, dec) }
func (c LLLVAR) Decode(h Hasher, enc []byte) ([]byte, error) { return c.Dec(h, c, enc) }

// Read reads a field with variable length (VAR).
func (c LLLVAR) Read(r io.Reader) ([]byte, error) {
	return varRead(r, 3, c.MaxLen, c.ChrSet.Dec)
}

// Write writes a field with variable length (VAR).
func (c LLLVAR) Write(w io.Writer, data []byte) error {
	return varWrite(w, data, 3, c.MaxLen, c.ChrSet.Enc)
}

// varRead reads a field with variable length (VAR).
func varRead(r io.Reader, lenOfLen, maxLen int, dec DecodeCharsetFunc) ([]byte, error) {
	raw := make([]byte, lenOfLen)
	_, err := r.Read(raw)
	if err != nil {
		return nil, err
	}
	if dec != nil {
		raw, err = dec(raw)
		if err != nil {
			return nil, err
		}
	}
	var length int
	length, err = strconv.Atoi(string(raw))
	if err != nil {
		return nil, err
	}
	if length < fieldMinLen {
		return nil, fmt.Errorf("VAR field length too small to read: %d < %d", length, fieldMinLen)
	}
	if length > maxLen {
		return nil, fmt.Errorf("VAR field length too big to read: %d > %d", length, maxLen)
	}
	data := make([]byte, length)
	_, err = r.Read(data)
	return data, err
}

// varWrite writes a field with variable length (VAR).
func varWrite(w io.Writer, data []byte, lenOfLen, maxLen int, enc EncodeCharsetFunc) error {
	if len(data) < fieldMinLen {
		return fmt.Errorf("field value length is too small to write: %d < %d", len(data), fieldMinLen)
	}
	if len(data) > maxLen || maxLen > fieldMaxLen {
		return fmt.Errorf("field value length is too big to write: %d > %d", len(data), maxLen)
	}
	rawLengthOfLength := []byte(fmt.Sprintf("%0"+strconv.Itoa(lenOfLen)+"d", len(data)))
	var err error
	if enc != nil {
		rawLengthOfLength, err = enc(rawLengthOfLength)
		if err != nil {
			return err
		}
	}
	_, err = w.Write(rawLengthOfLength)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(data))
	return err
}

type Hasher interface {
	Sum256(data []byte) ([]byte, error)
}

func NewHashify() hashify { return hashify{} }

type hashify struct{}

var pool = sync.Pool{New: func() interface{} { return sha256.New() }}

func (h hashify) Sum256(data []byte) ([]byte, error) {
	h256 := pool.Get().(hash.Hash)
	h256.Reset()
	defer pool.Put(h256)

	_, err := h256.Write(data)
	if err != nil {
		return nil, err
	}
	return h256.Sum(nil), nil
}

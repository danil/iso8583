package codec8583

import (
	"golang.org/x/text/encoding/charmap"
)

var (
	// ASCII encodes/decodes ASCII charset of the ISO 8583 MTI/bitmaps/fields.
	ASCII = NOPCharset
	// EBCDIC encodes/decodes EBCDIC charset of the ISO 8583 MTI/bitmaps/fields.
	EBCDIC = Charset{EncodeEbcdic, DecodeEbcdic}
	// NOPCharset is a "no operation" charset encoder/decoder.
	NOPCharset = Charset{}
)

type (
	// EncodeCharsetFunc encodes charset of the ISO 8583 MTI/bitmaps/fields.
	EncodeCharsetFunc func([]byte) ([]byte, error)
	// DecodeCharsetFunc decodes charset of the ISO 8583 MTI/bitmaps/fields.
	DecodeCharsetFunc func([]byte) ([]byte, error)
)

// Charset encodes/decodes charset of the ISO 8583 MTI/bitmaps/fields.
type Charset struct {
	Enc EncodeCharsetFunc
	Dec DecodeCharsetFunc
}

func (c Charset) Encode(data []byte) ([]byte, error) {
	if c.Enc == nil {
		return data, nil
	}
	return c.Enc(data)
}

func (c Charset) Decode(data []byte) ([]byte, error) {
	if c.Dec == nil {
		return data, nil
	}
	return c.Dec(data)
}

// EncodeEbcdic intends to encode all bytes in EBCDIC character encoding except the null character.
func EncodeEbcdic(src []byte) ([]byte, error) {
	enc := charmap.CodePage037.NewEncoder()
	dst := make([]byte, len(src)) // *8)
	nDst, _, err := enc.Transform(dst, src, true)
	if err != nil {
		return nil, err
	}
	return dst[:nDst], nil
}

// DecodeEbcdic intends to decode all bytes in EBCDIC character encoding except the null character.
func DecodeEbcdic(src []byte) ([]byte, error) {
	dec := charmap.CodePage037.NewDecoder()
	dst := make([]byte, len(src)) // *8)
	nDst, _, err := dec.Transform(dst, src, true)
	if err != nil {
		return nil, err
	}
	return dst[:nDst], nil
}

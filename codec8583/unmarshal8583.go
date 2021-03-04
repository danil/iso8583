package codec8583

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/danil/iso8583/bitmap64"
)

type Unmarshaler interface {
	Unmarshal(raw []byte, v interface{}) error
}

type unmarshal struct {
	format  Format
	source  io.Reader
	target  reflect.Value
	hashify Hasher
}

func NewUnmarshaler(f Format) unmarshal { return unmarshal{format: f, hashify: NewHashify()} }

func (umrs unmarshal) Unmarshal(raw []byte, v interface{}) error {
	err := umrs.unmarshal(raw, v)
	if err != nil {
		return fmt.Errorf("ISO 8583 unmarshal: %w, message: %#v", err, raw)
	}
	return nil
}

func (umrs unmarshal) unmarshal(raw []byte, v interface{}) error {
	mtiLen := umrs.format[-1].Len()
	bmpLen := umrs.format[0].Len()
	if len(raw) < mtiLen+bmpLen {
		return fmt.Errorf("message too small to read: %d < %d", len(raw), mtiLen+bmpLen)
	}
	umrs.source = bytes.NewReader(raw)
	umrs.target = reflect.ValueOf(v).Elem()
	err := umrs.decodeMTI()
	if err != nil {
		return fmt.Errorf("decode MTI: %w", err)
	}
	fldNumbers, err := umrs.decodeBitmaps()
	if err != nil {
		return fmt.Errorf("decode bitmaps: %w", err)
	}
	return umrs.decodeFields(fldNumbers)
}

func (umrs *unmarshal) decodeMTI() error {
	mtiCodec := umrs.format[-1]
	raw, err := mtiCodec.Read(umrs.source)
	if err != nil {
		return err
	}
	mti, err := mtiCodec.Decode(umrs.hashify, raw)
	if err != nil {
		return err
	}
	var mtiVal reflect.Value
	for i := 0; i < umrs.target.NumField(); i++ {
		fld := umrs.target.Type().Field(i)
		tag := strings.Split(fld.Tag.Get(Tag8583), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if tag == "MTI" {
			mtiVal = umrs.target.Field(i)
			break
		}
	}
	if !mtiVal.IsValid() {
		return errors.New("struct must have MTI field")
	}
	mtiVal.SetString(string(mti))
	return nil
}

// decodeBitmaps decode all bitmaps and return ISO 8583 field numbers.
func (umrs *unmarshal) decodeBitmaps() ([]int, error) {
	bmpCodec := umrs.format[0]
	raw, err := bmpCodec.Read(umrs.source)
	if err != nil {
		return nil, err
	}
	priBmp, err := bmpCodec.Decode(umrs.hashify, raw)
	if err != nil {
		return nil, err
	}
	var fldNumbers []int
	bmp := bitmap64.New(priBmp)
	for i := 0; i < 64; i++ {
		if bmp.Get(i + 1) {
			fldNumbers = append(fldNumbers, i+1)
		}
	}
	if !bmp.Get(1) {
		return fldNumbers, nil
	}
	bmpCodec = umrs.format[1]
	raw, err = bmpCodec.Read(umrs.source)
	if err != nil {
		return nil, err
	}
	secBmp, err := bmpCodec.Decode(umrs.hashify, raw)
	if err != nil {
		return nil, err
	}
	bmp = bitmap64.New(secBmp)
	for i := 0; i < 64; i++ {
		if bmp.Get(i + 1) {
			fldNumbers = append(fldNumbers, i+64+1)
		}
	}
	return fldNumbers, nil
}

func (umrs *unmarshal) decodeFields(fldNumbers []int) error {
	fldToTgtFld, err := decodingFieldToTargetFields(umrs.target)
	if err != nil {
		return err
	}
	for _, fld := range fldNumbers {
		if fld == 1 { // ISO 8583 first field (secondary bitmap) already read.
			continue
		}
		err = umrs.decodeField(fld, fldToTgtFld)
		if err != nil {
			return fmt.Errorf("decode field: %d, %w, all fields: %v", fld, err, fldNumbers)
		}
	}
	return nil
}

func (umrs *unmarshal) decodeField(fld int, fldToTgtFld map[int]int) error {
	fldCodec := umrs.format[fld]
	raw, err := fldCodec.Read(umrs.source)
	if err != nil {
		return err
	}
	val, err := fldCodec.Decode(umrs.hashify, raw)
	if err != nil {
		return err
	}
	fieldIndex, ok := fldToTgtFld[fld]
	if !ok {
		return fmt.Errorf("struct does not have a field with tag: %#[1]v, the message in field: %#[1]v, has a value: %#[2]v", fld, string(val))
	}
	valueField := umrs.target.Field(fieldIndex)
	valueField.SetString(string(val))
	return nil
}

// decodingFieldToTargetFields returns mapping of the ISO 8583 feilds to the target fields .
func decodingFieldToTargetFields(target reflect.Value) (map[int]int, error) {
	fldToTgtFld := make(map[int]int, target.NumField())
	for i := 0; i < target.NumField(); i++ {
		fld := target.Type().Field(i)
		tag := fld.Tag.Get(Tag8583)
		if tag == "" || tag == "MTI" {
			continue
		}
		idx, err := strconv.Atoi(tag)
		if err != nil {
			return nil, err
		}
		fldToTgtFld[idx] = i
	}
	return fldToTgtFld, nil
}

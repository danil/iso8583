package codec8583

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/danil/iso8583/bitmap64"
)

type Marshaler interface {
	Marshal(interface{}) ([]byte, error)
}

type marshal struct {
	format  Format
	source  reflect.Value
	target  io.Writer
	hashify Hasher
}

func NewMarshaler(f Format) marshal { return marshal{format: f, hashify: NewHashify()} }

func (mrs marshal) Marshal(v interface{}) ([]byte, error) {
	enc, err := mrs.marshal(v)
	if err != nil {
		return enc, fmt.Errorf("ISO 8583 marshal: %w, struct: %#v", err, v)
	}
	return enc, nil
}

func (mrs marshal) marshal(v interface{}) ([]byte, error) {
	var target bytes.Buffer
	mrs.source = reflect.ValueOf(v).Elem()
	mrs.target = &target
	var mtiVal reflect.Value
	for i := 0; i < mrs.source.NumField(); i++ {
		fld := mrs.source.Type().Field(i)
		tag := strings.Split(fld.Tag.Get(Tag8583), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if tag == "MTI" {
			mtiVal = mrs.source.Field(i)
			break
		}
	}
	if !mtiVal.IsValid() {
		return nil, errors.New("must have MTI field")
	}
	fldNums, fldVals := mrs.fieldNumbersAndValues()
	priBmp, secBmp := mrs.bitmaps(fldNums)
	if secBmp != [8]byte{} {
		fldVals[1] = secBmp[:]
		fldNums = append([]int{1}, fldNums...)
	}
	err := mrs.encodeMTI([]byte(mtiVal.String()))
	if err != nil {
		return nil, fmt.Errorf("encode MTI: %w", err)
	}
	err = mrs.encodePrimaryBitmap(priBmp)
	if err != nil {
		return nil, fmt.Errorf("encode primary bitmap: %w", err)
	}
	err = mrs.encodeFields(fldNums, fldVals)
	if err != nil {
		return nil, err
	}
	return target.Bytes(), nil
}

func (mrs *marshal) bitmaps(fldNums []int) ([8]byte, [8]byte) {
	priBmp := bitmap64.Bitmap([8]byte{})
	i := 0
	for i < len(fldNums) {
		n := fldNums[i]
		if n > 64 {
			break
		}
		priBmp.Set(n)
		if i == len(fldNums)-1 {
			return priBmp, [8]byte{}
		}
		i++
	}
	priBmp.Set(1)
	secBmp := bitmap64.Bitmap([8]byte{})
	for ; i < len(fldNums); i++ {
		secBmp.Set(fldNums[i] - 64)
	}
	return priBmp, secBmp
}

func (mrs *marshal) fieldNumbersAndValues() ([]int, map[int][]byte) {
	fldVals := map[int][]byte{}
	for i := 0; i < mrs.source.NumField(); i++ {
		fldVal := mrs.source.Field(i)
		fld := mrs.source.Type().Field(i)
		tag := fld.Tag.Get(Tag8583)
		if tag == "" {
			continue
		}
		idx, err := strconv.ParseInt(tag, 10, 64)
		if err != nil {
			continue
		}
		f := fldVal.Interface()
		val := reflect.ValueOf(f).String()
		if val == "" {
			continue
		}
		fldVals[int(idx)] = []byte(val)
	}
	fldNums := make([]int, 0, len(fldVals))
	for fld := range fldVals {
		fldNums = append(fldNums, fld)
	}
	sort.Ints(fldNums)
	return fldNums, fldVals
}

func (mrs *marshal) encodeMTI(mti []byte) error {
	mtiCodec := mrs.format[-1]
	raw, err := mtiCodec.Encode(mrs.hashify, mti)
	if err != nil {
		return err
	}
	return mtiCodec.Write(mrs.target, raw)
}

func (mrs *marshal) encodePrimaryBitmap(priBmp [8]byte) error {
	bmpCodec := mrs.format[1]
	raw, err := bmpCodec.Encode(mrs.hashify, priBmp[:])
	if err != nil {
		return err
	}
	return bmpCodec.Write(mrs.target, raw)
}

func (mrs *marshal) encodeFields(fldNums []int, fldVals map[int][]byte) error {
	for _, fld := range fldNums {
		val := fldVals[fld]
		err := mrs.encodeField(fld, val)
		if err != nil {
			return fmt.Errorf("encode field: %w, field: %d, value: %#v, all fields %v", err, fld, val, fldNums)
		}
	}
	return nil
}

func (mrs *marshal) encodeField(fld int, val []byte) error {
	fldCodec := mrs.format[fld]
	raw, err := fldCodec.Encode(mrs.hashify, val)
	if err != nil {
		return err
	}
	return fldCodec.Write(mrs.target, raw)
}

package codec8583_test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"github.com/danil/iso8583/codec8583"
)

var MastercardTruncateFormatTestCases = []struct {
	input     iso8583Msg
	truncate  iso8583Msg
	line      int
	benchmark bool
}{
	{
		input:     iso8583Msg{MessageTypeIndicator: "0100", PrimaryAccountNumber: "1234567890123456", ProcessingCode: "000000", AmountTransaction: "000000020126", AmountCardholderBilling: "000000000310", TransmissionDateTime: "1008053137", ConversionRateCardholderBilling: "70154259", SystemTraceAuditNumber: "196842", LocalTransactionTime: "083137", LocalTransactionDate: "1008", ExpirationDate: "2312", SettlementDate: "1008", CurrencyConversionDate: "1008", MerchantTypeOrMerchantCategoryCode: "5411", PointOfServiceEntryMode: "071", ApplicationPANSequenceNumber: "001", AcquiringInstitutionIdentificationCode: "005037", ForwardingInstitutionIdentificationCode: "200154", Track2Data: "1234567890123456D99122011969100000377", RetrievalReferenceNumber: "519939727023", CardAcceptorTerminalIdentification: "10747007", CardAcceptorIdentificationCode: "990000026622   ", CardAcceptorNameLocation: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction: "643", CurrencyCodeCardholderBilling: "840", ICCData: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1: "0000010000300643142103", ReservedPrivate3: "TNWGAP405"},
		truncate:  iso8583Msg{MessageTypeIndicator: "0100", PrimaryAccountNumber: "1234XXXXXXXX3456", ProcessingCode: "000000", AmountTransaction: "000000020126", AmountCardholderBilling: "000000000310", TransmissionDateTime: "1008053137", ConversionRateCardholderBilling: "70154259", SystemTraceAuditNumber: "196842", LocalTransactionTime: "083137", LocalTransactionDate: "1008", ExpirationDate: "2312", SettlementDate: "1008", CurrencyConversionDate: "1008", MerchantTypeOrMerchantCategoryCode: "5411", PointOfServiceEntryMode: "071", ApplicationPANSequenceNumber: "001", AcquiringInstitutionIdentificationCode: "005037", ForwardingInstitutionIdentificationCode: "200154", Track2Data: "1234XXXXXXXX3456D99122011969100000377", RetrievalReferenceNumber: "519939727023", CardAcceptorTerminalIdentification: "10747007", CardAcceptorIdentificationCode: "990000026622   ", CardAcceptorNameLocation: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction: "643", CurrencyCodeCardholderBilling: "840", ICCData: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1: "0000010000300643142103", ReservedPrivate3: "TNWGAP405"},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		input:    iso8583Msg{MessageTypeIndicator: "0100", PrimaryAccountNumber: "1234567890123456"},
		truncate: iso8583Msg{MessageTypeIndicator: "0100", PrimaryAccountNumber: "1234XXXXXXXX3456"},
		line:     func() int { _, _, l, _ := runtime.Caller(1); return l }(),
	},
}

func TestMastercardTruncateFormat(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range MastercardTruncateFormatTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			enc, err := codec8583.MastercardMarshaler.Marshal(&tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			dec := iso8583Msg{}
			err = codec8583.MastercardUnmarshalerTruncator.Unmarshal(enc, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != tc.truncate {
				t.Errorf("[truncate message] expected: %#v, received: %#v - %s", tc.truncate, dec, linkToExample)
			}
			enc, err = codec8583.MastercardMarshaler.Marshal(&dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			dec = iso8583Msg{}
			err = codec8583.MastercardUnmarshalerTruncator.Unmarshal(enc, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != tc.truncate {
				t.Errorf("[truncate message] expected: %#v, received: %#v - %s", tc.truncate, dec, linkToExample)
			}
		})
	}
}

func BenchmarkMastercardUnmarshalerTruncatorUnmarshal(b *testing.B) {
	dec := iso8583Msg{}
	err := codec8583.MastercardUnmarshalerTruncator.Unmarshal(fakeMessages[0], &dec)
	if err != nil {
		fmt.Println(err)
	}
}

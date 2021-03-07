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
		input:     iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "1234567890123456", ProcessingCode3: "000000", AmountTransaction4: "000000020126", AmountCardholderBilling6: "000000000310", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "196842", LocalTransactionTime12: "083137", LocalTransactionDate13: "1008", ExpirationDate14: "2312", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5411", POSPointOfServiceEntryMode22: "071", ApplicationPANSequenceNumber23: "001", AcquiringInstitutionIdentificationCode32: "005037", ForwardingInstitutionIdentificationCode33: "200154", Track2Data35: "1234567890123456D99122011969100000377", RetrievalReferenceNumber37: "519939727023", CardAcceptorTerminalIdentification41: "10747007", CardAcceptorIdentificationCode42: "990000026622   ", CardAcceptorNameLocation43: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate48: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1_61: "0000010000300643142103", ReservedPrivate3_63: "TNWGAP405"},
		truncate:  iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "1234XXXXXXXX3456", ProcessingCode3: "000000", AmountTransaction4: "000000020126", AmountCardholderBilling6: "000000000310", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "196842", LocalTransactionTime12: "083137", LocalTransactionDate13: "1008", ExpirationDate14: "2312", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5411", POSPointOfServiceEntryMode22: "071", ApplicationPANSequenceNumber23: "001", AcquiringInstitutionIdentificationCode32: "005037", ForwardingInstitutionIdentificationCode33: "200154", Track2Data35: "1234XXXXXXXX3456D99122011969100000377", RetrievalReferenceNumber37: "519939727023", CardAcceptorTerminalIdentification41: "10747007", CardAcceptorIdentificationCode42: "990000026622   ", CardAcceptorNameLocation43: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate48: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1_61: "0000010000300643142103", ReservedPrivate3_63: "TNWGAP405"},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		input:    iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "1234567890123456"},
		truncate: iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "1234XXXXXXXX3456"},
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
	err := codec8583.MastercardUnmarshalerTruncator.Unmarshal(iso8583Bytes[0], &dec)
	if err != nil {
		fmt.Println(err)
	}
}

package codec8583_test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"github.com/danil/iso8583/codec8583"
)

var MastercardHash256FormatTestCases = []struct {
	decoded iso8583Msg
	hashed  iso8583Msg
	line    int
}{
	{
		decoded: iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "1234567890123456", ProcessingCode3: "000000", AmountTransaction4: "000000020126", AmountCardholderBilling6: "000000000310", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "196842", LocalTransactionTime12: "083137", LocalTransactionDate13: "1008", ExpirationDate14: "2312", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5411", POSPointOfServiceEntryMode22: "071", ApplicationPANSequenceNumber23: "001", AcquiringInstitutionIdentificationCode32: "005037", ForwardingInstitutionIdentificationCode33: "200154", Track2Data35: "1234567890123456D99122011969100000377", RetrievalReferenceNumber37: "519939727023", CardAcceptorTerminalIdentification41: "10747007", CardAcceptorIdentificationCode42: "990000026622   ", CardAcceptorNameLocation43: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate48: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1_61: "0000010000300643142103", ReservedPrivate3_63: "TNWGAP405"},
		hashed:  iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "7a51d064a1a216a692f753fcdab276e4ff201a01d8b66f56d50d4d719fd0dc87", ProcessingCode3: "000000", AmountTransaction4: "000000020126", AmountCardholderBilling6: "000000000310", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "196842", LocalTransactionTime12: "083137", LocalTransactionDate13: "1008", ExpirationDate14: "2312", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5411", POSPointOfServiceEntryMode22: "071", ApplicationPANSequenceNumber23: "001", AcquiringInstitutionIdentificationCode32: "005037", ForwardingInstitutionIdentificationCode33: "200154", Track2Data35: "9a903b0d368e7fb03ee7bc44d84a621526eb03947f51f400ad858598ec9a4497", RetrievalReferenceNumber37: "519939727023", CardAcceptorTerminalIdentification41: "10747007", CardAcceptorIdentificationCode42: "990000026622   ", CardAcceptorNameLocation43: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate48: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1_61: "0000010000300643142103", ReservedPrivate3_63: "TNWGAP405"},
		line:    func() int { _, _, l, _ := runtime.Caller(1); return l }(),
	},
	{
		decoded: iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "1234567890123456"},
		hashed:  iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "7a51d064a1a216a692f753fcdab276e4ff201a01d8b66f56d50d4d719fd0dc87"},
		line:    func() int { _, _, l, _ := runtime.Caller(1); return l }(),
	},
}

func TestMastercardHash256Format(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range MastercardHash256FormatTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			enc, err := codec8583.MastercardMarshaler.Marshal(&tc.decoded)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			dec := iso8583Msg{}
			err = codec8583.MastercardUnmarshalerHasher256.Unmarshal(enc, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != tc.hashed {
				t.Errorf("[hashed message] expected: %#v, received: %#v - %s", tc.hashed, dec, linkToExample)
			}
			enc, err = codec8583.MastercardHashed256Marshaler.Marshal(&dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			dec = iso8583Msg{}
			err = codec8583.MastercardHashed256Unmarshaler.Unmarshal(enc, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != tc.hashed {
				t.Errorf("[hashed message] expected: %#v, received: %#v - %s", tc.hashed, dec, linkToExample)
			}
		})
	}
}

func BenchmarkMastercardUnmarshalerHasher256Unmarshal(b *testing.B) {
	dec := iso8583Msg{}
	err := codec8583.MastercardUnmarshalerHasher256.Unmarshal(iso8583Bytes[0], &dec)
	if err != nil {
		fmt.Println(err)
	}
}

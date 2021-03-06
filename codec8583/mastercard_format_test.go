package codec8583_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"testing"

	"github.com/danil/iso8583/codec8583"
)

var MastercardFormatTestCases = []struct {
	encoded   []byte
	decoded   iso8583Msg
	line      int
	benchmark bool
}{
	{
		encoded:   iso8583Bytes[0],
		decoded:   iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "5321XXXXXXXX6334", ProcessingCode3: "000000", AmountTransaction4: "000000020126", AmountCardholderBilling6: "000000000310", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "196842", LocalTransactionTime12: "083137", LocalTransactionDate13: "1008", ExpirationDate14: "2312", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5411", POSPointOfServiceEntryMode22: "071", ApplicationPANSequenceNumber23: "001", AcquiringInstitutionIdentificationCode32: "005037", ForwardingInstitutionIdentificationCode33: "200154", Track2Data35: "5321XXXXXXXX6334D23122011969100000377", RetrievalReferenceNumber37: "519939727023", CardAcceptorTerminalIdentification41: "10747007", CardAcceptorIdentificationCode42: "990000026622   ", CardAcceptorNameLocation43: "PYATEROCHKA 7904       Podolsk       RUS", AdditionalDataPrivate48: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAAgEmnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJggCXxqRiyDckZ8nAYCfMwPgCMifNAMfAwCfNgIBYJ83BM8lmh8=", ReservedPrivate1_61: "0000010000300643142103", ReservedPrivate3_63: "TNWGAP405"},
		line:      func() int { _, _, l, _ := runtime.Caller(1); return l }(),
		benchmark: true,
	},
	{
		encoded: []byte{0xf0, 0xf1, 0xf0, 0xf0, 0x76, 0x7f, 0x46, 0x1, 0xa8, 0xe1, 0xa2, 0xa, 0xf1, 0xf6, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf9, 0xf4, 0xf1, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf2, 0xf3, 0xf9, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf9, 0xf1, 0xf1, 0xf0, 0xf0, 0xf8, 0xf0, 0xf5, 0xf3, 0xf1, 0xf3, 0xf7, 0xf7, 0xf0, 0xf1, 0xf5, 0xf4, 0xf2, 0xf5, 0xf9, 0xf1, 0xf9, 0xf6, 0xf8, 0xf5, 0xf4, 0xf0, 0xf8, 0xf3, 0xf1, 0xf3, 0xf7, 0xf1, 0xf0, 0xf0, 0xf8, 0xf2, 0xf3, 0xf0, 0xf9, 0xf1, 0xf0, 0xf0, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf5, 0xf4, 0xf1, 0xf1, 0xf0, 0xf7, 0xf1, 0xf0, 0xf0, 0xf1, 0xf0, 0xf6, 0xf0, 0xf0, 0xf5, 0xf0, 0xf3, 0xf7, 0xf0, 0xf6, 0xf2, 0xf0, 0xf0, 0xf1, 0xf5, 0xf4, 0xf3, 0xf7, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf9, 0xf4, 0xf1, 0xf2, 0xc4, 0xf2, 0xf3, 0xf0, 0xf9, 0xf2, 0xf0, 0xf1, 0xf1, 0xf5, 0xf6, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf4, 0xf9, 0xf5, 0xf5, 0xf1, 0xf9, 0xf9, 0xf3, 0xf9, 0xf7, 0xf2, 0xf7, 0xf0, 0xf4, 0xf7, 0xf1, 0xf1, 0xf3, 0xf3, 0xf7, 0xf2, 0xf8, 0xf1, 0xf9, 0xf9, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf1, 0xf6, 0xf5, 0xf0, 0x40, 0x40, 0x40, 0xd4, 0xc1, 0xc7, 0xd5, 0xc9, 0xe3, 0x40, 0xd4, 0xd4, 0x40, 0xc1, 0xd9, 0xc1, 0xd2, 0xc3, 0xc8, 0xc9, 0xd5, 0xd2, 0xc1, 0x40, 0x40, 0x40, 0xd6, 0x94, 0xa2, 0x92, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd9, 0xe4, 0xe2, 0xf0, 0xf1, 0xf6, 0xd9, 0xf2, 0xf3, 0xf0, 0xf2, 0xf0, 0xf0, 0xf6, 0xf1, 0xf0, 0xf5, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf6, 0xf4, 0xf3, 0xf8, 0xf4, 0xf0, 0xf1, 0xf1, 0xf6, 0x5f, 0x2a, 0x2, 0x6, 0x43, 0x82, 0x2, 0x19, 0x80, 0x84, 0x7, 0xa0, 0x0, 0x0, 0x0, 0x4, 0x10, 0x10, 0x95, 0x5, 0x0, 0x0, 0x0, 0x80, 0x0, 0x9a, 0x3, 0x19, 0x10, 0x8, 0x9c, 0x1, 0x0, 0x9f, 0x2, 0x6, 0x0, 0x0, 0x0, 0x1, 0x23, 0x90, 0x9f, 0x3, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9f, 0x10, 0x12, 0x1, 0x10, 0xa0, 0x40, 0x3, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x9f, 0x1a, 0x2, 0x6, 0x43, 0x9f, 0x26, 0x8, 0xee, 0xf1, 0x98, 0xfd, 0x3d, 0x57, 0xd7, 0xd, 0x9f, 0x27, 0x1, 0x80, 0x9f, 0x33, 0x3, 0xe0, 0x8, 0xc8, 0x9f, 0x34, 0x3, 0x1f, 0x3, 0x0, 0x9f, 0x36, 0x2, 0x2, 0x3, 0x9f, 0x37, 0x4, 0x64, 0x75, 0x32, 0xb8, 0xf0, 0xf2, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf0, 0xf0, 0xf6, 0xf4, 0xf3, 0xf6, 0xf4, 0xf4, 0xf0, 0xf4, 0xf7, 0xf0, 0xf0, 0xf9, 0xe3, 0xd5, 0xe6, 0xd8, 0xc2, 0xd7, 0xc5, 0xf0, 0xf0},
		decoded: iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "5321XXXXXXXX9412", ProcessingCode3: "000000", AmountTransaction4: "000000012390", AmountCardholderBilling6: "000000000191", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "196854", LocalTransactionTime12: "083137", LocalTransactionDate13: "1008", ExpirationDate14: "2309", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5411", POSPointOfServiceEntryMode22: "071", ApplicationPANSequenceNumber23: "001", AcquiringInstitutionIdentificationCode32: "005037", ForwardingInstitutionIdentificationCode33: "200154", Track2Data35: "5321XXXXXXXX9412D23092011562000000495", RetrievalReferenceNumber37: "519939727047", CardAcceptorTerminalIdentification41: "11337281", CardAcceptorIdentificationCode42: "990000031650   ", CardAcceptorNameLocation43: "MAGNIT MM ARAKCHINKA   Omsk          RUS", AdditionalDataPrivate48: "2fLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "XyoCBkOCAhmAhAegAAAABBAQlQUAAACAAJoDGRAInAEAnwIGAAAAASOQnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJgju8Zj9PVfXDZ8nAYCfMwPgCMifNAMfAwCfNgICA583BGR1Mrg=", ReservedPrivate1_61: "0000010000300643644047", ReservedPrivate3_63: "TNWQBPE00"},
		line:    func() int { _, _, l, _ := runtime.Caller(1); return l }(),
	},
	{
		encoded: []byte{0xf0, 0xf1, 0xf0, 0xf0, 0x76, 0x7f, 0x44, 0x1, 0x88, 0xe1, 0xa1, 0xa, 0xf1, 0xf6, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf9, 0xf6, 0xf5, 0xf3, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf4, 0xf9, 0xf1, 0xf0, 0xf0, 0xf8, 0xf0, 0xf5, 0xf3, 0xf1, 0xf3, 0xf7, 0xf7, 0xf0, 0xf1, 0xf5, 0xf4, 0xf2, 0xf5, 0xf9, 0xf3, 0xf3, 0xf1, 0xf2, 0xf0, 0xf5, 0xf0, 0xf9, 0xf2, 0xf1, 0xf1, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf2, 0xf4, 0xf0, 0xf1, 0xf1, 0xf0, 0xf0, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf4, 0xf1, 0xf3, 0xf1, 0xf0, 0xf7, 0xf2, 0xf0, 0xf6, 0xf0, 0xf0, 0xf7, 0xf3, 0xf2, 0xf6, 0xf0, 0xf6, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf9, 0xf0, 0xf0, 0xf8, 0xf9, 0xf3, 0xf2, 0xf6, 0xf3, 0xf1, 0xf2, 0xf0, 0xf5, 0xf1, 0xf4, 0xf9, 0xf6, 0xf8, 0xf3, 0x40, 0x40, 0xf1, 0xf4, 0xf9, 0xf6, 0xf8, 0xf3, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xc9, 0xe2, 0xc5, 0xe3, 0x40, 0xe3, 0xd9, 0xc1, 0xd5, 0xe2, 0xd7, 0xd6, 0xd9, 0xe3, 0x40, 0xf8, 0xf2, 0xf4, 0x40, 0x40, 0x40, 0x40, 0x40, 0xc5, 0xd2, 0xc1, 0xe3, 0xc5, 0xd9, 0xc9, 0xd5, 0xc2, 0xe4, 0xd9, 0xc7, 0x40, 0x40, 0xd9, 0xe4, 0xe2, 0xf1, 0xf6, 0xf6, 0xe7, 0xf2, 0xf6, 0xf0, 0xf3, 0xf1, 0xf0, 0xf3, 0xf3, 0xf0, 0xf4, 0xf4, 0xc6, 0xf9, 0xa7, 0xc1, 0xd9, 0x98, 0xd1, 0xd3, 0x81, 0xc5, 0xa4, 0xc6, 0xa2, 0xc1, 0xe6, 0x91, 0xa3, 0xf7, 0xa6, 0xd1, 0xe6, 0xf6, 0xe4, 0xf1, 0x4e, 0x92, 0x93, 0xa3, 0xd7, 0xa5, 0xf5, 0xe4, 0x99, 0xe6, 0x81, 0x91, 0xa5, 0xc2, 0xf1, 0x86, 0xf0, 0xc1, 0xa2, 0x7e, 0xf3, 0xf3, 0xf4, 0xf8, 0xf0, 0xf1, 0xf0, 0xf1, 0xc3, 0xf0, 0xf2, 0xf1, 0xf6, 0xf5, 0xf4, 0xf4, 0xf3, 0xf3, 0xf1, 0xf9, 0xf5, 0xf0, 0xf7, 0xf1, 0xf9, 0xf0, 0xf3, 0xf3, 0xf6, 0xf0, 0xf3, 0xf0, 0xf4, 0xf2, 0xf2, 0xf0, 0xf9, 0xf0, 0xf6, 0xf1, 0xf1, 0xf5, 0xf0, 0xf1, 0xf1, 0xf0, 0xf0, 0xf3, 0xf0, 0xf2, 0xf7, 0xf3, 0xf3, 0xf4, 0xf1, 0xf1, 0xf0, 0xf0, 0xf1, 0xf4, 0xf3, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xe6, 0xf6, 0xf1, 0xf0, 0xf5, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf6, 0xf4, 0xf0, 0xf4, 0xf0, 0xf3, 0xf0, 0xf0, 0xf7, 0xf1, 0xf0, 0xf8, 0xf5, 0xf0, 0xc3, 0x40, 0xf5, 0xf1, 0xe5, 0x40, 0xf7, 0xf2, 0xf1, 0xf0, 0x10, 0x76, 0xa0, 0x98, 0x3d, 0x46, 0xb3, 0xcc, 0x0, 0x10, 0xf6, 0xf4, 0xf3, 0xf8, 0xf4, 0xf0, 0xf0, 0xf3, 0xf7, 0xf0, 0xf1, 0xf3, 0xf3, 0xf0, 0xf1, 0xf2, 0xf9, 0xf5, 0xf0, 0xf0, 0xf1, 0xc3, 0xf8, 0xd1, 0xd4, 0xc7, 0xf0, 0xf3, 0xd8, 0xd6, 0xd6, 0xe5, 0xe9, 0xf2, 0xf1, 0xf4, 0xc5, 0xc7, 0xc6, 0xe8, 0xe3, 0xe3, 0xf2, 0xc7, 0xc2, 0xc3, 0xf0, 0xf2, 0xf2, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf0, 0xf0, 0xf6, 0xf4, 0xf3, 0xf6, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf9, 0xe3, 0xd5, 0xe6, 0xd8, 0xc4, 0xc4, 0xd9, 0xf0, 0xf4},
		decoded: iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "5321XXXXXXXX9653", ProcessingCode3: "000000", AmountTransaction4: "000000003200", AmountCardholderBilling6: "000000000049", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "331205", LocalTransactionTime12: "092118", LocalTransactionDate13: "1008", ExpirationDate14: "2401", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "4131", POSPointOfServiceEntryMode22: "072", AcquiringInstitutionIdentificationCode32: "007326", ForwardingInstitutionIdentificationCode33: "200009", RetrievalReferenceNumber37: "008932631205", CardAcceptorTerminalIdentification41: "149683  ", CardAcceptorIdentificationCode42: "149683         ", CardAcceptorNameLocation43: "ISET TRANSPORT 824     EKATERINBURG  RUS", AdditionalDataPrivate48: "5/L28PPx8PPz8PT0xvmnwdmY0dOBxaTGosHmkaP3ptHm9uTxTpKTo9el9eSZ5oGRpcLxhvDBon7z8/T48PHw8cPw8vH29fT08/Px+fXw9/H58PPz9vDz8PTy8vD58Pbx8fXw8fHw8PPw8vfz8/Tx8fDw8fTz8PDw8PDm9vHw9fDw8PDx9vTw9PDz8PD38fD49fDDQPXx5UD38vHwEHagmD1Gs8wAEA==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ReservedISO56: "013301295001C8JMG03QOOVZ214EGFYTT2GBC", ReservedPrivate1_61: "1000000000300643620000", ReservedPrivate3_63: "TNWQDDR04"},
		line:    func() int { _, _, l, _ := runtime.Caller(1); return l }(),
	},
	{
		encoded: []byte{0xf0, 0xf1, 0xf0, 0xf0, 0x76, 0x7f, 0x46, 0x1, 0xa8, 0xe1, 0xa2, 0xa, 0xf1, 0xf6, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf3, 0xf4, 0xf9, 0xf4, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf7, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf6, 0xf2, 0xf1, 0xf0, 0xf0, 0xf8, 0xf0, 0xf5, 0xf3, 0xf1, 0xf3, 0xf7, 0xf7, 0xf0, 0xf1, 0xf5, 0xf4, 0xf2, 0xf5, 0xf9, 0xf7, 0xf4, 0xf1, 0xf8, 0xf4, 0xf4, 0xf0, 0xf8, 0xf3, 0xf1, 0xf3, 0xf3, 0xf1, 0xf0, 0xf0, 0xf8, 0xf2, 0xf3, 0xf0, 0xf1, 0xf1, 0xf0, 0xf0, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf5, 0xf8, 0xf1, 0xf4, 0xf0, 0xf7, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf6, 0xf0, 0xf1, 0xf4, 0xf1, 0xf1, 0xf3, 0xf0, 0xf6, 0xf2, 0xf0, 0xf0, 0xf3, 0xf9, 0xf2, 0xf3, 0xf7, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf3, 0xf4, 0xf9, 0xf4, 0xc4, 0xf2, 0xf3, 0xf0, 0xf1, 0xf2, 0xf0, 0xf1, 0xf1, 0xf1, 0xf6, 0xf3, 0xf8, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf9, 0xf9, 0xf1, 0xf1, 0xf0, 0xf0, 0xf8, 0xf6, 0xf0, 0xf5, 0xf2, 0xf5, 0xf2, 0xf5, 0xf4, 0xf9, 0xf1, 0xf0, 0xf0, 0xf2, 0xf6, 0xf0, 0xf5, 0xf9, 0xf1, 0xf0, 0xf0, 0xf1, 0xf3, 0xf4, 0xf3, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xc3, 0xc1, 0xc6, 0xc5, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xe2, 0xc1, 0xd5, 0xd2, 0xe3, 0x60, 0xd7, 0xc5, 0xe3, 0xc5, 0xd9, 0xc2, 0xe4, 0x40, 0xd9, 0xe4, 0xe2, 0xf0, 0xf1, 0xf6, 0xc6, 0xf2, 0xf3, 0xf0, 0xf2, 0xf0, 0xf0, 0xf6, 0xf1, 0xf0, 0xf5, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf6, 0xf4, 0xf3, 0xf8, 0xf4, 0xf0, 0xf1, 0xf2, 0xf6, 0x82, 0x2, 0x19, 0x80, 0x95, 0x5, 0x0, 0x0, 0x0, 0x80, 0x1, 0x9a, 0x3, 0x19, 0x10, 0x8, 0x9c, 0x1, 0x0, 0x5f, 0x2a, 0x2, 0x6, 0x43, 0x9f, 0x2, 0x6, 0x0, 0x0, 0x0, 0x1, 0x70, 0x0, 0x9f, 0x3, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9f, 0x10, 0x12, 0x1, 0x10, 0xa0, 0x40, 0x3, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x9f, 0x1a, 0x2, 0x6, 0x43, 0x9f, 0x26, 0x8, 0xf6, 0x36, 0x25, 0xf7, 0xa6, 0x4c, 0x4c, 0xe8, 0x9f, 0x27, 0x1, 0x80, 0x9f, 0x33, 0x3, 0xe0, 0x8, 0xc8, 0x9f, 0x36, 0x2, 0x9, 0x2c, 0x9f, 0x37, 0x4, 0x1c, 0x44, 0xa1, 0xfd, 0x9f, 0x41, 0x3, 0x2, 0x52, 0x54, 0x9f, 0x34, 0x3, 0x1f, 0x3, 0x2, 0x9f, 0x35, 0x1, 0x22, 0x84, 0x7, 0xa0, 0x0, 0x0, 0x0, 0x4, 0x10, 0x10, 0xf0, 0xf2, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf0, 0xf0, 0xf6, 0xf4, 0xf3, 0xf1, 0xf9, 0xf9, 0xf0, 0xf4, 0xf4, 0xf0, 0xf0, 0xf9, 0xe3, 0xd5, 0xe6, 0xd8, 0xf2, 0xe9, 0xe8, 0xf0, 0xf4},
		decoded: iso8583Msg{MTIMessageTypeIndicator: "0100", PANPrimaryAccountNumber2: "5321XXXXXXXX3494", ProcessingCode3: "000000", AmountTransaction4: "000000017000", AmountCardholderBilling6: "000000000262", TransmissionDateTime7: "1008053137", ConversionRateCardholderBilling10: "70154259", SystemTraceAuditNumber11: "741844", LocalTransactionTime12: "083133", LocalTransactionDate13: "1008", ExpirationDate14: "2301", SettlementDate15: "1008", CurrencyConversionDate16: "1008", MerchantTypeOrMerchantCategoryCode18: "5814", POSPointOfServiceEntryMode22: "072", ApplicationPANSequenceNumber23: "000", AcquiringInstitutionIdentificationCode32: "014113", ForwardingInstitutionIdentificationCode33: "200392", Track2Data35: "5321XXXXXXXX3494D23012011163800000991", RetrievalReferenceNumber37: "100860525254", CardAcceptorTerminalIdentification41: "91002605", CardAcceptorIdentificationCode42: "91001343       ", CardAcceptorNameLocation43: "CAFE                   SANKT-PETERBU RUS", AdditionalDataPrivate48: "xvLz8PLw8Pbx8PXw8PDw8Q==", CurrencyCodeTransaction49: "643", CurrencyCodeCardholderBilling51: "840", ICCData55: "ggIZgJUFAAAAgAGaAxkQCJwBAF8qAgZDnwIGAAAAAXAAnwMGAAAAAAAAnxASARCgQAMiAAAAAAAAAAAAAAD/nxoCBkOfJgj2NiX3pkxM6J8nAYCfMwPgCMifNgIJLJ83BBxEof2fQQMCUlSfNAMfAwKfNQEihAegAAAABBAQ", ReservedPrivate1_61: "0000010000300643199044", ReservedPrivate3_63: "TNWQ2ZY04"},
		line:    func() int { _, _, l, _ := runtime.Caller(1); return l }(),
	},
}

func TestMastercardFormat(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range MastercardFormatTestCases {
		tc := tc
		t.Run(strconv.Itoa(tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			dec := iso8583Msg{}
			err := codec8583.MastercardUnmarshaler.Unmarshal(tc.encoded, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != tc.decoded {
				t.Errorf("[decoded message] expected: %#v, received: %#v - %s", tc.decoded, dec, linkToExample)
			}
			enc, err := codec8583.MastercardMarshaler.Marshal(&dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if !bytes.Equal(enc, tc.encoded) {
				t.Errorf("[encoded message] expected %#v, received: %#v - %s", tc.encoded, enc, linkToExample)
			}
		})
	}
}

func TestMastercardFormatByMastercardJSON(t *testing.T) {
	jsonFile := "mastercard.json"
	p, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	var messages []iso8583Msg
	if err := json.Unmarshal(p, &messages); err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	for i, msg := range messages {
		msg := msg
		linkToExample := fmt.Sprintf("%s:%d", jsonFile, i+1)
		t.Run(linkToExample, func(t *testing.T) {
			t.Parallel()
			enc, err := codec8583.MastercardMarshaler.Marshal(&msg)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			dec := iso8583Msg{}
			err = codec8583.MastercardUnmarshaler.Unmarshal(enc, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != msg {
				t.Errorf("[decoded message] expected: %#v, received: %#v - %s", msg, dec, linkToExample)
			}
		})
	}
}

func TestMastercardFormatByNSPKJSON(t *testing.T) {
	jsonFile := "nspk.json"
	p, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	var messages []iso8583Msg
	if err := json.Unmarshal(p, &messages); err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	for i, msg := range messages {
		msg := msg
		linkToExample := fmt.Sprintf("%s:%d", jsonFile, i+1)
		t.Run(linkToExample, func(t *testing.T) {
			t.Parallel()
			enc, err := codec8583.MastercardMarshaler.Marshal(&msg)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			dec := iso8583Msg{}
			err = codec8583.MastercardUnmarshaler.Unmarshal(enc, &dec)
			if err != nil {
				t.Fatalf("unexpected error: %#v - %s", err, linkToExample)
			}
			if dec != msg {
				t.Errorf("[decoded message] expected: %#v, received: %#v - %s", msg, dec, linkToExample)
			}
		})
	}
}

func BenchmarkMastercardMarshalerMarshal(b *testing.B) {
	for _, tc := range MastercardFormatTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := codec8583.MastercardMarshaler.Marshal(&tc.decoded)
				if err != nil {
					fmt.Println(err)
				}
			}
		})
	}
}

func BenchmarkMastercardUnmarshalerUnmarshal(b *testing.B) {
	for _, tc := range MastercardFormatTestCases {
		if !tc.benchmark {
			continue
		}
		b.Run(strconv.Itoa(tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dec := iso8583Msg{}
				err := codec8583.MastercardUnmarshaler.Unmarshal(tc.encoded, &dec)
				if err != nil {
					fmt.Println(err)
				}
			}
		})
	}
}

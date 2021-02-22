package codec8583

import (
	"github.com/danil/iso8583/tsysbitmap64r"
)

var (
	TSYSMarshaler   = NewMarshaler(TSYS)
	TSYSUnmarshaler = NewUnmarshaler(TSYS)
)

// TSYS is a format of the ISO 8583 message which maps the codecs to MIT/bitmaps/each individual field.
var TSYS = Format{
	-1:  MTIAsciiCodec,                                        // Message type indicator (MTI)
	0:   BitmapCodec,                                          // Primary Bitmap
	1:   BitmapCodec,                                          // Second Bitmap
	2:   LLVAR{19, ASCII, EncN, DecN},                         // Primary account number (PAN)
	3:   FIX{6, ASCII, EncN, DecN},                            // Processing code
	4:   FIX{12, ASCII, EncN, DecN},                           // Amount, transaction
	5:   FIX{12, ASCII, EncN, DecN},                           // Amount, settlement
	6:   FIX{12, ASCII, EncN, DecN},                           // Amount, cardholder billing
	7:   FIX{10, ASCII, EncN, DecN},                           // Transmission date & time
	8:   FIX{8, ASCII, EncN, DecN},                            // Amount, cardholder billing fee
	9:   FIX{8, ASCII, EncN, DecN},                            // Conversion rate, settlement
	10:  FIX{8, ASCII, EncN, DecN},                            // Conversion rate, cardholder billing
	11:  FIX{6, ASCII, EncN, DecN},                            // System trace audit number (STAN)
	12:  FIX{6, ASCII, EncN, DecN},                            // Local transaction time (hhmmss)
	13:  FIX{4, ASCII, EncN, DecN},                            // Local transaction date (MMDD)
	14:  FIX{4, ASCII, EncN, DecN},                            // Expiration date
	15:  FIX{4, ASCII, EncN, DecN},                            // Settlement date
	16:  FIX{4, ASCII, EncN, DecN},                            // Currency conversion date
	17:  FIX{4, ASCII, EncN, DecN},                            // Capture date
	18:  FIX{4, ASCII, EncN, DecN},                            // Merchant type, or merchant category code
	19:  FIX{3, ASCII, EncN, DecN},                            // Acquiring institution (country code)
	20:  FIX{3, ASCII, EncN, DecN},                            // PAN extended (country code)
	21:  FIX{3, ASCII, EncN, DecN},                            // Forwarding institution (country code)
	22:  FIX{4, ASCII, EncN, DecN},                            // Point of service entry mode
	23:  FIX{3, ASCII, EncN, DecN},                            // Application PAN sequence number
	24:  FIX{3, ASCII, EncN, DecN},                            // Function code (ISO 8583:1993), or network international identifier (NII)
	25:  FIX{2, ASCII, EncN, DecN},                            // Point of service condition code
	26:  FIX{2, ASCII, EncN, DecN},                            // Point of service capture code
	27:  FIX{1, ASCII, EncN, DecN},                            // Authorizing identification response length
	28:  FIX{9, ASCII, EncN, DecN},                            // Amount, transaction fee
	29:  FIX{8, ASCII, EncN, DecN},                            // Amount, settlement fee
	30:  FIX{8, ASCII, EncN, DecN},                            // Amount, transaction processing fee
	31:  FIX{15, ASCII, EncN, DecN},                           // Amount, settlement processing fee
	32:  LLVAR{11, ASCII, EncN, DecN},                         // Acquiring institution identification code
	33:  LLVAR{11, ASCII, EncN, DecN},                         // Forwarding institution identification code
	34:  LLVAR{28, ASCII, EncN, DecN},                         // Primary account number, extended
	35:  LLVAR{37, ASCII, EncANS, DecANS},                     // Track 2 data (PAN with expiry date)
	36:  LLLVAR{104, ASCII, EncN, DecN},                       // Track 3 data
	37:  FIX{12, ASCII, EncAN, DecAN},                         // Retrieval reference number
	38:  FIX{6, ASCII, EncAN, DecAN},                          // Authorization identification response
	39:  FIX{2, ASCII, EncAN, DecAN},                          // Response code
	40:  FIX{3, ASCII, EncAN, DecAN},                          // Service restriction code
	41:  FIX{8, ASCII, EncANS, DecANS},                        // Card acceptor terminal identification
	42:  FIX{15, ASCII, EncANS, DecANS},                       // Card acceptor identification code
	43:  FIX{40, ASCII, EncANS, DecANS},                       // Card acceptor name/location (1–23 street address, –36 city, –38 state, 39–40 country)
	44:  LLVAR{25, ASCII, EncAN, DecAN},                       // Additional response data
	45:  LLVAR{76, ASCII, EncAN, DecAN},                       // Track 1 data
	46:  LLLVAR{999, ASCII, EncAN, DecAN},                     // Additional data (ISO)
	47:  LLLVAR{999, ASCII, EncAN, DecAN},                     // Additional data (national)
	48:  LLLVAR{999, ASCII, EncAN, DecAN},                     // Additional data (private) (DE 48: Subelement 10 — Encrypted PIN Block Key, Subelement 80 — PIN Service Code)
	49:  FIX{3, ASCII, EncA, DecA},                            // Currency code, transaction
	50:  FIX{3, ASCII, EncAN, DecAN},                          // Currency code, settlement
	51:  FIX{3, ASCII, EncA, DecA},                            // Currency code, cardholder billing
	52:  FIX{8, ASCII, EncB, DecB},                            // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	53:  FIX{18, ASCII, EncN, DecN},                           // Security related control information (DE 53: Security-Related Control Information, Subfield 1 — PIN Security Type Code, Subfield 2 — PIN Encryption Type Code, Subfield 3 — PIN Block Format Code, Subfield 4—PIN Key Index Number
	54:  LLLVAR{120, ASCII, EncAN, DecAN},                     // Additional amounts
	55:  LLLVAR{255, ASCII, EncB, DecB},                       // ICC data – EMV having multiple tags
	56:  LLVAR{35, ASCII, EncANS, DecANS},                     // Reserved (ISO)
	57:  LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved (national)
	58:  LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved (national)
	59:  LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved (national)
	60:  LLVAR{8, ASCII, EncAN, DecAN},                        // Reserved (national) (e.g. settlement request: batch number, advice transactions: original transaction amount, batch upload: original MTI plus original RRN plus original STAN, etc.)
	61:  LLVAR{36, ASCII, EncANS, DecANS},                     // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	62:  LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	63:  LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved (private)
	64:  FIX{16, ASCII, EncB, DecB},                           // Message authentication code (MAC)
	65:  FIX{16, ASCII, EncB, DecB},                           // Extended bitmap indicator
	66:  FIX{1, ASCII, EncN, DecN},                            // Settlement code
	67:  FIX{2, ASCII, EncN, DecN},                            // Extended payment code
	68:  FIX{3, ASCII, EncN, DecN},                            // Receiving institution country code
	69:  FIX{3, ASCII, EncN, DecN},                            // Settlement institution country code
	70:  FIX{3, ASCII, EncN, DecN},                            // Network management information code
	71:  FIX{4, ASCII, EncN, DecN},                            // Message number
	72:  LLLVAR{999, ASCII, EncANS, DecANS},                   // Last message's number
	73:  FIX{6, ASCII, EncN, DecN},                            // Action date (YYMMDD)
	74:  FIX{10, ASCII, EncN, DecN},                           // Number of credits
	75:  FIX{10, ASCII, EncN, DecN},                           // Credits, reversal number
	76:  FIX{10, ASCII, EncN, DecN},                           // Number of debits
	77:  FIX{10, ASCII, EncN, DecN},                           // Debits, reversal number
	78:  FIX{10, ASCII, EncN, DecN},                           // Transfer number
	79:  FIX{10, ASCII, EncN, DecN},                           // Transfer, reversal number
	80:  FIX{10, ASCII, EncN, DecN},                           // Number of inquiries
	81:  FIX{10, ASCII, EncN, DecN},                           // Number of authorizations
	82:  FIX{12, ASCII, EncN, DecN},                           // Credits, processing fee amount
	83:  FIX{12, ASCII, EncN, DecN},                           // Credits, transaction fee amount
	84:  FIX{12, ASCII, EncN, DecN},                           // Debits, processing fee amount
	85:  FIX{12, ASCII, EncN, DecN},                           // Debits, transaction fee amount
	86:  FIX{15, ASCII, EncN, DecN},                           // Total amount of credits
	87:  FIX{15, ASCII, EncN, DecN},                           // Credits, reversal amount
	88:  FIX{15, ASCII, EncN, DecN},                           // Total amount of debits
	89:  FIX{15, ASCII, EncN, DecN},                           // Debits, reversal amount
	90:  FIX{42, ASCII, EncN, DecN},                           // Original data elements
	91:  FIX{1, ASCII, EncAN, DecAN},                          // File update code
	92:  FIX{2, ASCII, EncN, DecN},                            // File security code
	93:  FIX{4, ASCII, EncN, DecN},                            // Response indicator
	94:  FIX{2, ASCII, EncAN, DecAN},                          // Service indicator
	95:  FIX{42, ASCII, EncAN, DecAN},                         // Replacement amounts
	96:  FIX{8, ASCII, EncAN, DecAN},                          // Message security code
	97:  FIX{16, ASCII, EncN, DecN},                           // Net settlement amount
	98:  FIX{25, ASCII, EncANS, DecANS},                       // Payee
	99:  FIX{1, ASCII, EncN, DecN},                            // Settlement institution identification code
	100: LLVAR{11, ASCII, EncAN, DecAN},                       // Receiving institution identification code
	101: FIX{17, ASCII, EncANS, DecANS},                       // File name
	102: LLVAR{28, ASCII, EncANS, DecANS},                     // Account identification 1
	103: LLVAR{28, ASCII, EncANS, DecANS},                     // Account identification 2
	104: LLVAR{28, ASCII, EncANS, DecANS},                     // Transaction description
	105: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	106: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	107: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	108: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	109: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	110: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	111: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for ISO use
	112: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for national use
	113: LLVAR{11, ASCII, EncN, DecN},                         // Reserved for national use
	114: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for national use
	115: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for national use
	116: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for national use
	117: LLLVAR{120, ASCII, EncANS, DecANS},                   // Reserved for national use
	118: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for national use
	119: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for national use
	120: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for private use
	121: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for private use
	122: LLVAR{78, ASCII, EncANS, DecANS},                     // Reserved for private use
	123: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for private use
	124: LLLVAR{255, ASCII, EncANS, DecANS},                   // Reserved for private use
	125: LLVAR{28, ASCII, EncANS, DecANS},                     // Reserved for private use
	126: LLVAR{99, ASCII, EncTSYSBitmap64R, DecTSYSBitmap64R}, // Reserved for private use
	127: LLLVAR{999, ASCII, EncANS, DecANS},                   // Reserved for private use
	128: FIX{16, ASCII, EncB, DecB},                           // Message authentication code
}

// EncTSYSBitmap64R intends to encode TSYS specific bitmap
// the size of which is 8 bytes (64 bit)
// plus one additional optional byte for the "R" flag.
func EncTSYSBitmap64R(_ Hasher, _ Codec, data []byte) ([]byte, error) {
	b, err := tsysbitmap64r.NewString(string(data))
	if err != nil {
		return nil, err
	}
	return b.MarshalISO8583()
}

// DecTSYSBitmap64R intends to decode TSYS specific bitmap
// the size of which is 8 bytes (64 bit)
// plus one additional optional byte for the "R" flag.
func DecTSYSBitmap64R(_ Hasher, _ Codec, data []byte) ([]byte, error) {
	return []byte(tsysbitmap64r.New(data).String()), nil
}

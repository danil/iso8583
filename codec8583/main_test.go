package codec8583_test

var iso8583Bytes = [][]byte{
	[]byte{0xf0, 0xf1, 0xf0, 0xf0, 0x76, 0x7f, 0x46, 0x1, 0xa8, 0xe1, 0xa2, 0xa, 0xf1, 0xf6, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf6, 0xf3, 0xf3, 0xf4, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf0, 0xf1, 0xf2, 0xf6, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf1, 0xf0, 0xf1, 0xf0, 0xf0, 0xf8, 0xf0, 0xf5, 0xf3, 0xf1, 0xf3, 0xf7, 0xf7, 0xf0, 0xf1, 0xf5, 0xf4, 0xf2, 0xf5, 0xf9, 0xf1, 0xf9, 0xf6, 0xf8, 0xf4, 0xf2, 0xf0, 0xf8, 0xf3, 0xf1, 0xf3, 0xf7, 0xf1, 0xf0, 0xf0, 0xf8, 0xf2, 0xf3, 0xf1, 0xf2, 0xf1, 0xf0, 0xf0, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf5, 0xf4, 0xf1, 0xf1, 0xf0, 0xf7, 0xf1, 0xf0, 0xf0, 0xf1, 0xf0, 0xf6, 0xf0, 0xf0, 0xf5, 0xf0, 0xf3, 0xf7, 0xf0, 0xf6, 0xf2, 0xf0, 0xf0, 0xf1, 0xf5, 0xf4, 0xf3, 0xf7, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf6, 0xf3, 0xf3, 0xf4, 0xc4, 0xf2, 0xf3, 0xf1, 0xf2, 0xf2, 0xf0, 0xf1, 0xf1, 0xf9, 0xf6, 0xf9, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf7, 0xf7, 0xf5, 0xf1, 0xf9, 0xf9, 0xf3, 0xf9, 0xf7, 0xf2, 0xf7, 0xf0, 0xf2, 0xf3, 0xf1, 0xf0, 0xf7, 0xf4, 0xf7, 0xf0, 0xf0, 0xf7, 0xf9, 0xf9, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf6, 0xf6, 0xf2, 0xf2, 0x40, 0x40, 0x40, 0xd7, 0xe8, 0xc1, 0xe3, 0xc5, 0xd9, 0xd6, 0xc3, 0xc8, 0xd2, 0xc1, 0x40, 0xf7, 0xf9, 0xf0, 0xf4, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd7, 0x96, 0x84, 0x96, 0x93, 0xa2, 0x92, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd9, 0xe4, 0xe2, 0xf0, 0xf1, 0xf6, 0xd9, 0xf2, 0xf3, 0xf0, 0xf2, 0xf0, 0xf0, 0xf6, 0xf1, 0xf0, 0xf5, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf6, 0xf4, 0xf3, 0xf8, 0xf4, 0xf0, 0xf1, 0xf1, 0xf6, 0x5f, 0x2a, 0x2, 0x6, 0x43, 0x82, 0x2, 0x19, 0x80, 0x84, 0x7, 0xa0, 0x0, 0x0, 0x0, 0x4, 0x10, 0x10, 0x95, 0x5, 0x0, 0x0, 0x0, 0x80, 0x0, 0x9a, 0x3, 0x19, 0x10, 0x8, 0x9c, 0x1, 0x0, 0x9f, 0x2, 0x6, 0x0, 0x0, 0x0, 0x2, 0x1, 0x26, 0x9f, 0x3, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9f, 0x10, 0x12, 0x1, 0x10, 0xa0, 0x40, 0x3, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x9f, 0x1a, 0x2, 0x6, 0x43, 0x9f, 0x26, 0x8, 0x2, 0x5f, 0x1a, 0x91, 0x8b, 0x20, 0xdc, 0x91, 0x9f, 0x27, 0x1, 0x80, 0x9f, 0x33, 0x3, 0xe0, 0x8, 0xc8, 0x9f, 0x34, 0x3, 0x1f, 0x3, 0x0, 0x9f, 0x36, 0x2, 0x1, 0x60, 0x9f, 0x37, 0x4, 0xcf, 0x25, 0x9a, 0x1f, 0xf0, 0xf2, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf0, 0xf0, 0xf6, 0xf4, 0xf3, 0xf1, 0xf4, 0xf2, 0xf1, 0xf0, 0xf3, 0xf0, 0xf0, 0xf9, 0xe3, 0xd5, 0xe6, 0xc7, 0xc1, 0xd7, 0xf4, 0xf0, 0xf5},
	[]byte{0x30, 0x32, 0x30, 0x30, 0xf2, 0x3a, 0xc4, 0x81, 0x28, 0xe0, 0x80, 0x10, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x31, 0x37, 0x31, 0x34, 0x31, 0x37, 0x32, 0x38, 0x30, 0x30, 0x30, 0x30, 0x36, 0x34, 0x31, 0x33, 0x35, 0x32, 0x35, 0x33, 0x30, 0x35, 0x31, 0x36, 0x30, 0x35, 0x31, 0x37, 0x30, 0x35, 0x31, 0x36, 0x36, 0x30, 0x31, 0x30, 0x39, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x36, 0x39, 0x39, 0x39, 0x39, 0x30, 0x35, 0x33, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x31, 0x33, 0x36, 0x31, 0x39, 0x39, 0x30, 0x39, 0x34, 0x34, 0x38, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x20, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x4d, 0x6f, 0x73, 0x63, 0x6f, 0x77, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x55, 0x36, 0x34, 0x33, 0x30, 0x38, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4d, 0x43, 0x54, 0x49, 0x44, 0x48, 0x49, 0x31, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x36, 0x35, 0x55, 0x44, 0x30, 0x36, 0x30, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x30, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x32, 0x30, 0x35, 0x30, 0x31, 0x30, 0x30, 0x36, 0x30, 0x31, 0x30, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x45, 0x52, 0x30, 0x30, 0x35, 0x34, 0x30, 0x30, 0x37, 0x33},
}

type iso8583Msg struct {
	MTIMessageTypeIndicator                   string `iso8583:"MTI"` // Message type indicator (MTI)
	PANPrimaryAccountNumber2                  string `iso8583:"2"`   // Primary account number (PAN)
	ProcessingCode3                           string `iso8583:"3"`   // Processing code
	AmountTransaction4                        string `iso8583:"4"`   // Amount, transaction
	AmountSettlement5                         string `iso8583:"5"`   // Amount, settlement
	AmountCardholderBilling6                  string `iso8583:"6"`   // Amount, cardholder billing
	TransmissionDateTime7                     string `iso8583:"7"`   // Transmission date & time
	AmountCardholderBillingFee8               string `iso8583:"8"`   // Amount, cardholder billing fee
	ConversionRateSettlement9                 string `iso8583:"9"`   // Conversion rate, settlement
	ConversionRateCardholderBilling10         string `iso8583:"10"`  // Conversion rate, cardholder billing
	SystemTraceAuditNumber11                  string `iso8583:"11"`  // System trace audit number (STAN)
	LocalTransactionTime12                    string `iso8583:"12"`  // Local transaction time (hhmmss)
	LocalTransactionDate13                    string `iso8583:"13"`  // Local transaction date (MMDD)
	ExpirationDate14                          string `iso8583:"14"`  // Expiration date
	SettlementDate15                          string `iso8583:"15"`  // Settlement date
	CurrencyConversionDate16                  string `iso8583:"16"`  // Currency conversion date
	CaptureDate17                             string `iso8583:"17"`  // Capture date
	MerchantTypeOrMerchantCategoryCode18      string `iso8583:"18"`  // Merchant type, or merchant category code
	AcquiringInstitutionCountryCode19         string `iso8583:"19"`  // Acquiring institution (country code)
	PANExtendedCountryCode20                  string `iso8583:"20"`  // PAN extended (country code)
	ForwardingInstitutionCountryCode21        string `iso8583:"21"`  // Forwarding institution (country code)
	POSPointOfServiceEntryMode22              string `iso8583:"22"`  // Point of service entry mode
	ApplicationPANSequenceNumber23            string `iso8583:"23"`  // Application PAN sequence number
	FunctionCode24                            string `iso8583:"24"`  // Function code (ISO 8583:1993), or network international identifier (NII)
	POSPointOfServiceConditionCode25          string `iso8583:"25"`  // Point of service condition code
	POSPointOfServiceCaptureCode26            string `iso8583:"26"`  // Point of service capture code
	AuthorizingIdentificationResponseLength27 string `iso8583:"27"`  // Authorizing identification response length
	AmountTransactionFee28                    string `iso8583:"28"`  // Amount, transaction fee
	AmountSettlementFee29                     string `iso8583:"29"`  // Amount, settlement fee
	AmountTransactionProcessingFee30          string `iso8583:"30"`  // Amount, transaction processing fee
	AmountSettlementProcessingFee31           string `iso8583:"31"`  // Amount, settlement processing fee
	AcquiringInstitutionIdentificationCode32  string `iso8583:"32"`  // Acquiring institution identification code
	ForwardingInstitutionIdentificationCode33 string `iso8583:"33"`  // Forwarding institution identification code
	PrimaryAccountNumberExtended34            string `iso8583:"34"`  // Primary account number, extended
	Track2Data35                              string `iso8583:"35"`  // Track 2 data (PAN with expiry date)
	Track3Data36                              string `iso8583:"36"`  // Track 3 data
	RetrievalReferenceNumber37                string `iso8583:"37"`  // Retrieval reference number
	AuthorizationIdentificationResponse38     string `iso8583:"38"`  // Authorization identification response
	ResponseCode39                            string `iso8583:"39"`  // Response code
	ServiceRestrictionCode40                  string `iso8583:"40"`  // Service restriction code
	CardAcceptorTerminalIdentification41      string `iso8583:"41"`  // Card acceptor terminal identification
	CardAcceptorIdentificationCode42          string `iso8583:"42"`  // Card acceptor identification code
	CardAcceptorNameLocation43                string `iso8583:"43"`  // Card acceptor name/location (1–23 street address, –36 city, –38 state, 39–40 country)
	AdditionalResponseData44                  string `iso8583:"44"`  // Additional response data
	Track1Data45                              string `iso8583:"45"`  // Track 1 data
	AdditionalDataISO46                       string `iso8583:"46"`  // Additional data (ISO)
	AdditionalDataNational47                  string `iso8583:"47"`  // Additional data (national)
	AdditionalDataPrivate48                   string `iso8583:"48"`  // Additional data (private) (DE 48: Subelement 10 — Encrypted PIN Block Key, Subelement 80 — PIN Service Code)
	CurrencyCodeTransaction49                 string `iso8583:"49"`  // Currency code, transaction
	CurrencyCodeSettlement50                  string `iso8583:"50"`  // Currency code, settlement
	CurrencyCodeCardholderBilling51           string `iso8583:"51"`  // Currency code, cardholder billing
	PersonalIdentificationNumberData52        string `iso8583:"52"`  // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	SecurityRelatedControlInformation53       string `iso8583:"53"`  // Security related control information (DE 53: Security-Related Control Information, Subfield 1 — PIN Security Type Code, Subfield 2 — PIN Encryption Type Code, Subfield 3 — PIN Block Format Code, Subfield 4—PIN Key Index Number)
	AdditionalAmounts54                       string `iso8583:"54"`  // Additional amounts
	ICCData55                                 string `iso8583:"55"`  // ICC data – EMV having multiple tags
	ReservedISO56                             string `iso8583:"56"`  // Reserved (ISO)
	ReservedNational1_57                      string `iso8583:"57"`  // Reserved (national)
	ReservedNational2_58                      string `iso8583:"58"`  // Reserved (national)
	ReservedNational3_59                      string `iso8583:"59"`  // Reserved (national)
	ReservedNational4_60                      string `iso8583:"60"`  // Reserved (national) (e.g. settlement request: batch number, advice transactions: original transaction amount, batch upload: original MTI plus original RRN plus original STAN, etc.)
	ReservedPrivate1_61                       string `iso8583:"61"`  // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	ReservedPrivate2_62                       string `iso8583:"62"`  // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	ReservedPrivate3_63                       string `iso8583:"63"`  // Reserved (private)
	MessageAuthenticationCode1_64             string `iso8583:"64"`  // Message authentication code (MAC)
	ExtendedBitmapIndicator65                 string `iso8583:"65"`  // Extended bitmap indicator
	SettlementCode66                          string `iso8583:"66"`  // Settlement code
	ExtendedPaymentCode67                     string `iso8583:"67"`  // Extended payment code
	ReceivingInstitutionCountryCode68         string `iso8583:"68"`  // Receiving institution country code
	SettlementInstitutionCountryCode69        string `iso8583:"69"`  // Settlement institution country code
	NetworkManagementInformationCode70        string `iso8583:"70"`  // Network management information code
	MessageNumber71                           string `iso8583:"71"`  // Message number
	LastMessagesNumber72                      string `iso8583:"72"`  // Last message's number
	ActionDate73                              string `iso8583:"73"`  // Action date (YYMMDD)
	NumberOfCredits74                         string `iso8583:"74"`  // Number of credits
	CreditsReversalNumber75                   string `iso8583:"75"`  // Credits, reversal number
	NumberOfDebits76                          string `iso8583:"76"`  // Number of debits
	DebitsReversalNumber77                    string `iso8583:"77"`  // Debits, reversal number
	TransferNumber78                          string `iso8583:"78"`  // Transfer number
	TransferReversalNumber79                  string `iso8583:"79"`  // Transfer, reversal number
	NumberOfInquiries80                       string `iso8583:"80"`  // Number of inquiries
	NumberOfAuthorizations81                  string `iso8583:"81"`  // Number of authorizations
	CreditsProcessingFeeAmount82              string `iso8583:"82"`  // Credits, processing fee amount
	CreditsTransactionFeeAmount83             string `iso8583:"83"`  // Credits, transaction fee amount
	DebitsProcessingFeeAmount84               string `iso8583:"84"`  // Debits, processing fee amount
	DebitsTransactionFeeAmount85              string `iso8583:"85"`  // Debits, transaction fee amount
	TotalAmountOfCredits86                    string `iso8583:"86"`  // Total amount of credits
	CreditsReversalAmount87                   string `iso8583:"87"`  // Credits, reversal amount
	TotalAmountOfDebits88                     string `iso8583:"88"`  // Total amount of debits
	DebitsReversalAmount89                    string `iso8583:"89"`  // Debits, reversal amount
	OriginalDataElements90                    string `iso8583:"90"`  // Original data elements
	FileUpdateCode91                          string `iso8583:"91"`  // File update code
	FileSecurityCode92                        string `iso8583:"92"`  // File security code
	ResponseIndicator93                       string `iso8583:"93"`  // Response indicator
	ServiceIndicator94                        string `iso8583:"94"`  // Service indicator
	ReplacementAmounts95                      string `iso8583:"95"`  // Replacement amounts
	MessageSecurityCode96                     string `iso8583:"96"`  // Message security code
	NetSettlementAmount97                     string `iso8583:"97"`  // Net settlement amount
	Payee98                                   string `iso8583:"98"`  // Payee
	SettlementInstitutionIdentificationCode99 string `iso8583:"99"`  // Settlement institution identification code
	ReceivingInstitutionIdentificationCode100 string `iso8583:"100"` // Receiving institution identification code
	FileName101                               string `iso8583:"101"` // File name
	AccountIdentification1_102                string `iso8583:"102"` // Account identification 1
	AccountIdentification2_103                string `iso8583:"103"` // Account identification 2
	TransactionDescription_104                string `iso8583:"104"` // Transaction description
	ReservedForISOUse1_105                    string `iso8583:"105"` // Reserved for ISO use
	ReservedForISOUse2_106                    string `iso8583:"106"` // Reserved for ISO use
	ReservedForISOUse3_107                    string `iso8583:"107"` // Reserved for ISO use
	ReservedForISOUse4_108                    string `iso8583:"108"` // Reserved for ISO use
	ReservedForISOUse5_109                    string `iso8583:"109"` // Reserved for ISO use
	ReservedForISOUse6_110                    string `iso8583:"110"` // Reserved for ISO use
	ReservedForISOUse7_111                    string `iso8583:"111"` // Reserved for ISO use
	ReservedForNationalUse1_112               string `iso8583:"112"` // Reserved for national use
	ReservedForNationalUse2_113               string `iso8583:"113"` // Reserved for national use
	ReservedForNationalUse3_114               string `iso8583:"114"` // Reserved for national use
	ReservedForNationalUse4_115               string `iso8583:"115"` // Reserved for national use
	ReservedForNationalUse5_116               string `iso8583:"116"` // Reserved for national use
	ReservedForNationalUse6_117               string `iso8583:"117"` // Reserved for national use
	ReservedForNationalUse7_118               string `iso8583:"118"` // Reserved for national use
	ReservedForNationalUse8_119               string `iso8583:"119"` // Reserved for national use
	ReservedForPrivateUse1_120                string `iso8583:"120"` // Reserved for private use
	ReservedForPrivateUse2_121                string `iso8583:"121"` // Reserved for private use
	ReservedForPrivateUse3_122                string `iso8583:"122"` // Reserved for private use
	ReservedForPrivateUse4_123                string `iso8583:"123"` // Reserved for private use
	ReservedForPrivateUse5_124                string `iso8583:"124"` // Reserved for private use
	ReservedForPrivateUse6_125                string `iso8583:"125"` // Reserved for private use
	ReservedForPrivateUse7_126                string `iso8583:"126"` // Reserved for private use
	ReservedForPrivateUse8_127                string `iso8583:"127"` // Reserved for private use
	MessageAuthenticationCode2_128            string `iso8583:"128"` // Message authentication code
}

type iso8583Msg2 struct {
	MTIMessageTypeIndicator              string `iso8583:"MTI"`
	PANPrimaryAccountNumber2             string `iso8583:"2"`
	ProcessingCode3                      string `iso8583:"3"`
	AmountOrig4                          string `iso8583:"4"`
	Amount6                              string `iso8583:"6"`
	TransmissionDateTime7                string `iso8583:"7"`
	BillingRate10                        string `iso8583:"10"`
	TraceNumber11                        string `iso8583:"11"`
	LocalTime12                          string `iso8583:"12"`
	LocalDate13                          string `iso8583:"13"`
	DateExpiration14                     string `iso8583:"14"`
	DateSettlement15                     string `iso8583:"15"`
	DateCapture17                        string `iso8583:"17"`
	MerchantType18                       string `iso8583:"18"`
	AcquiringInstitutionCountryCode19    string `iso8583:"19"`
	POSDataCode22                        string `iso8583:"22"`
	POSPointOfServiceConditionCode25     string `iso8583:"25"`
	TransactionFee28                     string `iso8583:"28"`
	ONLINEIssuerAuthorizationFeeAmount31 string `iso8583:"31"`
	AcquirerInstitutionID32              string `iso8583:"32"`
	TrackData35                          string `iso8583:"35"`
	RetrievalReference37                 string `iso8583:"37"`
	AuthIDCode38                         string `iso8583:"38"`
	RespCode39                           string `iso8583:"39"`
	CardAccptrTermnlID41                 string `iso8583:"41"`
	CardAccptrIDCode42                   string `iso8583:"42"`
	CardAccptrNameLoc43                  string `iso8583:"43"`
	AdditionalResponseData44             string `iso8583:"44"`
	CurrencyOrig49                       string `iso8583:"49"`
	Currency51                           string `iso8583:"51"`
	PersonalIdentificationNumberData52   string `iso8583:"52"`
	SecurityRelatedControlInformation53  string `iso8583:"53"`
	AddtnlAmounts54                      string `iso8583:"54"`
	ICCRelatedData55                     string `iso8583:"55"`
	OriginalDataSerials56                string `iso8583:"56"`
	AdditionalInformation60              string `iso8583:"60"`
	OtherAmtTrans61                      string `iso8583:"61"`
	NetworkManagementInformationCode70   string `iso8583:"70"`
	BusinessDate73                       string `iso8583:"73"`
	OrigDataElemts90                     string `iso8583:"90"`
	NumberOfAccounts93                   string `iso8583:"93"`
	QuerySequence94                      string `iso8583:"94"`
	ReplacementAmount95                  string `iso8583:"95"`
	MoreFlag99                           string `iso8583:"99"`
	MessageOriginator100                 string `iso8583:"100"`
	AccountFrom102                       string `iso8583:"102"`
	AccountTo103                         string `iso8583:"103"`
	PrivateData104                       string `iso8583:"104"`
	AdditionalInformationPart2_116       string `iso8583:"116"`
	AdditionalAmountAccountTo117         string `iso8583:"117"`
	AdditionalInformationPart1_120       string `iso8583:"120"`
	Transfercurrencies122                string `iso8583:"122"`
	CardholderUtilityAccount125          string `iso8583:"125"`
	PrivateUseFields126                  string `iso8583:"126"`
}

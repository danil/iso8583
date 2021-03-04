package codec8583_test

var iso8583Bytes = [][]byte{
	[]byte{0xf0, 0xf1, 0xf0, 0xf0, 0x76, 0x7f, 0x46, 0x1, 0xa8, 0xe1, 0xa2, 0xa, 0xf1, 0xf6, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf6, 0xf3, 0xf3, 0xf4, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf0, 0xf1, 0xf2, 0xf6, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf1, 0xf0, 0xf1, 0xf0, 0xf0, 0xf8, 0xf0, 0xf5, 0xf3, 0xf1, 0xf3, 0xf7, 0xf7, 0xf0, 0xf1, 0xf5, 0xf4, 0xf2, 0xf5, 0xf9, 0xf1, 0xf9, 0xf6, 0xf8, 0xf4, 0xf2, 0xf0, 0xf8, 0xf3, 0xf1, 0xf3, 0xf7, 0xf1, 0xf0, 0xf0, 0xf8, 0xf2, 0xf3, 0xf1, 0xf2, 0xf1, 0xf0, 0xf0, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf5, 0xf4, 0xf1, 0xf1, 0xf0, 0xf7, 0xf1, 0xf0, 0xf0, 0xf1, 0xf0, 0xf6, 0xf0, 0xf0, 0xf5, 0xf0, 0xf3, 0xf7, 0xf0, 0xf6, 0xf2, 0xf0, 0xf0, 0xf1, 0xf5, 0xf4, 0xf3, 0xf7, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf6, 0xf3, 0xf3, 0xf4, 0xc4, 0xf2, 0xf3, 0xf1, 0xf2, 0xf2, 0xf0, 0xf1, 0xf1, 0xf9, 0xf6, 0xf9, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf7, 0xf7, 0xf5, 0xf1, 0xf9, 0xf9, 0xf3, 0xf9, 0xf7, 0xf2, 0xf7, 0xf0, 0xf2, 0xf3, 0xf1, 0xf0, 0xf7, 0xf4, 0xf7, 0xf0, 0xf0, 0xf7, 0xf9, 0xf9, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf6, 0xf6, 0xf2, 0xf2, 0x40, 0x40, 0x40, 0xd7, 0xe8, 0xc1, 0xe3, 0xc5, 0xd9, 0xd6, 0xc3, 0xc8, 0xd2, 0xc1, 0x40, 0xf7, 0xf9, 0xf0, 0xf4, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd7, 0x96, 0x84, 0x96, 0x93, 0xa2, 0x92, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd9, 0xe4, 0xe2, 0xf0, 0xf1, 0xf6, 0xd9, 0xf2, 0xf3, 0xf0, 0xf2, 0xf0, 0xf0, 0xf6, 0xf1, 0xf0, 0xf5, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf6, 0xf4, 0xf3, 0xf8, 0xf4, 0xf0, 0xf1, 0xf1, 0xf6, 0x5f, 0x2a, 0x2, 0x6, 0x43, 0x82, 0x2, 0x19, 0x80, 0x84, 0x7, 0xa0, 0x0, 0x0, 0x0, 0x4, 0x10, 0x10, 0x95, 0x5, 0x0, 0x0, 0x0, 0x80, 0x0, 0x9a, 0x3, 0x19, 0x10, 0x8, 0x9c, 0x1, 0x0, 0x9f, 0x2, 0x6, 0x0, 0x0, 0x0, 0x2, 0x1, 0x26, 0x9f, 0x3, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9f, 0x10, 0x12, 0x1, 0x10, 0xa0, 0x40, 0x3, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x9f, 0x1a, 0x2, 0x6, 0x43, 0x9f, 0x26, 0x8, 0x2, 0x5f, 0x1a, 0x91, 0x8b, 0x20, 0xdc, 0x91, 0x9f, 0x27, 0x1, 0x80, 0x9f, 0x33, 0x3, 0xe0, 0x8, 0xc8, 0x9f, 0x34, 0x3, 0x1f, 0x3, 0x0, 0x9f, 0x36, 0x2, 0x1, 0x60, 0x9f, 0x37, 0x4, 0xcf, 0x25, 0x9a, 0x1f, 0xf0, 0xf2, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf0, 0xf0, 0xf6, 0xf4, 0xf3, 0xf1, 0xf4, 0xf2, 0xf1, 0xf0, 0xf3, 0xf0, 0xf0, 0xf9, 0xe3, 0xd5, 0xe6, 0xc7, 0xc1, 0xd7, 0xf4, 0xf0, 0xf5},
	[]byte{0x30, 0x32, 0x30, 0x30, 0xf2, 0x3a, 0xc4, 0x81, 0x28, 0xe0, 0x80, 0x10, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x31, 0x37, 0x31, 0x34, 0x31, 0x37, 0x32, 0x38, 0x30, 0x30, 0x30, 0x30, 0x36, 0x34, 0x31, 0x33, 0x35, 0x32, 0x35, 0x33, 0x30, 0x35, 0x31, 0x36, 0x30, 0x35, 0x31, 0x37, 0x30, 0x35, 0x31, 0x36, 0x36, 0x30, 0x31, 0x30, 0x39, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x36, 0x39, 0x39, 0x39, 0x39, 0x30, 0x35, 0x33, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x31, 0x33, 0x36, 0x31, 0x39, 0x39, 0x30, 0x39, 0x34, 0x34, 0x38, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x20, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x4d, 0x6f, 0x73, 0x63, 0x6f, 0x77, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x55, 0x36, 0x34, 0x33, 0x30, 0x38, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4d, 0x43, 0x54, 0x49, 0x44, 0x48, 0x49, 0x31, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x36, 0x35, 0x55, 0x44, 0x30, 0x36, 0x30, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x30, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x32, 0x30, 0x35, 0x30, 0x31, 0x30, 0x30, 0x36, 0x30, 0x31, 0x30, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x45, 0x52, 0x30, 0x30, 0x35, 0x34, 0x30, 0x30, 0x37, 0x33},
}

type iso8583Msg struct {
	MessageTypeIndicator                    string `iso8583:"MTI"` // Message type indicator (MTI)
	PrimaryAccountNumber                    string `iso8583:"2"`   // Primary account number (PAN)
	ProcessingCode                          string `iso8583:"3"`   // Processing code
	AmountTransaction                       string `iso8583:"4"`   // Amount, transaction
	AmountSettlement                        string `iso8583:"5"`   // Amount, settlement
	AmountCardholderBilling                 string `iso8583:"6"`   // Amount, cardholder billing
	TransmissionDateTime                    string `iso8583:"7"`   // Transmission date & time
	AmountCardholderBillingFee              string `iso8583:"8"`   // Amount, cardholder billing fee
	ConversionRateSettlement                string `iso8583:"9"`   // Conversion rate, settlement
	ConversionRateCardholderBilling         string `iso8583:"10"`  // Conversion rate, cardholder billing
	SystemTraceAuditNumber                  string `iso8583:"11"`  // System trace audit number (STAN)
	LocalTransactionTime                    string `iso8583:"12"`  // Local transaction time (hhmmss)
	LocalTransactionDate                    string `iso8583:"13"`  // Local transaction date (MMDD)
	ExpirationDate                          string `iso8583:"14"`  // Expiration date
	SettlementDate                          string `iso8583:"15"`  // Settlement date
	CurrencyConversionDate                  string `iso8583:"16"`  // Currency conversion date
	CaptureDate                             string `iso8583:"17"`  // Capture date
	MerchantTypeOrMerchantCategoryCode      string `iso8583:"18"`  // Merchant type, or merchant category code
	AcquiringInstitutionCountryCode         string `iso8583:"19"`  // Acquiring institution (country code)
	PANExtendedCountryCode                  string `iso8583:"20"`  // PAN extended (country code)
	ForwardingInstitutionCountryCode        string `iso8583:"21"`  // Forwarding institution (country code)
	PointOfServiceEntryMode                 string `iso8583:"22"`  // Point of service entry mode
	ApplicationPANSequenceNumber            string `iso8583:"23"`  // Application PAN sequence number
	FunctionCode                            string `iso8583:"24"`  // Function code (ISO 8583:1993), or network international identifier (NII)
	PointOfServiceConditionCode             string `iso8583:"25"`  // Point of service condition code
	PointOfServiceCaptureCode               string `iso8583:"26"`  // Point of service capture code
	AuthorizingIdentificationResponseLength string `iso8583:"27"`  // Authorizing identification response length
	AmountTransactionFee                    string `iso8583:"28"`  // Amount, transaction fee
	AmountSettlementFee                     string `iso8583:"29"`  // Amount, settlement fee
	AmountTransactionProcessingFee          string `iso8583:"30"`  // Amount, transaction processing fee
	AmountSettlementProcessingFee           string `iso8583:"31"`  // Amount, settlement processing fee
	AcquiringInstitutionIdentificationCode  string `iso8583:"32"`  // Acquiring institution identification code
	ForwardingInstitutionIdentificationCode string `iso8583:"33"`  // Forwarding institution identification code
	PrimaryAccountNumberExtended            string `iso8583:"34"`  // Primary account number, extended
	Track2Data                              string `iso8583:"35"`  // Track 2 data (PAN with expiry date)
	Track3Data                              string `iso8583:"36"`  // Track 3 data
	RetrievalReferenceNumber                string `iso8583:"37"`  // Retrieval reference number
	AuthorizationIdentificationResponse     string `iso8583:"38"`  // Authorization identification response
	ResponseCode                            string `iso8583:"39"`  // Response code
	ServiceRestrictionCode                  string `iso8583:"40"`  // Service restriction code
	CardAcceptorTerminalIdentification      string `iso8583:"41"`  // Card acceptor terminal identification
	CardAcceptorIdentificationCode          string `iso8583:"42"`  // Card acceptor identification code
	CardAcceptorNameLocation                string `iso8583:"43"`  // Card acceptor name/location (1–23 street address, –36 city, –38 state, 39–40 country)
	AdditionalResponseData                  string `iso8583:"44"`  // Additional response data
	Track1Data                              string `iso8583:"45"`  // Track 1 data
	AdditionalDataISO                       string `iso8583:"46"`  // Additional data (ISO)
	AdditionalDataNational                  string `iso8583:"47"`  // Additional data (national)
	AdditionalDataPrivate                   string `iso8583:"48"`  // Additional data (private) (DE 48: Subelement 10 — Encrypted PIN Block Key, Subelement 80 — PIN Service Code)
	CurrencyCodeTransaction                 string `iso8583:"49"`  // Currency code, transaction
	CurrencyCodeSettlement                  string `iso8583:"50"`  // Currency code, settlement
	CurrencyCodeCardholderBilling           string `iso8583:"51"`  // Currency code, cardholder billing
	PersonalIdentificationNumberData        string `iso8583:"52"`  // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	SecurityRelatedControlInformation       string `iso8583:"53"`  // Security related control information (DE 53: Security-Related Control Information, Subfield 1 — PIN Security Type Code, Subfield 2 — PIN Encryption Type Code, Subfield 3 — PIN Block Format Code, Subfield 4—PIN Key Index Number)
	AdditionalAmounts                       string `iso8583:"54"`  // Additional amounts
	ICCData                                 string `iso8583:"55"`  // ICC data – EMV having multiple tags
	ReservedISO                             string `iso8583:"56"`  // Reserved (ISO)
	ReservedNational1                       string `iso8583:"57"`  // Reserved (national)
	ReservedNational2                       string `iso8583:"58"`  // Reserved (national)
	ReservedNational3                       string `iso8583:"59"`  // Reserved (national)
	ReservedNational4                       string `iso8583:"60"`  // Reserved (national) (e.g. settlement request: batch number, advice transactions: original transaction amount, batch upload: original MTI plus original RRN plus original STAN, etc.)
	ReservedPrivate1                        string `iso8583:"61"`  // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	ReservedPrivate2                        string `iso8583:"62"`  // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	ReservedPrivate3                        string `iso8583:"63"`  // Reserved (private)
	MessageAuthenticationCode1              string `iso8583:"64"`  // Message authentication code (MAC)
	ExtendedBitmapIndicator                 string `iso8583:"65"`  // Extended bitmap indicator
	SettlementCode                          string `iso8583:"66"`  // Settlement code
	ExtendedPaymentCode                     string `iso8583:"67"`  // Extended payment code
	ReceivingInstitutionCountryCode         string `iso8583:"68"`  // Receiving institution country code
	SettlementInstitutionCountryCode        string `iso8583:"69"`  // Settlement institution country code
	NetworkManagementInformationCode        string `iso8583:"70"`  // Network management information code
	MessageNumber                           string `iso8583:"71"`  // Message number
	LastMessagesNumber                      string `iso8583:"72"`  // Last message's number
	ActionDate                              string `iso8583:"73"`  // Action date (YYMMDD)
	NumberOfCredits                         string `iso8583:"74"`  // Number of credits
	CreditsReversalNumber                   string `iso8583:"75"`  // Credits, reversal number
	NumberOfDebits                          string `iso8583:"76"`  // Number of debits
	DebitsReversalNumber                    string `iso8583:"77"`  // Debits, reversal number
	TransferNumber                          string `iso8583:"78"`  // Transfer number
	TransferReversalNumber                  string `iso8583:"79"`  // Transfer, reversal number
	NumberOfInquiries                       string `iso8583:"80"`  // Number of inquiries
	NumberOfAuthorizations                  string `iso8583:"81"`  // Number of authorizations
	CreditsProcessingFeeAmount              string `iso8583:"82"`  // Credits, processing fee amount
	CreditsTransactionFeeAmount             string `iso8583:"83"`  // Credits, transaction fee amount
	DebitsProcessingFeeAmount               string `iso8583:"84"`  // Debits, processing fee amount
	DebitsTransactionFeeAmount              string `iso8583:"85"`  // Debits, transaction fee amount
	TotalAmountOfCredits                    string `iso8583:"86"`  // Total amount of credits
	CreditsReversalAmount                   string `iso8583:"87"`  // Credits, reversal amount
	TotalAmountOfDebits                     string `iso8583:"88"`  // Total amount of debits
	DebitsReversalAmount                    string `iso8583:"89"`  // Debits, reversal amount
	OriginalDataElements                    string `iso8583:"90"`  // Original data elements
	FileUpdateCode                          string `iso8583:"91"`  // File update code
	FileSecurityCode                        string `iso8583:"92"`  // File security code
	ResponseIndicator                       string `iso8583:"93"`  // Response indicator
	ServiceIndicator                        string `iso8583:"94"`  // Service indicator
	ReplacementAmounts                      string `iso8583:"95"`  // Replacement amounts
	MessageSecurityCode                     string `iso8583:"96"`  // Message security code
	NetSettlementAmount                     string `iso8583:"97"`  // Net settlement amount
	Payee                                   string `iso8583:"98"`  // Payee
	SettlementInstitutionIdentificationCode string `iso8583:"99"`  // Settlement institution identification code
	ReceivingInstitutionIdentificationCode  string `iso8583:"100"` // Receiving institution identification code
	FileName                                string `iso8583:"101"` // File name
	AccountIdentification1                  string `iso8583:"102"` // Account identification 1
	AccountIdentification2                  string `iso8583:"103"` // Account identification 2
	TransactionDescription                  string `iso8583:"104"` // Transaction description
	ReservedForISOUse1                      string `iso8583:"105"` // Reserved for ISO use
	ReservedForISOUse2                      string `iso8583:"106"` // Reserved for ISO use
	ReservedForISOUse3                      string `iso8583:"107"` // Reserved for ISO use
	ReservedForISOUse4                      string `iso8583:"108"` // Reserved for ISO use
	ReservedForISOUse5                      string `iso8583:"109"` // Reserved for ISO use
	ReservedForISOUse6                      string `iso8583:"110"` // Reserved for ISO use
	ReservedForISOUse7                      string `iso8583:"111"` // Reserved for ISO use
	ReservedForNationalUse1                 string `iso8583:"112"` // Reserved for national use
	ReservedForNationalUse2                 string `iso8583:"113"` // Reserved for national use
	ReservedForNationalUse3                 string `iso8583:"114"` // Reserved for national use
	ReservedForNationalUse4                 string `iso8583:"115"` // Reserved for national use
	ReservedForNationalUse5                 string `iso8583:"116"` // Reserved for national use
	ReservedForNationalUse6                 string `iso8583:"117"` // Reserved for national use
	ReservedForNationalUse7                 string `iso8583:"118"` // Reserved for national use
	ReservedForNationalUse8                 string `iso8583:"119"` // Reserved for national use
	ReservedForPrivateUse1                  string `iso8583:"120"` // Reserved for private use
	ReservedForPrivateUse2                  string `iso8583:"121"` // Reserved for private use
	ReservedForPrivateUse3                  string `iso8583:"122"` // Reserved for private use
	ReservedForPrivateUse4                  string `iso8583:"123"` // Reserved for private use
	ReservedForPrivateUse5                  string `iso8583:"124"` // Reserved for private use
	ReservedForPrivateUse6                  string `iso8583:"125"` // Reserved for private use
	ReservedForPrivateUse7                  string `iso8583:"126"` // Reserved for private use
	ReservedForPrivateUse8                  string `iso8583:"127"` // Reserved for private use
	MessageAuthenticationCode2              string `iso8583:"128"` // Message authentication code
}

type iso8583Msg2 struct {
	MTI                                string `iso8583:"MTI"`
	PrimaryAccountNumber               string `iso8583:"2"`
	ProcessingCode                     string `iso8583:"3"`
	AmountOrig                         string `iso8583:"4"`
	Amount                             string `iso8583:"6"`
	TransmissionDateTime               string `iso8583:"7"`
	BillingRate                        string `iso8583:"10"`
	TraceNumber                        string `iso8583:"11"`
	LocalTime                          string `iso8583:"12"`
	LocalDate                          string `iso8583:"13"`
	DateExpiration                     string `iso8583:"14"`
	DateSettlement                     string `iso8583:"15"`
	DateCapture                        string `iso8583:"17"`
	MerchantType                       string `iso8583:"18"`
	AcquiringInstitutionCountryCode    string `iso8583:"19"`
	POSDataCode                        string `iso8583:"22"`
	PointOfServiceConditionCode        string `iso8583:"25"`
	TransactionFee                     string `iso8583:"28"`
	ONLINEIssuerAuthorizationFeeAmount string `iso8583:"31"`
	AcquirerInstitutionID              string `iso8583:"32"`
	TrackData                          string `iso8583:"35"`
	RetrievalReference                 string `iso8583:"37"`
	AuthIDCode                         string `iso8583:"38"`
	RespCode                           string `iso8583:"39"`
	CardAccptrTermnlID                 string `iso8583:"41"`
	CardAccptrIDCode                   string `iso8583:"42"`
	CardAccptrNameLoc                  string `iso8583:"43"`
	AdditionalResponseData             string `iso8583:"44"`
	CurrencyOrig                       string `iso8583:"49"`
	Currency                           string `iso8583:"51"`
	PersonalIdentificationNumberData   string `iso8583:"52"`
	SecurityRelatedControlInformation  string `iso8583:"53"`
	AddtnlAmounts                      string `iso8583:"54"`
	ICCRelatedData                     string `iso8583:"55"`
	OriginalDataSerials                string `iso8583:"56"`
	AdditionalInformation              string `iso8583:"60"`
	OtherAmtTrans                      string `iso8583:"61"`
	NetworkManagementInformationCode   string `iso8583:"70"`
	BusinessDate                       string `iso8583:"73"`
	OrigDataElemts                     string `iso8583:"90"`
	NumberOfAccounts                   string `iso8583:"93"`
	QuerySequence                      string `iso8583:"94"`
	ReplacementAmount                  string `iso8583:"95"`
	MoreFlag                           string `iso8583:"99"`
	MessageOriginator                  string `iso8583:"100"`
	AccountFrom                        string `iso8583:"102"`
	AccountTo                          string `iso8583:"103"`
	PrivateData                        string `iso8583:"104"`
	AdditionalInformationPart2         string `iso8583:"116"`
	AdditionalAmountAccountTo          string `iso8583:"117"`
	AdditionalInformationPart1         string `iso8583:"120"`
	Transfercurrencies                 string `iso8583:"122"`
	CardholderUtilityAccount           string `iso8583:"125"`
	PrivateUseFields                   string `iso8583:"126"`
}

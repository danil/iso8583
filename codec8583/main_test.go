package codec8583_test

var iso8583Bytes = [][]byte{
	[]byte{0xf0, 0xf1, 0xf0, 0xf0, 0x76, 0x7f, 0x46, 0x1, 0xa8, 0xe1, 0xa2, 0xa, 0xf1, 0xf6, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf6, 0xf3, 0xf3, 0xf4, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf0, 0xf1, 0xf2, 0xf6, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf1, 0xf0, 0xf1, 0xf0, 0xf0, 0xf8, 0xf0, 0xf5, 0xf3, 0xf1, 0xf3, 0xf7, 0xf7, 0xf0, 0xf1, 0xf5, 0xf4, 0xf2, 0xf5, 0xf9, 0xf1, 0xf9, 0xf6, 0xf8, 0xf4, 0xf2, 0xf0, 0xf8, 0xf3, 0xf1, 0xf3, 0xf7, 0xf1, 0xf0, 0xf0, 0xf8, 0xf2, 0xf3, 0xf1, 0xf2, 0xf1, 0xf0, 0xf0, 0xf8, 0xf1, 0xf0, 0xf0, 0xf8, 0xf5, 0xf4, 0xf1, 0xf1, 0xf0, 0xf7, 0xf1, 0xf0, 0xf0, 0xf1, 0xf0, 0xf6, 0xf0, 0xf0, 0xf5, 0xf0, 0xf3, 0xf7, 0xf0, 0xf6, 0xf2, 0xf0, 0xf0, 0xf1, 0xf5, 0xf4, 0xf3, 0xf7, 0xf5, 0xf3, 0xf2, 0xf1, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xe7, 0xf6, 0xf3, 0xf3, 0xf4, 0xc4, 0xf2, 0xf3, 0xf1, 0xf2, 0xf2, 0xf0, 0xf1, 0xf1, 0xf9, 0xf6, 0xf9, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf7, 0xf7, 0xf5, 0xf1, 0xf9, 0xf9, 0xf3, 0xf9, 0xf7, 0xf2, 0xf7, 0xf0, 0xf2, 0xf3, 0xf1, 0xf0, 0xf7, 0xf4, 0xf7, 0xf0, 0xf0, 0xf7, 0xf9, 0xf9, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf6, 0xf6, 0xf2, 0xf2, 0x40, 0x40, 0x40, 0xd7, 0xe8, 0xc1, 0xe3, 0xc5, 0xd9, 0xd6, 0xc3, 0xc8, 0xd2, 0xc1, 0x40, 0xf7, 0xf9, 0xf0, 0xf4, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd7, 0x96, 0x84, 0x96, 0x93, 0xa2, 0x92, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xd9, 0xe4, 0xe2, 0xf0, 0xf1, 0xf6, 0xd9, 0xf2, 0xf3, 0xf0, 0xf2, 0xf0, 0xf0, 0xf6, 0xf1, 0xf0, 0xf5, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf6, 0xf4, 0xf3, 0xf8, 0xf4, 0xf0, 0xf1, 0xf1, 0xf6, 0x5f, 0x2a, 0x2, 0x6, 0x43, 0x82, 0x2, 0x19, 0x80, 0x84, 0x7, 0xa0, 0x0, 0x0, 0x0, 0x4, 0x10, 0x10, 0x95, 0x5, 0x0, 0x0, 0x0, 0x80, 0x0, 0x9a, 0x3, 0x19, 0x10, 0x8, 0x9c, 0x1, 0x0, 0x9f, 0x2, 0x6, 0x0, 0x0, 0x0, 0x2, 0x1, 0x26, 0x9f, 0x3, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9f, 0x10, 0x12, 0x1, 0x10, 0xa0, 0x40, 0x3, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x9f, 0x1a, 0x2, 0x6, 0x43, 0x9f, 0x26, 0x8, 0x2, 0x5f, 0x1a, 0x91, 0x8b, 0x20, 0xdc, 0x91, 0x9f, 0x27, 0x1, 0x80, 0x9f, 0x33, 0x3, 0xe0, 0x8, 0xc8, 0x9f, 0x34, 0x3, 0x1f, 0x3, 0x0, 0x9f, 0x36, 0x2, 0x1, 0x60, 0x9f, 0x37, 0x4, 0xcf, 0x25, 0x9a, 0x1f, 0xf0, 0xf2, 0xf2, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf3, 0xf0, 0xf0, 0xf6, 0xf4, 0xf3, 0xf1, 0xf4, 0xf2, 0xf1, 0xf0, 0xf3, 0xf0, 0xf0, 0xf9, 0xe3, 0xd5, 0xe6, 0xc7, 0xc1, 0xd7, 0xf4, 0xf0, 0xf5},
	[]byte{0x30, 0x32, 0x30, 0x30, 0xf2, 0x3a, 0xc4, 0x81, 0x28, 0xe0, 0x80, 0x10, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x1, 0x0, 0x31, 0x36, 0x35, 0x33, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x31, 0x37, 0x31, 0x34, 0x31, 0x37, 0x32, 0x38, 0x30, 0x30, 0x30, 0x30, 0x36, 0x34, 0x31, 0x33, 0x35, 0x32, 0x35, 0x33, 0x30, 0x35, 0x31, 0x36, 0x30, 0x35, 0x31, 0x37, 0x30, 0x35, 0x31, 0x36, 0x36, 0x30, 0x31, 0x30, 0x39, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x36, 0x39, 0x39, 0x39, 0x39, 0x30, 0x35, 0x33, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x31, 0x33, 0x36, 0x31, 0x39, 0x39, 0x30, 0x39, 0x34, 0x34, 0x38, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x38, 0x30, 0x31, 0x31, 0x30, 0x30, 0x30, 0x31, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x20, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x4d, 0x6f, 0x73, 0x63, 0x6f, 0x77, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x55, 0x36, 0x34, 0x33, 0x30, 0x38, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x4d, 0x43, 0x54, 0x49, 0x44, 0x48, 0x49, 0x31, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31, 0x31, 0x36, 0x34, 0x33, 0x30, 0x36, 0x35, 0x55, 0x44, 0x30, 0x36, 0x30, 0x43, 0x49, 0x30, 0x34, 0x35, 0x30, 0x31, 0x30, 0x31, 0x30, 0x30, 0x32, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x31, 0x30, 0x30, 0x34, 0x30, 0x31, 0x32, 0x30, 0x35, 0x30, 0x31, 0x30, 0x30, 0x36, 0x30, 0x31, 0x30, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x38, 0x30, 0x31, 0x30, 0x30, 0x39, 0x30, 0x31, 0x30, 0x45, 0x52, 0x30, 0x30, 0x35, 0x34, 0x30, 0x30, 0x37, 0x33},
}

type iso8583Msg struct {
	MessageTypeIndicator                    string `iso8538:"MTI"` // Message type indicator (MTI)
	PrimaryAccountNumber                    string `iso8538:"2"`   // Primary account number (PAN)
	ProcessingCode                          string `iso8538:"3"`   // Processing code
	AmountTransaction                       string `iso8538:"4"`   // Amount, transaction
	AmountSettlement                        string `iso8538:"5"`   // Amount, settlement
	AmountCardholderBilling                 string `iso8538:"6"`   // Amount, cardholder billing
	TransmissionDateTime                    string `iso8538:"7"`   // Transmission date & time
	AmountCardholderBillingFee              string `iso8538:"8"`   // Amount, cardholder billing fee
	ConversionRateSettlement                string `iso8538:"9"`   // Conversion rate, settlement
	ConversionRateCardholderBilling         string `iso8538:"10"`  // Conversion rate, cardholder billing
	SystemTraceAuditNumber                  string `iso8538:"11"`  // System trace audit number (STAN)
	LocalTransactionTime                    string `iso8538:"12"`  // Local transaction time (hhmmss)
	LocalTransactionDate                    string `iso8538:"13"`  // Local transaction date (MMDD)
	ExpirationDate                          string `iso8538:"14"`  // Expiration date
	SettlementDate                          string `iso8538:"15"`  // Settlement date
	CurrencyConversionDate                  string `iso8538:"16"`  // Currency conversion date
	CaptureDate                             string `iso8538:"17"`  // Capture date
	MerchantTypeOrMerchantCategoryCode      string `iso8538:"18"`  // Merchant type, or merchant category code
	AcquiringInstitutionCountryCode         string `iso8538:"19"`  // Acquiring institution (country code)
	PANExtendedCountryCode                  string `iso8538:"20"`  // PAN extended (country code)
	ForwardingInstitutionCountryCode        string `iso8538:"21"`  // Forwarding institution (country code)
	PointOfServiceEntryMode                 string `iso8538:"22"`  // Point of service entry mode
	ApplicationPANSequenceNumber            string `iso8538:"23"`  // Application PAN sequence number
	FunctionCode                            string `iso8538:"24"`  // Function code (ISO 8583:1993), or network international identifier (NII)
	PointOfServiceConditionCode             string `iso8538:"25"`  // Point of service condition code
	PointOfServiceCaptureCode               string `iso8538:"26"`  // Point of service capture code
	AuthorizingIdentificationResponseLength string `iso8538:"27"`  // Authorizing identification response length
	AmountTransactionFee                    string `iso8538:"28"`  // Amount, transaction fee
	AmountSettlementFee                     string `iso8538:"29"`  // Amount, settlement fee
	AmountTransactionProcessingFee          string `iso8538:"30"`  // Amount, transaction processing fee
	AmountSettlementProcessingFee           string `iso8538:"31"`  // Amount, settlement processing fee
	AcquiringInstitutionIdentificationCode  string `iso8538:"32"`  // Acquiring institution identification code
	ForwardingInstitutionIdentificationCode string `iso8538:"33"`  // Forwarding institution identification code
	PrimaryAccountNumberExtended            string `iso8538:"34"`  // Primary account number, extended
	Track2Data                              string `iso8538:"35"`  // Track 2 data (PAN with expiry date)
	Track3Data                              string `iso8538:"36"`  // Track 3 data
	RetrievalReferenceNumber                string `iso8538:"37"`  // Retrieval reference number
	AuthorizationIdentificationResponse     string `iso8538:"38"`  // Authorization identification response
	ResponseCode                            string `iso8538:"39"`  // Response code
	ServiceRestrictionCode                  string `iso8538:"40"`  // Service restriction code
	CardAcceptorTerminalIdentification      string `iso8538:"41"`  // Card acceptor terminal identification
	CardAcceptorIdentificationCode          string `iso8538:"42"`  // Card acceptor identification code
	CardAcceptorNameLocation                string `iso8538:"43"`  // Card acceptor name/location (1–23 street address, –36 city, –38 state, 39–40 country)
	AdditionalResponseData                  string `iso8538:"44"`  // Additional response data
	Track1Data                              string `iso8538:"45"`  // Track 1 data
	AdditionalDataISO                       string `iso8538:"46"`  // Additional data (ISO)
	AdditionalDataNational                  string `iso8538:"47"`  // Additional data (national)
	AdditionalDataPrivate                   string `iso8538:"48"`  // Additional data (private) (DE 48: Subelement 10 — Encrypted PIN Block Key, Subelement 80 — PIN Service Code)
	CurrencyCodeTransaction                 string `iso8538:"49"`  // Currency code, transaction
	CurrencyCodeSettlement                  string `iso8538:"50"`  // Currency code, settlement
	CurrencyCodeCardholderBilling           string `iso8538:"51"`  // Currency code, cardholder billing
	PersonalIdentificationNumberData        string `iso8538:"52"`  // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	SecurityRelatedControlInformation       string `iso8538:"53"`  // Security related control information (DE 53: Security-Related Control Information, Subfield 1 — PIN Security Type Code, Subfield 2 — PIN Encryption Type Code, Subfield 3 — PIN Block Format Code, Subfield 4—PIN Key Index Number)
	AdditionalAmounts                       string `iso8538:"54"`  // Additional amounts
	ICCData                                 string `iso8538:"55"`  // ICC data – EMV having multiple tags
	ReservedISO                             string `iso8538:"56"`  // Reserved (ISO)
	ReservedNational1                       string `iso8538:"57"`  // Reserved (national)
	ReservedNational2                       string `iso8538:"58"`  // Reserved (national)
	ReservedNational3                       string `iso8538:"59"`  // Reserved (national)
	ReservedNational4                       string `iso8538:"60"`  // Reserved (national) (e.g. settlement request: batch number, advice transactions: original transaction amount, batch upload: original MTI plus original RRN plus original STAN, etc.)
	ReservedPrivate1                        string `iso8538:"61"`  // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	ReservedPrivate2                        string `iso8538:"62"`  // Reserved (private) (e.g. transactions: invoice number, key exchange transactions: TPK key, etc.)
	ReservedPrivate3                        string `iso8538:"63"`  // Reserved (private)
	MessageAuthenticationCode1              string `iso8538:"64"`  // Message authentication code (MAC)
	ExtendedBitmapIndicator                 string `iso8538:"65"`  // Extended bitmap indicator
	SettlementCode                          string `iso8538:"66"`  // Settlement code
	ExtendedPaymentCode                     string `iso8538:"67"`  // Extended payment code
	ReceivingInstitutionCountryCode         string `iso8538:"68"`  // Receiving institution country code
	SettlementInstitutionCountryCode        string `iso8538:"69"`  // Settlement institution country code
	NetworkManagementInformationCode        string `iso8538:"70"`  // Network management information code
	MessageNumber                           string `iso8538:"71"`  // Message number
	LastMessagesNumber                      string `iso8538:"72"`  // Last message's number
	ActionDate                              string `iso8538:"73"`  // Action date (YYMMDD)
	NumberOfCredits                         string `iso8538:"74"`  // Number of credits
	CreditsReversalNumber                   string `iso8538:"75"`  // Credits, reversal number
	NumberOfDebits                          string `iso8538:"76"`  // Number of debits
	DebitsReversalNumber                    string `iso8538:"77"`  // Debits, reversal number
	TransferNumber                          string `iso8538:"78"`  // Transfer number
	TransferReversalNumber                  string `iso8538:"79"`  // Transfer, reversal number
	NumberOfInquiries                       string `iso8538:"80"`  // Number of inquiries
	NumberOfAuthorizations                  string `iso8538:"81"`  // Number of authorizations
	CreditsProcessingFeeAmount              string `iso8538:"82"`  // Credits, processing fee amount
	CreditsTransactionFeeAmount             string `iso8538:"83"`  // Credits, transaction fee amount
	DebitsProcessingFeeAmount               string `iso8538:"84"`  // Debits, processing fee amount
	DebitsTransactionFeeAmount              string `iso8538:"85"`  // Debits, transaction fee amount
	TotalAmountOfCredits                    string `iso8538:"86"`  // Total amount of credits
	CreditsReversalAmount                   string `iso8538:"87"`  // Credits, reversal amount
	TotalAmountOfDebits                     string `iso8538:"88"`  // Total amount of debits
	DebitsReversalAmount                    string `iso8538:"89"`  // Debits, reversal amount
	OriginalDataElements                    string `iso8538:"90"`  // Original data elements
	FileUpdateCode                          string `iso8538:"91"`  // File update code
	FileSecurityCode                        string `iso8538:"92"`  // File security code
	ResponseIndicator                       string `iso8538:"93"`  // Response indicator
	ServiceIndicator                        string `iso8538:"94"`  // Service indicator
	ReplacementAmounts                      string `iso8538:"95"`  // Replacement amounts
	MessageSecurityCode                     string `iso8538:"96"`  // Message security code
	NetSettlementAmount                     string `iso8538:"97"`  // Net settlement amount
	Payee                                   string `iso8538:"98"`  // Payee
	SettlementInstitutionIdentificationCode string `iso8538:"99"`  // Settlement institution identification code
	ReceivingInstitutionIdentificationCode  string `iso8538:"100"` // Receiving institution identification code
	FileName                                string `iso8538:"101"` // File name
	AccountIdentification1                  string `iso8538:"102"` // Account identification 1
	AccountIdentification2                  string `iso8538:"103"` // Account identification 2
	TransactionDescription                  string `iso8538:"104"` // Transaction description
	ReservedForISOUse1                      string `iso8538:"105"` // Reserved for ISO use
	ReservedForISOUse2                      string `iso8538:"106"` // Reserved for ISO use
	ReservedForISOUse3                      string `iso8538:"107"` // Reserved for ISO use
	ReservedForISOUse4                      string `iso8538:"108"` // Reserved for ISO use
	ReservedForISOUse5                      string `iso8538:"109"` // Reserved for ISO use
	ReservedForISOUse6                      string `iso8538:"110"` // Reserved for ISO use
	ReservedForISOUse7                      string `iso8538:"111"` // Reserved for ISO use
	ReservedForNationalUse1                 string `iso8538:"112"` // Reserved for national use
	ReservedForNationalUse2                 string `iso8538:"113"` // Reserved for national use
	ReservedForNationalUse3                 string `iso8538:"114"` // Reserved for national use
	ReservedForNationalUse4                 string `iso8538:"115"` // Reserved for national use
	ReservedForNationalUse5                 string `iso8538:"116"` // Reserved for national use
	ReservedForNationalUse6                 string `iso8538:"117"` // Reserved for national use
	ReservedForNationalUse7                 string `iso8538:"118"` // Reserved for national use
	ReservedForNationalUse8                 string `iso8538:"119"` // Reserved for national use
	ReservedForPrivateUse1                  string `iso8538:"120"` // Reserved for private use
	ReservedForPrivateUse2                  string `iso8538:"121"` // Reserved for private use
	ReservedForPrivateUse3                  string `iso8538:"122"` // Reserved for private use
	ReservedForPrivateUse4                  string `iso8538:"123"` // Reserved for private use
	ReservedForPrivateUse5                  string `iso8538:"124"` // Reserved for private use
	ReservedForPrivateUse6                  string `iso8538:"125"` // Reserved for private use
	ReservedForPrivateUse7                  string `iso8538:"126"` // Reserved for private use
	ReservedForPrivateUse8                  string `iso8538:"127"` // Reserved for private use
	MessageAuthenticationCode2              string `iso8538:"128"` // Message authentication code
}

type iso8583Msg2 struct {
	MTI                                string `iso8538:"MTI"`
	PrimaryAccountNumber               string `iso8538:"2"`
	ProcessingCode                     string `iso8538:"3"`
	AmountOrig                         string `iso8538:"4"`
	Amount                             string `iso8538:"6"`
	TransmissionDateTime               string `iso8538:"7"`
	BillingRate                        string `iso8538:"10"`
	TraceNumber                        string `iso8538:"11"`
	LocalTime                          string `iso8538:"12"`
	LocalDate                          string `iso8538:"13"`
	DateExpiration                     string `iso8538:"14"`
	DateSettlement                     string `iso8538:"15"`
	DateCapture                        string `iso8538:"17"`
	MerchantType                       string `iso8538:"18"`
	AcquiringInstitutionCountryCode    string `iso8538:"19"`
	POSDataCode                        string `iso8538:"22"`
	PointOfServiceConditionCode        string `iso8538:"25"`
	TransactionFee                     string `iso8538:"28"`
	ONLINEIssuerAuthorizationFeeAmount string `iso8538:"31"`
	AcquirerInstitutionID              string `iso8538:"32"`
	TrackData                          string `iso8538:"35"`
	RetrievalReference                 string `iso8538:"37"`
	AuthIDCode                         string `iso8538:"38"`
	RespCode                           string `iso8538:"39"`
	CardAccptrTermnlID                 string `iso8538:"41"`
	CardAccptrIDCode                   string `iso8538:"42"`
	CardAccptrNameLoc                  string `iso8538:"43"`
	AdditionalResponseData             string `iso8538:"44"`
	CurrencyOrig                       string `iso8538:"49"`
	Currency                           string `iso8538:"51"`
	PersonalIdentificationNumberData   string `iso8538:"52"`
	SecurityRelatedControlInformation  string `iso8538:"53"`
	AddtnlAmounts                      string `iso8538:"54"`
	ICCRelatedData                     string `iso8538:"55"`
	OriginalDataSerials                string `iso8538:"56"`
	AdditionalInformation              string `iso8538:"60"`
	OtherAmtTrans                      string `iso8538:"61"`
	NetworkManagementInformationCode   string `iso8538:"70"`
	BusinessDate                       string `iso8538:"73"`
	OrigDataElemts                     string `iso8538:"90"`
	NumberOfAccounts                   string `iso8538:"93"`
	QuerySequence                      string `iso8538:"94"`
	ReplacementAmount                  string `iso8538:"95"`
	MoreFlag                           string `iso8538:"99"`
	MessageOriginator                  string `iso8538:"100"`
	AccountFrom                        string `iso8538:"102"`
	AccountTo                          string `iso8538:"103"`
	PrivateData                        string `iso8538:"104"`
	AdditionalInformationPart2         string `iso8538:"116"`
	AdditionalAmountAccountTo          string `iso8538:"117"`
	AdditionalInformationPart1         string `iso8538:"120"`
	Transfercurrencies                 string `iso8538:"122"`
	CardholderUtilityAccount           string `iso8538:"125"`
	PrivateUseFields                   string `iso8538:"126"`
}

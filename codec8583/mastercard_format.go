package codec8583

var (
	MastercardMarshaler   = NewMarshaler(Mastercard)
	MastercardUnmarshaler = NewUnmarshaler(Mastercard)

	// MastercardTruncate decoder and truncator (masker) of the decoded ISO 8583 message.
	MastercardTruncate             = make(Format)
	MastercardUnmarshalerTruncator Unmarshaler

	// MastercardHash256 decoder and hasher (obfuscatator) of the ISO 8583 message.
	MastercardHash256              = make(Format)
	MastercardUnmarshalerHasher256 Unmarshaler

	// MastercardHashed256 coder/decoder of the ISO 8583 obfuscatated message.
	MastercardHashed256            = make(Format)
	MastercardHashed256Marshaler   Marshaler
	MastercardHashed256Unmarshaler Unmarshaler
)

func init() {
	for k, v := range Mastercard {
		MastercardTruncate[k] = v
	}
	MastercardTruncate[2] = LLVAR{19, EBCDIC, EncN, DecPANTruncate}         // Primary account number (PAN)
	MastercardTruncate[35] = LLVAR{37, EBCDIC, EncANS, DecFirstPANTruncate} // Track 2 data (PAN with expiry date)
	MastercardTruncate[52] = FIX{8, EBCDIC, EncB, DecNullify}               // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	MastercardUnmarshalerTruncator = NewUnmarshaler(MastercardTruncate)

	for k, v := range Mastercard {
		MastercardHash256[k] = v
	}
	MastercardHash256[2] = LLVAR{64, EBCDIC, EncN, DecHash256}    // Primary account number (PAN) (maximum PAN length is 19 characters but we use the maximum length of 64 to fit the SHA 256 hash sum of the PAN)
	MastercardHash256[35] = LLVAR{64, EBCDIC, EncANS, DecHash256} // Track 2 data (PAN with expiry date) (maximum PAN with expiry date length is 37 characters but we use the maximum length of 64 to fit the SHA 256 hash sum of the PAN with expiry date)
	MastercardHash256[52] = FIX{8, EBCDIC, EncB, DecNullify}      // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	MastercardUnmarshalerHasher256 = NewUnmarshaler(MastercardHash256)

	for k, v := range MastercardHash256 {
		MastercardHashed256[k] = v
	}
	MastercardHashed256[2] = LLVAR{64, EBCDIC, EncN, DecN}      // Primary account number (PAN) (maximum PAN length is 19 characters but we use the maximum length of 64 to fit the SHA 256 hash sum of the PAN)
	MastercardHashed256[35] = LLVAR{64, EBCDIC, EncANS, DecANS} // Track 2 data (PAN with expiry date) (maximum PAN with expiry date length is 37 characters but we use the maximum length of 64 to fit the SHA 256 hash sum of the PAN with expiry date)
	MastercardHashed256[52] = FIX{8, EBCDIC, EncB, EncB}        // Personal identification number data (DE 52: Personal ID Number (PIN) Data)
	MastercardHashed256Marshaler = NewMarshaler(MastercardHashed256)
	MastercardHashed256Unmarshaler = NewUnmarshaler(MastercardHashed256)
}

// Mastercard is a format of the ISO 8583 message which maps the codecs to MIT/bitmaps/each individual field.
var Mastercard = Format{
	-1:  MTIEbcdicCodec,                      // Message type indicator (MTI). Is a four-digit numeric data element describing the type of message being interpreted. The MTI is required and must be present as the first data element of each Authorization Platform message.
	0:   BitmapCodec,                         // Primary Bitmap. Must always be present in a message. The most frequently used data elements are indexed from DE 1 (Bit Map, Secondary) through DE 64 (Message Authentication Code [MAC]). Infrequently used data elements are indexed from the DE 66 (Settlement Code) through DE 128 (Message Authentication Code [MAC]).
	1:   BitmapCodec,                         // Second Bitmap. Is a series of 64 bits that identifies the presence (1) or absence (0) of each data element in the second segment of a message. This would include DE 65 (Bit Map, Extended) through DE 128 (Message Authentication Code [MAC]).
	2:   LLVAR{19, EBCDIC, EncN, DecN},       // Primary account number (PAN). Is a series of digits used to identify a customer account or relationship.
	3:   FIX{6, EBCDIC, EncN, DecN},          // Processing code. Describes the effect of a transaction on the customer account and the type of accounts affected.
	4:   FIX{12, EBCDIC, EncN, DecN},         // Amount, transaction. Is the amount of funds the cardholder requested in the local currency of the acquirer or source location of the transaction.
	5:   FIX{12, EBCDIC, EncN, DecN},         // Amount, settlement. Is the amount of funds to be transferred between the acquirer and the issuer equal to DE 4 (Amount, Transaction) in the settlement currency. Mastercard programs and services use U.S. dollars as the currency of settlement.
	6:   FIX{12, EBCDIC, EncN, DecN},         // Amount, cardholder billing. Indicates the transaction amount in the issuer’s currency. It is the amount billed to the cardholder in the cardholder account currency, excluding cardholder billing fees.
	7:   FIX{10, EBCDIC, EncN, DecN},         // Transmission date & time. Is the date and time that a message is entered into the Mastercard Network. Date and time must be expressed in Coordinated Universal Time (UTC).
	8:   FIX{8, EBCDIC, EncN, DecN},          // Amount, cardholder billing fee. Is the fee the issuer is to bill to the cardholder in the same currency as DE 6 (Amount, Cardholder Billing).
	9:   FIX{8, EBCDIC, EncN, DecN},          // Conversion rate, settlement. Is the factor used in the conversion from transaction to settlement amount. DE 4 (Amount, Transaction) is multiplied by DE 9 to determine DE 5 (Amount, Settlement).
	10:  FIX{8, EBCDIC, EncN, DecN},          // Conversion rate, cardholder billing. Is the factor used in the conversion from transaction to cardholder billing amount. DE 4 (Amount, Transaction) is multiplied by DE 10 to determine DE 6 (Amount, Cardholder Billing).
	11:  FIX{6, EBCDIC, EncN, DecN},          // System trace audit number (STAN). Is a number a message initiator assigns to uniquely identify a transaction.
	12:  FIX{6, EBCDIC, EncN, DecN},          // Local transaction time (hhmmss). Is the local time at which the transaction takes place at the point of card acceptor location.
	13:  FIX{4, EBCDIC, EncN, DecN},          // Local transaction date (MMDD). Is the local month and day on which the transaction takes place at the point of card acceptor location.
	14:  FIX{4, EBCDIC, EncN, DecN},          // Expiration date. Specifies the year and month after which an issuer designates a cardholder's card to be "expired".
	15:  FIX{4, EBCDIC, EncN, DecN},          // Settlement date. Is the date (month and day) that funds will be transferred between an acquirer and an issuer or an appropriate intermediate network facility (INF).
	16:  FIX{4, EBCDIC, EncN, DecN},          // Currency conversion date. Indicates the effective date of DE 9 (Conversion Rate, Settlement) and also DE 10 (Conversion Rate, Cardholder Billing) whenever these data elements are present within a message.
	17:  FIX{4, EBCDIC, EncN, DecN},          // Capture date. Is the month and day the acquirer processed the transaction data.
	18:  FIX{4, EBCDIC, EncN, DecN},          // Merchant type, or merchant category code. Is the classification (card acceptor business code/merchant category code [MCC]) of the merchant's type of business or service.
	19:  FIX{3, EBCDIC, EncN, DecN},          // Acquiring institution (country code). Is the code of the country where the acquiring institution is located. Refer to the ISO 3166 specification for more information.
	20:  FIX{3, EBCDIC, EncN, DecN},          // PAN extended (country code). Is a code identifying the country where the card issuer is located.
	21:  FIX{3, EBCDIC, EncN, DecN},          // Forwarding institution (country code). Is the code of the country where the forwarding institution is located.
	22:  FIX{3, EBCDIC, EncN, DecN},          // Point of service (POS) entry mode. Consists of numeric codes to indicate the method by which the PAN was entered into the interchange system and to indicate the POS terminal PIN entry capabilities.
	23:  FIX{3, EBCDIC, EncN, DecN},          // Application PAN sequence number. Distinguishes among separate cards having the same DE 2 (Primary Account Number [PAN]) or DE 34 (Primary Account Number [PAN] Extended). Issuers may encode chip cards with Card Sequence Numbers. Acquirers with chip-reading capability may pass this information encoded on the chip in DE 23 of Authorization Request/0100 messages.
	24:  FIX{3, EBCDIC, EncN, DecN},          // Function code (ISO 8583:1993), or network international identifier (NII). Identifies a single international network of card issuers.
	25:  FIX{2, EBCDIC, EncN, DecN},          // Point of service (POS) condition code. Is an ID of the condition under which the transaction takes place at the point of interaction.
	26:  FIX{2, EBCDIC, EncN, DecN},          // Point of service (POS) capture code. Indicates the technique, maximum number, or both of PIN characters that can be accepted by the POS terminal used to construct the PIN data.
	27:  FIX{1, EBCDIC, EncN, DecN},          // Authorizing identification response length. Is the maximum length of the authorization response that the acquirer can accommodate. The issuer or its agent is expected to limit response to this length.
	28:  FIX{9, EBCDIC, EncAN, DecAN},        // Amount, transaction fee. Is the fee charged (for example, by the acquirer) for transaction activity in the currency of DE 4 (Amount, Transaction).
	29:  FIX{9, EBCDIC, EncAN, DecAN},        // Amount, settlement fee. Is the fee to be transferred between the acquirer and the issuer equal to DE 28 (Amount, Transaction Fee) in the currency of DE 5 (Amount, Settlement).
	30:  FIX{9, EBCDIC, EncAN, DecAN},        // Amount, transaction processing fee. Is the fee charged (for example, by the acquirer, issuer, or INF) for the handling and routing of messages in the currency of DE 4 (Amount, Transaction).
	31:  FIX{9, EBCDIC, EncAN, DecAN},        // Amount, settlement processing fee. Is the fee charged (for example, by the acquirer, issuer, or INF) for the handling and routing of messages in the currency of DE 5 (Amount, Settlement).
	32:  LLVAR{6, EBCDIC, EncN, DecN},        // Acquiring institution identification code. Identifies the acquiring institution (for example, merchant bank) or its agent.
	33:  LLVAR{6, EBCDIC, EncN, DecN},        // Forwarding institution identification code. Identifies the institution forwarding a Request or Advice message in an interchange system if it is not the same institution as specified in DE 32 (Acquiring Institution ID Code). DE 33 is used within a message to contain the Mastercard six- digit customer ID number of the CPS or INF responsible for directly routing that message to the Authorization Platform.
	34:  LLVAR{28, EBCDIC, EncANS, DecANS},   // Primary account number (PAN), extended. Identifies a customer account or relationship, and is used only when PAN begins with a 59 BIN.
	35:  LLVAR{37, EBCDIC, EncANS, DecANS},   // Track 2 data (PAN with expiry date). Is the information encoded on track 2 of the card magnetic stripe as defined in the ISO 7813 specification, including data element separators but excluding beginning and ending sentinels and longitudinal redundancy check (LRC) characters as defined therein.
	36:  LLLVAR{104, EBCDIC, EncANS, DecANS}, // Track 3 data. Is the information encoded on track 3 of the card magnetic stripe as defined in the ISO 4909–1986 specification, including data element separators but excluding beginning and ending sentinels and LRC characters as defined therein.
	37:  FIX{12, EBCDIC, EncAN, DecAN},       // Retrieval reference number. Is a document reference number supplied by the system retaining the original source document of the transaction and assists in locating that source document or a copy thereof. DE 37 is made available for use by automated merchant POS systems that may be interconnected into the interchange system. Merchant POS systems may assign a unique receipt or sales document ID to be used to satisfy regulatory or legal requirements when the merchant performs source document capture and truncation. DE 37 may be used to relay source document reference numbers to the issuer at the time each transaction is processed.
	38:  FIX{6, EBCDIC, EncANS, DecANS},      // Authorization identification response. Is a transaction response ID code that the authorizing institution assigns. DE 38 is used to transmit a card issuer's "authorization code" for Authorization transactions.
	39:  FIX{2, EBCDIC, EncAN, DecAN},        // Response code. Defines the disposition of a previous message or an action taken as a result of receipt of a previous message. Response codes also are used to indicate approval or decline of a transaction. In the event an authorization is declined, the response code indicates the reason for rejection and may indicate an action to be taken at the card acceptor (for example, capture card).
	40:  FIX{3, EBCDIC, EncAN, DecAN},        // Service restriction code. Identifies geographic or service availability.
	41:  FIX{8, EBCDIC, EncANS, DecANS},      // Card acceptor terminal identification. Uniquely identifies a terminal at the card acceptor location of acquiring institutions or merchant POS systems. The terminal ID should be printed on all transaction receipts in ATM and POS transactions where the terminal is capable of generating customer receipts.
	42:  FIX{15, EBCDIC, EncANS, DecANS},     // Card acceptor identification code. Identifies the card acceptor that defines the point of the transaction in both local and interchange environments. DE 42 is used as a merchant ID to uniquely identify the merchant in a POS transaction.
	43:  FIX{40, EBCDIC, EncANS, DecANS},     // Card acceptor name/location (1–23 street address, –36 city, –38 state, 39–40 country). Contains the name and location of the card acceptor that defines the point of interaction in both local and interchange environments (excluding ATM and Card-Activated Public Phones).
	44:  LLVAR{25, EBCDIC, EncANS, DecANS},   // Additional response data. Provides other supplemental data that may be required in response to an authorization or other type of transaction request. This data element may also be present in any response message when DE 39 (Response Code) contains the value 30, indicating that a Format Error condition was detected in the preceding message. In this case, the first three bytes of DE 44 (if present) will contain a three-digit numeric value indicating the data element number where the format error occurred.
	45:  LLVAR{76, EBCDIC, EncANS, DecANS},   // Track 1 data. Is the information encoded on track 1 of the card’s magnetic stripe as defined in the ISO 7813 specification, including data element separators but excluding beginning and ending sentinels and LRC characters as defined in this data element definition.
	46:  LLLVAR{999, EBCDIC, EncANS, DecANS}, // Additional data (ISO). Provides data supplemental to that already conveyed in the specific data elements in the message.
	47:  LLLVAR{999, EBCDIC, EncANS, DecANS}, // Additional data (national). Is reserved for national organizations to define data unique to country applications.
	48:  LLLVAR{999, EBCDIC, EncB, DecB},     // Additional data (private) (DE 48: Subelement 10 — Encrypted PIN Block Key, Subelement 80 — PIN Service Code). Is reserved for private organizations to define data unique to specific networks or specific programs and services. DE 48 provides other supplemental data in a message when a specific ISO-designated data element is not available. It is a free-format, variable-length data element that may be used for multiple purposes.
	49:  FIX{3, EBCDIC, EncN, DecN},          // Currency code, transaction. Is the local currency of the acquirer or source location of the transaction. It specifies the currency used in DE 4 (Amount, Transaction).
	50:  FIX{3, EBCDIC, EncN, DecN},          // Currency code, settlement. Defines the currency of DE 5 (Amount, Settlement) and DE 29 (Amount, Settlement Fee).
	51:  FIX{3, EBCDIC, EncN, DecN},          // Currency code, cardholder billing. Defines the currency of DE 6 (Amount, Cardholder Billing) and DE 8 (Amount, Cardholder Billing Fee).
	52:  FIX{8, EBCDIC, EncB, DecB},          // Personal identification number data (DE 52: Personal ID Number (PIN) Data). Contains a number assigned to a cardholder intended to uniquely identify that cardholder at the point of interaction. The use of the PIN is subject to bilateral agreement. The data element may contain the PIN itself or a derivative. This data element transmits PIN information from acquirers to issuers (or to the network) for PIN verification or validation.
	53:  FIX{16, EBCDIC, EncN, DecN},         // Security related control information (DE 53: Security-Related Control Information, Subfield 1 — PIN Security Type Code, Subfield 2 — PIN Encryption Type Code, Subfield 3 — PIN Block Format Code, Subfield 4—PIN Key Index Number). Is used with PIN data to provide specific information about PIN block encoding and PIN data encryption to assist the issuer (or its agent) in processing PINs entered at the point of interaction.
	54:  LLLVAR{120, EBCDIC, EncAN, DecAN},   // Additional amounts. Provides information on up to two amount types and related account data.
	55:  LLLVAR{255, EBCDIC, EncB, DecB},     // Integrated Circuit Card (ICC) data – EMV having multiple tags. Contains binary data that only the issuer, the issuer agent, or MDES processes; it is used locally by the payment application on the chip at a chip-capable terminal. This data element is present in chip full-grade transactions and can be present in DSRP transactions.
	56:  LLLVAR{37, EBCDIC, EncAN, DecAN},    // Payment Account Data. Contains unique, non-financial reference information associated with the PAN or token used to initiate the transaction.
	57:  LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved (national). Are reserved for future use.
	58:  LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved (national). Are reserved for future use.
	59:  LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved (national). Are reserved for future use.
	60:  LLLVAR{60, EBCDIC, EncANS, DecANS},  // Advice Reason Code. Indicates to the receiver of an Advice message the specific reason for the transmission of the Advice message.
	61:  LLLVAR{26, EBCDIC, EncANS, DecANS},  // Point of Service (POS) Data. Supersedes and replaces the ISO-specified DE 25 (Point-of- Service [POS] Condition Code) that customers must not use in the Authorization Request/ 0100. DE 61 indicates the conditions that exist at the point of service at the time of the transaction.
	62:  LLLVAR{100, EBCDIC, EncANS, DecANS}, // Intermediate Network Facility (INF) Data. Contains "acquiring network trace information" that INFs may require to quickly and accurately route Administrative Advice/0620 messages back to the original acquiring institution. DE 62 assists acquiring INF facilities that connect directly to the Authorization Platform. It allows these INFs to maintain sufficient information within a message to permit immediate online routing of chargebacks and retrieval requests without the requirement of maintaining large online reference databases containing the original transactions.
	63:  LLLVAR{50, EBCDIC, EncAN, DecAN},    // Network Data. Is generated by the Authorization Platform for each originating message routed through the network. The receiver must retain the data element and use it in any response or acknowledgement message associated with the originating message.
	64:  FIX{8, EBCDIC, EncB, DecB},          // Message authentication code (MAC). Validates the source and the text of the message between the sender and the receiver.
	65:  FIX{8, EBCDIC, EncB, DecB},          // Extended bitmap indicator. Is a series of eight bytes (64 bits) used to identify the presence (denoted by 1) or the absence (denoted by 0) of each data element in an extended (third) message segment.
	66:  FIX{1, EBCDIC, EncN, DecN},          // Settlement code. Indicates the result of a reconciliation request.
	67:  FIX{2, EBCDIC, EncN, DecN},          // Extended payment code. Indicates the number of months that the cardholder prefers to pay for an item (the item purchased during the course of this transaction) if permitted by the card issuer.
	68:  FIX{3, EBCDIC, EncN, DecN},          // Receiving institution country code. Is the code of the country where the receiving institution is located.
	69:  FIX{3, EBCDIC, EncN, DecN},          // Settlement institution country code. Is the code of the country where the settlement institution is located.
	70:  FIX{3, EBCDIC, EncN, DecN},          // Network management information code. Identifies network status.
	71:  FIX{4, EBCDIC, EncN, DecN},          // Message number. Is a sequential, cyclic number the message initiator assigns to a message. Message Number is used to monitor the integrity of interchange.
	72:  FIX{4, EBCDIC, EncN, DecN},          // Last message's number. Is a sequential, cyclic number the message initiator assigns to a message, used to monitor the integrity of interchange.
	73:  FIX{6, EBCDIC, EncN, DecN},          // Action date (YYMMDD). Specifies the date (year, month, and day) of a future action. In addition, a message originator may use it as a static time such as a birthdate.
	74:  FIX{10, EBCDIC, EncN, DecN},         // Number of credits. Is the numeric sum of credit transactions processed.
	75:  FIX{10, EBCDIC, EncN, DecN},         // Credits, reversal number. Is the sum number of reversal credit transactions.
	76:  FIX{10, EBCDIC, EncN, DecN},         // Number of debits. Is the sum number of debit transactions processed.
	77:  FIX{10, EBCDIC, EncN, DecN},         // Debits, reversal number. Is the sum number of reversal debit transactions.
	78:  FIX{10, EBCDIC, EncN, DecN},         // Transfer number. Is the sum number of all transfer transactions processed.
	79:  FIX{10, EBCDIC, EncN, DecN},         // Transfer, reversal number. Is the sum number of all transfer reversal transactions processed.
	80:  FIX{10, EBCDIC, EncN, DecN},         // Number of inquiries. Is the sum number of inquiry transaction requests processed.
	81:  FIX{10, EBCDIC, EncN, DecN},         // Number of authorizations. Is the sum number of Authorization Request/0100 and Authorization Advice/0120 messages processed.
	82:  FIX{12, EBCDIC, EncN, DecN},         // Credits, processing fee amount. Is the sum of all processing fees due to an institution or customer for services associated with handling and routing transactions. This Mastercard definition replaces the ISO standard definition.
	83:  FIX{12, EBCDIC, EncN, DecN},         // Credits, transaction fee amount. Is the sum of all transaction fees due to an institution or customer for processing interchange transactions. This Mastercard definition replaces the ISO standard definition.
	84:  FIX{12, EBCDIC, EncN, DecN},         // Debits, processing fee amount. Is the sum of all processing fees due from an institution or customer for services associated with handling and routing transactions. This Mastercard definition replaces the ISO standard definition.
	85:  FIX{12, EBCDIC, EncN, DecN},         // Debits, transaction fee amount. Is the sum of all transaction fees due from an institution or customer for processing interchange transactions. This Mastercard definition replaces the ISO standard definition.
	86:  FIX{16, EBCDIC, EncN, DecN},         // Total amount of credits. Is the sum amount of all credit transactions processed exclusive of any fees.
	87:  FIX{16, EBCDIC, EncN, DecN},         // Credits, reversal amount. Is the sum amount of reversal credits processed exclusive of any fees.
	88:  FIX{16, EBCDIC, EncN, DecN},         // Total amount of debits. Is the sum amount of all debit transactions processed exclusive of any fees.
	89:  FIX{16, EBCDIC, EncN, DecN},         // Debits, reversal amount. Is the sum amount of reversal debits processed exclusive of any fees.
	90:  FIX{42, EBCDIC, EncN, DecN},         // Original data elements. Is the data elements in the original message, intended to identify a transaction for correction or reversal.
	91:  FIX{1, EBCDIC, EncAN, DecAN},        // File update code. Indicates to the system maintaining a file which procedure to follow.
	92:  FIX{2, EBCDIC, EncAN, DecAN},        // File security code. Is an Issuer File Update security code used to indicate that a message originator is authorized to update a file.
	93:  FIX{5, EBCDIC, EncN, DecN},          // Response indicator. Indicates the update action a POS system takes.
	94:  FIX{7, EBCDIC, EncANS, DecANS},      // Service indicator. Indicates the service a message recipient requires.
	95:  FIX{42, EBCDIC, EncN, DecN},         // Replacement amounts. Contains the “actual amount” subfields necessary to perform a partial or full reversal of a financial transaction.
	96:  FIX{8, EBCDIC, EncN, DecN},          // Message security code. Is a verification between a card acceptor and a card issuer that a message is authorized to update or modify a special file.
	97:  FIX{17, EBCDIC, EncAN, DecAN},       // Net settlement amount. Is the net value of all gross amounts.
	98:  FIX{25, EBCDIC, EncANS, DecANS},     // Payee. Is the third-party beneficiary in a payment transaction.
	99:  LLVAR{11, EBCDIC, EncN, DecN},       // Settlement institution identification code. Identifies the settlement institution or its agent.
	100: LLVAR{11, EBCDIC, EncN, DecN},       // Receiving institution identification code. Is the identity of the institution receiving a Request or Advice message in an interchange system if not the same as identified in DE 2 (Primary Account Number [PAN]) or DE 34 (Primary Account Number [PAN], Extended). The Authorization Platform uses DE 100 to determine the destination routing of Administrative/ 06xx messages. For these messages, DE 33 (Forwarding Institution ID Code) identifies the sender of the message; DE 100 identifies the receiver of the message.
	101: LLVAR{17, EBCDIC, EncANS, DecANS},   // File name. Is the actual or abbreviated name of the file that the issuer accesses. DE 101 is used in Issuer File Update/03xx messages to identify the specific name of an Authorization Platform data file or program parameter table that is being updated by a customer’s Issuer File Update Request/0302.
	102: LLVAR{28, EBCDIC, EncANS, DecANS},   // Account identification 1. Is a series of digits that identify a customer account or relationship. Customers primarily use it for the "from" account in a transfer transaction. DE 102 may be used in Authorization Request Response/0110 messages to identify the specific "from" account that the transaction affected. DE 102 may be used for printing on cardholder transaction receipts.
	103: LLVAR{28, EBCDIC, EncANS, DecANS},   // Account identification 2. Is a series of digits that identify a customer account or relationship. Customers primarily use it for the "to" account in a transfer transaction.
	104: LLLVAR{100, EBCDIC, EncANS, DecANS}, // Transaction description. Describes additional characteristics of the transaction for billing purposes
	105: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved for Mastercard Use
	106: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved for Mastercard Use
	107: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved for Mastercard Use
	108: LLLVAR{999, EBCDIC, EncANS, DecANS}, // MoneySend Reference Data. Provides the capability for the acquirers to send in Sender, Receiver, and Transaction level data to the issuer in MoneySend Payment Transactions or MoneySend Funding Transactions.
	109: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved for ISO use
	110: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved for ISO use
	111: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Reserved for ISO use
	112: LLLVAR{782, EBCDIC, EncANS, DecANS}, // Additional Data (National Use). Is reserved for national organizations to define data unique to specific networks or specific programs and services. DE 112 provides other supplemental data in a message when a specific ISO-designated data element is not available. It is a free-format, variable-length, alphanumeric data element used for information on transactions between customers.
	113: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard recommends that DE 113 contain Application Generic Data and Application Banking Data. This data element is typically present for consumer and business application requests, counteroffer replies, and pre-approved offer inquiries.
	114: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard recommends that DE 114 contain Consumer Application Data or Consumer Maintenance Data. This data element is typically present for consumer application requests, application status inquiries, preapproved offer inquiries, or consumer maintenance requests as well as consumer application or consumer maintenance responses. DE 114 also may be present for business application requests that require a personal guarantee.
	115: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard recommends that DE 115 contain Business Application Data or Business Maintenance Data. This data element is typically present for business application requests, application status inquiries, preapproved offer inquiries, or business maintenance requests as well as business application or business maintenance responses.
	116: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard recommends that DE 116 contain Consumer User Lookup Data and Consumer Account Lookup Data. This data element is typically present to request consumer user and account information and provide consumer user account information.
	117: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard recommends that DE 117 contain Business User Lookup Data and Business Account Lookup Data. This data element is typically present to request business user and account information and provide business user account information.
	118: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard recommends that DE 118 contain Authorized Users. This data element may be present for consumer and business application requests and lookup responses.
	119: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Mastercard is reserving DE 119 for customer-specific data and is not recommending any particular usage.
	120: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Record Data. Is a variable-length data element used for transmitting file record data or textual character string data in various message types.
	121: LLLVAR{6, EBCDIC, EncN, DecN},       // Authorizing Agent ID Code. When used, must contain the appropriate Mastercard- assigned customer ID number that uniquely identifies the Authorization Platform Stand-In processing facility or alternate routing CPS responsible for performing Stand-In processing on- behalf of the issuer.
	122: LLLVAR{999, EBCDIC, EncANS, DecANS}, // Is a free-format, variable-length data element used for transmitting file record data in various message types. When used in Issuer File Update Request Response/0312 messages, this data element contains additional record data for file inquiry requests.
	123: LLLVAR{512, EBCDIC, EncANS, DecANS}, // Receipt Free Text. Only applies to the Swedish Domestic Authorization Switching Service (SASS), Peru, and the Mastercard Installment Payment Service. For SASS, DE 123 contains information to be printed on a receipt (not displayed on the terminal screen) for balance inquiry and ATM transactions (where DE 3 [Processing Code] is value 01 [Withdrawal] or value 30 [Balance Inquiry]). For Peru and Mastercard Installment Payment Service, DE 123 contains a text message to be printed on point-of-sale (POS) sales receipts.
	124: LLLVAR{299, EBCDIC, EncANS, DecANS}, // Member-defined Data—General Use. May be used to submit up to 299 bytes of customer-defined data. DE 124 can contain program-specific data as defined by the DE 124 subelements.
	125: FIX{8, EBCDIC, EncB, DecB},          // New PIN Data. Consists of a binary block containing a derived encrypted value calculated from the new PIN introduced by the cardholder at the ATM offering the PIN change service.
	126: LLLVAR{100, EBCDIC, EncANS, DecANS}, // Private Data. Is reserved for future use.
	127: LLLVAR{100, EBCDIC, EncANS, DecANS}, // Private Data. May contain any private-use data that the customer may want to include in a message. Any Authorization Platform message originator may use DE 127.
	128: FIX{8, EBCDIC, EncB, DecB},          // Message authentication code (MAC). Validates the source and the text of the message between the sender and the receiver.
}

// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 with prefix 'ac'
package camt_035_001_05

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 FinInstnId"`
}

type Case5 struct {
	Id    Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Id"`
	Cretr Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Cretr"`
}

type Case5TCH struct {
	Id    Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Id"`
	Cretr Party40ChoiceTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Cretr"`
}

type CaseAssignment5 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Id"`
	Assgnr  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgnr"`
	Assgne  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgne"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 CreDtTm"`
}

type CaseAssignment5TCH struct {
	Id      Max35TextTCH     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Id"`
	Assgnr  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgnr"`
	Assgne  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgne"`
	CreDtTm rtp.ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 CreDtTm"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 MmbId"`
}

type DocumentTCH struct {
	XMLName           xml.Name
	PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryFrmtInvstgtn"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 ClrSysMmbId"`
}

type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Agt,omitempty"`
}

type Party40ChoiceTCH struct {
	Agt *BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Agt,omitempty"`
}

type Party40ChoiceTCH2 struct {
	Pty *PartyIdentification135TCH                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Agt,omitempty"`
}

type PartyIdentification135 struct {
	Nm Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Nm"`
}

type PartyIdentification135TCH struct {
	Nm Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Nm"`
}

type ProprietaryData6Reduced struct {
	Ustrd    *Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Ustrd,omitempty"`
	OrigRefs TransactionReferences8Reduced `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 OrigRefs"`
}

type ProprietaryData6ReducedTCH struct {
	Ustrd    *Max140Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Ustrd,omitempty"`
	OrigRefs TransactionReferences8ReducedTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 OrigRefs"`
}

type ProprietaryData7TCH struct {
	Tp   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Tp"`
	Data ProprietaryData6Reduced `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Data"`
}

type ProprietaryData7TCHTCH struct {
	Tp   Max35TextTCH3              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Tp"`
	Data ProprietaryData6ReducedTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Data"`
}

type ProprietaryFormatInvestigationV05 struct {
	Assgnmt   CaseAssignment5     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgnmt"`
	Case      Case5               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Case"`
	PrtryData ProprietaryData7TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryData"`
}

type ProprietaryFormatInvestigationV05TCH struct {
	Assgnmt   CaseAssignment5TCH     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgnmt"`
	Case      Case5TCH               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Case"`
	PrtryData ProprietaryData7TCHTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryData"`
}

type TransactionReferences8Reduced struct {
	InstrId    Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 InstrId"`
	EndToEndId Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 EndToEndId"`
	TxId       Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 TxId"`
	UETR       *UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 UETR,omitempty"`
}

type TransactionReferences8ReducedTCH struct {
	InstrId    Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 InstrId"`
	EndToEndId Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 EndToEndId"`
	TxId       Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 TxId"`
	UETR       *UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 UETR,omitempty"`
}

// XSD SimpleType declarations

type Max140Text string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type Max35TextTCH3 string

const Max35TextTCH3Ack Max35TextTCH3 = "ACK"
const Max35TextTCH3Acwp Max35TextTCH3 = "ACWP"

type UUIDv4Identifier string
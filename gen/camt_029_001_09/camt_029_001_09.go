// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 with prefix 'tr'
package camt_029_001_09

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 FinInstnId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "tr:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 FinInstnId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "tr:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CancellationStatusReason3Choice struct {
	Cd *ExternalPaymentCancellationRejection1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cd,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CancellationStatusReason3Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "tr:Cd"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CancellationStatusReason3ChoiceTCH struct {
	Cd *ExternalPaymentCancellationRejection1CodeTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cd,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CancellationStatusReason3ChoiceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "tr:Cd"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CancellationStatusReason4 struct {
	Rsn CancellationStatusReason3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Rsn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CancellationStatusReason4) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "tr:Rsn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CancellationStatusReason4TCH struct {
	Rsn CancellationStatusReason3ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Rsn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CancellationStatusReason4TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "tr:Rsn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Case5 struct {
	Id    Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	Cretr Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cretr"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Case5) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "tr:Id"}})
	e.EncodeElement(v.Cretr, xml.StartElement{Name: xml.Name{Local: "tr:Cretr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Case5TCH struct {
	Id    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	Cretr Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Cretr"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Case5TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "tr:Id"}})
	e.EncodeElement(v.Cretr, xml.StartElement{Name: xml.Name{Local: "tr:Cretr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CaseAssignment5 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	Assgnr  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnr"`
	Assgne  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgne"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CreDtTm"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CaseAssignment5) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "tr:Id"}})
	e.EncodeElement(v.Assgnr, xml.StartElement{Name: xml.Name{Local: "tr:Assgnr"}})
	e.EncodeElement(v.Assgne, xml.StartElement{Name: xml.Name{Local: "tr:Assgne"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "tr:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CaseAssignment5TCH struct {
	Id      Max35TextTCH     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Id"`
	Assgnr  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnr"`
	Assgne  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgne"`
	CreDtTm rtp.ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CreDtTm"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CaseAssignment5TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "tr:Id"}})
	e.EncodeElement(v.Assgnr, xml.StartElement{Name: xml.Name{Local: "tr:Assgnr"}})
	e.EncodeElement(v.Assgne, xml.StartElement{Name: xml.Name{Local: "tr:Assgne"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "tr:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Charges7) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Amt, xml.StartElement{Name: xml.Name{Local: "tr:Amt"}})
	e.EncodeElement(v.Agt, xml.StartElement{Name: xml.Name{Local: "tr:Agt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Charges7TCH struct {
	Amt ActiveOrHistoricCurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Amt"`
	Agt BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Charges7TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Amt, xml.StartElement{Name: xml.Name{Local: "tr:Amt"}})
	e.EncodeElement(v.Agt, xml.StartElement{Name: xml.Name{Local: "tr:Agt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "tr:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "tr:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Document struct {
	XMLName         xml.Name
	RsltnOfInvstgtn ResolutionOfInvestigationV09 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnOfInvstgtn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RsltnOfInvstgtn, xml.StartElement{Name: xml.Name{Local: "tr:RsltnOfInvstgtn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type DocumentTCH struct {
	XMLName         xml.Name
	RsltnOfInvstgtn ResolutionOfInvestigationV09TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnOfInvstgtn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RsltnOfInvstgtn, xml.StartElement{Name: xml.Name{Local: "tr:RsltnOfInvstgtn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "tr:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "tr:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type InvestigationStatus5Choice struct {
	Conf *ExternalInvestigationExecutionConfirmation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Conf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v InvestigationStatus5Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Conf, xml.StartElement{Name: xml.Name{Local: "tr:Conf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type OriginalGroupHeader14 struct {
	RslvdCase    Case5                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RslvdCase"`
	OrgnlMsgId   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlCreDtTm"`
	CxlStsRsnInf *CancellationStatusReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlStsRsnInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupHeader14) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RslvdCase, xml.StartElement{Name: xml.Name{Local: "tr:RslvdCase"}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlCreDtTm"}})
	e.EncodeElement(v.CxlStsRsnInf, xml.StartElement{Name: xml.Name{Local: "tr:CxlStsRsnInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type OriginalGroupHeader14TCH struct {
	RslvdCase    Case5TCH                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RslvdCase"`
	OrgnlMsgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlCreDtTm"`
	CxlStsRsnInf *CancellationStatusReason4TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlStsRsnInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupHeader14TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RslvdCase, xml.StartElement{Name: xml.Name{Local: "tr:RslvdCase"}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlCreDtTm"}})
	e.EncodeElement(v.CxlStsRsnInf, xml.StartElement{Name: xml.Name{Local: "tr:CxlStsRsnInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Party40Choice struct {
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party40Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Agt, xml.StartElement{Name: xml.Name{Local: "tr:Agt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Party40ChoiceTCH struct {
	Agt *BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Agt,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party40ChoiceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Agt, xml.StartElement{Name: xml.Name{Local: "tr:Agt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PaymentTransaction102 struct {
	CxlStsId     *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlStsId,omitempty"`
	OrgnlUETR    *UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlUETR,omitempty"`
	RsltnRltdInf *ResolutionData1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnRltdInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.CxlStsId, xml.StartElement{Name: xml.Name{Local: "tr:CxlStsId"}})
	e.EncodeElement(v.OrgnlUETR, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlUETR"}})
	e.EncodeElement(v.RsltnRltdInf, xml.StartElement{Name: xml.Name{Local: "tr:RsltnRltdInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PaymentTransaction102TCH struct {
	CxlStsId     *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlStsId,omitempty"`
	OrgnlUETR    *UUIDv4Identifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlUETR,omitempty"`
	RsltnRltdInf *ResolutionData1TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnRltdInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction102TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.CxlStsId, xml.StartElement{Name: xml.Name{Local: "tr:CxlStsId"}})
	e.EncodeElement(v.OrgnlUETR, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlUETR"}})
	e.EncodeElement(v.RsltnRltdInf, xml.StartElement{Name: xml.Name{Local: "tr:RsltnRltdInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ResolutionData1 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 IntrBkSttlmAmt"`
	IntrBkSttlmDt  *rtp.ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 IntrBkSttlmDt,omitempty"`
	ClrChanl       *ClearingChannel2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrChanl,omitempty"`
	Chrgs          *Charges7                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Chrgs,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ResolutionData1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "tr:IntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "tr:IntrBkSttlmDt"}})
	e.EncodeElement(v.ClrChanl, xml.StartElement{Name: xml.Name{Local: "tr:ClrChanl"}})
	e.EncodeElement(v.Chrgs, xml.StartElement{Name: xml.Name{Local: "tr:Chrgs"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ResolutionData1TCH struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 IntrBkSttlmAmt"`
	IntrBkSttlmDt  *rtp.ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 IntrBkSttlmDt,omitempty"`
	ClrChanl       *ClearingChannel2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClrChanl,omitempty"`
	Chrgs          *Charges7TCH                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Chrgs,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ResolutionData1TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "tr:IntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "tr:IntrBkSttlmDt"}})
	e.EncodeElement(v.ClrChanl, xml.StartElement{Name: xml.Name{Local: "tr:ClrChanl"}})
	e.EncodeElement(v.Chrgs, xml.StartElement{Name: xml.Name{Local: "tr:Chrgs"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ResolutionOfInvestigationV09 struct {
	Assgnmt CaseAssignment5            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnmt"`
	Sts     InvestigationStatus5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Sts"`
	CxlDtls UnderlyingTransaction22    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlDtls"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ResolutionOfInvestigationV09) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Assgnmt, xml.StartElement{Name: xml.Name{Local: "tr:Assgnmt"}})
	e.EncodeElement(v.Sts, xml.StartElement{Name: xml.Name{Local: "tr:Sts"}})
	e.EncodeElement(v.CxlDtls, xml.StartElement{Name: xml.Name{Local: "tr:CxlDtls"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ResolutionOfInvestigationV09TCH struct {
	Assgnmt CaseAssignment5TCH         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnmt"`
	Sts     InvestigationStatus5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Sts"`
	CxlDtls UnderlyingTransaction22TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlDtls"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ResolutionOfInvestigationV09TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Assgnmt, xml.StartElement{Name: xml.Name{Local: "tr:Assgnmt"}})
	e.EncodeElement(v.Sts, xml.StartElement{Name: xml.Name{Local: "tr:Sts"}})
	e.EncodeElement(v.CxlDtls, xml.StartElement{Name: xml.Name{Local: "tr:CxlDtls"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type UnderlyingTransaction22 struct {
	OrgnlGrpInfAndSts OriginalGroupHeader14  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlGrpInfAndSts"`
	TxInfAndSts       *PaymentTransaction102 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 TxInfAndSts,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v UnderlyingTransaction22) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlGrpInfAndSts, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlGrpInfAndSts"}})
	e.EncodeElement(v.TxInfAndSts, xml.StartElement{Name: xml.Name{Local: "tr:TxInfAndSts"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type UnderlyingTransaction22TCH struct {
	OrgnlGrpInfAndSts OriginalGroupHeader14TCH  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 OrgnlGrpInfAndSts"`
	TxInfAndSts       *PaymentTransaction102TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 TxInfAndSts,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v UnderlyingTransaction22TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlGrpInfAndSts, xml.StartElement{Name: xml.Name{Local: "tr:OrgnlGrpInfAndSts"}})
	e.EncodeElement(v.TxInfAndSts, xml.StartElement{Name: xml.Name{Local: "tr:TxInfAndSts"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType rtp.Amount

func (a ActiveOrHistoricCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return rtp.Amount(a).MarshalText()
}

type ActiveOrHistoricCurrencyCode string

const ActiveOrHistoricCurrencyCodeUsd ActiveOrHistoricCurrencyCode = "USD"

type ClearingChannel2Code string

const ClearingChannel2CodeMpns ClearingChannel2Code = "MPNS"
const ClearingChannel2CodeRtgs ClearingChannel2Code = "RTGS"
const ClearingChannel2CodeRtns ClearingChannel2Code = "RTNS"

type ExternalInvestigationExecutionConfirmation1Code string

const ExternalInvestigationExecutionConfirmation1CodeIpay ExternalInvestigationExecutionConfirmation1Code = "IPAY"
const ExternalInvestigationExecutionConfirmation1CodePecr ExternalInvestigationExecutionConfirmation1Code = "PECR"
const ExternalInvestigationExecutionConfirmation1CodeRjcr ExternalInvestigationExecutionConfirmation1Code = "RJCR"

type ExternalPaymentCancellationRejection1Code string

const ExternalPaymentCancellationRejection1CodeAc04 ExternalPaymentCancellationRejection1Code = "AC04"
const ExternalPaymentCancellationRejection1CodeAm04 ExternalPaymentCancellationRejection1Code = "AM04"
const ExternalPaymentCancellationRejection1CodeArdt ExternalPaymentCancellationRejection1Code = "ARDT"
const ExternalPaymentCancellationRejection1CodeCust ExternalPaymentCancellationRejection1Code = "CUST"
const ExternalPaymentCancellationRejection1CodeLegl ExternalPaymentCancellationRejection1Code = "LEGL"
const ExternalPaymentCancellationRejection1CodeNoas ExternalPaymentCancellationRejection1Code = "NOAS"
const ExternalPaymentCancellationRejection1CodeNoor ExternalPaymentCancellationRejection1Code = "NOOR"

type ExternalPaymentCancellationRejection1CodeTCH string

const ExternalPaymentCancellationRejection1CodeTCHAc04 ExternalPaymentCancellationRejection1CodeTCH = "AC04"
const ExternalPaymentCancellationRejection1CodeTCHAm04 ExternalPaymentCancellationRejection1CodeTCH = "AM04"
const ExternalPaymentCancellationRejection1CodeTCHArdt ExternalPaymentCancellationRejection1CodeTCH = "ARDT"
const ExternalPaymentCancellationRejection1CodeTCHCust ExternalPaymentCancellationRejection1CodeTCH = "CUST"
const ExternalPaymentCancellationRejection1CodeTCHLegl ExternalPaymentCancellationRejection1CodeTCH = "LEGL"
const ExternalPaymentCancellationRejection1CodeTCHNoas ExternalPaymentCancellationRejection1CodeTCH = "NOAS"
const ExternalPaymentCancellationRejection1CodeTCHNoor ExternalPaymentCancellationRejection1CodeTCH = "NOOR"

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type OrigMsgName string

const OrigMsgNameCamt05600105 OrigMsgName = "camt.056.001.05"
const OrigMsgNameCamt05600108 OrigMsgName = "camt.056.001.08"

type UUIDv4Identifier string

// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 with prefix 'c9'
package pacs_009_001_08

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Othr,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v AccountIdentification4Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Othr, xml.StartElement{Name: xml.Name{Local: "c9:Othr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ActiveCurrencyAndAmount struct {
	Value ActiveCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCode                `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 FinInstnId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "c9:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 FinInstnId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "c9:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CashAccount38 struct {
	Id AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Id"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CashAccount38) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "c9:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemIdentification3Choice struct {
	Cd *ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cd,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification3Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "c9:Cd"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "c9:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH3 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "c9:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CreditTransferTransaction36 struct {
	PmtId          PaymentIdentification7                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtId"`
	PmtTpInf       PaymentTypeInformation28                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtTpInf"`
	IntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmAmt"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstdAgt"`
	Dbtr           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Dbtr"`
	DbtrAcct       *CashAccount38                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 DbtrAcct,omitempty"`
	Cdtr           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cdtr"`
	CdtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtrAcct"`
	RmtInf         *RemittanceInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 RmtInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditTransferTransaction36) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.PmtId, xml.StartElement{Name: xml.Name{Local: "c9:PmtId"}})
	e.EncodeElement(v.PmtTpInf, xml.StartElement{Name: xml.Name{Local: "c9:PmtTpInf"}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "c9:IntrBkSttlmAmt"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "c9:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "c9:InstdAgt"}})
	e.EncodeElement(v.Dbtr, xml.StartElement{Name: xml.Name{Local: "c9:Dbtr"}})
	e.EncodeElement(v.DbtrAcct, xml.StartElement{Name: xml.Name{Local: "c9:DbtrAcct"}})
	e.EncodeElement(v.Cdtr, xml.StartElement{Name: xml.Name{Local: "c9:Cdtr"}})
	e.EncodeElement(v.CdtrAcct, xml.StartElement{Name: xml.Name{Local: "c9:CdtrAcct"}})
	e.EncodeElement(v.RmtInf, xml.StartElement{Name: xml.Name{Local: "c9:RmtInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CreditTransferTransaction36TCH struct {
	PmtId          PaymentIdentification7TCH                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtId"`
	PmtTpInf       PaymentTypeInformation28TCH                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtTpInf"`
	IntrBkSttlmAmt ActiveCurrencyAndAmount                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmAmt"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstdAgt"`
	Dbtr           BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Dbtr"`
	DbtrAcct       *CashAccount38                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 DbtrAcct,omitempty"`
	Cdtr           BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cdtr"`
	CdtrAcct       CashAccount38                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtrAcct"`
	RmtInf         *RemittanceInformation2TCH                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 RmtInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditTransferTransaction36TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.PmtId, xml.StartElement{Name: xml.Name{Local: "c9:PmtId"}})
	e.EncodeElement(v.PmtTpInf, xml.StartElement{Name: xml.Name{Local: "c9:PmtTpInf"}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "c9:IntrBkSttlmAmt"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "c9:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "c9:InstdAgt"}})
	e.EncodeElement(v.Dbtr, xml.StartElement{Name: xml.Name{Local: "c9:Dbtr"}})
	e.EncodeElement(v.DbtrAcct, xml.StartElement{Name: xml.Name{Local: "c9:DbtrAcct"}})
	e.EncodeElement(v.Cdtr, xml.StartElement{Name: xml.Name{Local: "c9:Cdtr"}})
	e.EncodeElement(v.CdtrAcct, xml.StartElement{Name: xml.Name{Local: "c9:CdtrAcct"}})
	e.EncodeElement(v.RmtInf, xml.StartElement{Name: xml.Name{Local: "c9:RmtInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type DocumentTCH struct {
	XMLName  xml.Name
	FICdtTrf FinancialInstitutionCreditTransferV08TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 FICdtTrf"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FICdtTrf, xml.StartElement{Name: xml.Name{Local: "c9:FICdtTrf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionCreditTransferV08 struct {
	GrpHdr      GroupHeader93               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 GrpHdr"`
	CdtTrfTxInf CreditTransferTransaction36 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtTrfTxInf"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionCreditTransferV08) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "c9:GrpHdr"}})
	e.EncodeElement(v.CdtTrfTxInf, xml.StartElement{Name: xml.Name{Local: "c9:CdtTrfTxInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionCreditTransferV08TCH struct {
	GrpHdr      GroupHeader93TCH               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 GrpHdr"`
	CdtTrfTxInf CreditTransferTransaction36TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtTrfTxInf"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionCreditTransferV08TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "c9:GrpHdr"}})
	e.EncodeElement(v.CdtTrfTxInf, xml.StartElement{Name: xml.Name{Local: "c9:CdtTrfTxInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "c9:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "c9:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GenericAccountIdentification1 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Id"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GenericAccountIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "c9:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GroupHeader93 struct {
	MsgId             Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MsgId"`
	CreDtTm           rtp.ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CreDtTm"`
	NbOfTxs           Max1NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 NbOfTxs"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TtlIntrBkSttlmAmt"`
	IntrBkSttlmDt     rtp.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmDt"`
	SttlmInf          SettlementInstruction7  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmInf"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader93) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "c9:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "c9:CreDtTm"}})
	e.EncodeElement(v.NbOfTxs, xml.StartElement{Name: xml.Name{Local: "c9:NbOfTxs"}})
	e.EncodeElement(v.TtlIntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "c9:TtlIntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "c9:IntrBkSttlmDt"}})
	e.EncodeElement(v.SttlmInf, xml.StartElement{Name: xml.Name{Local: "c9:SttlmInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GroupHeader93TCH struct {
	MsgId             Max35TextTCH              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MsgId"`
	CreDtTm           rtp.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CreDtTm"`
	NbOfTxs           Max1NumericText           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 NbOfTxs"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TtlIntrBkSttlmAmt"`
	IntrBkSttlmDt     rtp.ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmDt"`
	SttlmInf          SettlementInstruction7TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmInf"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader93TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "c9:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "c9:CreDtTm"}})
	e.EncodeElement(v.NbOfTxs, xml.StartElement{Name: xml.Name{Local: "c9:NbOfTxs"}})
	e.EncodeElement(v.TtlIntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "c9:TtlIntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "c9:IntrBkSttlmDt"}})
	e.EncodeElement(v.SttlmInf, xml.StartElement{Name: xml.Name{Local: "c9:SttlmInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PaymentIdentification7 struct {
	InstrId    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstrId"`
	EndToEndId Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 EndToEndId"`
	TxId       Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TxId"`
	ClrSysRef  *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysRef,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentIdentification7) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.InstrId, xml.StartElement{Name: xml.Name{Local: "c9:InstrId"}})
	e.EncodeElement(v.EndToEndId, xml.StartElement{Name: xml.Name{Local: "c9:EndToEndId"}})
	e.EncodeElement(v.TxId, xml.StartElement{Name: xml.Name{Local: "c9:TxId"}})
	e.EncodeElement(v.ClrSysRef, xml.StartElement{Name: xml.Name{Local: "c9:ClrSysRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PaymentIdentification7TCH struct {
	InstrId    Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstrId"`
	EndToEndId Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 EndToEndId"`
	TxId       Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TxId"`
	ClrSysRef  *Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysRef,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentIdentification7TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.InstrId, xml.StartElement{Name: xml.Name{Local: "c9:InstrId"}})
	e.EncodeElement(v.EndToEndId, xml.StartElement{Name: xml.Name{Local: "c9:EndToEndId"}})
	e.EncodeElement(v.TxId, xml.StartElement{Name: xml.Name{Local: "c9:TxId"}})
	e.EncodeElement(v.ClrSysRef, xml.StartElement{Name: xml.Name{Local: "c9:ClrSysRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PaymentTypeInformation28 struct {
	SvcLvl ServiceLevel8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SvcLvl"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTypeInformation28) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.SvcLvl, xml.StartElement{Name: xml.Name{Local: "c9:SvcLvl"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PaymentTypeInformation28TCH struct {
	SvcLvl ServiceLevel8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SvcLvl"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTypeInformation28TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.SvcLvl, xml.StartElement{Name: xml.Name{Local: "c9:SvcLvl"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type RemittanceInformation2 struct {
	Ustrd *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Ustrd,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceInformation2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Ustrd, xml.StartElement{Name: xml.Name{Local: "c9:Ustrd"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type RemittanceInformation2TCH struct {
	Ustrd *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Ustrd,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceInformation2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Ustrd, xml.StartElement{Name: xml.Name{Local: "c9:Ustrd"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ServiceLevel8Choice struct {
	Cd *ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cd,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ServiceLevel8Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "c9:Cd"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SettlementInstruction7 struct {
	SttlmMtd SettlementMethod1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSys"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SettlementInstruction7) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.SttlmMtd, xml.StartElement{Name: xml.Name{Local: "c9:SttlmMtd"}})
	e.EncodeElement(v.ClrSys, xml.StartElement{Name: xml.Name{Local: "c9:ClrSys"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SettlementInstruction7TCH struct {
	SttlmMtd SettlementMethod1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSys"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SettlementInstruction7TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.SttlmMtd, xml.StartElement{Name: xml.Name{Local: "c9:SttlmMtd"}})
	e.EncodeElement(v.ClrSys, xml.StartElement{Name: xml.Name{Local: "c9:ClrSys"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type ActiveCurrencyAndAmountSimpleType rtp.Amount

func (a ActiveCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return rtp.Amount(a).MarshalText()
}

type ActiveCurrencyCode string

const ActiveCurrencyCodeUsd ActiveCurrencyCode = "USD"

type ExternalCashClearingSystem1Code string

const ExternalCashClearingSystem1CodeTch ExternalCashClearingSystem1Code = "TCH"

type ExternalServiceLevel1Code string

const ExternalServiceLevel1CodeSdva ExternalServiceLevel1Code = "SDVA"

type Max140Text string

type Max1NumericText string

type Max34Text string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type Max35TextTCH3 string

type SettlementMethod1Code string

const SettlementMethod1CodeClrg SettlementMethod1Code = "CLRG"

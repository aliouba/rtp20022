// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 with prefix 'rf'
package admn_004_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification4ADMN struct {
	FinInstnId FinancialInstitutionIdentification7ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 FinInstnId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification4ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "rf:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2ADMN struct {
	MmbId Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "rf:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Document struct {
	XMLName         xml.Name
	AdmnSignOffResp SignOffResponse `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 AdmnSignOffResp"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.AdmnSignOffResp, xml.StartElement{Name: xml.Name{Local: "rf:AdmnSignOffResp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type DocumentTCH struct {
	XMLName         xml.Name
	AdmnSignOffResp SignOffResponseTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 AdmnSignOffResp"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.AdmnSignOffResp, xml.StartElement{Name: xml.Name{Local: "rf:AdmnSignOffResp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification7ADMN struct {
	ClrSysMmbId ClearingSystemMemberIdentification2ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification7ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "rf:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GrpHdr struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 CreDtTm"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GrpHdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "rf:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "rf:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GrpHdrTCH struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 CreDtTm"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GrpHdrTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "rf:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "rf:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SignOffResp struct {
	OrgnlInstrId Max35Text                                        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 OrgnlInstrId"`
	InstgAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 InstdAgt"`
	Sts          TransactionGroupStatus3CodeAdmin                 `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 Sts"`
	StsRsnInf    *StatusReasonInformation8                        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 StsRsnInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignOffResp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "rf:OrgnlInstrId"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "rf:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "rf:InstdAgt"}})
	e.EncodeElement(v.Sts, xml.StartElement{Name: xml.Name{Local: "rf:Sts"}})
	e.EncodeElement(v.StsRsnInf, xml.StartElement{Name: xml.Name{Local: "rf:StsRsnInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SignOffResponse struct {
	GrpHdr      GrpHdr      `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 GrpHdr"`
	SignOffResp SignOffResp `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 SignOffResp"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignOffResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "rf:GrpHdr"}})
	e.EncodeElement(v.SignOffResp, xml.StartElement{Name: xml.Name{Local: "rf:SignOffResp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SignOffResponseTCH struct {
	GrpHdr      GrpHdrTCH      `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 GrpHdr"`
	SignOffResp SignOffRespTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 SignOffResp"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignOffResponseTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "rf:GrpHdr"}})
	e.EncodeElement(v.SignOffResp, xml.StartElement{Name: xml.Name{Local: "rf:SignOffResp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SignOffRespTCH struct {
	OrgnlInstrId Max35Text                                        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 OrgnlInstrId"`
	InstgAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 InstdAgt"`
	Sts          TransactionGroupStatus3CodeAdmin                 `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 Sts"`
	StsRsnInf    *StatusReasonInformation8TCH                     `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 StsRsnInf,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignOffRespTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "rf:OrgnlInstrId"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "rf:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "rf:InstdAgt"}})
	e.EncodeElement(v.Sts, xml.StartElement{Name: xml.Name{Local: "rf:Sts"}})
	e.EncodeElement(v.StsRsnInf, xml.StartElement{Name: xml.Name{Local: "rf:StsRsnInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type StatusReason6Choice struct {
	Prtry *ProprietaryReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 Prtry,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReason6Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Prtry, xml.StartElement{Name: xml.Name{Local: "rf:Prtry"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type StatusReasonInformation8 struct {
	Rsn StatusReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 Rsn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReasonInformation8) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "rf:Rsn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type StatusReasonInformation8TCH struct {
	Rsn StatusReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01 Rsn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReasonInformation8TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "rf:Rsn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type Max35Text string

type Min11Max11Text string

type ProprietaryReasonCode string

const ProprietaryReasonCode9946 ProprietaryReasonCode = "9946"
const ProprietaryReasonCode9948 ProprietaryReasonCode = "9948"
const ProprietaryReasonCode9964 ProprietaryReasonCode = "9964"
const ProprietaryReasonCodeDs0H ProprietaryReasonCode = "DS0H"
const ProprietaryReasonCodeRc02 ProprietaryReasonCode = "RC02"

type TransactionGroupStatus3CodeAdmin string

const TransactionGroupStatus3CodeAdminActc TransactionGroupStatus3CodeAdmin = "ACTC"
const TransactionGroupStatus3CodeAdminRjct TransactionGroupStatus3CodeAdmin = "RJCT"

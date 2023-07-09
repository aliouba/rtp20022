// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:head.001.001.01 with prefix 'head'
package head_001_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/gen/xmldsig"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification5BAH struct {
	FinInstnId FinancialInstitutionIdentification8BAH `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 FinInstnId"`
	BrnchId    *BranchData2BAH                        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 BrnchId,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification5BAH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "head:FinInstnId"}})
	e.EncodeElement(v.BrnchId, xml.StartElement{Name: xml.Name{Local: "head:BrnchId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BranchAndFinancialInstitutionIdentification5BAHTCH struct {
	FinInstnId FinancialInstitutionIdentification8BAH `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 FinInstnId"`
	BrnchId    *BranchData2BAHTCH                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 BrnchId,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification5BAHTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "head:FinInstnId"}})
	e.EncodeElement(v.BrnchId, xml.StartElement{Name: xml.Name{Local: "head:BrnchId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BranchData2BAH struct {
	Id Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Id"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchData2BAH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "head:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BranchData2BAHTCH struct {
	Id Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Id"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchData2BAHTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "head:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BusinessApplicationHeaderV01 struct {
	Fr        Party9ChoiceBAH           `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Fr"`
	To        Party9ChoiceBAH           `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 To"`
	BizMsgIdr Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 BizMsgIdr"`
	MsgDefIdr OrigMsgName               `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 MsgDefIdr"`
	CreDt     rtp.ISONormalisedDateTime `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CreDt"`
	CpyDplct  *CopyDuplicate1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CpyDplct,omitempty"`
	Sgntr     *Sgntr                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Sgntr,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BusinessApplicationHeaderV01) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Fr, xml.StartElement{Name: xml.Name{Local: "head:Fr"}})
	e.EncodeElement(v.To, xml.StartElement{Name: xml.Name{Local: "head:To"}})
	e.EncodeElement(v.BizMsgIdr, xml.StartElement{Name: xml.Name{Local: "head:BizMsgIdr"}})
	e.EncodeElement(v.MsgDefIdr, xml.StartElement{Name: xml.Name{Local: "head:MsgDefIdr"}})
	e.EncodeElement(v.CreDt, xml.StartElement{Name: xml.Name{Local: "head:CreDt"}})
	e.EncodeElement(v.CpyDplct, xml.StartElement{Name: xml.Name{Local: "head:CpyDplct"}})
	e.EncodeElement(v.Sgntr, xml.StartElement{Name: xml.Name{Local: "head:Sgntr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type BusinessApplicationHeaderV01TCH struct {
	Fr        Party9ChoiceBAHTCH        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Fr"`
	To        Party9ChoiceBAHTCH        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 To"`
	BizMsgIdr Max35TextTCH              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 BizMsgIdr"`
	MsgDefIdr OrigMsgName               `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 MsgDefIdr"`
	CreDt     rtp.ISONormalisedDateTime `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CreDt"`
	CpyDplct  *CopyDuplicate1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CpyDplct,omitempty"`
	Sgntr     *Sgntr                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Sgntr,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BusinessApplicationHeaderV01TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Fr, xml.StartElement{Name: xml.Name{Local: "head:Fr"}})
	e.EncodeElement(v.To, xml.StartElement{Name: xml.Name{Local: "head:To"}})
	e.EncodeElement(v.BizMsgIdr, xml.StartElement{Name: xml.Name{Local: "head:BizMsgIdr"}})
	e.EncodeElement(v.MsgDefIdr, xml.StartElement{Name: xml.Name{Local: "head:MsgDefIdr"}})
	e.EncodeElement(v.CreDt, xml.StartElement{Name: xml.Name{Local: "head:CreDt"}})
	e.EncodeElement(v.CpyDplct, xml.StartElement{Name: xml.Name{Local: "head:CpyDplct"}})
	e.EncodeElement(v.Sgntr, xml.StartElement{Name: xml.Name{Local: "head:Sgntr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2ADMN struct {
	MmbId Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "head:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification8BAH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification8BAH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "head:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Party9ChoiceBAH struct {
	FIId *BranchAndFinancialInstitutionIdentification5BAH `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 FIId,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party9ChoiceBAH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FIId, xml.StartElement{Name: xml.Name{Local: "head:FIId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Party9ChoiceBAHTCH struct {
	FIId *BranchAndFinancialInstitutionIdentification5BAHTCH `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 FIId,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party9ChoiceBAHTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FIId, xml.StartElement{Name: xml.Name{Local: "head:FIId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type Sgntr struct {
	Signature *xmldsig.Signature
}

// XSD SimpleType declarations

type CopyDuplicate1Code string

const CopyDuplicate1CodeCopy CopyDuplicate1Code = "COPY"
const CopyDuplicate1CodeDupl CopyDuplicate1Code = "DUPL"

type Max35Text string

type Max35TextTCH string

type Min11Max11Text string

type OrigMsgName string

const OrigMsgNameAcmt02200102 OrigMsgName = "acmt.022.001.02"
const OrigMsgNameAdmi00200101 OrigMsgName = "admi.002.001.01"
const OrigMsgNameAdmi00400102 OrigMsgName = "admi.004.001.02"
const OrigMsgNameAdmn00100101 OrigMsgName = "admn.001.001.01"
const OrigMsgNameAdmn00200101 OrigMsgName = "admn.002.001.01"
const OrigMsgNameAdmn00300101 OrigMsgName = "admn.003.001.01"
const OrigMsgNameAdmn00400101 OrigMsgName = "admn.004.001.01"
const OrigMsgNameAdmn00500101 OrigMsgName = "admn.005.001.01"
const OrigMsgNameAdmn00600101 OrigMsgName = "admn.006.001.01"
const OrigMsgNameAdmn00700101 OrigMsgName = "admn.007.001.01"
const OrigMsgNameAdmn00800101 OrigMsgName = "admn.008.001.01"
const OrigMsgNameCamt02600107 OrigMsgName = "camt.026.001.07"
const OrigMsgNameCamt02800109 OrigMsgName = "camt.028.001.09"
const OrigMsgNameCamt02900109 OrigMsgName = "camt.029.001.09"
const OrigMsgNameCamt03500105 OrigMsgName = "camt.035.001.05"
const OrigMsgNameCamt05600108 OrigMsgName = "camt.056.001.08"
const OrigMsgNamePacs00200110 OrigMsgName = "pacs.002.001.10"
const OrigMsgNamePacs00800108 OrigMsgName = "pacs.008.001.08"
const OrigMsgNamePacs00900108 OrigMsgName = "pacs.009.001.08"
const OrigMsgNamePacs02800103 OrigMsgName = "pacs.028.001.03"
const OrigMsgNamePain01300107 OrigMsgName = "pain.013.001.07"
const OrigMsgNamePain01400107 OrigMsgName = "pain.014.001.07"
const OrigMsgNameRemt00100104 OrigMsgName = "remt.001.001.04"

// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 with prefix 'mr'
package admi_002_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type DocumentTCH struct {
	XMLName      xml.Name
	Admi00200101 MessageRejectV01TCH `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 admi.002.001.01"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Admi00200101, xml.StartElement{Name: xml.Name{Local: "mr:admi.002.001.01"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type MessageReference struct {
	Ref Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 Ref"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v MessageReference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Ref, xml.StartElement{Name: xml.Name{Local: "mr:Ref"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type MessageReferenceTCH struct {
	Ref Max35TextTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 Ref"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v MessageReferenceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Ref, xml.StartElement{Name: xml.Name{Local: "mr:Ref"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type MessageRejectV01 struct {
	RltdRef MessageReference `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 RltdRef"`
	Rsn     RejectionReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 Rsn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v MessageRejectV01) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RltdRef, xml.StartElement{Name: xml.Name{Local: "mr:RltdRef"}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "mr:Rsn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type MessageRejectV01TCH struct {
	RltdRef MessageReferenceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 RltdRef"`
	Rsn     RejectionReason2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 Rsn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v MessageRejectV01TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RltdRef, xml.StartElement{Name: xml.Name{Local: "mr:RltdRef"}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "mr:Rsn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type RejectionReason2 struct {
	RjctgPtyRsn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 RjctgPtyRsn"`
	AddtlData   *rtp.Cdata `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 AddtlData,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RejectionReason2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RjctgPtyRsn, xml.StartElement{Name: xml.Name{Local: "mr:RjctgPtyRsn"}})
	e.EncodeElement(v.AddtlData, xml.StartElement{Name: xml.Name{Local: "mr:AddtlData"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type RejectionReason2TCH struct {
	RjctgPtyRsn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 RjctgPtyRsn"`
	AddtlData   *rtp.Cdata `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 AddtlData,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RejectionReason2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RjctgPtyRsn, xml.StartElement{Name: xml.Name{Local: "mr:RjctgPtyRsn"}})
	e.EncodeElement(v.AddtlData, xml.StartElement{Name: xml.Name{Local: "mr:AddtlData"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type Max35Text string

type Max35TextTCH string

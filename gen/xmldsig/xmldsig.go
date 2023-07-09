// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for http://www.w3.org/2000/09/xmldsig# with prefix 'xml_dsig'
package xmldsig

import (
	"encoding/xml"
)

// XSD Elements

type Signature struct {
	XMLName        xml.Name           `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	Id             *string            `xml:"Id,attr,omitempty"`
	SignedInfo     SignedInfoType     `xml:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"SignatureValue"`
	KeyInfo        *KeyInfoType       `xml:"KeyInfo,omitempty"`
	Object         []*ObjectType      `xml:"Object,omitempty"`
}

// XSD ComplexType declarations

type SignatureValueType struct {
	Id   *string `xml:"Id,attr,omitempty"`
	Data string  `xml:",chardata"`
}

type SignedInfoType struct {
	Id                     *string                    `xml:"Id,attr,omitempty"`
	CanonicalizationMethod CanonicalizationMethodType `xml:""`
	SignatureMethod        SignatureMethodType        `xml:""`
	Reference              []ReferenceType            `xml:""`
}

type CanonicalizationMethodType struct {
	Algorithm string  `xml:"Algorithm,attr"`
	Item      *string `xml:",any,omitempty"`
}

type SignatureMethodType struct {
	Algorithm        string                `xml:"Algorithm,attr"`
	HMACOutputLength *HMACOutputLengthType `xml:"HMACOutputLength,omitempty"`
	Item             *string               `xml:",any,omitempty"`
}

type ReferenceType struct {
	Id           *string          `xml:"Id,attr,omitempty"`
	URI          *string          `xml:"URI,attr,omitempty"`
	Type         *string          `xml:"Type,attr,omitempty"`
	Transforms   *TransformsType  `xml:",omitempty"`
	DigestMethod DigestMethodType `xml:""`
	DigestValue  DigestValueType  `xml:""`
}

type TransformsType struct {
	Transform []TransformType `xml:""`
}

type TransformType struct {
	Algorithm string    `xml:"Algorithm,attr"`
	XPath     []*string `xml:"XPath,omitempty"`
}

type DigestMethodType struct {
	Algorithm string  `xml:"Algorithm,attr"`
	Item      *string `xml:",any,omitempty"`
}

type KeyInfoType struct {
	Id              *string                `xml:"Id,attr,omitempty"`
	KeyName         []*string              `xml:",omitempty"`
	KeyValue        []*KeyValueType        `xml:",omitempty"`
	RetrievalMethod []*RetrievalMethodType `xml:",omitempty"`
	X509Data        []*X509DataType        `xml:",omitempty"`
	PGPData         []*PGPDataType         `xml:",omitempty"`
	SPKIData        []*SPKIDataType        `xml:",omitempty"`
	MgmtData        []*string              `xml:",omitempty"`
}

type KeyValueType struct {
	DSAKeyValue *DSAKeyValueType `xml:",omitempty"`
	RSAKeyValue *RSAKeyValueType `xml:",omitempty"`
}

type RetrievalMethodType struct {
	URI        string          `xml:"URI,attr"`
	Type       *string         `xml:"Type,attr,omitempty"`
	Transforms *TransformsType `xml:",omitempty"`
}

type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial,omitempty"`
	X509SKI          *string               `xml:"X509SKI,omitempty"`
	X509SubjectName  *string               `xml:"X509SubjectName,omitempty"`
	X509Certificate  *string               `xml:"X509Certificate,omitempty"`
	X509CRL          *string               `xml:"X509CRL,omitempty"`
}

type X509IssuerSerialType struct {
	X509IssuerName   string `xml:"X509IssuerName"`
	X509SerialNumber string `xml:"X509SerialNumber"`
}

type PGPDataType struct {
	PGPKeyID     *string `xml:"PGPKeyID,omitempty"`
	PGPKeyPacket *string `xml:"PGPKeyPacket,omitempty"`
}

type SPKIDataType struct {
	SPKISexp string  `xml:"SPKISexp"`
	Item     *string `xml:",any,omitempty"`
}

type ObjectType struct {
	Id       *string `xml:"Id,attr,omitempty"`
	MimeType *string `xml:"MimeType,attr,omitempty"`
	Encoding *string `xml:"Encoding,attr,omitempty"`
	Item     string  `xml:",any"`
}

type ManifestType struct {
	Id        *string         `xml:"Id,attr,omitempty"`
	Reference []ReferenceType `xml:""`
}

type SignaturePropertiesType struct {
	Id                *string                 `xml:"Id,attr,omitempty"`
	SignatureProperty []SignaturePropertyType `xml:""`
}

type SignaturePropertyType struct {
	Target string  `xml:"Target,attr"`
	Id     *string `xml:"Id,attr,omitempty"`
}

type DSAKeyValueType struct {
	G *CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary  `xml:"Y"`
	J *CryptoBinary `xml:"J,omitempty"`
}

type RSAKeyValueType struct {
	Modulus  CryptoBinary `xml:"Modulus"`
	Exponent CryptoBinary `xml:"Exponent"`
}

// XSD SimpleType declarations

type CryptoBinary string

type DigestValueType string

type HMACOutputLengthType int64

// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:tch
package messages

import (
	"encoding/xml"
	"fmt"

	"github.com/moov-io/rtp20022/gen/acmt_022_001_02"
	"github.com/moov-io/rtp20022/gen/admi_002_001_01"
	"github.com/moov-io/rtp20022/gen/admi_004_001_02"
	"github.com/moov-io/rtp20022/gen/admn_001_001_01"
	"github.com/moov-io/rtp20022/gen/admn_002_001_01"
	"github.com/moov-io/rtp20022/gen/admn_003_001_01"
	"github.com/moov-io/rtp20022/gen/admn_004_001_01"
	"github.com/moov-io/rtp20022/gen/admn_005_001_01"
	"github.com/moov-io/rtp20022/gen/admn_006_001_01"
	"github.com/moov-io/rtp20022/gen/admn_007_001_01"
	"github.com/moov-io/rtp20022/gen/admn_008_001_01"
	"github.com/moov-io/rtp20022/gen/camt_026_001_07"
	"github.com/moov-io/rtp20022/gen/camt_028_001_09"
	"github.com/moov-io/rtp20022/gen/camt_029_001_09"
	"github.com/moov-io/rtp20022/gen/camt_035_001_05"
	"github.com/moov-io/rtp20022/gen/camt_056_001_08"
	"github.com/moov-io/rtp20022/gen/head_001_001_01"
	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/gen/pacs_009_001_08"
	"github.com/moov-io/rtp20022/gen/pacs_028_001_03"
	"github.com/moov-io/rtp20022/gen/pain_013_001_07"
	"github.com/moov-io/rtp20022/gen/pain_014_001_07"
	"github.com/moov-io/rtp20022/gen/remt_001_001_04"
)

// XSD Elements

type Message struct {
	XMLName                       xml.Name                                        `xml:"Message"`
	Xmlns                         []xml.Attr                                      `xml:",attr"`
	AppHdr                        head_001_001_01.BusinessApplicationHeaderV01TCH `xml:"urn:tch AppHdr"`
	CreditTransfer                *pacs_008_001_08.DocumentTCH                    `xml:"urn:tch CreditTransfer,omitempty"`
	MessageStatusReport           *pacs_002_001_10.DocumentTCH                    `xml:"urn:tch MessageStatusReport,omitempty"`
	FICreditTransfer              *pacs_009_001_08.DocumentTCH                    `xml:"urn:tch FICreditTransfer,omitempty"`
	Acknowledgement               *camt_035_001_05.DocumentTCH                    `xml:"urn:tch Acknowledgement,omitempty"`
	ResponseRequestForInformation *camt_028_001_09.DocumentTCH                    `xml:"urn:tch ResponseRequestForInformation,omitempty"`
	RequestForInformation         *camt_026_001_07.DocumentTCH                    `xml:"urn:tch RequestForInformation,omitempty"`
	ReturnOfFunds                 *camt_056_001_08.DocumentTCH                    `xml:"urn:tch ReturnOfFunds,omitempty"`
	PaymentRequest                *pain_013_001_07.DocumentTCH                    `xml:"urn:tch PaymentRequest,omitempty"`
	ResponsePaymentRequest        *pain_014_001_07.DocumentTCH                    `xml:"urn:tch ResponsePaymentRequest,omitempty"`
	ResponseReturnOfFunds         *camt_029_001_09.DocumentTCH                    `xml:"urn:tch ResponseReturnOfFunds,omitempty"`
	EchoRequest                   *admn_005_001_01.DocumentTCH                    `xml:"urn:tch EchoRequest,omitempty"`
	EchoResponse                  *admn_006_001_01.DocumentTCH                    `xml:"urn:tch EchoResponse,omitempty"`
	SignOffRequest                *admn_003_001_01.DocumentTCH                    `xml:"urn:tch SignOffRequest,omitempty"`
	SignOffResponse               *admn_004_001_01.DocumentTCH                    `xml:"urn:tch SignOffResponse,omitempty"`
	SignOnRequest                 *admn_001_001_01.DocumentTCH                    `xml:"urn:tch SignOnRequest,omitempty"`
	SignOnResponse                *admn_002_001_01.DocumentTCH                    `xml:"urn:tch SignOnResponse,omitempty"`
	StandaloneRemittance          *remt_001_001_04.DocumentTCH                    `xml:"urn:tch StandaloneRemittance,omitempty"`
	SystemNotificationEvent       *admi_004_001_02.DocumentTCH                    `xml:"urn:tch SystemNotificationEvent,omitempty"`
	MessageReject                 *admi_002_001_01.DocumentTCH                    `xml:"urn:tch MessageReject,omitempty"`
	TokenIdentification           *acmt_022_001_02.DocumentTCH                    `xml:"urn:tch TokenIdentification,omitempty"`
	ParticipantReport             *admn_007_001_01.DocumentTCH                    `xml:"urn:tch ParticipantReport,omitempty"`
	ParticipantReportResponse     *admn_008_001_01.DocumentTCH                    `xml:"urn:tch ParticipantReportResponse,omitempty"`
	PaymentStatusRequest          *pacs_028_001_03.DocumentTCH                    `xml:"urn:tch PaymentStatusRequest,omitempty"`
}

// UnmarshalXML is a custom unmarshaller that allows us to capture the xmlns attributes
func (v *Message) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if (attr.Name.Space == "" && attr.Name.Local == "xmlns") || (attr.Name.Space == "xmlns") {
			newAttr := xml.Attr{}
			newAttr.Value = attr.Value
			newAttr.Name = xml.Name{}
			if attr.Name.Space == "" {
				newAttr.Name.Local = attr.Name.Local
			} else {
				newAttr.Name.Local = fmt.Sprintf("%s:%s", attr.Name.Space, attr.Name.Local)
			}
			v.Xmlns = append(v.Xmlns, newAttr)
		}
	}

	// Go on with unmarshalling.
	type vv Message
	return d.DecodeElement((*vv)(v), &start)
}

// XSD ComplexType declarations

type HdrAndDataTCH struct {
	AppHdr                        head_001_001_01.BusinessApplicationHeaderV01TCH `xml:"urn:tch AppHdr"`
	CreditTransfer                *pacs_008_001_08.DocumentTCH                    `xml:"urn:tch CreditTransfer,omitempty"`
	MessageStatusReport           *pacs_002_001_10.DocumentTCH                    `xml:"urn:tch MessageStatusReport,omitempty"`
	FICreditTransfer              *pacs_009_001_08.DocumentTCH                    `xml:"urn:tch FICreditTransfer,omitempty"`
	Acknowledgement               *camt_035_001_05.DocumentTCH                    `xml:"urn:tch Acknowledgement,omitempty"`
	ResponseRequestForInformation *camt_028_001_09.DocumentTCH                    `xml:"urn:tch ResponseRequestForInformation,omitempty"`
	RequestForInformation         *camt_026_001_07.DocumentTCH                    `xml:"urn:tch RequestForInformation,omitempty"`
	ReturnOfFunds                 *camt_056_001_08.DocumentTCH                    `xml:"urn:tch ReturnOfFunds,omitempty"`
	PaymentRequest                *pain_013_001_07.DocumentTCH                    `xml:"urn:tch PaymentRequest,omitempty"`
	ResponsePaymentRequest        *pain_014_001_07.DocumentTCH                    `xml:"urn:tch ResponsePaymentRequest,omitempty"`
	ResponseReturnOfFunds         *camt_029_001_09.DocumentTCH                    `xml:"urn:tch ResponseReturnOfFunds,omitempty"`
	EchoRequest                   *admn_005_001_01.DocumentTCH                    `xml:"urn:tch EchoRequest,omitempty"`
	EchoResponse                  *admn_006_001_01.DocumentTCH                    `xml:"urn:tch EchoResponse,omitempty"`
	SignOffRequest                *admn_003_001_01.DocumentTCH                    `xml:"urn:tch SignOffRequest,omitempty"`
	SignOffResponse               *admn_004_001_01.DocumentTCH                    `xml:"urn:tch SignOffResponse,omitempty"`
	SignOnRequest                 *admn_001_001_01.DocumentTCH                    `xml:"urn:tch SignOnRequest,omitempty"`
	SignOnResponse                *admn_002_001_01.DocumentTCH                    `xml:"urn:tch SignOnResponse,omitempty"`
	StandaloneRemittance          *remt_001_001_04.DocumentTCH                    `xml:"urn:tch StandaloneRemittance,omitempty"`
	SystemNotificationEvent       *admi_004_001_02.DocumentTCH                    `xml:"urn:tch SystemNotificationEvent,omitempty"`
	MessageReject                 *admi_002_001_01.DocumentTCH                    `xml:"urn:tch MessageReject,omitempty"`
	TokenIdentification           *acmt_022_001_02.DocumentTCH                    `xml:"urn:tch TokenIdentification,omitempty"`
	ParticipantReport             *admn_007_001_01.DocumentTCH                    `xml:"urn:tch ParticipantReport,omitempty"`
	ParticipantReportResponse     *admn_008_001_01.DocumentTCH                    `xml:"urn:tch ParticipantReportResponse,omitempty"`
	PaymentStatusRequest          *pacs_028_001_03.DocumentTCH                    `xml:"urn:tch PaymentStatusRequest,omitempty"`
}

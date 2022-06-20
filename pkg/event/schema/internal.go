package schema

import "errors"

const (
	MobilisimQueueName = "mobilisimSMSQueue"

	MobilisimOneToNEventType = "oneToN"

	MobilisimRejectedStatus      = "REJECTED"
	MobilisimUndeliverableStatus = "UNDELIVERABLE"
	MobilisimPendingStatus       = "PENDING"
)

const (
	MobilisimSystemError        = "SYSTEM"
	MobilisimUnauthorizedError  = "UNAUTHORIZED"
	MobilisimInvalidSenderError = "SENDER ERROR"
	MobilisimCreditError        = "CREDIT"
)

const (
	MobilisimEnglishMessageDecoder = "default"
	MobilisimUnicodeMessageDecoder = "unicode"
	MobilisimTurkishMessageDecoder = "tr"
)

var (
	MobilisimUnexpectedEventType   = errors.New("event: unexpected event type")
	MobilisimUnexpectedHandlerType = errors.New("handler: unexpected event type")
)

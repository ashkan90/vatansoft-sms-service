package schema

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
	MobilisimTurkishMessageDecoder = "tr"
)

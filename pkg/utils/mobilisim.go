package utils

import (
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/event/schema"
)

var mobilisimRequestErrors = map[string]string{
	schema.MobilisimSystemError:        constants.MobilisimSystemError,
	schema.MobilisimInvalidSenderError: constants.MobilisimInvalidSenderError,
	schema.MobilisimUnauthorizedError:  constants.MobilisimUnauthorizedError,
	schema.MobilisimCreditError:        constants.MobilisimCreditError,
}

func GetErrorDescription(mID string) string {
	err, ok := mobilisimRequestErrors[mID]
	if !ok {
		return constants.MobilisimErrorStatus
	}

	return err
}

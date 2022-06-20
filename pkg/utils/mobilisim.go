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

func GetMessageType(mType string) string {
	if mType == constants.MobilisimEnglishMessageDecoder {
		return schema.MobilisimEnglishMessageDecoder
	}

	if mType == constants.MobilisimUnicodeMessageDecoder {
		return schema.MobilisimUnicodeMessageDecoder
	}

	return schema.MobilisimTurkishMessageDecoder
}

func GetMessageQuantity(message, mType string) int {
	var mLength = GetMessageLength(message, mType)

	return MessageQuantity(mType, mLength)
}

func GetMessageLength(message, mType string) int {
	const sthMessageLn = 5

	var plusContainer []string
	var counter = 0

	if GetMessageType(mType) == schema.MobilisimEnglishMessageDecoder {
		plusContainer = []string{"\r", "\n", "€", "{", "}", "[", "~", "]", "^", "|", "ç", "ş", "ğ", "ı", "Ş", "İ", "Ğ"}
		for _, chr := range message {
			if InArray(plusContainer, string(chr)) {
				counter += 2
			} else {
				counter++
			}
		}
	}

	if GetMessageType(mType) == schema.MobilisimTurkishMessageDecoder {
		plusContainer = []string{"\r", "\n", "€", "{", "}", "[", "~", "]", "^", "|"}
		for _, chr := range message {
			if InArray(plusContainer, string(chr)) {
				counter += 2
			} else {
				counter++
			}
		}
	}

	return counter + sthMessageLn
}

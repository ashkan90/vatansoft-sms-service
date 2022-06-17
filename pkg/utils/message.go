package utils

import (
	"strings"
	"vatansoft-sms-service/pkg/constants"
)

var (
	mergeChars = []string{
		"Ç", "C", "ç", "c",
		"Ğ", "G", "ğ", "g",
		"İ", "I", "Ö", "O",
		"ö", "o", "Ş", "S",
		"ş", "s", "Ü", "U",
		"ü", "u",
	}
)

var (
	quantityEnglishMap = map[int]int{
		1:   1,
		160: 2,
		306: 3,
		459: 4,
		612: 5,
		765: 6,
		918: 0,
	}
	quantityTurkishMap = map[int]int{
		1:   1,
		155: 2,
		294: 3,
		441: 4,
		588: 5,
		735: 6,
		882: 0,
	}
)

func RecomposeMessage(message, messageType string) string {
	if messageType == constants.MobilisimEnglishMessageDecoder {
		message = strings.NewReplacer(mergeChars...).Replace(message)
	}

	return message
}

func MessageQuantity(mType string, mLen int) int {
	if mType == constants.MobilisimEnglishMessageDecoder {
		return OrderedComparison(quantityEnglishMap, mLen).(int)
	}

	return OrderedComparison(quantityTurkishMap, mLen).(int)
}

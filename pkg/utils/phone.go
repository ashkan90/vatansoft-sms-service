package utils

import "regexp"

func CleanupPhone(number string) string {
	return regexp.MustCompile(`/[^.%0-9]/`).ReplaceAllString(number, "")
}

func IsGSMValid(number string) bool {
	var validGSMCodes = []string{
		"50", "51", "53", "54", "55", "56",
	}
	var gsm = number[:2]

	return Index(validGSMCodes, gsm) != -1
}

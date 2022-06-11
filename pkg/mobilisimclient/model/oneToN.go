package model

type RequestOneToN struct {
	Messages []OneToNMessage `json:"messages"`
}

type OneToNMessage struct {
	From            string               `json:"from"`
	Text            string               `json:"text"`
	CallbackData    string               `json:"callbackData"`
	Transliteration string               `json:"transliteration"`
	ValidityPeriod  int                  `json:"validityPeriod"`
	Destinations    []MessageDestination `json:"destinations"`
	Language        MessageLanguage      `json:"language"`
}

type MessageDestination struct {
	To        string `json:"to"`
	MessageID string `json:"messageId,omitempty"`
}

type MessageLanguage struct {
	LanguageCode string `json:"languageCode"`
	SingleShift  bool   `json:"singleShift"`
	LockingShift bool   `json:"lockingShift"`
}

type ResourceOneToN struct{}

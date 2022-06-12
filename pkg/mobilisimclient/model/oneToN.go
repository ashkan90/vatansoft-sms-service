package model

import "vatansoft-sms-service/pkg/event/schema"

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

type ResourceOneToN struct {
	Messages []ResourceOneToNMessages
}

type ResourceOneToNMessages struct {
	To        string               `json:"to"`
	Status    ResourceOneToNStatus `json:"status"`
	SMSLength int                  `json:"smsCount"`
}

type ResourceOneToNStatus struct {
	Action      string `json:"action,omitempty"`
	Status      string `json:"groupName"`
	StatusField string `json:"name"`
	Description string `json:"description"`
}

func (r ResourceOneToNMessages) IsRejected() bool {
	return r.Status.Status == schema.MobilisimRejectedStatus
}

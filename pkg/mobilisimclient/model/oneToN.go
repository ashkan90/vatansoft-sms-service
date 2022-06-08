package model

type RequestOneToN struct {
	Messages OneToNContainer `json:"messages"`
}

type OneToNContainer []OneToNMessage
type OneToNMessage struct {
	From            string
	Text            string
	CallbackData    string
	Transliteration string
	ValidityPeriod  int
	Destinations    []MessageDestination
	Language        MessageLanguage
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

func (c OneToNContainer) SetMessage(m OneToNMessage) {
	c[0] = m
}

func (c OneToNContainer) AddReceiver(phone string) {
	c[0].Destinations = append(c[0].Destinations, MessageDestination{
		To: phone,
	})
}

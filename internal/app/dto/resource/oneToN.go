package resource

import (
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

type OneToNResource struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	ReportID    int    `json:"id"`
	NumberCount int    `json:"numberCount"`
	SMSQuantity int    `json:"quantity"`
}

func NewOneToNResource(r *model.ResourceOneToN) OneToNResource {
	var res OneToNResource
	for _, message := range r.Messages {
		// reject olmayan responseları handleliyoruz ama kimin reject olduğunu hiç bir şekilde öğrenemeyecek.
		// Bunun rapor edilmesi gerektiği aşikar fakat nasıl bi yol izlenmeli fikrim yok.
		if !message.IsRejected() {
			res.NumberCount += 1
			res.SMSQuantity += message.SMSLength
		}
	}

	return res
}

package resource

import (
	"vatansoft-sms-service/pkg/constants"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
	"vatansoft-sms-service/pkg/utils"
)

type OneToNResource struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	ReportID    int    `json:"id"`
	NumberCount int    `json:"numberCount"`
	SMSQuantity int    `json:"quantity"`
}

func NewOneToNResource(r *model.ResourceOneToN, mText, mType string) OneToNResource {
	var res = OneToNResource{
		Status:      constants.MobilisimSuccessStatus,
		Description: constants.MobilisimSuccessDescription,
	}

	if err := r.Error(); err != "" {
		res.Status = constants.MobilisimErrorStatus
		res.Description = utils.GetErrorDescription(err)
		return res
	}

	for _, message := range r.Messages {
		// reject olmayan responseları handleliyoruz ama kimin reject olduğunu hiç bir şekilde öğrenemeyecek.
		// Bunun rapor edilmesi gerektiği aşikar fakat nasıl bi yol izlenmeli fikrim yok.
		if !message.IsRejected() {
			res.NumberCount += 1
			res.SMSQuantity += utils.GetMessageQuantity(mText, mType)
		}
	}

	return res
}

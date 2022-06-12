package response

type MobilisimSuccessResponse struct {
	msg     string
	msgType string
}

func NewMobilisimSuccessResponse(msg, msgType string) MobilisimSuccessResponse {
	return MobilisimSuccessResponse{
		msg:     msg,
		msgType: msgType,
	}
}

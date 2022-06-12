package response

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func NewSuccessResponse(d interface{}) SuccessResponse {
	return SuccessResponse{
		Data: d,
	}
}

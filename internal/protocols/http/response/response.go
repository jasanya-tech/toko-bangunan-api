package response

type Response struct {
	Status  *string `json:"status"`
	Message *any    `json:"message"`
	Data    *any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  *string `json:"status"`
	Message *any    `json:"message"`
}

func NewResponse(status string, code int, message any, data any) any {
	if data != nil {
		return &Response{
			Status:  &status,
			Message: &message,
			Data:    &data,
		}
	} else {
		return &ResponseWithoutData{
			Status:  &status,
			Message: &message,
		}
	}
}

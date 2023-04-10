package helpers

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func NewResponse(status string, code int, message any, data any) any {
	if data != nil {
		return &Response{
			Status:  status,
			Code:    code,
			Message: message,
			Data:    data,
		}
	} else {
		return &ResponseWithoutData{
			Status:  status,
			Code:    code,
			Message: message,
		}
	}
}

package response

import "toko-bangunan/internal/helpers/pagination/entities"

type Response struct {
	Status  *string `json:"status"`
	Message *any    `json:"message"`
	Data    *any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  *string `json:"status"`
	Message *any    `json:"message"`
}

type ResponseWithPaginate struct {
	Status   *string                  `json:"status"`
	Message  *any                     `json:"message"`
	Data     *any                     `json:"data"`
	Paginate *entities.PaginationInfo `json:"paginate_info"`
}

func NewResponse(status string, code int, message any, data any, paginate *entities.PaginationInfo) any {
	if data != nil {
		if paginate != nil {
			return &ResponseWithPaginate{
				Status:   &status,
				Message:  &message,
				Data:     &data,
				Paginate: paginate,
			}
		}
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

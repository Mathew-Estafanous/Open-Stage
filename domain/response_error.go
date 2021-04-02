package domain

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponseError(msg string, sts int) *ResponseError {
	return &ResponseError{
		Message: msg,
		Status:  sts,
	}
}

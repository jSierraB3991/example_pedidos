package dto

type ErrorResponse struct {
	code    int
	message string
}

func NewError(code int, message string) ErrorResponse {
	return ErrorResponse{
		code:    code,
		message: message,
	}
}

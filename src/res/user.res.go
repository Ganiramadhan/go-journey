package res

type Response struct {
	Status  string      `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse untuk response sukses
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  "success",
		Success: true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse untuk response error
func ErrorResponse(message string, err error) Response {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return Response{
		Status:  "error",
		Success: false,
		Message: message,
		Error:   errMsg,
	}
}

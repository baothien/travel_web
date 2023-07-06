package api_response

type ErrorInfo struct {
	Code    int      `json:"code,omitempty"`
	Message []string `json:"message,omitempty"`
}

type Response struct {
	ErrorInfo *ErrorInfo  `json:"error_info,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data interface{}, message string) *Response {
	return &Response{Data: data, Message: message}
}

func SimpleSuccessResponse(data interface{}) *Response {
	return &Response{Data: data}
}

func SimpleSuccessResponseWithMsg(msg string) *Response {
	return &Response{Message: msg}
}

func SimpleErrorResponse(code int, message ...string) *Response {
	return &Response{ErrorInfo: &ErrorInfo{Code: code, Message: message}}
}

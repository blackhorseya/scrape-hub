package http

// Response 定義標準 API 回應格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse 建立標準回應物件
func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewSuccessResponse 建立成功回應
func NewSuccessResponse(data interface{}) *Response {
	return NewResponse(200, "Success", data)
}

// NewErrorResponse 建立錯誤回應
func NewErrorResponse(code int, message string) *Response {
	return NewResponse(code, message, nil)
}

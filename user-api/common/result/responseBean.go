package result

type ResponseSuccessBean struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

type ResponseErrorBean struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg, nil}
}

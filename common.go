package gintool

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespOK(code int, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: CodeMsg(code),
		Data:    data,
	}
}

func RespError(code int, err error) *Response {
	resp := &Response{
		Code:    code,
		Message: CodeMsg(code),
	}
	if err != nil {
		resp.Message += " - " + err.Error()
	}
	return resp
}

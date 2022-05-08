package serialize

const (
	SUCCESS             = 200
	SERVER_ERR          = 500
	BADREQUEST          = 400
	NOT_AUTHENTICATION  = 401
	DEFAULT_SUCCESS_MSG = "ok"
	DEFAULT_ERROR_MSG   = "服务器异常"
)

type Response struct {
	Status int16       `json:"status"`
	Msg    string      `json:"msg"`
	ErrMsg error       `json:"errorMsg"`
	Data   interface{} `json:"data"`
}

// 请求成功
func OK() Response {
	return Response{
		Status: SUCCESS,
		Msg:    DEFAULT_SUCCESS_MSG,
	}
}

// 请求成功，返回data数据
func OKData(data interface{}) Response {
	return Response{
		Status: SUCCESS,
		Msg:    DEFAULT_SUCCESS_MSG,
		Data:   data,
	}
}

// 请求失败
func Fail(msg string, err error) Response {
	return Response{
		Status: SERVER_ERR,
		Msg:    msg,
		ErrMsg: err,
	}
}

// 坏的请求，参数异常
func BadReq(msg string) Response {
	if msg == "" {
		msg = "参数异常"
	}
	return Response{
		Status: BADREQUEST,
		Msg:    msg,
	}
}

package serialize

const ()

type Response struct {
	Status int16
	Msg    string
	ErrMsg error
	Data   interface{}
}

// 请求成功
func (res Response) OK() Response {
	return Response{
		Status: 200,
		Msg:    "success",
	}
}

// 请求成功，返回data数据
func (res Response) OKData(data interface{}) Response {
	return Response{
		Status: 200,
		Msg:    "success",
		Data:   data,
	}
}

// 请求失败
func (res Response) Fail(msg string, err error) Response {
	return Response{
		Status: 500,
		Msg:    msg,
		ErrMsg: err,
	}
}

// 坏的请求，参数异常
func (res Response) BadReq() Response {
	return Response{
		Status: 400,
		Msg:    "参数异常",
	}
}

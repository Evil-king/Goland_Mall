package serializer

type Result struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Error  error   `json:"error"`
	Data interface{} `json:"data"`
}

func SuccessMsg(data interface{},err error) Result {
	return Result{
		Data: data,
		Code: 200,
		Msg: "success",
		Error: err,
	}
}

func Success() Result {
	return Result{
		Data: nil,
		Code: 200,
		Msg: "success",
		Error: nil,
	}
}

func SuccessData(data interface{}) Result {
	return Result{
		Data: data,
		Code: 200,
		Msg: "success",
		Error: nil,
	}
}

func Fail(data interface{},err error) Result  {
	return Result{
		Data: data,
		Msg: "fail",
		Error: err,
	}
}

func FailMsg(msg string) Result  {
	return Result{
		Data: nil,
		Msg: msg,
		Error: nil,
	}
}

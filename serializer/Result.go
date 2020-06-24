package serializer

type Result struct {
	Msg string `json:"msg"`
	Error  error   `json:"error"`
	Data interface{} `json:"data"`
}

func Success(data interface{},err error) Result {
	return Result{
		Data: data,
		Msg: "成功",
		Error: err,
	}
}

func Fail(data interface{},err error) Result  {
	return Result{
		Data: data,
		Msg: "失败",
		Error: err,
	}
}

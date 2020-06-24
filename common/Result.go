package common

type Result struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessToData(data interface{}) Result {
	return Result{
		Code: "200",
		Data: data,
		Msg: "成功",
	}
}

func Success() Result {
	return Result{
		Code: "200",
		Msg: "成功",
	}
}

func Fail() Result  {
	return Result{
		Code: "400",
		Msg: "失败",
	}
}

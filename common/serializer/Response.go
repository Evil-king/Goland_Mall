package serializer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Result struct {
//	Msg string `json:"msg"`
//	Code int `json:"code"`
//	Error  error   `json:"error"`
//	Data interface{} `json:"data"`
//}
//
//func SuccessMsg(data interface{},err error) Result {
//	return Result{
//		Data: data,
//		Code: 200,
//		Msg: "success",
//		Error: err,
//	}
//}
//
//func Success() Result {
//	return Result{
//		Data: nil,
//		Code: 200,
//		Msg: "success",
//		Error: nil,
//	}
//}
//
//func SuccessData(data interface{}) Result {
//	return Result{
//		Data: data,
//		Code: 200,
//		Msg: "success",
//		Error: nil,
//	}
//}
//
//func Fail(data interface{},err error) Result  {
//	return Result{
//		Data: data,
//		Msg: "fail",
//		Error: err,
//	}
//}
//
//func FailMsg(msg string) Result  {
//	return Result{
//		Data: nil,
//		Msg: msg,
//		Error: nil,
//	}
//}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
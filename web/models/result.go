package models

import "github.com/kataras/iris/v12"

//统一返回
type Result struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Data interface{} `json:"data"`
}

type PageResult struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Data interface{} `json:"data"`
	Total int `json:"total"`
}

//成功返回时
func ResponseOk(ctx iris.Context,data interface{},msg string) {
	ctx.StatusCode(iris.StatusOK)//写状态码
	ctx.JSON(Result{Code:1,Msg:msg,Data:data})
	return
}

//返回报错
func ResponseError(ctx iris.Context,errcode int,errmsg string) {
	ctx.StatusCode(errcode)//写状态码
	ctx.JSON(Result{Code:-1,Msg:errmsg})
	return
}
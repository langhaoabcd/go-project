package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-project/utils/jwtHelper"
	"go-project/web/models"
	"go-project/web/service"
)

type UserController struct {
	Ctx iris.Context
	Service service.UserService//接口类型
}

func NewUserController() *UserController  {
	return &UserController{Service:service.NewUserService()}
}

//控制器前置路由
func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	//b.Handle("GET", "/test", "GetTest")// method,path,funcName,middleware
}

// @Summary 测试
// @Description 测试
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Result
// @Router /user/test [get]
// @Security ApiKeyAuth
func (c *UserController) GetTest() (res models.Result) {

	tokenStr:=jwtHelper.GetToken(c.Ctx)
	_,_=jwtHelper.ParseToken(tokenStr)
	userId:=jwtHelper.GetUserId(tokenStr)
	res =models.Result{Data:userId,Msg:"ss",Code:1}
	c.Ctx.StatusCode(iris.StatusUnauthorized)
	return
}

// @Summary 新增
// @Description 新增
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Result
// @Router /user/add [post]
// @Security ApiKeyAuth
func (c *UserController) PostAdd() (res models.Result)  {
	ret:=c.Service.Add(1)
	return ret
}
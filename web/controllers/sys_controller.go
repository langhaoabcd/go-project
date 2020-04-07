package controllers

import (
"github.com/kataras/iris/v12"
"github.com/kataras/iris/v12/mvc"
"go-project/web/models"
"go-project/web/service"
)

type SysController struct {
	Ctx iris.Context
	Service service.SysService//接口类型
}

func NewSysController() *SysController  {
	return &SysController{Service:service.NewSysService()}
}

//控制器前置路由
func (m *SysController) BeforeActivation(b mvc.BeforeActivation) {
	//b.Handle("GET", "/test", "GetTest")// method,path,funcName,middleware
}

// @Summary 系统登录
// @Description 系统登录
// @Tags Sys
// @Accept  json
// @Produce  json
// @Param account query string true "账号"
// @Param password query string true "密码"
// @Success 200 {object} dtos.UserInfo
// @Router /sys/login [post]
func (c *SysController) PostLogin() (res models.Result)  {
	var m map[string]string
	err:=c.Ctx.ReadJSON(&m)
	if err!=nil {
	}
	ret:=c.Service.Login(m)
	return ret
}

// @Summary 试验
// @Description 测试参数错误400和内部错误500
// @Tags Sys
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Result
// @Router /sys/diy [get]
func (c *SysController) GetDiy() {
	//c.Ctx.StatusCode(iris.StatusBadRequest)
	//c.Ctx.JSON(models.Result{Msg:"服务器内部错误",Code:-1})
	models.ResponseError(c.Ctx,iris.StatusInternalServerError,"服务器内部错误")
}
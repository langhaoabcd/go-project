package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"              // swagger middleware for Iris
	"github.com/iris-contrib/swagger/v12/swaggerFiles" // swagger embed files
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	//"github.com/kataras/iris/v12/middleware/recover"
	conf "go-project/configs"
	"go-project/datasource"
	_ "go-project/docs"
	"go-project/web/router"
)

// @title Golang Restful API v1.0
// @version 1.0
// @description Go-Iris-Gorm示例项目

// @contact.name hzh
// @contact.url http://cube.newtouch.com
// @contact.email 1198754948@qq.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /api
func main() {
	//iris.RegisterOnInterrupt(func() {
	//	datasource.GetDB().Close()
	//})
	defer datasource.GetDB().Close()
	app :=createApp()
	router.Register(app)
	err:=app.Run(iris.Addr(":"+conf.GetConf().Server.Port)) //监听端口
	if err!= nil{
		panic(err)
	}
}

func createApp() *iris.Application  {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	//打印控制台日志到终端
	app.Logger().SetLevel("debug")
	app.Use(logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            true,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc:            nil,
		LogFuncCtx:         nil,
		Skippers:           nil,
	}))
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)

	//配置Swagger
	swgconfig := &swagger.Config{
		URL: "http://"+conf.GetConf().Server.Host+":"+conf.GetConf().Server.Port+"/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swgconfig, swaggerFiles.Handler))
	return app
}
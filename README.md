# go-project

## 项目介绍  
最近全球病毒肆虐，在疫情期间学了下Golang。
在Iris Web Framework的基础上花了两天搭建一套Golang Restful Api，Golang这门语言真心不错，感觉相见恨晚...

## 框架&功能  
1.Iris mvc --号称最快的golang web框架  
2.Gorm --本来想试试xorm，因被墙get不下来  
3.Swagger展示接口文档   
4.Jwt安全验证   
5.Go mod管理依赖

## 后续计划   
1.集成Redis,Mongo,RabbitMq等中间件   
2.集成WeSocket   
。。。   

## 项目结构 
 
configs                 --存配置文件及解析  
datasource              --读取数据库  
utils                   --帮助类  
docs                    --自动生成swag文档  
web                     --项目主目录   
    controllers      --控制器  
    router           --路由  
    middleware       --中间件  
    models           --模型放表实体，枚举，dto  
    service          --业务层  
    repositories     --数据仓储层  


## 启动项目  
1.安装依赖 go get  
2.go run main.go  


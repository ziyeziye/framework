Golang+Gin+Gorm  
  
目录  
app  
    ├── api             接口  
    ├── bin             命令  
    ├── conf            用于存储配置文件  
    ├── controllers     控制器  
    ├── middleware      应用中间件  
    ├── models          应用数据库模型  
    ├── pkg             扩展工具包  
    ├── routers         路由逻辑处理  
    ├── runtime         临时文件  
    ├── services        服务层/逻辑  
    ├── spider          爬虫  
    ├── test            测试文件  
    └── view            视图模板

Framework-gen  
github.com/ziyeziye/framework-gen 是一个可以通过数据库生成对应framework的models,struct以及相应的restful api的工具。  
  
执行
go run main.go
访问
127.0.0.1:80 访问模板首页
127.0.0.1:80/api/test 访问接口api/test


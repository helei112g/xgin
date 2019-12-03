# xgin

使用gin，搭建的一个基础化的go web项目结构，方便以后使用

# 目录结构

```
├── app
│   ├── config          // 配置文件
│   ├── controller      // 控制器
│   ├── middleware      // 中间件
│   └── model           // model
├── conf
│   └── app.toml        // 配置文件
├── router
│   └── router.go       // 路由
├── test
│   ├── api             // 功能测试
│   └── benchmark       // 压力测试
├── utils               // 工具包
├── LICENSE             // 开源协议
├── README.md
├── go.mod              // mod 包管理
├── go.sum
├── main.go             // 入口文件
└── vendor
```

# 测试

功能测试：

```sh
cd test/api
go test
```

性能测试

```sh
cd test/benchmark
go test -bench=.
```

# 记录日志

request 日志
[request_time] [request_method] [request_uri] [request_proto] [request_ua] [request_referer] [request_post_data] [client_ip]

response 日志
[response_time] [response_code] [response_msg] [response_data]

cost_time 话费时间
# xgin

使用gin，搭建的一个基础化的go web项目结构，方便以后使用

# 目录结构

```
├── LICENSE         // 开源协议
├── README.md
├── api             // 对外开放的api包
│   └── v1          // 版本控制
├── conf            // 配置文件
│   └── app.toml
├── framework       // 一些框架用到的包
│   ├── config      // 配置文件
│   ├── e           // 错误码管理
│   ├── store       // 数据库 or cache初始化
│   └── util        // 工具类
├── go.mod
├── go.sum
├── main.go         // 入口文件
├── middleware      // 中间件
├── model           // 对应底层model
├── router          // 路由
│   └── router.go
├── runtime         // 运行时存储数据
│   └── logs
├── test            // 测试
│   ├── api         // http 形式的api测试
│   └── benchmark   // 压力测试
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
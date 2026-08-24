go 项目创建

```shell
$ mkdir goroapp
$ cd goroapp
$ git init

# 初始化
$ go mod init github.com/rocwg/goroapp
```

安装依赖

```shell
$ go get github.com/jmoiron/sqlx # sqlx（轻量封装 SQL，推荐）
$ go get github.com/lib/pq       # # PostgreSQL driver
$ go get github.com/google/uuid  # # uuid
```


```shell
$ go mod tidy
```


运行

```shell
$ docker compose up -d     ## 启动 pg 数据库
$ go run cmd/app/main.go
```

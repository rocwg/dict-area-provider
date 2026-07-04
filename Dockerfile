# ==================== Stage 1: Build ====================
FROM golang:1.26-alpine AS builder

# 设置 Go 环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

# 先复制 go.mod 和 go.sum 锁依赖
COPY go.mod go.sum ./
# 注意：由于你在本地开发使用了 replace 指向了本地的 ../grpc-contracts 目录
# 如果要在 Docker 独立编译，有两种标准做法：
# 1. 临时注释掉 replace，让它直接去 GitHub 拉取最新的线上契约包
# 2. 或者把契约仓也一起 COPY 进来。
# 为了让你今天最快闭环，建议你在 Docker 编译前，先将本地最新契约 git push 到 GitHub 线上。
RUN go mod download

# 复制业务源码
COPY . .

# 编译出名为 app 的静态二进制文件
RUN go build -ldflags="-s -w" -o app main.go

# ==================== Stage 2: Run ====================
FROM alpine:3.20

# 安装时区数据包，确保容器内时间是中国上海
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

# 从 builder 阶段把编译好的二进制文件偷过来
COPY --from=builder /build/app .

# 暴露出 gRPC 监听端口
EXPOSE 50051

# 启动命令
CMD ["./app"]

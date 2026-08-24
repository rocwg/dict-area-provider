package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rocwg/dict-area-provider/service"
	pb "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func main() {
	// 1. 初始化 PostgreSQL 连接 (请根据你的本地实际 PG 密码和端口修改 dsn)
	//dsn := "host=127.0.0.1 user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=192.168.1.60 user=postgres password=cs1@gis dbname=forestry port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ 数据库连接失败: %v", err)
	}
	log.Println("🚀 数据库连接成功!")

	// 2. 监听本地端口
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("❌ 监听端口失败: %v", err)
	}

	// 3. 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 4. 注册我们写好的服务实现类
	srv := &service.DictAreaServiceServer{DB: db}
	pb.RegisterDictAreaServiceServer(grpcServer, srv)

	// 5. 开启 gRPC 反射服务 (核心！免去 grpcurl 导入 proto 的痛苦)
	reflection.Register(grpcServer)

	log.Printf("🤖 Dict-Area gRPC 服务已在端口 [%d] 成功启动...", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ 服务启动异常: %v", err)
	}
}

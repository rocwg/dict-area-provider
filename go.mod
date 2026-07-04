module github.com/rocwg/dict-area-service

go 1.26.4

require (
	// 引入契约仓的根模块（不要带子路径，版本号用 v0.0.0 占位即可）
	github.com/rocwg/grpc-contracts v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.82.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.2
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.10.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260630182238-925bb5da69e7 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

// 本地开发时启用，提交代码前注释掉或删除
// 本地联动指向：直接指向根目录级别
replace github.com/rocwg/grpc-contracts => ../grpc-contracts

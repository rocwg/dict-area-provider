module github.com/rocwg/dict-area-provider

go 1.27

require (
	// 引入契约仓的根模块（不要带子路径，版本号用 v0.0.0 占位即可）
	github.com/rocwg/grpc-contracts v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.83.2
	gorm.io/driver/postgres v1.6.2
	gorm.io/gorm v1.31.2
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.10.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260825221802-da73d73af1c5 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

// 本地开发时启用，提交代码前注释掉或删除
// 本地联动指向：直接指向根目录级别
replace github.com/rocwg/grpc-contracts => ../../grpc-contracts

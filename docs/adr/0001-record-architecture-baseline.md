# ADR 0001: 2026年5月项目初始架构规范与命名决策

- **日期**: 2026-05-17
- **状态**: Accepted (已通过/执行中) / Proposed (草案) / Superseded (被取代)
- **作者**: [你的名字]

## 1. 背景上下文 (Context)
项目初始化时命名为 `go-postgres-demo`，采用了标准的 `cmd/` + `internal/model/` + `internal/service/` 结构。
考虑到项目需要具备长期的生命周期，并且后续会随着 Go 官方版本的发布（如新特性引入）和云原生部署环境的变化而演进，需要对初始的命名和目录拓扑进行前瞻性规范。

## 2. 决策思考与痛点 (Decision & Reasoning)
- **关于命名**：技术栈命名（Go+Postgres）会导致后续引入 Redis 或更换底层时产生命名耦合，不符合业务导向趋势。
- **关于目录**：原 `docker/` 目录过于局限于 Docker Compose，无法包容未来可能引入的 Kubernetes/Helm 等编排。
- **关于数据层**：传统 `model` 容易变成贫血模型，且强耦合具体数据库驱动，不利于进行高并发下的单元测试（Unit Testing）。

## 3. 最终方案 (Consequences)
1. **重命名意识**：对外和 Go Module 统一采用业务导向命名（如 `identity-service`），根目录脚手架暂时通用化为 `go-service-template`。
2. **拓扑调整**：
   - 将 `docker/` 升级为 `deployments/`。
   - 引入 `internal/domain/` 替代 `model`，在其中通过 `interface` 隔离具体数据库实现。
3. **演进追踪**：后续若因 Go 新版本发布（如泛型优化、标准库变更）导致的架构调整，统一在此目录下新建 ADR 记录。

## 4. 补充思考/备忘 (Notes)
- [2026-05-17]：目前重点观察 Go 社区对于结构化日志（slog）和原生并发控制（如 errgroup）的最新最佳实践，预计在下一个版本升级时引入评估。


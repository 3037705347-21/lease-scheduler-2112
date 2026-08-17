# lease-scheduler__005 Docker 交付说明

## 项目概览
- Lease Scheduler is a Go library for coordinating short-lived ownership of named
- Go module: `example.com/lease-scheduler`

## 标准命令

```bash
go build ./...
go test ./...
```

## Docker 构建

```bash
./build_benzhi_docker.sh lease-scheduler__005-benzhi linux/amd64
docker run --rm -it lease-scheduler__005-benzhi bash
```

## 环境

- 基础镜像: `golang:1.23`
- 依赖在镜像构建阶段预下载，容器内可直接执行 Go 构建和测试命令。
- 代码目录: `/app`

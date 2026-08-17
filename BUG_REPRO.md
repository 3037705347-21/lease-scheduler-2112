# 修复前故障复现（Docker）

## 项目与标准命令

该项目提供进程内的短租约调度能力。使用 `go test -count=20 ./...` 验证活动租约按资源键稳定升序返回。

## 环境构建与编译

在当前 Windows Docker 环境中，使用仓库提供的 `benzhi.Dockerfile` 构建镜像；容器内执行 `go build ./...` 可以完成编译。

## 故障触发步骤

1. 创建资源键分别为 `worker-z`、`worker-m` 和 `worker-a` 的活动租约。
2. 请求活动租约快照。
3. 执行 `go test -count=20 ./...`。

## 实际错误输出

```text
--- FAIL: TestSnapshotOrdersActiveLeasesByAscendingKey (0.00s)
    snapshot_contract_test.go:16: Snapshot() = []lease.Lease{lease.Lease{Key:"worker-z"}, lease.Lease{Key:"worker-m"}, lease.Lease{Key:"worker-a"}}, want ascending keys
FAIL
FAIL    example.com/lease-scheduler    0.099s
FAIL
```

## 期望行为

活动租约快照应以资源键从小到大的稳定顺序返回，即 `worker-a`、`worker-m`、`worker-z`。

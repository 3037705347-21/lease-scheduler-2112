# 修复前故障复现（Docker）

## 项目与标准命令

该项目提供进程内的短租约调度能力。使用 `go test -count=20 ./...` 验证同一资源在重新获取租约后，旧令牌不能释放新租约。

## 环境构建与编译

在当前 Windows Docker 环境中，使用仓库提供的 `benzhi.Dockerfile` 构建镜像；容器内执行 `go build ./...` 可以完成编译。

## 故障触发步骤

1. 节点为 `worker-b` 获取一份租约并保存令牌。
2. 租约过期后，同一节点重新获取该资源并得到新令牌。
3. 使用旧令牌释放当前租约。
4. 执行 `go test -count=20 ./...`。

## 实际错误输出

```text
--- FAIL: TestReleaseRejectsAStaleTokenAfterReacquire (0.00s)
    token_contract_test.go:16: Release() error = <nil>, want ErrTokenMismatch
FAIL
FAIL    example.com/lease-scheduler    0.004s
FAIL
```

## 期望行为

重新获取资源后，只有新租约的令牌可以释放该资源；旧令牌必须被拒绝，且新租约保持有效。

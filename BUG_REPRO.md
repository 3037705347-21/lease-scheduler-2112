# 修复前故障复现（Docker）

## 项目与标准命令

该项目提供进程内的短租约调度能力。使用 `go test -count=20 ./...` 验证租约在到期时是否可被新的持有者接管。

## 环境构建与编译

在当前 Windows Docker 环境中，使用仓库提供的 `benzhi.Dockerfile` 构建镜像；容器内执行 `go build ./...` 可以完成编译。

## 故障触发步骤

1. 创建一个持续一分钟的 `worker-a` 租约。
2. 将当前时间固定为该租约的精确到期时刻。
3. 让另一个节点为同一个资源申请新的租约。
4. 执行 `go test -count=20 ./...`。

## 实际错误输出

```text
--- FAIL: TestAcquireReclaimsLeaseAtItsExactExpiry (0.00s)
    expiry_contract_test.go:17: Acquire() at expiry error = lease is held by another holder
FAIL
FAIL    example.com/lease-scheduler    0.088s
FAIL
```

## 期望行为

租约到达其到期时刻后应不再属于旧节点，新的节点应取得新的租约和新的令牌。

# 修复前故障复现（Docker）

## 项目与标准命令

该项目提供进程内的短租约调度能力。使用 `go test -count=20 ./...` 验证续期后的截止时间从续期时刻开始计算。

## 环境构建与编译

在当前 Windows Docker 环境中，使用仓库提供的 `benzhi.Dockerfile` 构建镜像；容器内执行 `go build ./...` 可以完成编译。

## 故障触发步骤

1. 创建在 12:01:00 到期的租约。
2. 在 12:00:30 请求两分钟续期。
3. 执行 `go test -count=20 ./...`。

## 实际错误输出

```text
--- FAIL: TestRenewUsesTheRenewalTimeAsItsDeadlineBase (0.00s)
    renew_contract_test.go:17: Renew() expiry = 2026-08-17 12:03:00 +0000 UTC, want 2026-08-17 12:02:30 +0000 UTC
--- FAIL: TestSchedulerLifecycle (0.00s)
    scheduler_test.go:26: Renew() expiry = 2026-08-17 12:03:00 +0000 UTC
FAIL
FAIL    example.com/lease-scheduler    0.102s
FAIL
```

## 期望行为

续期后的截止时间应等于续期请求时刻加上请求时长，本例中应为 12:02:30。

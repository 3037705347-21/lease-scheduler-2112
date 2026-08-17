# 修复前故障复现（Docker）

## 项目与标准命令

该项目提供进程内的短租约调度能力。使用 `go test -count=20 ./...` 验证未提供存储时的默认初始化路径。

## 环境构建与编译

在当前 Windows Docker 环境中，使用仓库提供的 `benzhi.Dockerfile` 构建镜像；容器内执行 `go build ./...` 可以完成编译。

## 故障触发步骤

1. 初始化调度器时不提供底层存储。
2. 读取活动租约。
3. 执行 `go test -count=20 ./...`。

## 实际错误输出

```text
--- FAIL: TestNewSchedulerCreatesAStoreWhenNoneIsProvided (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x0]
FAIL
FAIL    example.com/lease-scheduler    0.107s
FAIL
```

## 期望行为

未提供存储时，调度器应建立可用的默认存储；读取活动租约应返回空结果，不应发生 panic。

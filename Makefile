# 单元测试
.PHONY: ut
ut:
	@#go test -race ./...   -race  支持数据竞争检测
	@go test ./...
.PHONY: e2e

SHELL := /bin/bash

# 根模块路径
ROOT_MOD = github.com/your-username/your-project-name

# RPC 模块列表，例如：RPC_MOD=product order
RPC_MOD = product

# 生成 RPC 客户端和服务端代码
.PHONY: gen-rpc
gen-rpc:
	@# 生成客户端
	@cd rpc_gen && cwgo client --type RPC --service ${service} --module ${ROOT_MOD}/rpc_gen -I ../idl/rpc --idl ../idl/rpc/${service}.proto && go mod tidy
	@# 生成服务端
	@cd app/${service} && cwgo server --type RPC --service ${service} --module ${ROOT_MOD}/app/${service} -I ../../idl/rpc --idl ../../idl/rpc/${service}.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" && go mod tidy

# 批量生成 RPC 模块
.PHONY: gen-rpc-all
gen-rpc-all:
	@SERVICE=${RPC_MOD}; \
	for svr in "$${SERVICE[@]}"; do \
		make gen-rpc service=$$svr; \
	done

# 运行 RPC 模块的测试
.PHONY: test-rpc
test-rpc:
	@cd app/${service}/biz/service && go test -v

# 批量测试 RPC 模块
.PHONY: test-rpc-all
test-rpc-all:
	@clear
	@SERVICE=${RPC_MOD}; \
	for svr in "$${SERVICE[@]}"; do \
		make test-rpc service=$$svr;\
	done

# 生成 API 代码
.PHONY: gen-api
gen-api:
	@cd app/hertz && cwgo server --type HTTP --service hertz --module ${ROOT_MOD}/app/hertz --idl ../../idl/hertz/${service}.proto && go mod tidy

# 监控开发环境，自动重载服务
.PHONY: watch
watch:
	@cd app/hertz && air

# 运行服务
.PHONY: run-service
run-service:
	@cd app/${service} && air

# 启动所有服务
.PHONY: bootstrap
bootstrap:
	@SERVICE=${RPC_MOD}; \
	for svr in "$${SERVICE[@]}"; do \
		make run-service service=$$svr; \
	done

# 清理依赖
.PHONY: tidy
tidy:
	@go install github.com/air-verse/air@latest

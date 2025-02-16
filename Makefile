SHELL := /bin/bash
ROOT_MOD = github.com/MelodyDeep/TikTok-E-commerce
# RPC module list, for example: RPC_MOD= user product order checkout
RPC_MOD= product 


.PHONY: gen-rpc
gen-rpc:
	@# generate client
	@cd rpc_gen && cwgo client --type RPC --service ${service} --module ${ROOT_MOD}/rpc_gen -I ../idl/rpc --idl ../idl/rpc/${service}.proto && go mod tidy
	@# generate server
	@cd app/${service} && cwgo server --type RPC --service ${service} --module ${ROOT_MOD}/app/${service} -I ../../idl/rpc --idl ../../idl/rpc/${service}.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" && go mod tidy

.PHONY: gen-rpc-all
gen-rpc-all:
	@SERVICE=${RPC_MOD}; \
	for svr in "$${SERVICE[@]}"; do \
		make gen-rpc service=$$svr; \
	done


.PHONY: test-rpc
test-rpc:
	@cd app/${service}/biz/service && go test -v

.PHONY: test-rpc-all
test-rpc-all:
	@clear
	@SERVICE=${RPC_MOD}; \
	for svr in "$${SERVICE[@]}"; do \
		make test-rpc service=$$svr;\
	done

.PHONY: gen-api
gen-api:
	@cd app/hertz && cwgo server --type HTTP --service hertz --module ${ROOT_MOD}/app/hertz --idl ../../idl/hertz/${service}.proto && go mod tidy

.PHONY: watch
watch:
	@cd app/hertz && air
.PHONY:build build-ui build-backend help

build: build-ui build-backend


build-ui:
	cd CasaOS-UI && yarn install && yarn build

build-backend:
	export CGO_ENABLED=1;export CGO_LDFLAGS=-static;go build -o ./6dainas main.go;upx --lzma --best 6dainas

help:
	@echo "call john"

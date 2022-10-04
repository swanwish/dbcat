appName=dbcat

amd64-build:
	@echo "build ${appName} for android emulator"
	@CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" -o ${appName}_amd64 main.go
	@upx ${appName}_amd64

arm64-build:
	@echo "build ${appName} for android physical machine"
	@CC=aarch64-linux-musl-gcc CXX=aarch64-linux-musl-g++ GOARCH=arm64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" -o ${appName}_arm64 main.go
	@upx ${appName}_arm64

mac-build:
	@echo "build ${appName} for mac"
	@go build -o ${appName}_mac main.go

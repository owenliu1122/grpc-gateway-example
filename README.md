# grpc-gateway 使用指南

## 安装 protobuf

grpc-gateway 依赖 protobuf, 所以需要先安装 protobuf

### 下载源码

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protobuf-all-3.7.1.tar.gz
tar -xzvf protobuf-all-3.7.1.tar.gz
cd protobuf-3.7.1/

```

### 安装 automake 和 libtool

```bash
brew install automake
brew install libtool
```

### 运行自动生成脚本

```bash
./autogen.sh
```

### 安装

```bash
./configure
make check
make
sudo make install
```

### 查看版本

```bash
protoc --version
```

## 安装 protoc-gen-go

```bash
# 下载 googleapis
mkdir -p $GOPATH/src/github.com/googleapis
git clone https://github.com/googleapis/googleapis.git $GOPATH/src/github.com/googleapis/googleapis

go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
```

## 安装 protoc-gen-govalidators

这个插件和 grpc-gateway 无关，可选安装，我这里只是为了测试

```bash
go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
```

## 安装 grpc-gateway

```bash
# 下载 genproto
mkdir -p $GOPATH/src/google.golang.org
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

## 生成 protobuf 代码

```bash
make proto
```

## 编译代码

```bash
make build
```

## 运行 server

```bash
./grpc-gateway-example server --config config-example.yaml
```

## 运行 client

```bash
./grpc-gateway-example client --config config-example.yaml
```

## 运行 gateway

```bash
./grpc-gateway-example gateway --config config-example.yaml
```
# Setup GoLang

## Install GoLang
```bash
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go 
sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
```

## Set GOPATH
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

# Download protoc utility
```
https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-win64.zip
```


# Import gRPC and other dependancies
```
go get -u github.com/golang/protobuf/protoc-gen-go

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
// google.golang.org/protobuf
//go install google.golang.org/grpc@latest

go get google.golang.org/grpc
go get google.golang.org/grpc/codes
go get google.golang.org/grpc/status
go get google.golang.org/protobuf/reflect/protoreflect
go get google.golang.org/protobuf/runtime/protoimpl
```


# Build client and service from protobuf
```bash
PROTO_FILE=gosync.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $PROTO_FILE
```
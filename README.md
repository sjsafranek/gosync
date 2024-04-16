# Install GoLang
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go 
sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

# Set GOPATH
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
// export PATH="$PATH:$(go env GOPATH)/bin"

# Import gRPC
go get -u github.com/golang/protobuf/protoc-gen-go

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go install google.golang.org/grpc@latest




PROTO_FILE=gosync.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $PROTO_FILE
# gosync

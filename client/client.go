package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/sjsafranek/gosync/service"
	pb "github.com/sjsafranek/gosync/gosync"
	"github.com/sjsafranek/logger"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

const (
	DEFAULT_CHUNK_SIZE int    = 64
	DEFAULT_HOST       string = "localhost"
	DEFAULT_PORT       int    = 9622
	DEFAULT_FORCE      bool   = false
)

func uploadFileToServer(client pb.GoSyncServiceClient, filename string, chunk_size int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return uploadFileToServerWithContext(ctx, client, filename, chunk_size)
}

func uploadFileToServerWithContext(ctx context.Context, client pb.GoSyncServiceClient, filename string, chunk_size int) error {
	logger.Infof("Starting file upload: %v", filename)

	// Start streaming client
	stream, err := client.UploadFile(ctx)
	if nil != err {
		return err
	}
	
	// Upload file to server
	err = service.StreamFile(stream, filename, chunk_size, true)
	if nil != err {
		return err
	}

	reply, err := stream.CloseAndRecv()
	if nil != err && io.EOF != err {
		return err
	}
	logger.Info(reply)

	return err
}

func main() {
	var upload_file string
	var chunk_size int
	var host string
	var port int
	var force bool
	flag.StringVar(&upload_file, "file", "", "File to upload")
	flag.BoolVar(&force, "force", DEFAULT_FORCE, "Force")
	flag.IntVar(&chunk_size, "s", DEFAULT_CHUNK_SIZE, "Chunk size")
	flag.StringVar(&host, "h", DEFAULT_HOST, "Server Host")
	flag.IntVar(&port, "p", DEFAULT_PORT, "Server Port")
	flag.Parse()

	// Setup connection
	var opts []grpc.DialOption
	credentials := insecure.NewCredentials()
	opts = append(opts, grpc.WithTransportCredentials(credentials))

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Setup client
	client := pb.NewGoSyncServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = uploadFileToServerWithContext(ctx, client, upload_file, chunk_size)
	if nil != err {
		panic(err)
	}

	// // UPLOAD FILE TO SERVER
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// // Start streaming client
	// stream, err := client.UploadFile(ctx)
	// if nil != err {
	// 	logger.Error(err)
	// 	panic(err) // dont use panic in your real project
	// }
	
	// // Upload file to server
	// err = service.StreamFile(stream, upload_file, chunk_size, true)
	// if nil != err {
	// 	logger.Error(err)
	// 	panic(err) // dont use panic in your real project
	// }

	// reply, err := stream.CloseAndRecv()
	// if nil != err && io.EOF != err {
	// 	panic(err)
	// }
	// logger.Info(reply)
}



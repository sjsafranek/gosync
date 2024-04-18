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
	DEFAULT_CHUNK_SIZE int32    = 64
	DEFAULT_HOST       string = "localhost"
	DEFAULT_PORT       int    = 9622
	DEFAULT_FORCE      bool   = false
)

func uploadFileToServer(client pb.GoSyncServiceClient, filename string, chunk_size int32) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return uploadFileToServerWithContext(ctx, client, filename, chunk_size)
}

func uploadFileToServerWithContext(ctx context.Context, client pb.GoSyncServiceClient, filename string, chunk_size int32) error {
	logger.Infof("Starting file upload: %v", filename)

	// Start streaming client
	stream, err := client.UploadFile(ctx)
	if nil != err {
		return err
	}
	
	// Upload file to server
	err = service.SendFile(stream, filename, chunk_size, true)
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

func downloadFileFromServer(client pb.GoSyncServiceClient, filename string, chunk_size int32) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return downloadFileFromServerWithContext(ctx, client, filename, chunk_size)
}

func downloadFileFromServerWithContext(ctx context.Context, client pb.GoSyncServiceClient, filename string, chunk_size int32) error {
	stream, err := client.DownloadFile(ctx, &pb.FilePayload{
		FileDetails: &pb.FileDetails{
			Filename: filename,
		},
		FileOptions: &pb.FileOptions{
			ChunkSize: chunk_size,
			Encryption: false,
		},
	})
	if nil != err {
		return err
	}
	err = service.RecvFile(stream, "./")
	if err == io.EOF {
		return nil
	}
	return err
}

func main() {
	var filename string
	var _chunk_size int
	var chunk_size int32
	var host string
	var port int
	var force bool
	flag.StringVar(&filename, "file", "", "File to upload")
	flag.BoolVar(&force, "force", DEFAULT_FORCE, "Force")
	flag.IntVar(&_chunk_size, "s", int(DEFAULT_CHUNK_SIZE), "Chunk size")
	flag.StringVar(&host, "h", DEFAULT_HOST, "Server Host")
	flag.IntVar(&port, "p", DEFAULT_PORT, "Server Port")
	flag.Parse()

	chunk_size = int32(_chunk_size)

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

	err = uploadFileToServerWithContext(ctx, client, filename, chunk_size)
	if nil != err {
		panic(err)
	}

	err = downloadFileFromServerWithContext(ctx, client, filename, chunk_size)
	if nil != err {
		panic(err)
	}

	response, err := client.GetFileDetails(ctx, &pb.FilePayload{
		FileDetails: &pb.FileDetails{
			Filename: filename,
		},
		FileOptions: &pb.FileOptions{
			ChunkSize: chunk_size,
		},
	})
	if nil != err {
		panic(err)
	}
	logger.Info(response)
}



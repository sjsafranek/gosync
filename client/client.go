package main

import (
	"context"
	"fmt"
	"log"
	"sync"
    
    "github.com/sjsafranek/logger"
    "github.com/schollz/progressbar/v3"
	pb "github.com/sjsafranek/gosync/gosync"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

const (
	DEFAULT_HOST string = "localhost"
	DEFAULT_PORT int    = 9622
)

var (
	host   string = DEFAULT_HOST
	port   int    = DEFAULT_PORT
	client pb.GoSyncServiceClient
)

func uploadFile(ctx context.Context, filename string) error {
	// Start file transfer
	total_size := getFileSize(filename)
	resp, err := client.StartTransfer(ctx, &pb.StartRequest{
		Filename:    filename,
		Md5Checksum: getMD5Checksum(filename),
		Size:        total_size,
	})
	if nil != err {
		return err
	}
	if pb.State_Continue != resp.State {
		return nil
	}
	transfer_id := resp.TransferId

	// Create progress bar
	progress := progressbar.DefaultBytes(total_size, filename)

	// Upload file in chunks
	queue, err := readFileInChunks(filename, 8)
	if nil != err {
		return err
	}

	var offset int64 = 0
	for chunk := range queue {
		progress.Add(len(chunk))
		resp, err := client.UploadChunk(ctx, &pb.ChunkRequest{
			TransferId: transfer_id,
			Chunk:      chunk,
			Offset:     offset,
		})
		if nil != err {
			return err
		}
		if pb.State_Continue != resp.State {
			return nil
		}
		offset += int64(len(chunk))
	}

	return nil
}

func main() {

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
	client = pb.NewGoSyncServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filename := "test.txt"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := uploadFile(ctx, filename)
		if nil != err {
			logger.Error(err)
		}
	}()
	wg.Wait()
}

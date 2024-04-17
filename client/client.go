package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/schollz/progressbar/v3"
	"github.com/sjsafranek/gosync/fileutils"
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

	// UPLOAD FILE TO SERVER
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Collect file metadata
	total_size := fileutils.GetFileSize(upload_file)
	checksum := fileutils.GetMD5Checksum(upload_file)

	// Read file in chunks
	queue, err := fileutils.ReadFileInChunks(upload_file, chunk_size)
	if nil != err {
		logger.Error(err)
		panic(err)
	}

	// Start streaming client
	stream, err := client.UploadFile(ctx)
	if err != nil {
		logger.Error(err)
		panic(err) // dont use panic in your real project
	}

	// Create progress bar
	progress := progressbar.DefaultBytes(total_size, upload_file)

	// Start file transfer
	request := &pb.Request{
		Filename:    upload_file,
		Md5Checksum: checksum,
		TotalSize:   total_size,
		Overwrite:   force,
	}
	stream.Send(request)

	var offset int64 = 0
	for chunk := range queue {
		progress.Add(len(chunk))
		err = stream.Send(&pb.Request{
			Filename: upload_file,
			Chunk:    chunk,
			Offset:   offset,
		})
		if nil != err {
			panic(err)
		}
		offset += int64(len(chunk))
	}

	reply, err := stream.CloseAndRecv()
	if nil != err && io.EOF != err {
		panic(err)
	}
	logger.Info(reply)

}

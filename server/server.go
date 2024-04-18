package main

import (
	// "errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	// "time"
	"errors"
	"path"
	"context"
	"os/signal"
	"syscall"

	// "github.com/google/uuid"
	grpc "google.golang.org/grpc"
	// "google.golang.org/grpc/peer"

	"github.com/sjsafranek/logger"
	"github.com/sjsafranek/gosync/fileutils"
	"github.com/sjsafranek/gosync/crypto"
    "github.com/sjsafranek/gosync/service"
	pb "github.com/sjsafranek/gosync/gosync"
)

const (
	DEFAULT_HOST string = "localhost"
	DEFAULT_PORT int    = 9622
)

var (
	host string = DEFAULT_HOST
	port int    = DEFAULT_PORT
	data_directory = "./data"
)

type server struct {
	pb.UnimplementedGoSyncServiceServer
}

func (self *server) UploadFile(stream pb.GoSyncService_UploadFileServer) error {
	err := service.RecvFile(stream, data_directory)
	if err == io.EOF {
		return stream.SendAndClose(&pb.FilePayload{
			Status: pb.Status_Ok,
		})
	}
	return err
}

func (self *server) DownloadFile(request *pb.FilePayload, stream pb.GoSyncService_DownloadFileServer) error {
	filename := request.FileDetails.Filename
	filename = path.Join(data_directory, filename)
	if !fileutils.Exists(filename) {
		logger.Warnf("Not Found: %v", filename)
		return errors.New("Not Found")
	}
    return service.SendFile(stream, filename, request.FileOptions.ChunkSize, false) 
}

func (self *server) GetFileDetails (ctx context.Context, request *pb.FilePayload) (*pb.FilePayload, error) {
	filename := request.FileDetails.Filename
	filename = path.Join(data_directory, filename)
	if !fileutils.Exists(filename) {
		return nil, errors.New("Not Found")
	}

	logger.Infof("Collecting file details: %v", filename)

	total_size := fileutils.GetFileSize(filename)
	checksum := fileutils.GetMD5Checksum(filename)

	chunk_size := request.FileOptions.ChunkSize
    if 0 >= chunk_size {
        chunk_size = service.DEFAULT_CHUNK_SIZE
    }
	queue, err := fileutils.ReadFileInChunks(filename, chunk_size)
	if nil != err {
		return nil, err
	}

	var offset int64 = 0
	chunks := []*pb.FileChunk{} 
	for chunk := range queue {
		chunks = append(chunks, &pb.FileChunk{
			Offset: offset,
			MD5Checksum: crypto.MD5FromBytes(chunk),
		})
		offset += int64(len(chunk))
	}

	response := &pb.FilePayload{
		FileDetails: &pb.FileDetails{
			Filename: filename,
			MD5Checksum: checksum,
			Size:   total_size,
		},
		FileOptions: &pb.FileOptions{
			ChunkSize: chunk_size,
		},
		FileChunks: chunks,
	}

	return response, nil
}

func new() *server {
	return &server{}
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// Setup server
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterGoSyncServiceServer(server, new())

	// Handle graceful shutdowns
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		logger.Warnf("got signal %v, attempting graceful shutdown", s)
		//cancel()
		server.GracefulStop()
		wg.Done()
	}()

	// Hookup TCP Listener
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	logger.Infof("Listening on %v", address)

	// Serve gRPC server
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
	wg.Wait()

	logger.Info("clean shutdown")
}

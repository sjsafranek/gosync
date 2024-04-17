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
	//"context"
	"os/signal"
	"syscall"

	// "github.com/google/uuid"
	grpc "google.golang.org/grpc"
	// "google.golang.org/grpc/peer"

	"github.com/sjsafranek/logger"
	// "github.com/sjsafranek/gosync/fileutils"
	// "github.com/sjsafranek/gosync/crypto"
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
)

type server struct {
	pb.UnimplementedGoSyncServiceServer
}

func (self *server) UploadFile(stream pb.GoSyncService_UploadFileServer) error {
	err := service.RecvFile(stream)
	if err == io.EOF {
		return stream.SendAndClose(&pb.FilePayload{
			Status: pb.Status_Ok,
		})
	}
	return err
}

func (self *server) DownloadFile(request *pb.FilePayload, stream pb.GoSyncService_DownloadFileServer) error {
    return service.SendFile(stream, request.FileDetails.Filename, request.FileOptions.ChunkSize, false) 
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

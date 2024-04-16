package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "sync"
    "time"
    "errors"
    "context"
    "os/signal"
    "syscall"

    "github.com/sjsafranek/logger"
    "github.com/google/uuid"
    grpc "google.golang.org/grpc"
    pb "github.com/sjsafranek/gosync/gosync"
)

const (
    DEFAULT_HOST string = "localhost"
    DEFAULT_PORT int = 9622
)

var (
    host string = DEFAULT_HOST
    port int = DEFAULT_PORT
)

type transfer struct {
    Request *pb.StartRequest
    Size int64
    File *os.File
    StartTime time.Time
}

type service struct {
    pb.UnimplementedGoSyncServiceServer
    lock sync.RWMutex
    transfers map[string]*transfer
}

func(self *service) getTransferById(transfer_id string) *transfer {
    self.lock.RLock()
    defer self.lock.RUnlock()
    transfer, _ := self.transfers[transfer_id]
    return transfer
}

func(self *service) StartTransfer(ctx context.Context, req *pb.StartRequest) (*pb.Response, error) {
    logger.Info(req);
    transfer_id := uuid.New().String()
    
    file, err := os.OpenFile(req.Filename,  os.O_WRONLY|os.O_CREATE, 0600)
    if nil != err {
        return nil, err
    }

    logger.Infof("Starting Transfer(%v)", transfer_id)
    self.lock.Lock()
    defer self.lock.Unlock()
    self.transfers[transfer_id] = &transfer{
        Request: req,
        Size: 0,
        File: file,
        StartTime: time.Now(),
    }

    return &pb.Response{State: pb.State_Continue, TransferId: transfer_id}, nil
}

func(self *service) UploadChunk(ctx context.Context, req *pb.ChunkRequest) (*pb.Response, error) {
    logger.Info(req);

    transfer_id := req.TransferId;

    transfer := self.getTransferById(transfer_id)
    if nil == transfer {
        return nil, errors.New("Transfer does not exist")
    }

    logger.Debugf("Writing Transfer(%v) chunk", transfer_id)
    n, err := transfer.File.WriteAt(req.Chunk, req.Offset)
    if nil != err {
        return nil, errors.New("Unable to write to file")
    }
    transfer.Size += int64(n)

    // Transfer has been completed
    if transfer.Request.Size == transfer.Size {
        logger.Infof("Transfer(%v) complete", req.TransferId)
        self.lock.Lock()
        transfer.File.Close();
        delete(self.transfers, req.TransferId)
        defer self.lock.Unlock()
        return &pb.Response{State: pb.State_Complete}, nil
    }

    return &pb.Response{State: pb.State_Continue}, nil
}

func newServiceServer() *service {
    return &service{transfers: make(map[string]*transfer)}
} 

func main() {
    // ctx, cancel := context.WithCancel(context.Background())
    // defer cancel()
    
    // Setup server
    var opts []grpc.ServerOption
    server := grpc.NewServer(opts...)
    pb.RegisterGoSyncServiceServer(server, newServiceServer())

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
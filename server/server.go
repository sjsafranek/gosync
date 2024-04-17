package main

import (
    "os"
    "io"
    "fmt"
    "log"
    "net"
    "sync"
    "time"
    "errors"
    //"context"
    "os/signal"
    "syscall"

    // "github.com/google/uuid"
    grpc "google.golang.org/grpc"
    // "google.golang.org/grpc/peer"

    "github.com/sjsafranek/logger"
    // "github.com/sjsafranek/gosync/fileutils"
    "github.com/sjsafranek/gosync/crypto"
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
    File *os.File
    StartTime time.Time
    UpdateTime time.Time
    BytesExpected int64
    BytesWritten int64
}

func newResponse(state pb.State) *pb.Response {
    return &pb.Response{State: state}
}

type service struct {
    pb.UnimplementedGoSyncServiceServer
    lock sync.RWMutex
    transfers map[string]*transfer
}

func(self *service) getTransferById(transfer_id string) (*transfer, error) {
    self.lock.RLock()
    defer self.lock.RUnlock()
    transfer, ok := self.transfers[transfer_id]
    if !ok {
        return nil, errors.New("Transfer does not exist")
    }    
    return transfer, nil
}

func (self *service) transferExists(transfer_id string) bool {
    self.lock.RLock()
    defer self.lock.RUnlock()
    _, ok := self.transfers[transfer_id]
    return ok
}

func (self *service) createTransferIfNotExists(request *pb.Request) error {
    transfer_id := crypto.MD5(request.Filename)

    // Check if transfer already exists
    if self.transferExists(transfer_id) {
        return nil
    }

    // // Check if file already exists
    // filename := request.Filename
    // if !request.Overwrite {
    //     if fileutils.Exists(filename) {
    //         if fileutils.GetMD5Checksum(filename) == request.Md5Checksum {
    //             logger.Warn("File already exists")
    //             return errors.New("File already exists")
    //         }
    //     }
    // }

    // Create empty file
    file, err := os.OpenFile(request.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
    if nil != err {
        return err
    }

    // Create new transfer
    self.lock.Lock()
    defer self.lock.Unlock()
    self.transfers[transfer_id] = &transfer{
        File: file,
        StartTime: time.Now(),
        UpdateTime: time.Now(),
        BytesExpected: request.TotalSize,
        BytesWritten: 0,
    }

    return nil
}

func (self *service) deleteTransferByTransferIdIfExists(transfer_id string) error {
    transfer, _ := self.getTransferById(transfer_id)
    if nil != transfer {
        logger.Debugf("Transfer(%v) complete", transfer_id)
        self.lock.Lock()
        defer self.lock.Unlock()
        transfer.File.Close();
        delete(self.transfers, transfer_id)
    }
    return nil
}

func (self *service) UploadFile(stream pb.GoSyncService_UploadFileServer) error {
    tracking := make(map[string]bool)
    defer func() {
        for transfer_id := range tracking {
            self.deleteTransferByTransferIdIfExists(transfer_id)
        }
    }()

    for {
        request, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.Response{
                State: pb.State_Complete,
            })
        } else if nil != err {
            return err
        }

        // Get transfer id
        transfer_id := crypto.MD5(request.Filename)
        tracking[transfer_id] = true

        // Start a new transfer if needed
        err = self.createTransferIfNotExists(request)
        if nil != err {
            return err
        }
        
        // Check if data has been recieved
        if nil != request.Chunk {
            // Get transfer job
            transfer, err := self.getTransferById(transfer_id)
            if nil != err {
                return err
            }

            // Write Chunk
            logger.Debugf("Writing Transfer(%v) chunk", transfer_id)
            n, err := transfer.File.WriteAt(request.Chunk, request.Offset)
            if nil != err {
                return err
            }
            transfer.BytesWritten += int64(n)
            transfer.UpdateTime = time.Now()

            if transfer.BytesWritten == transfer.BytesExpected {
                logger.Infof("Transfer(%v) complete", transfer_id)
                self.deleteTransferByTransferIdIfExists(transfer_id)
            } else if transfer.BytesExpected < transfer.BytesWritten {
                return errors.New("File size mismatch")
            }
        }

    }
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
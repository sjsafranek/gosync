package service

import (
	"time"
	"errors"

	"github.com/sjsafranek/logger"

	"github.com/sjsafranek/gosync/crypto"
	pb "github.com/sjsafranek/gosync/gosync"
)

type iReciever interface {
	Recv() (*pb.FilePayload, error)
}

func RecvFile(stream iReciever) error {

	tracking := make(map[string]bool)
	defer func() {
		for transfer_id := range tracking {
			manager.DeleteIfExists(transfer_id)
		}
	}()

	for {
		// Read incoming messages
		request, err := stream.Recv()
		if nil != err {
			return err
		}
		
		// Get transfer id
		transfer_id := crypto.MD5(request.FileDetails.Filename)
		tracking[transfer_id] = true

		// Start a new transfer if needed
		err = manager.CreateIfNotExists(request)
		if nil != err {
			return err
		}

		// Check if data has been recieved
		if nil != request.FileChunk.Chunk {
			// Get transfer job
			transfer, err := manager.Get(transfer_id)
			if nil != err {
				return err
			}

			// Write Chunk
			logger.Debugf("Writing Transfer(%v) chunk", transfer_id)
			n, err := transfer.File.WriteAt(request.FileChunk.Chunk, request.FileChunk.Offset)
			if nil != err {
				return err
			}
			transfer.BytesWritten += int64(n)
			transfer.UpdateTime = time.Now()

			if transfer.BytesWritten == transfer.BytesExpected {
				logger.Infof("Transfer(%v) complete", transfer_id)
				manager.DeleteIfExists(transfer_id)
				delete(tracking, transfer_id)
			} else if transfer.BytesExpected < transfer.BytesWritten {
				return errors.New("File size mismatch")
			}
		}
	}
}

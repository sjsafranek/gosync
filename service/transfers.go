package service

import (
	"os"
	"path"
	"path/filepath"
	"time"
	"sync"
	"errors"

	"github.com/sjsafranek/logger"
	// "github.com/sjsafranek/gosync/fileutils"
	"github.com/sjsafranek/gosync/crypto"
	pb "github.com/sjsafranek/gosync/gosync"
)

type transferManager struct {
	transfers map[string]*transfer
	lock      sync.RWMutex
}

func (self *transferManager) Get(transfer_id string) (*transfer, error) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	transfer, ok := self.transfers[transfer_id]
	if !ok {
		return nil, errors.New("Transfer does not exist")
	}
	return transfer, nil
}

func (self *transferManager) Exists(transfer_id string) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	_, ok := self.transfers[transfer_id]
	return ok
}

func (self *transferManager) CreateIfNotExists(request *pb.FilePayload, directory string) error {
	transfer_id := crypto.MD5FromString(request.FileDetails.Filename)

	// Check if transfer already exists
	if self.Exists(transfer_id) {
		return nil
	}

	// // Create output directory if needed
	// err := fileutils.MakeDirectoryIfNotExists(directory)
	// if nil != err {
	// 	return err
	// }

	// Determine file path
	filename := path.Join(directory, request.FileDetails.Filename)

	logger.Info(filename)

	// Make sure all directories exist
	os.MkdirAll(filepath.Dir(filename), 0700)

	// Create empty file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if nil != err {
		return err
	}

	// Create new transfer
	self.lock.Lock()
	defer self.lock.Unlock()
	self.transfers[transfer_id] = &transfer{
		File:          file,
		StartTime:     time.Now(),
		UpdateTime:    time.Now(),
		BytesExpected: request.FileDetails.Size,
		BytesWritten:  0,
	}

	return nil
}

func (self *transferManager) DeleteIfExists(transfer_id string) error {
	transfer, _ := self.Get(transfer_id)
	if nil != transfer {
		logger.Debugf("Transfer(%v) complete", transfer_id)
		self.lock.Lock()
		defer self.lock.Unlock()
		transfer.File.Close()
		delete(self.transfers, transfer_id)
	}
	return nil
}
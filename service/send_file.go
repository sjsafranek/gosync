package service

import (
	//"github.com/sjsafranek/logger"
	"github.com/schollz/progressbar/v3"

	"github.com/sjsafranek/gosync/fileutils"
	pb "github.com/sjsafranek/gosync/gosync"
)

type iSender interface {
	Send(*pb.FilePayload) error
}

func SendFile(stream iSender, filename string, chunk_size int32, show_progress bool) error {
	// Check parameters
    if 0 >= chunk_size {
        chunk_size = DEFAULT_CHUNK_SIZE
    }	

	// Collect file metadata
	total_size := fileutils.GetFileSize(filename)
	checksum := fileutils.GetMD5Checksum(filename)

	// Read file in chunks
	queue, err := fileutils.ReadFileInChunks(filename, chunk_size)
	if nil != err {
		return err
	}

	// Create progress bar
	var progress *progressbar.ProgressBar
	if show_progress {
		progress = progressbar.DefaultBytes(total_size, filename)
	}

	// Stream file to server
	var offset int64 = 0
	for chunk := range queue {
		if nil != progress {
			progress.Add(len(chunk))
		}
		err = stream.Send(&pb.FilePayload{
			FileDetails: &pb.FileDetails{
				Filename: filename,
				MD5Checksum: checksum,
				Size:   total_size,
			},
			FileChunk: &pb.FileChunk{
				Chunk:    chunk,
				Offset:   offset,
			},
		})
		if nil != err {
			return err
		}
		offset += int64(len(chunk))
	}

	return nil
}
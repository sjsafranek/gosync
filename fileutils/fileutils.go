package fileutils

import (
	"bufio"
	"crypto/md5"	
	"encoding/hex"
	"io"
	"log"
	"os"
)

func MakeDirectoryIfNotExists(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
    	return os.Mkdir(directory, os.ModeDir)
	}
	return nil
}

// Exists reports whether the named file or directory exists.
func Exists(pathname string) bool {
	if _, err := os.Stat(pathname); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetMD5Checksum(filename string) string {
	h := md5.New()
	f, err := os.Open(filename)
	if nil != err {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := io.Copy(h, f); nil != err {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func GetFileSize(filename string) int64 {
	info, err := os.Stat(filename)
	if nil != err {
		log.Fatal(err)
	}
	return info.Size()
}

func ReadFileInChunks(filename string, chunk_size int32) (chan []byte, error) {
	file, err := os.Open(filename)
	if nil != err {
		return nil, err
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, chunk_size)
	queue := make(chan []byte, 4)

	go func() {
		defer file.Close()
		for {
			n, err := reader.Read(buffer)
			if nil != err {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			chunk := make([]byte, len(buffer[0:n]))
			copy(chunk, buffer[0:n])
			queue <- chunk
		}
		close(queue)
	}()

	return queue, nil
}

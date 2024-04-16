package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
        "bufio"
)

func getMD5Checksum(filename string) string {
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

func getFileSize(filename string) int64 {
	info, err := os.Stat(filename)
	if nil != err {
		log.Fatal(err)
	}
	return info.Size()
}

func readFileInChunks(filename string, chunk_size int) (chan []byte, error) {
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

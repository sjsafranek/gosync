package compress

import (
	"bytes"
	"compress/flate"
	"io"
)

// CompressWithOption returns compressed data using the specified level
func CompressWithOption(src []byte, level int) []byte {
	compressedData := new(bytes.Buffer)
	err := compress(src, compressedData, level)
	if nil != err {
		return err
	}
	return compressedData.Bytes()
}

// Compress returns a compressed byte slice.
func Compress(src []byte) []byte {
	compressedData := new(bytes.Buffer)
	err := compress(src, compressedData, -2)
	if nil != err {
		return err
	}
	return compressedData.Bytes()
}

// Decompress returns a decompressed byte slice.
func Decompress(src []byte) []byte {
	compressedData := bytes.NewBuffer(src)
	deCompressedData := new(bytes.Buffer)
	decompress(compressedData, deCompressedData)
	return deCompressedData.Bytes()
}

// compress uses flate to compress a byte slice to a corresponding level
func compress(src []byte, dest io.Writer, level int) error {
	compressor, err := flate.NewWriter(dest, level)
	if nil != err {
		return err
	}
	if _, err := compressor.Write(src); nil != err {
		return err
	}
	return compressor.Close()
}

// compress uses flate to decompress an io.Reader
func decompress(src io.Reader, dest io.Writer) error {
	decompressor := flate.NewReader(src)
	if _, err := io.Copy(dest, decompressor); nil != err {
		return err
	}
	return decompressor.Close()
}

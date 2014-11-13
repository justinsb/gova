package sources

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/justinsb/gova/files"
)

type ByteSource interface {
	Open() (io.ReadCloser, error)
	Size() (int64, error)
}

func ReadToString(byteSource ByteSource) (string, error) {
	in, err := byteSource.Open()
	if err != nil {
		return "", err
	}

	defer in.Close()

	s, err := ioutil.ReadAll(in)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

type FileByteSource struct {
	path string
}

func NewFileByteSource(path string) *FileByteSource {
	self := &FileByteSource{}
	self.path = path
	return self
}

func (self *FileByteSource) Exists() (bool, error) {
	return files.Exists(self.path)
}

func (self *FileByteSource) Open() (io.ReadCloser, error) {
	f, err := os.Open(self.path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (self *FileByteSource) Size() (int64, error) {
	fi, err := os.Stat(self.path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

type ArrayToByteSource struct {
	data []byte
}

func NewArrayToByteSource(data []byte) *ArrayToByteSource {
	self := &ArrayToByteSource{}
	self.data = data
	return self
}

func (self *ArrayToByteSource) Open() (io.ReadCloser, error) {
	reader := bytes.NewReader(self.data)
	return ioutil.NopCloser(reader), nil
}

func (self *ArrayToByteSource) Size() (int64, error) {
	n := len(self.data)
	return int64(n), nil
}
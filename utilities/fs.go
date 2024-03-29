package utilities

import (
	"io"
	"net/http"
	"os"
)

// https://stackoverflow.com/a/49592238/13155318

type JustFilesFilesystem struct {
	FS http.FileSystem
	// readDirBatchSize - configuration parameter for `Readdir` func
	ReadDirBatchSize int
}

func (fs JustFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.FS.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredStatFile{File: f, readDirBatchSize: fs.ReadDirBatchSize}, nil
}

type neuteredStatFile struct {
	http.File
	readDirBatchSize int
}

func (e neuteredStatFile) Stat() (os.FileInfo, error) {
	s, err := e.File.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
	LOOP:
		for {
			fl, err := e.File.Readdir(e.readDirBatchSize)
			switch err {
			case io.EOF:
				break LOOP
			case nil:
				for _, f := range fl {
					if f.Name() == "index.html" {
						return s, err
					}
				}
			default:
				return nil, err
			}
		}
		return nil, os.ErrNotExist
	}
	return s, err
}

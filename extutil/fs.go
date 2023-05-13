package extutil

import (
	"net/http"
	"strings"
)

// https://gist.github.com/hauxe/f88a87f4037bca23f04f6d100f6e08d4#file-http_static_custom_http_server-go

// FileSystem custom file system handler
type FileSystem struct {
	Fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

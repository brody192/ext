package middleware

// adapted from https://github.com/labstack/echo/blob/c0c00e6241a5950075e5c5f12b2e66a42cf0348b/middleware/body_limit.go for use with net/http

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
)

type limitedReader struct {
	limitBytes int64
	reader     io.ReadCloser
	read       int64
	context    context.Context
}

// BodyLimit returns a BodyLimit middleware.
//
// BodyLimit middleware sets the maximum allowed size for a request body, if the size exceeds the configured limit, it
// sends "413 - Request Entity Too Large" response. The BodyLimit is determined based on both `Content-Length` request
// header and actual content read, which makes it super secure.
func LimitBytes(limitBytes int64) func(http.Handler) http.Handler {
	var pool = sync.Pool{New: func() any {
		return &limitedReader{limitBytes: limitBytes}
	}}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.ContentLength > limitBytes {
				http.Error(w, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)
				return
			}

			var p = pool.Get().(*limitedReader)
			p.Reset(r.Context(), r.Body)
			defer pool.Put(p)
			r.Body = p

			next.ServeHTTP(w, r)
		})
	}

}

func (r *limitedReader) Read(b []byte) (n int, err error) {
	n, err = r.reader.Read(b)
	r.read += int64(n)
	if r.read > r.limitBytes {
		return n, errors.New(http.StatusText(http.StatusRequestEntityTooLarge))
	}
	return
}

func (r *limitedReader) Close() error {
	return r.reader.Close()
}

func (r *limitedReader) Reset(context context.Context, reader io.ReadCloser) {
	r.reader = reader
	r.context = context
	r.read = 0
}

package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type dumpMiddleware struct {
	path string
	next http.Handler
}

func NewDumpMiddleware(
	ctx context.Context, path string, next http.Handler,
) (
	http.Handler, error,
) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf(
			"cleaning dir name %q: %w",
			path,
			err,
		)
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		return nil, fmt.Errorf(
			"preparing dump directory %q: %w",
			path,
			err,
		)
	}

	return &dumpMiddleware{
		path: path,
		next: next,
	}, nil
}

func (dump *dumpMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	dumpReq := req.Clone(context.TODO())
	dumpBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("cannot dump body!", err)
		return
	}
	dumpReq.Body = io.NopCloser(bytes.NewBuffer(dumpBody))
	go dump.dumpRequest(dumpReq)

	req.Body = io.NopCloser(bytes.NewBuffer(dumpBody))
	dump.next.ServeHTTP(w, req)
}

func (dump *dumpMiddleware) dumpRequest(req *http.Request) {
	if req.Method != http.MethodPost {
		return
	}

	timeString := time.Now().UTC().Format(time.RFC3339)
	fileName := strings.ReplaceAll(timeString, ":", "-")
	fileName = filepath.Join(dump.path, fileName+".http")

	fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("cannot create dump file!", err)
		return
	}
	defer fd.Close()

	fmt.Fprintf(fd, "### %s\n", timeString)
	fmt.Fprintf(fd, "%s %s\n", req.Method, req.URL)

	for key, values := range req.Header {
		for _, value := range values {
			fmt.Fprintf(fd, "%s: %s\n", key, value)
		}
	}
	fmt.Fprintln(fd)

	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("cannot dump body!", err)
		return
	}

	if len(data) > 0 {
		fmt.Fprintln(fd, string(data))
	}
}

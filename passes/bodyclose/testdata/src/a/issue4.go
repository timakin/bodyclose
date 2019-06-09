package a

import (
	"io"
	"net/http"
)

func issue4_1() {
	resp, _ := http.Get("https://example.com") // want "response body must be closed"

	foo(resp.Body)
}

func foo(r io.ReadCloser) {}

func issue4_2() {
	resp, _ := http.Get("https://example.com") // want "response body must be closed"

	_ = http.MaxBytesReader(nil, resp.Body, 1024)
}

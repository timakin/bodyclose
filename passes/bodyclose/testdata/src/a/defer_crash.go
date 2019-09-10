package a

import (
	"io"
	"net/http"
)

func testNoCrashOnDefer() {
	resp, _ := http.Get("https://example.com") // want "response body must be closed"
	defer func(body io.ReadCloser) {}(resp.Body)
}

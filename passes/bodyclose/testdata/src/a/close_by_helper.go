package a

import (
	"io"
	"net/http"
)

// closeByHelperOK passes resp.Body to a helper that closes it.
func closeByHelperOK() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		return
	}

	defer safeClose("response body", resp.Body)
}

// closeByHelperDeferOK uses a helper that takes an io.Closer.
func closeByHelperDeferOK() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		return
	}

	defer safeCloseCloser(resp.Body)
}

// closeByHelperNotClosed passes resp.Body to a function that does NOT close it.
func closeByHelperNotClosed() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		return
	}

	defer consumeBody(resp.Body)
}

func safeClose(label string, body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		panic(err)
	}
	_ = label
}

func safeCloseCloser(c io.Closer) {
	_ = c.Close()
}

func consumeBody(body io.ReadCloser) {
	_, _ = io.ReadAll(body)
}

package a

import (
	"net/http"
)

func f1() {
	_, _ = http.Get("http://example.com/") // want "response body must be closed"
}

func f2() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close() // OK
}


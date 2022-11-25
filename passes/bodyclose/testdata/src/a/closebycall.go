package a

import (
	"io"
	"net/http"
)

func closeByCall() {
	res, err := http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		panic(err)
	}

	process(res)
}

func closeByCallDeferOK() {
	res, err := http.Get("http://example.com/") // OK
	if err != nil {
		panic(err)
	}

	defer close(res)
}

func process(res *http.Response) {
	_, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
}

func close(res *http.Response) {
	process(res)

	defer res.Body.Close()
}

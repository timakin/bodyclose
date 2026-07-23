package a

import (
	"net/http"

	"handledresponse"
)

//bodyclose:handled
func handledResponse() (*http.Response, error) {
	return http.Get("http://example.com/")
}

func openResponse() (*http.Response, error) {
	return http.Get("http://example.com/")
}

func responseHandledDirectiveCallSites() {
	_, _ = handledResponse()
	_, _ = handledresponse.Get()
	_, _ = openResponse() // want "response body must be closed"
}

package handledresponse

import "net/http"

// bodyclose:handled
func Get() (*http.Response, error) {
	return http.Get("http://example.com/")
}

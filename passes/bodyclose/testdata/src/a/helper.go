package a

import "net/http"

func doRequestWithoutClose() (*http.Response, error) {
	return http.Get("https://example.com")
}

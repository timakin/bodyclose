package a

import (
	"net/http"
)

type MyResponse struct {
	Original *http.Response
}

func (r *MyResponse) Response() *http.Response {
	return r.Original
}

// issue42_1 is case when http.Response from struct field
func issue42_1() {
	r := &MyResponse{}
	r.Original, _ = http.Get("http://example.com/") // OK
	_ = r.Original.Body.Close()
}

// issue42_2 is case when http.Response from a function
func issue42_2() {
	r := &MyResponse{}
	r.Original, _ = http.Get("http://example.com/") // OK
	_ = r.Response().Body.Close()
}

package a

import (
	"net/http"
)

func issue61() {
	var resp *http.Response
	if true {
		resp, _ = http.Get("http://example.com") // OK
	}
	defer resp.Body.Close()
}

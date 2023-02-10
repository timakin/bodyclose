package a

import "net/http"

var resp *http.Response

func issue36_1() {
	resp, _ = http.Get("https://example.com") // OK
	resp.Body.Close()
}

func issue36_2() {
	// Also OK. Responses stored in global variables are not checked.
	resp, _ = http.Get("https://example.com")
}

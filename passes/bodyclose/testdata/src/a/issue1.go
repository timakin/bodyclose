package a

import "net/http"

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}

func main() {
	resp := get()
	resp.Body.Close()
}

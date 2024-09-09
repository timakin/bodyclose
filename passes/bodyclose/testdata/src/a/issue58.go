package a

import (
	"net/http"
)

func issue58_1() {
	resp, _ := http.DefaultClient.Get("https://example.com") // want "response body must be closed"
	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
	}
}

func issue58_2() {
	resp, _ := http.DefaultClient.Get("https://example.com") // want "response body must be closed"
	if resp.StatusCode >= http.StatusInternalServerError {
		defer resp.Body.Close()
	} else if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
	}
}

func issue58_3() {
	resp, _ := http.DefaultClient.Get("https://example.com") // OK
	if resp.StatusCode >= http.StatusInternalServerError {
		defer resp.Body.Close()
	} else if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
	} else {
		defer resp.Body.Close()
	}
}

func issue58_4() {
	resp, err := http.DefaultClient.Get("https://example.com") // OK
	if err != nil {
		// handle error
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		defer resp.Body.Close()
	} else {
		defer resp.Body.Close()
	}
}

func issue58_5() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		panic(err) // handle error
	}
	resp.Body.Close()

	resp2, err := http.Get("http://example.com/") // OK
	if err != nil {
		panic(err) // handle error
	}
	resp2.Body.Close()
}

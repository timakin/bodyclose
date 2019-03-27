package a

import (
	"fmt"
	"net/http"
)

func f1() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		// handle error
	}
	resp.Body.Close()

	resp2, err := http.Get("http://example.com/") // OK
	if err != nil {
		// handle error
	}
	resp2.Body.Close()
}

func f2() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		// handle error
	}
	body := resp.Body
	body.Close()

	resp2, err := http.Get("http://example.com/") // OK
	body2 := resp2.Body
	body2.Close()
	if err != nil {
		// handle error
	}
}

func f3() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}

func f4() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		// handle error
	}
	fmt.Print(resp)

	resp, err = http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		// handle error
	}
	fmt.Print(resp.Status)

	resp, err = http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		// handle error
	}
	fmt.Print(resp.Body)
	return
}

func f5() {
	_, err := http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		// handle error
	}
}

func f6() {
	http.Get("http://example.com/") // want "response body must be closed"
}

func f7() {
	res, _ := http.Get("http://example.com/") // OK
	resCloser := func() {
		res.Body.Close()
	}
	resCloser()
}

func f8() {
	res, _ := http.Get("http://example.com/") // want "response body must be closed"
	_ = func() {
		res.Body.Close()
	}
}

func f9() {
	_ = func() {
		res, _ := http.Get("http://example.com/") // OK
		res.Body.Close()
	}
}

func f10() {
	res, _ := http.Get("http://example.com/") // OK
	resCloser := func(res *http.Response) {
		res.Body.Close()
	}
	resCloser(res)
}

func handleResponse(res *http.Response) {
	res.Body.Close()
}

func f11() {
	res, _ := http.Get("http://example.com/") // OK
	handleResponse(res)
}

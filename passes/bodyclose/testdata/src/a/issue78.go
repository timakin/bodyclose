package a

import (
	"net/http"
)

type CustomClient struct {
	*http.Client
}

func issue78HeadPackage() {
	resp, err := http.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	_ = resp.Header
}

func issue78HeadPackageUnassigned() {
	http.Head("http://example.com/") // OK
}

func issue78HeadPackageDiscard() {
	_, err := http.Head("http://example.com/") // OK
	if err != nil {
		return
	}
}

func issue78HeadClientPointer() {
	client := &http.Client{}
	resp, err := client.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	_ = resp.Header
}

func issue78HeadClientValue() {
	var client http.Client
	resp, err := client.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	_ = resp.Header
}

func issue78HeadDefaultClient() {
	resp, err := http.DefaultClient.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	_ = resp.Header
}

func issue78HeadClientUnassigned() {
	client := &http.Client{}
	client.Head("http://example.com/") // OK
}

func issue78HeadWithClose() {
	resp, err := http.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func issue78HeadCustomClient() {
	cc := &CustomClient{Client: &http.Client{}}
	resp, err := cc.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	_ = resp.Header
}

func issue78HeadReadBody() {
	resp, err := http.Head("http://example.com/") // OK
	if err != nil {
		return
	}
	buf := make([]byte, 10)
	resp.Body.Read(buf)
}

func issue78HeadInClosure() {
	go func() {
		resp, err := http.Head("http://example.com/") // OK
		if err != nil {
			return
		}
		_ = resp.StatusCode
	}()
}

func issue78HeadLoop() {
	for i := 0; i < 3; i++ {
		resp, err := http.Head("http://example.com/") // OK
		if err != nil {
			continue
		}
		_ = resp.StatusCode
	}
}

func issue78GetStillReported() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		return
	}
	_ = resp.Header
}

func issue78ClientGetStillReported() {
	client := &http.Client{}
	resp, err := client.Get("http://example.com/") // want "response body must be closed"
	if err != nil {
		return
	}
	_ = resp.Header
}

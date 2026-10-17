package a

import (
	"fmt"
	"net/http"
)

type HeadClient interface {
	Head(url string) (*http.Response, error)
}

func issue78() {
	// Call http.Head directly without assigning return values
	http.Head("http://example.com") // OK

	// Call http.Head ignoring return values
	_, _ = http.Head("http://example.com") // OK

	// Call http.Head and inspect response without closing body
	resp, err := http.Head("http://example.com") // OK
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)

	// Call http.Head and close body explicitly
	respWithClose, err := http.Head("http://example.com") // OK
	if err != nil {
		return
	}
	defer respWithClose.Body.Close()

	// Call client.Head on *http.Client without closing body
	client := &http.Client{}
	respClient, err := client.Head("http://example.com") // OK
	if err != nil {
		return
	}
	fmt.Println(respClient.Status)

	// Call client.Head directly without assigning
	client.Head("http://example.com") // OK

	// Call client.Head on http.Client value without closing body
	var valClient http.Client
	respValClient, err := valClient.Head("http://example.com") // OK
	if err != nil {
		return
	}
	fmt.Println(respValClient.Header)

	// Call http.DefaultClient.Head without closing body
	respDefault, err := http.DefaultClient.Head("http://example.com") // OK
	if err != nil {
		return
	}
	fmt.Println(respDefault.StatusCode)
}

func issue78Interface(c HeadClient) {
	resp, err := c.Head("http://example.com") // OK
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
}

func issue78RegressionGet() {
	resp, err := http.Get("http://example.com") // want "response body must be closed"
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
}

func issue78RegressionClientGet(client *http.Client) {
	resp, err := client.Get("http://example.com") // want "response body must be closed"
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
}

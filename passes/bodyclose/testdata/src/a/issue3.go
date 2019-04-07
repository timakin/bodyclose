package a

import (
	"fmt"
	"io"
	"net/http"
)

func closeBody(c io.Closer) {
	_ = c.Close()
}

func issue3_1() {
	resp, _ := http.Get("https://example.com")
	defer closeBody(resp.Body)
}

func issue3_2() {
	resp, _ := http.Get("https://example.com")
	defer func() {
		_ = resp.Body.Close()
	}()
}

func issue3_3() {
	resp, err := http.DefaultClient.Do(nil)
	if err != nil {
		// handle err
	}
	defer func() { fmt.Println(resp.Body.Close()) }()
}

func funcReceiver(msg string, er error) {
	fmt.Println(msg)
	if er != nil {
		fmt.Println(er)
	}
}

func issue3_4() {
	resp, _ := http.Get("https://example.com")
	defer func() { funcReceiver("test", resp.Body.Close()) }()
}

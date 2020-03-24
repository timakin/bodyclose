package a

import (
	"io"
	"net/http"
)

func issue27_1(url string) (io.ReadCloser, error) { // body should be closed by User
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

func issue27_2(url string) (io.Closer, error) { // body should be closed by User
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

func issue27_3(url string) (io.Closer, error) { // body should be closed by User
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

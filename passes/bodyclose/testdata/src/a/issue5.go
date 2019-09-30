package a

import (
	"io"
	"io/ioutil"
	"net/http"
)

func f12() {
	res, _ := http.Get("http://example.com/") // OK
	defer func() {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()
}

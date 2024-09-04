package a

import (
	"io"
	_ "net/http"
	"net/http/httptest"
)

func f() {
	w := httptest.NewRecorder()
	resp := w.Result() // OK
	_, _ = io.ReadAll(resp.Body)
}

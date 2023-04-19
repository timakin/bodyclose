package a

import (
	"io"
	"net/http/httptest"
)

func issues11() {
	w := httptest.NewRecorder()
	resp := w.Result()
	defer func() {
		_ = resp.Body.Close()
	}()
	_, _ = io.ReadAll(resp.Body)
}

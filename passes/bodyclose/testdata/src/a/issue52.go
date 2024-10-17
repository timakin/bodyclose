package a

import (
	"io"
	"net/http"
	"time"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	rc := http.NewResponseController(w) // OK
	_ = rc.SetWriteDeadline(time.Time{})
	_, _ = io.Copy(w, r.Body)
}

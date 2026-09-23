package a

import (
	"context"
	"net/http"
)

// Simulates a retryablehttp.CheckRetry policy builder.
// Returns func(context.Context, *http.Response, error) (bool, error).
// call.Type() is *types.Signature — not *http.Response or a tuple containing it.
func NewRetryPolicy(isKnown func(*http.Response, []byte) bool) func(context.Context, *http.Response, error) (bool, error) {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return false, nil
	}
}

func useFunctionTypes() {
	_ = NewRetryPolicy(nil) // OK - returns a function type, not *http.Response

	// Real http call should still be detected
	resp, _ := http.Get("http://example.com/") // want "response body must be closed"
	_ = resp
}

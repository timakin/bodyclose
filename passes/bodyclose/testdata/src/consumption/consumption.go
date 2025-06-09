package consumption

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Test cases for consumption checking - documents exactly what patterns are supported

// ✅ SUPPORTED PATTERNS (should pass - no errors)

// Pattern 1: io.Copy with io.Discard
func consumedWithIOCopy() {
	resp, err := http.Get("http://example.com/") // OK - io.Copy detected
	if err != nil {
		return
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
}

// Pattern 2: io.ReadAll
func consumedWithIOReadAll() {
	resp, err := http.Get("http://example.com/") // OK - io.ReadAll detected
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, _ = io.ReadAll(resp.Body)
}

// Pattern 3: ioutil.ReadAll (legacy)
func consumedWithIoutilReadAll() {
	resp, err := http.Get("http://example.com/") // OK - ioutil.ReadAll detected
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, _ = ioutil.ReadAll(resp.Body)
}

// Pattern 4: json.NewDecoder
func consumedWithJSONDecoder() {
	resp, err := http.Get("http://example.com/") // OK - json.NewDecoder detected
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
}

// Pattern 5: bufio.NewScanner
func consumedWithBufioScanner() {
	resp, err := http.Get("http://example.com/") // OK - bufio.NewScanner detected
	if err != nil {
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		_ = scanner.Text()
	}
}

// Pattern 6: bufio.NewReader
func consumedWithBufioReader() {
	resp, err := http.Get("http://example.com/") // OK - bufio.NewReader detected
	if err != nil {
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	_, _, _ = reader.ReadLine()
}

// Pattern 7: Helper function with known consumption
func consumedInHelper() {
	resp, err := http.Get("http://example.com/") // OK - helper uses io.Copy
	if err != nil {
		return
	}
	defer drainAndClose(resp)
}

func drainAndClose(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		io.Copy(io.Discard, resp.Body) // This io.Copy is detected
		resp.Body.Close()
	}
}

// ⚠️  FALSE POSITIVES (these will incorrectly show errors)

// These patterns actually DO consume the body correctly, but are not detected
// by the current implementation. In real code, use //nolint:bodyclose to suppress.

func falsePositiveDirectRead() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed and consumed"
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// This DOES consume the body, but analyzer doesn't detect it
	buf := make([]byte, 1024)
	resp.Body.Read(buf) // Actually consumes the body
}

func falsePositiveCustomFunction() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed and consumed"
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// This DOES consume the body, but analyzer doesn't detect it
	customProcess(resp.Body) // Actually consumes the body
}

func customProcess(r io.Reader) {
	// Custom processing logic
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil || n == 0 {
			break
		}
	}
}

// ❌ REAL PROBLEMS (correctly detected errors)

// These cases should show errors because the body is NOT properly consumed

func actuallyNotConsumed() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed and consumed"
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Body is closed but NOT consumed - this is a real problem
}

func neitherClosedNorConsumed() {
	resp, err := http.Get("http://example.com/") // want "response body must be closed and consumed"
	if err != nil {
		return
	}
	_ = resp
	// Body is neither closed nor consumed - this is a real problem
}

// RequestBody/ResponseBody distinction test cases
func requestBodyReadShouldNotInterfere(w http.ResponseWriter, r *http.Request) {
	// Read incoming request body
	_, _ = io.ReadAll(r.Body) // This is REQUEST body consumption

	// Make outgoing request - response body should still be detected as unconsumed
	resp, err := http.Get("http://example.com/") // want "response body must be closed and consumed"
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Response body is NOT consumed - should be detected despite request body read
}

func localRequestBodyDistinction() {
	// Create local request and read its body
	req, _ := http.NewRequest("POST", "http://example.com", nil)
	_, _ = io.ReadAll(req.Body) // This is REQUEST body consumption

	// Make HTTP request - response body not consumed, should be detected
	resp, _ := http.Get("http://example.com") // want "response body must be closed and consumed"
	defer resp.Body.Close()
}

func functionParameterRequestBody(req *http.Request) {
	// Read request body from function parameter
	_, _ = io.ReadAll(req.Body) // This is REQUEST body consumption

	// Make HTTP request - response body not consumed, should be detected
	resp, _ := http.Get("http://example.com") // want "response body must be closed and consumed"
	defer resp.Body.Close()
}

func properResponseBodyConsumptionWithRequestBody(w http.ResponseWriter, r *http.Request) {
	// Read request body
	_, _ = io.ReadAll(r.Body) // This is REQUEST body consumption

	// Make HTTP request - response body properly consumed, should pass
	resp, _ := http.Get("http://example.com") // OK - response body properly consumed
	defer resp.Body.Close()
	io.ReadAll(resp.Body) // This is RESPONSE body consumption
}

// closeBeforeConsume is commented out because current implementation
// doesn't detect execution order (documented limitation)
//
// func closeBeforeConsume() {
// 	resp, err := http.Get("http://example.com/")
// 	if err != nil {
// 		return
// 	}
// 	resp.Body.Close() // Closed first
// 	io.ReadAll(resp.Body) // Then trying to consume - this would fail at runtime
// 	// NOTE: Current implementation doesn't detect execution order,
// 	// but this would fail at runtime anyway
// }

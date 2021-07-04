package a

func doRequestInHelperFunc() {
	_, _ = doRequestWithoutClose() // want "response body must be closed"
}

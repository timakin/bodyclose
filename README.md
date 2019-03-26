# bodyclose

[![CircleCI](https://circleci.com/gh/timakin/bodyclose.svg?style=svg)](https://circleci.com/gh/timakin/bodyclose)

`bodyclose` is a static analysis tool which checks whether `res.Body` is correctly closed.

## Install

You can get `bodyclose` by `go get` command.

```bash
$ go get -u github.com/timakin/bodyclose
```

## How to use

`bodyclose` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which bodyclose) github.com/timakin/go_api/...
# github.com/timakin/go_api
internal/httpclient/httpclient.go:13:13: response body must be closed
```

When Go is lower than 1.12, just run `bodyclose` command with the package name (import path).

But it cannot accept some options such as `--tags`.

```bash
$ bodyclose github.com/timakin/go_api/...
~/go/src/github.com/timakin/api/internal/httpclient/httpclient.go:13:13: response body must be closed
```

## Analyzer

`bodyclose` validates whether [*net/http.Response](https://golang.org/pkg/net/http/#Response) of HTTP request calls method `Body.Close()` such as below code.

```go
resp, err := http.Get("http://example.com/") // Wrong case
if err != nil {
	// handle error
}
body, err := ioutil.ReadAll(resp.Body)
```

This code is wrong. You must call resp.Body.Close when finished reading resp.Body.

```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close() // OK
body, err := ioutil.ReadAll(resp.Body)
```

In the [GoDoc of Client.Do](https://golang.org/pkg/net/http/#Client.Do) this rule is clearly described.

If you forget this sentence, a HTTP client cannot re-use a persistent TCP connection to the server for a subsequent "keep-alive" request.

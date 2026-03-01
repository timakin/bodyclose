我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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

### Options

You can enable additional checks with the `-check-consumption` flag to also verify that response bodies are consumed:

```bash
$ go vet -vettool=$(which bodyclose) -check-consumption github.com/timakin/go_api/...
```

#### Supported Consumption Patterns

When `-check-consumption` is enabled, the following patterns are recognized as valid body consumption:

- `io.Copy(io.Discard, resp.Body)`
- `io.ReadAll(resp.Body)`
- `ioutil.ReadAll(resp.Body)` (legacy)
- `json.NewDecoder(resp.Body)`
- `bufio.NewScanner(resp.Body)`
- `bufio.NewReader(resp.Body)`

##### Limitations and False Positives

**Note**: Patterns not listed above may trigger false positives even when the body is properly consumed. Use `//nolint:bodyclose` to suppress warnings for custom consumption patterns that are not automatically detected.

Example of suppressing false positives:
```go
func customBodyProcessing() {
    resp, _ := http.Get("http://example.com/") //nolint:bodyclose
    defer resp.Body.Close()

    // Custom consumption logic that analyzer doesn't recognize
    buf := make([]byte, 1024)
    resp.Body.Read(buf) // This actually consumes the body
}
```

**Limitation**: The analyzer does not detect execution order, so patterns where `Close()` is called before consumption (which would fail at runtime) are not specifically flagged.

### Legacy Usage (Go < 1.12)

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

# www

Simple http client for golang with user-friendly interface.

## Features

- Chainable API
- Direct file upload
- Timeout
- Cookie
- GZIP
- Charset detection
- Cleaned http client

## Installation

```bash
go get github.com/GarryGaller/go-www
```

## Quick Start

```go
package main

import (
    "fmt"
    "net/url"

    "github.com/GarryGaller/go-www"
)

func main() {
    client := www.NewClient()
    req := www.NewRequest(client)
    resp := req.WithQuery(&url.Values{"key": {"value"}}).
        Get("https://httpbin.org/get")

    if resp.Error() != nil {
        fmt.Printf("%v", resp.Error())
    } else {
        fmt.Printf("%s\n", resp.Status)
        fmt.Printf("%s\n", resp.Text())
    }

    // or cleaned client and request in one step
    resp = www.New().
        WithQuery(&url.Values{"key": {"value"}}).
        Get("https://httpbin.org/get")

    fmt.Printf("%s\n", resp.Status)
    fmt.Printf("%s\n", resp.Text())
}
```

## Usage


### Client types

```go

// returns a new http.Client with similar default values to http.Client, but with a non-shared 
// Transport, idle connections disabled, and keepalives disabled.
cleanedClient := www.Cleaned()

//returns a new http.Client with similar default values to
// http.Client, but with a shared Transport. Do not use this function for
// transient clients as it can leak file descriptors over time. Only use this
// for clients that will be re-used for the same host(s).
sharedClient := www.Pooled()

// standard client &http.Client{}
defaultClient := www.Default()

// hand over your client or  or will be used &http.Client{}
client := www.NewClient(...)

// Returns the request object along with the cleaned client
req := www.New()
```


### Sending Request

```go
// get
req.Get("https://httpbin.org/get")

// get with query and headers
req.WithQuery(&url.Values{
        "q": []string{"go", "generics"}})
        .Get("https://httpbin.org/get", 
            http.Header{"User-Agent": {"Mozilla"}},
        )

// post
req.WithForm(&url.Values{"token": {"123456"}}).
    Post("https://httpbin.org/post")

// post file as data
req.WithFile(MustOpen(filePath), "text/plain; charset=utf-8").
    Post("https://httpbin.org/post")

// post file as multipart
req.AttachFile(MustOpen(filePath)).
    Post("https://httpbin.org/post"

// post files(multipart)
req.AttachFiles(map[string]interface{}{
    "file":  {MustOpen(filePath), "text/plain; charset=utf-8"},
    "file2": {MustOpen(filePath2),"text/plain; charset=utf-8"},
    "other": strings.NewReader("hello world!"),
    }).Post("https://httpbin.org/post")

// delete
req.Delete("http://httpbin.org/delete")

// patch
req.Head("http://httpbin.org/patch")

// put
req.Head("http://httpbin.org/put")

```

### Customize Client and Request

Before starting a new HTTP request, you can specify additional client options and add a query string or form data to the request object as well as new headers.
The client object can be used by sharing it between other requests.

```go
client := www.NewClient()
client.WithTimeout(2 * time.Second)
jar, _ := cookiejar.New(nil)
client.WithJar(jar)
fmt.Printf("%s\n", client.Timeout)
fmt.Printf("%#v\n", client.Jar)

client = www.Cleaned().With(2 *time.Second, jar)
fmt.Printf("%s\n", client.Timeout)
fmt.Printf("%#v\n", client.Jar)

req := www.NewRequest(client)
req.WithQuery(&url.Values{"q": {"generics"}, "l":{"go"}, "type":{"topics"}})
resp := req.Get("https://github.com/search",
        http.Header{
             "User-Agent": {"Mozilla"},
            //"Accept": {"application/vnd.github.v3+json"},
            //"Authorization": {"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}},
        })
fmt.Printf("%s\n", resp.Status)
fmt.Printf("%s\n", resp.Headers())
```

### Response

The `www.Response` is a thin wrap of `http.Response`.

```go

// response as text
resp = www.Get("https://httpbin.org/get")
bodyAsString = resp.Text()

// response as bytes
resp = www.Get("https://httpbin.org/get")
bodyAsBytes = resp.Content()

// response as map[key]interface{}
resp = www.WithJson(params).Post("https://httpbin.org/post")
bodyAsMap = resp.Json()

```

### Error Checking

```go

req := NewRequest(client)
if req.Error() != nil {
    fmt.Printf("%v\n", req.Error())
}

resp := req.Get("https://httpbin.org/get")
if resp.Error() != nil {
    fmt.Printf("%v\n", resp.Error())
}
```

### Handle Cookies

```go
// cookies
req.SetCookies(&http.Cookie{
                Name:   "token",
                Value:  "some_token",
                MaxAge: 300,
    }).Get("https://httpbin.org/cookies")
    
fmt.Printf("%s\n", req.Cookies())    
    
```    
    

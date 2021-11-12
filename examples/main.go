package main

import (
    "fmt"
    "net/url"
    "time"
    "net/http/cookiejar"
    
    "github.com/GarryGaller/go-www"
)

func main() {
    client := www.NewClient()
    client.With(2 *time.Second)
    jar, _ := cookiejar.New(nil)
    client.With(jar)
    
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
    fmt.Printf("%s\n", resp.Mime())
    //--------------------------------
    client = www.Cleaned().With(2 *time.Second, jar)
    fmt.Printf("%s\n", client.Timeout)
    fmt.Printf("%#v\n", client.Jar)
    
    resp = www.NewRequest(client).Get("http://www.tim.org")
    fmt.Printf("%s\n", resp.DetectCodePage())
    fmt.Printf("%s\n", resp.Mime())
    fmt.Printf("%s\n", resp.Charset())
    
}

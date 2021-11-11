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

	// or client and request in one step
	resp = www.New().
		WithQuery(&url.Values{"key": {"value"}}).
		Get("https://httpbin.org/get")

	fmt.Printf("%s\n", resp.Status)
	fmt.Printf("%s\n", resp.Text())
}

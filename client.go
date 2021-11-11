package www

// v.0.1.0

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type ClientOptions map[string]interface{}

func (opt *ClientOptions) Merge(other ClientOptions) {
	for key := range other {
		(*opt)[key] = other[key]
	}
}

type StandartClient struct {
	*http.Client
	err error
}

func New() *Request {
	client := NewClient()
	return NewRequest(client)
}

func NewClient(clients ...*http.Client) (cl *StandartClient) {

	var client *http.Client
	if len(clients) == 0 {
		client = &http.Client{}
	} else {
		client = clients[0]
	}

	return &StandartClient{client, nil}
}

func (client *StandartClient) Error() error {
	return client.err
}

func (client *StandartClient) With(options ClientOptions) *StandartClient {

	for key, val := range options {
		switch key {
		case "timeout":
			_val, ok := val.(time.Duration)
			if !ok {
				client.err = fmt.Errorf("value <%s> not time.Duration", key)
				return client
			}
			client.Timeout = _val

		case "transport":
			_val, ok := val.(http.RoundTripper)
			if !ok {
				client.err = fmt.Errorf("value <%s> not http.Transport", key)
				return client
			}
			client.Transport = _val

		case "jar":
			_val, ok := val.(http.CookieJar)
			if !ok {
				client.err = fmt.Errorf("value <%s> not http.CookieJar", key)
				return client
			}
			client.Jar = _val
		}
	}
	return client
}

func (client *StandartClient) WithTimeout(
	timeout time.Duration) *StandartClient {

	client.Timeout = timeout
	return client
}

func (client *StandartClient) WithJar(jar http.CookieJar) *StandartClient {

	client.Jar = jar
	return client
}

func (client *StandartClient) SetCookies(
	host string, cookies ...*http.Cookie) *StandartClient {

	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(nil)
	}

	u, _ := url.Parse(host)
	client.Jar.SetCookies(u, cookies)
	return client
}

func (client *StandartClient) Cookies(host string) []*http.Cookie {
	u, _ := url.Parse(host)
	return client.Jar.Cookies(u)
}

func (client *StandartClient) WithTransport(
	transport http.RoundTripper) *StandartClient {

	client.Transport = transport
	return client
}

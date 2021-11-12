package www

// v.0.2.0

import (
	//"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/hashicorp/go-cleanhttp"
)

type ClientOptions map[string]interface{}

func (opt *ClientOptions) Merge(other ClientOptions) {
	for key := range other {
		(*opt)[key] = other[key]
	}
}

type StandardClient struct {
	*http.Client
	Logger interface{}
	err    error
}

func New() *Request {
	return NewRequest(Cleaned())
}

func Default() (cl *StandardClient) {
	return &StandardClient{
		Client: &http.Client{},
		Logger: defaultLogger,
	}
}

func Cleaned() (cl *StandardClient) {
	return &StandardClient{
		Client: cleanhttp.DefaultClient(),
		Logger: defaultLogger,
	}
}

func Pooled() (cl *StandardClient) {
	return &StandardClient{
		Client: cleanhttp.DefaultPooledClient(),
		Logger: defaultLogger,
	}
}

func NewClient(clients ...*http.Client) (cl *StandardClient) {

	var client *http.Client
	if len(clients) == 0 {
		client = &http.Client{}
	} else {
		client = clients[0]
	}

	return &StandardClient{
		Client: client,
		Logger: defaultLogger,
	}
}

func (client *StandardClient) Error() error {
	return client.err
}

func (client *StandardClient) With(options ...interface{}) *StandardClient {

	for _, option := range options {
		switch option.(type) {
		case time.Duration:
			client.Timeout = option.(time.Duration)

		case http.RoundTripper:
			client.Transport = option.(http.RoundTripper)

		case http.CookieJar:
			client.Jar = option.(http.CookieJar)
		}
	}
	return client
}

func (client *StandardClient) WithTimeout(
	timeout time.Duration) *StandardClient {

	client.Timeout = timeout
	return client
}

func (client *StandardClient) WithJar(jar http.CookieJar) *StandardClient {

	client.Jar = jar
	return client
}

func (client *StandardClient) SetCookies(
	host string, cookies ...*http.Cookie) *StandardClient {

	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(nil)
	}

	u, _ := url.Parse(host)
	client.Jar.SetCookies(u, cookies)
	return client
}

func (client *StandardClient) Cookies(host string) []*http.Cookie {
	u, _ := url.Parse(host)
	return client.Jar.Cookies(u)
}

func (client *StandardClient) WithTransport(
	transport http.RoundTripper) *StandardClient {

	client.Transport = transport
	return client
}

func (client *StandardClient) WithLogger(logger Logger) *StandardClient {

	client.Logger = logger
	return client
}

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

func (c ClientOptions) Merge(other ClientOptions) {
	for key := range other {
		c[key] = other[key]
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

func Default() *StandardClient {
	return &StandardClient{
		Client: &http.Client{},
		Logger: defaultLogger,
	}
}

func Cleaned() *StandardClient {
	return &StandardClient{
		Client: cleanhttp.DefaultClient(),
		Logger: defaultLogger,
	}
}

func Pooled() *StandardClient {
	return &StandardClient{
		Client: cleanhttp.DefaultPooledClient(),
		Logger: defaultLogger,
	}
}

func NewClient(clients ...*http.Client) *StandardClient {
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

func (cl StandardClient) Error() error {
	return cl.err
}

func (cl *StandardClient) With(options ...interface{}) *StandardClient {
	for _, option := range options {
		switch option.(type) {
		case time.Duration:
			cl.Timeout = option.(time.Duration)

		case http.RoundTripper:
			cl.Transport = option.(http.RoundTripper)

		case http.CookieJar:
			cl.Jar = option.(http.CookieJar)
		}
	}
	return cl
}

func (cl *StandardClient) WithTimeout(timeout time.Duration) *StandardClient {
	cl.Timeout = timeout
	return cl
}

func (cl *StandardClient) WithJar(jar http.CookieJar) *StandardClient {
	cl.Jar = jar
	return cl
}

func (cl *StandardClient) SetCookies(host string, cookies ...*http.Cookie) *StandardClient {
	if cl.Jar == nil {
		cl.Jar, _ = cookiejar.New(nil)
	}

	u, _ := url.Parse(host)
	cl.Jar.SetCookies(u, cookies)
	return cl
}

func (cl StandardClient) Cookies(host string) []*http.Cookie {
	u, _ := url.Parse(host)
	return cl.Jar.Cookies(u)
}

func (cl *StandardClient) WithTransport(transport http.RoundTripper) *StandardClient {
	cl.Transport = transport
	return cl
}

func (cl *StandardClient) WithLogger(logger Logger) *StandardClient {
	cl.Logger = logger
	return cl
}

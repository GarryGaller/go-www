package www

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	//"fmt"
	"golang.org/x/net/html/charset"
)

type Response struct {
	*http.Response
	err     error
	content []byte
}

func (resp *Response) Error() error {
	return resp.err
}

func (resp *Response) Content() []byte {
	if resp.content == nil {
		resp.content = resp.readAll()
	}
	return resp.content
}

func (resp *Response) Text() string {
	if resp.content == nil {
		resp.content = resp.readAll()
	}
	return string(resp.content)
}

func (resp *Response) Headers() http.Header {

	return resp.Header
}

func (resp *Response) Json() (data map[string]interface{}) {

	contentType := resp.Header.Get("Content-Type")
	if contentType == "application/json" {
		if resp.content == nil {
			resp.content = resp.readAll()
		}
		if err := json.Unmarshal(resp.content, &data); err != nil {
			resp.err = err
			return
		}
	}
	return
}

func (resp *Response) readAll() (content []byte) {
	var reader io.Reader
	var err error
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			resp.err = err
			return
		}
	default:
		reader = resp.Body
	}

	defer resp.Body.Close()

	reader, err = charset.NewReader(
		reader,
		resp.Header.Get("Content-Type"),
	)
	if err != nil {
		resp.err = err
		return
	}

	content, err = ioutil.ReadAll(reader)
	if err != nil {
		resp.err = err
		return
	}

	return
}

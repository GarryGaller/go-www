package www

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	//"fmt"
	//"golang.org/x/net/html/charset"
	"github.com/softlandia/cpd"
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
		resp.content = resp.readAll(true)
	}

	return string(resp.content)
}
   
func (resp *Response) ContentType(contentTypes ...string) (mime, charset string) {
	var contentType string
	
	if len(contentTypes) > 0 {
		contentType = contentTypes[0]
	} else {    
		contentType = resp.Header.Get("Content-Type")
	}
    cp := strings.Split(contentType, ";")
	mime = cp[0]
    if len(cp) > 1 {
        cp = strings.Split(cp[1], "=")
        if len(cp) > 1 {
            charset = strings.TrimSpace(cp[1])  
        }
    }
    return 
}

func (resp *Response) Mime(contentTypes ...string) (mime string) {
    mime, _ = resp.ContentType(contentTypes...)
    return  
}
 
func (resp *Response) Charset(contentTypes ...string) (charset string) {
    _, charset = resp.ContentType(contentTypes...)
    return  
}



func (resp *Response) DetectCodePage() string {
	if resp.content == nil {
		resp.content = resp.readAll()
	}

	return cpd.CodepageAutoDetect(resp.content).String()
}


func (resp *Response) NewReader() (reader io.Reader) {
	//reader, err := charset.NewReader(reader,resp.Header.Get("Content-Type"))
	reader, err := cpd.NewReader(resp.Body) // need to be tested.
	if err != nil {
		resp.err = err
	}

	return
}

func (resp *Response) Headers() http.Header {

	return resp.Header
}

func (resp *Response) Json() (data map[string]interface{}) {

	contentType := resp.Header.Get("Content-Type")
	if contentType == "application/json" {
		if resp.content == nil {
			resp.content = resp.readAll(true)
		}
		if err := json.Unmarshal(resp.content, &data); err != nil {
			resp.err = err
		}
	}
	return
}

func (resp *Response) JSON() (data map[string]interface{}) {
	return resp.Json()
}

func (resp *Response) readAll(convertToUTF8 ...bool) (content []byte) {
	var reader io.Reader
	var err error
	//contentType := resp.Header.Get("Content-Type")
	contentEncoding := resp.Header.Get("Content-Encoding")

	switch contentEncoding {
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

	if len(convertToUTF8) > 0 && convertToUTF8[0] == true {
		reader = resp.NewReader()
	}

	content, err = ioutil.ReadAll(reader)
	if err != nil {
		resp.err = err
	}

	return
}

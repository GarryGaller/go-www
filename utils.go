package www

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"
)

func MustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func closeReader(r io.Reader, verbose ...bool) bool {
	rc, ok := r.(io.ReadCloser)
	if ok {
		rc.Close()
		if len(verbose) > 0 {
			if x, ok := r.(*os.File); ok {
				fmt.Printf("CLOSE:%s\n", x.Name())
			}
		}
	}

	return ok
}

var quoteEscapists = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscapists.Replace(s)
}

func CreateFormFile(w *multipart.Writer,
	fieldname, filename string,
	contentType ...string) (io.Writer, error) {

	defaultType := "application/octet-stream"
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname),
			escapeQuotes(filename)))
	if len(contentType) > 0 {
		defaultType = contentType[0]
	}
	h.Set("Content-Type", defaultType)
	return w.CreatePart(h)
}

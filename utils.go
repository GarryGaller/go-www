package www

import (
	"fmt"
	"io"
	"os"
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

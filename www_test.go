package www

import (
    "io"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "strings"
    "testing"
    "time"
)

func TestWWW(t *testing.T) {

    headers := http.Header{"User-Agent": {"Mozilla"}}
    params := &url.Values{"key": {"value"}}
    data := &url.Values{"key2": {"value2"}}

    fileName := `Эдгар Аллан По Сердце-обличитель.txt`
    fileName2 := `Edgar Allan Poe The Cask of Amontillado.txt`
    filePath := `testdata\` + fileName
    filePath2 := `testdata\` + fileName2

    cl := NewClient().WithTimeout(2 * time.Second)

    t.Run("DELETE", func(t *testing.T) {

        r := NewRequest(cl)

        resp := r.Delete("https://httpbin.org/delete", http.Header{
            "User-Agent": {"Mozilla"},
            "Accept":     {"application/json"},
        })

        if resp.Error() != nil {
            t.Errorf("%v", resp.Error())
        } else {
            if resp.StatusCode != 200 {
                t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
            } else {
                t.Logf("%s", resp.Status)
                t.Logf("%s", resp.Json())
                t.Logf("%s", r.Headers())
            }
        }
    })

    t.Run("PATCH", func(t *testing.T) {

        r := NewRequest(cl)

        resp := r.Patch("https://httpbin.org/patch", http.Header{
            "User-Agent": {"Mozilla"},
            "Accept":     {"application/json"},
        })

        if resp.Error() != nil {
            t.Errorf("%v", resp.Error())
        } else {
            if resp.StatusCode != 200 {
                t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
            } else {
                t.Logf("%s", resp.Status)
                t.Logf("%s", resp.Json())
                t.Logf("%s", r.Headers())
            }
        }
    })

    t.Run("PUT", func(t *testing.T) {

        r := NewRequest(cl)

        resp := r.Put("https://httpbin.org/put", http.Header{
            "User-Agent": {"Mozilla"},
            "Accept":     {"application/json"},
        })

        if resp.Error() != nil {
            t.Errorf("%v", resp.Error())
        } else {
            if resp.StatusCode != 200 {
                t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
            } else {
                t.Logf("%s", resp.Status)
                t.Logf("%s", resp.Json())
                t.Logf("%s", r.Headers())
            }
        }
    })

    t.Run("GET", func(t *testing.T) {

        r := NewRequest(cl)
        resp := r.WithQuery(params).
            Get("https://httpbin.org/get", headers)

        if resp.Error() != nil {
            t.Errorf("%v", resp.Error())
        } else {
            if resp.StatusCode != 200 {
                t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
            } else {
                t.Logf("%s", resp.Status)
                t.Logf("%s", resp.Text())
                t.Logf("%s", r.Headers())
            }
        }
    })

    t.Run("POST", func(t *testing.T) {

        t.Run("FORM", func(t *testing.T) {
            r := NewRequest(cl)
            resp := r.WithQuery(params).
                WithForm(data).
                Post("https://httpbin.org/post", headers)

            if resp.Error() != nil {
                t.Errorf("%v", resp.Error())
            } else {
                if resp.StatusCode != 200 {
                    t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
                } else {
                    t.Logf("%s", resp.Status)
                    t.Logf("%s", resp.Text())
                    t.Logf("%s", r.Headers())
                }
            }
        })

        t.Run("JSON", func(t *testing.T) {
            r := NewRequest(cl)
            resp := r.WithQuery(params).
                WithJson(params).
                Post("https://httpbin.org/post", headers)

            if resp.Error() != nil {
                t.Errorf("%v", resp.Error())
            } else {
                if resp.StatusCode != 200 {
                    t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
                } else {
                    t.Logf("%s", resp.Status)
                    t.Logf("%s", resp.Json())
                    t.Logf("%s", r.Headers())
                }
            }
        })

        t.Run("FILE", func(t *testing.T) {

            reader := MustOpen(filePath)

            r := NewRequest(cl)
            resp := r.WithQuery(params).
                WithFile(reader).
                Post("https://httpbin.org/post", headers)

            if resp.Error() != nil {
                t.Errorf("%v", resp.Error())
            } else {
                if resp.StatusCode != 200 {
                    t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
                } else {
                    t.Logf("%s", resp.Status)
                    t.Logf("%s", resp.Text())
                    t.Logf("%s", r.Headers())
                }
            }
        })

        t.Run("AttachFile", func(t *testing.T) {
            
            r := NewRequest(cl)
            resp := r.AttachFile(MustOpen(filePath)).
                Post("https://httpbin.org/post", headers)

            if resp.Error() != nil {
                t.Errorf("%v", resp.Error())
            } else {
                if resp.StatusCode != 200 {
                    t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
                } else {
                    t.Logf("%s", resp.Status)
                    t.Logf("%s", resp.Text())
                    t.Logf("%s", r.Headers())
                }
            }
        })

        t.Run("AttachFiles", func(t *testing.T) {

            values := map[string]io.Reader{
                "file":  MustOpen(filePath),
                "file2": MustOpen(filePath2),
                "other": strings.NewReader("hello world!"),
            }

            r := NewRequest(cl)
            resp := r.AttachFiles(values).
                Post("https://httpbin.org/post", headers)

            if resp.Error() != nil {
                t.Errorf("%v", resp.Error())
            } else {
                if resp.StatusCode != 200 {
                    t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
                } else {
                    t.Logf("%s", resp.Status)
                    t.Logf("%s", resp.Text())
                    t.Logf("%s", r.Headers())
                }
            }
        })

    })

    t.Run("COOKIES", func(t *testing.T) {

        jar, _ := cookiejar.New(nil)
        cl.WithJar(jar).SetCookies("https://httpbin.org/",
            &http.Cookie{
                Name:   "token",
                Value:  "some_token",
                MaxAge: 300,
            },
        )
        r := NewRequest(cl)

        cookies := []*http.Cookie{
            {
                Name:   "token1",
                Value:  "some_token1",
                MaxAge: 300,
            },

            {
                Name:   "token2",
                Value:  "some_token2",
                MaxAge: 300,
            },
        }
        resp := r.SetCookies(cookies...).
            Get("https://httpbin.org/cookies", headers)

        if resp.Error() != nil {
            t.Errorf("%v", resp.Error())
        } else {
            if resp.StatusCode != 200 {
                t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
            } else {
                t.Logf("%s", resp.Status)
                t.Logf("%s", resp.Text())
                t.Logf("%s", r.Cookies()) // returns the cookies that are sent with the header Cookie
                t.Logf("%s", r.Headers().Get("Cookie"))
            }
        }
    })

    t.Run("SET COOKIES", func(t *testing.T) {

        jar, _ := cookiejar.New(nil)
        cl.WithJar(jar)
        r := NewRequest(cl)

        resp := r.WithQuery(&url.Values{"name": {"token"}}).
            Get("https://httpbin.org/cookies/set", headers)

        if resp.Error() != nil {
            t.Errorf("%v", resp.Error())
        } else {
            if resp.StatusCode != 200 {
                t.Errorf("StatusCode:got %d, want 200", resp.StatusCode)
            } else {
                t.Logf("%s", resp.Status)
                t.Logf("%s", resp.Text())
                t.Logf("%s", resp.Cookies()) // returns the cookies set in the Set-Cookie headers
                t.Logf("%s", resp.Headers().Get("Set-Cookie"))
            }
        }
    })

    t.Run("GITHUB", func(t *testing.T) {

        client := NewClient()
        client.WithTimeout(2 * time.Second)
        jar, _ := cookiejar.New(nil)
        client.WithJar(jar)

        req := NewRequest(client)
        req.SetCookies(&http.Cookie{
            Name:   "Garry",
            Value:  "Galler",
            MaxAge: 300,
        })
        req.WithQuery(&url.Values{"q": {"generics"}, "l":{"go"}, "type":{"topics"}})
        resp := req.Get("https://github.com/search",
                http.Header{
                     "User-Agent": {"Mozilla"},
                    //"Accept": {"application/vnd.github.v3+json"},
                    //"Authorization": {"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}},
                })
        t.Logf("%s", resp.Status)
        t.Logf("%s", resp.Headers())
    })

}

func BenchmarkWWW(b *testing.B) {

    headers := http.Header{"User-Agent": {"Mozilla"}}

    fileName := `Эдгар Аллан По Сердце-обличитель.txt`
    filePath := `testdata\` + fileName

    cl := NewClient().WithTimeout(2 * time.Second)

    b.RunParallel(func(pb *testing.PB) {

        for pb.Next() {

            r := NewRequest(cl)
            resp := r.AttachFile(MustOpen(filePath)).
                Post("https://httpbin.org/post", headers)

            if resp.Error() != nil {
                b.Errorf("%v\n", resp.Error())
            } else {
                b.Logf("status:%s", resp.Status)
            }
        }
    })
}

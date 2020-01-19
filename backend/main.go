package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func newDirector(origin url.URL) func(*http.Request) {
	return func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}
}

func newReplacer(orig, replace string) func(resp *http.Response) error {
	return func(resp *http.Response) error {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = resp.Body.Close()
		if err != nil {
			return err
		}

		b = bytes.Replace(b, []byte(orig), []byte(replace), -1)
		body := ioutil.NopCloser(bytes.NewReader(b))

		resp.Body = body
		resp.ContentLength = int64(len(b))
		resp.Header.Set("Content-Length", strconv.Itoa(len(b)))

		return nil
	}
}

func Frontend(w http.ResponseWriter, r *http.Request) {
	origin, _ := url.Parse("http://localhost:5000/")
	director := newDirector(*origin)
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}

func liverload_js(w http.ResponseWriter, r *http.Request) {
	origin, _ := url.Parse("http://localhost:35729/")
	director := newDirector(*origin)
	modifier := newReplacer("this.port = 35729;", "this.port = 443;")
	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifier}
	proxy.ServeHTTP(w, r)
}

func liverload_ws(w http.ResponseWriter, r *http.Request) {
	origin, _ := url.Parse("http://localhost:35729/")
	director := newDirector(*origin)
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}

func Bundle_js(w http.ResponseWriter, r *http.Request) {
	origin, _ := url.Parse("http://localhost:5000/")
	director := newDirector(*origin)
	modifier := newReplacer(":35729/livereload.js?snipver=1", ":443/livereload.js?snipver=1")
	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifier}
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/build/bundle.js", Bundle_js)
	http.HandleFunc("/livereload.js", liverload_js)
	http.HandleFunc("/livereload", liverload_ws)
	http.HandleFunc("/", Frontend)
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}

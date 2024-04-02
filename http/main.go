package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

// my personal http client

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

// RoundTrip - This is Middleware method
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {

	//jar, err := cookiejar.New(nil)
	//jar.SetCookies()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { // Check redirect
			fmt.Println(req.Response.Status)
			fmt.Println("REDIRECT")
			return nil
		},
		Transport: &loggingRoundTripper{ // Middleware
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
	}

	resp, err := client.Get("https://google.org")
	//resp, err := http.DefaultClient.Get("https://google.org")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	fmt.Println("REQUEST STATUS", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", body)

	fmt.Println(runtime.GOMAXPROCS(0))
}

package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func formatRequest(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	return strings.Join(request, "\n")
}

func NewErgoProxy(config *Config) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		fmt.Printf("request: %v %v %v \n", req.Method, req.URL, req.Proto)

		if config.Verbose {
			fmt.Println(formatRequest(req))
		}

		service := config.GetService(req.URL.Host)
		if service != nil {
			target, _ := url.Parse(service.Url)
			targetQuery := target.RawQuery

			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)

			if targetQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = targetQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
			}

			if _, ok := req.Header["User-Agent"]; !ok {
				req.Header.Set("User-Agent", "")
			}
		}
	}

	return &httputil.ReverseProxy{Director: director}
}

func ServeProxy(config *Config) {
	http.HandleFunc("/proxy.pac", func(w http.ResponseWriter, r *http.Request) {
		content := `
		function FindProxyForURL (url, host) {
			if (dnsDomainIs(host, '` + config.Domain + `')) {
				return 'PROXY 127.0.0.1:` + config.Port + `';
			}

			return 'DIRECT';
		}
		`
		w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")
		w.Write([]byte(content))
	})

	http.Handle("/", NewErgoProxy(config))

	http.ListenAndServe(":"+config.Port, nil)
}

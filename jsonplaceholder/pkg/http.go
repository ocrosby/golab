package pkg

//go:generate mockgen -destination=./mock_http_client.go -package=pkg github.com/ocrosby/golab/pkg IHttpClient

import "net/http"

type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

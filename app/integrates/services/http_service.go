package services

import (
	"github.com/gofiber/fiber/v2/log"
	"io"
	"n4a3/clean-architecture/app/base"
	"net/http"
	"time"
)

type HttpService interface {
	HttpGet() string
}

type httpService struct {
}

func NewHttpService() HttpService {
	return &httpService{}
}

func (s httpService) HttpGet() string {
	client := getHttpClient()
	req, err := http.NewRequest("GET", "http://localhost:8083/api/master-data/service-provider-method", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error on request", err)
	}
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr, Timeout: time.Second * 5}
	return client
}

func getHttpRequest(client *http.Client, url string) base.Either[string, base.ErrContext] {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {

	}
	if resp != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {

		}
		return base.RightEither[string, base.ErrContext](string(body))
	}
	return base.RightEither[string, base.ErrContext]("")
}

package services

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"n4a3/clean-architecture/app/core"
	"net/http"
	"time"
)

type HttpMethod string

const (
	HttpGET    HttpMethod = "GET"
	HttpPOST   HttpMethod = "POST"
	HttpPUT    HttpMethod = "PUT"
	HttpPATCH  HttpMethod = "PATCH"
	HttpDELETE HttpMethod = "DELETE"
)

type HttpService interface {
	GetHttpRequest(url string, headers map[string]string) core.Either[[]byte, core.ErrContext]
	PostHttpRequest(url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext]
	PutHttpRequest(url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext]
	DeleteHttpRequest(url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext]
}

type httpService struct {
}

func NewHttpService() HttpService {
	return &httpService{}
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

func (s httpService) GetHttpRequest(url string, headers map[string]string) core.Either[[]byte, core.ErrContext] {
	client := getHttpClient()
	resp := httpRequest(client, HttpGET, url, nil, nil)
	return resp
}

func (s httpService) PostHttpRequest(url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext] {
	client := getHttpClient()
	return httpRequest(client, HttpPOST, url, nil, body)
}

func (s httpService) PutHttpRequest(url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext] {
	client := getHttpClient()
	return httpRequest(client, HttpPUT, url, nil, body)
}

func (s httpService) DeleteHttpRequest(url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext] {
	client := getHttpClient()
	return httpRequest(client, HttpDELETE, url, nil, body)
}

func httpRequest(client *http.Client, method HttpMethod, url string, headers map[string]string, body any) core.Either[[]byte, core.ErrContext] {
	b, err := json.Marshal(body)
	req, err := http.NewRequest(string(method), url, bytes.NewBuffer(b))
	if err != nil {
		return core.LeftEither[[]byte, core.ErrContext](core.NewErrorWithCode(core.Integration, err))
	}
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return core.LeftEither[[]byte, core.ErrContext](core.NewErrorWithCode(core.Integration, err))
	}
	if resp != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return core.LeftEither[[]byte, core.ErrContext](core.NewErrorWithMsg(core.Integration, "cannot read response body", err))
		}
		return core.RightEither[[]byte, core.ErrContext](body)
	}
	return core.RightEither[[]byte, core.ErrContext](nil)
}

func (s httpService) HttpGet(url string) string {
	client := getHttpClient()
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error on request", err)
		return ""
	}
	if resp.Body == nil {
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return string(body)
	}
	return ""
}

package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	jsonrpcversion = "2.0"
	contenttype    = "application/json"
)

type Service struct {
	client *Client
}

type Client struct {
	client *http.Client
	url    *url.URL
}

type request struct {
	JsonRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	ID      string   `json:"id"`
	Params  []string `json:"params"`
}

func newClient(rawurl string) (*Client, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	c := Client{
		client: http.DefaultClient,
		url:    url,
	}
	return &c, nil
}

func (c Client) send(contentType string, msg io.Reader) (*http.Response, error) {
	res, err := c.client.Post(c.url.String(), contentType, msg)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c Client) prepBody(method string, params ...string) (io.Reader, error) {
	body := request{
		JsonRPC: jsonrpcversion,
		Method:  method,
		ID:      "",
		Params:  []string{},
	}
	for _, param := range params {
		body.Params = append(body.Params, param)
	}
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(buf)
	return r, nil
}

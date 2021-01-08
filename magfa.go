package magfa

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is magfa service client that has multiple methods to communicate with magfa services
type Client struct {
	httpClient *http.Client
	baseUrl    string
	username   string
	password   string
	domain     string
}

// New magfa client
func New(username, domain, password string, timeout time.Duration) (*Client, error) {
	httpClient := &http.Client{
		Timeout: timeout,
	}

	return &Client{httpClient,
		"https://sms.magfa.com/api/http/sms/v2/",
		username,
		password,
		domain,
	}, nil
}

// SetBaseUrl to magfa client for change url address
func (c *Client) SetBaseUrl(url string) {
	c.baseUrl = url
}

func (c *Client) sendRequest(path string, method string, reqBody []byte) ([]byte, error) {
	url := c.baseUrl + path

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, fmt.Errorf("create request for ibsng failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("Cache-Control", "no-cache")

	reqCtx, cancel := context.WithTimeout(context.Background(), time.Minute)

	req = req.WithContext(reqCtx)
	req.SetBasicAuth(c.username+"/"+c.domain, c.password)

	defer cancel()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request to magfa failed: %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read magfa response body failed: %s", err.Error())
	}

	return body, nil
}

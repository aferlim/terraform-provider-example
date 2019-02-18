package iacitem

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// BaseClient is the base Client for items
type BaseClient struct {
	hostname   string
	port       int
	authToken  string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, port int, token string) *BaseClient {

	return &BaseClient{
		hostname:   hostname,
		port:       port,
		authToken:  token,
		httpClient: &http.Client{},
	}

}

func (c *BaseClient) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {

	req, err := http.NewRequest(method, c.requestPath(path), &body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.authToken)

	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}

	return resp.Body, nil
}

func (c *BaseClient) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
}

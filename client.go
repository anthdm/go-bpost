package bpost

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
)

const defaultBaseURL = "https://api.bpost.be/services/shm/"

// Client manages communication with the BPOST API.
type Client struct {
	// HTTP client to communicate with the request API.
	*http.Client

	// base URL for API requests.
	baseURL string

	// accountID used to authenticate against API requests.
	accountID string

	// passPhrase used to make a base64 string, set in the authorization header.
	passPhrase string
}

// NewClient returns a new BPOST API client.
func NewClient(httpClient *http.Client, id, passPhrase string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		Client:     httpClient,
		baseURL:    defaultBaseURL,
		accountID:  id,
		passPhrase: passPhrase,
	}
	return c
}

func (c *Client) NewRequest(method, url string, body interface{}) (*http.Request, error) {
	reqURL := path.Join(c.accountID, url)
	reqURL = c.baseURL + reqURL

	buf := &bytes.Buffer{}
	if err := xml.NewEncoder(buf).Encode(&body); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, reqURL, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.bpost.shm-order-v3+XML")

	token := fmt.Sprintf("%s:%s", c.accountID, c.passPhrase)
	base64Str := encodeBase64([]byte(token))
	req.Header.Set("Authorization", "Basic "+base64Str)
	return req, nil
}

func encodeBase64(p []byte) string {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(p)))
	base64.StdEncoding.Encode(base64Text, p)
	return string(base64Text)
}

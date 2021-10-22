package zoom

import (
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Client struct {
	endpoint  string
	secretKey string
	apiKey    string
}

var (
	once       sync.Once
	client     *Client
	httpClient = http.Client{Timeout: time.Second * 5, Transport: http.DefaultTransport}
)

func NewClient(apiKey, secretKey string) *Client {
	once.Do(func() {
		client = &Client{
			apiKey: apiKey, secretKey: secretKey, endpoint: API_ENDPOINT,
		}
	})
	return client
}

func jwtToken(key, secret string) (string, error) {
	standardClaims := jwt.StandardClaims{
		Issuer:    key,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(300)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)
	return token.SignedString([]byte(secret))
}

func (c *Client) createRequest(path, method string, data io.Reader) (*[]byte, error) {
	jwt, err := jwtToken(c.apiKey, c.secretKey)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, c.endpoint+path, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwt)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err2 := checkError(body)
	return &body, err2
}

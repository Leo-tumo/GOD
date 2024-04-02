package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const Url = "https://api.coincap.io/v2/assets"

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets() ([]Asset, error) {
	//jar, err := cookiejar.New(nil)
	//jar.SetCookies()

	resp, err := c.client.Get(Url)
	//resp, err := http.DefaultClient.Get("https://google.org")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r assetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	return r.Assets, nil

	//for _, asset := range r.Data {
	//	fmt.Println(asset)
	//}
	//
	//fmt.Println(time.Unix(r.Timestamp/1000, 0).Format("02/01/2006 15:04:05"))
}

func (c Client) GetAsset(name string) (Asset, error) {
	url := fmt.Sprintf("%s/%s", Url, name)
	resp, err := c.client.Get(url)

	if err != nil {
		return Asset{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r assetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	return r.Asset, nil
}

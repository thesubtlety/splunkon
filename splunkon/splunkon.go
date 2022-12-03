package splunkon

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Client struct {
	username string
	password string
	baseURL  string
}

func New(username string, password string, baseURL url.URL) *Client {
	return &Client{
		username: username,
		password: password,
		baseURL:  baseURL.String(),
	}
}

// func (client Client) Get(path string) (*http.Response, error) {

// 	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

// 	url := fmt.Sprintf("%s/%s", client.baseURL, path)
// 	//fmt.Printf("[dbg] %s\n", url)

// 	c := &http.Client{
// 		Timeout: time.Second * 5,
// 	}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s", err.Error())
// 	}
// 	req.Header.Set("User-Agent", "splunk-sdk-python/1.6.42")
// 	req.SetBasicAuth(client.username, client.password)

// 	response, err := c.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s", err.Error())
// 	}

// 	//bodyBytes, _ := io.ReadAll(response.Body)
// 	//fmt.Printf("[dbg] body %+s\n", string(bodyBytes))

// 	if response.StatusCode == http.StatusForbidden {
// 		log.Fatal("Forbidden")
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("Unexpected status, %v", response.StatusCode)
// 	}

// 	return response, nil
// }

func (client Client) Get(path string) ([]byte, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	uri, err := url.JoinPath(client.baseURL, path)
	if err != nil {
		return nil, err
	}
	uri, err = url.QueryUnescape(uri)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	req.Header.Set("User-Agent", "splunk-sdk-python/1.6.42")
	req.SetBasicAuth(client.username, client.password)

	response, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	if err := write(bodyBytes); err != nil {
		log.Fatal(err)
	}

	return bodyBytes, nil
}

func write(b []byte) error {
	f, err := os.OpenFile("splunkon.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}
	return nil
}

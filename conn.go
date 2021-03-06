package arin

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func httpClient() *http.Client {
	var d = &net.Dialer{
		Timeout: 5 * time.Second,
	}

	var tr = &http.Transport{
		Dial:                d.Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}
}

func makeRequest(resource, handle string) string {
	url := fmt.Sprintf("http://whois.arin.net/rest/%s/%s", resource, handle)

	return request(url)
}

func makeSubRequest(resource, handle, sub string) string {
	url := fmt.Sprintf("http://whois.arin.net/rest/%s/%s/%s", resource, handle, sub)

	return request(url)
}

func request(url string) string {
	client := httpClient()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Request Error: %v\n", err)
		return ""
	}

	req.Header.Set("Accept", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Response Error: %v\n", err)
		return ""
	}

	if resp.StatusCode == 404 {
		return ""
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Parse Error: %v\n", err)
		return ""
	}

	return string(data)
}

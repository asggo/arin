package arin

import (
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "time"
    "regexp"
    "strings"
)

var reKeyVal = regexp.MustCompile("(.*): +(.*)")

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
    client := httpClient()
    url := fmt.Sprintf("http://whois.arin.net/rest/%s/%s", resource, handle)

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

    defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Printf("Parse Error: %v\n", err)
		return ""
	}

    return string(data)
}

func parseRecord(record string) map[string]string {
    recMap := make(map[string]string)
    matches := reKeyVal.FindAllStringSubmatch(record, -1)

    for _, match := range matches {
        k, v := match[1], strings.Trim(match[2], "\r")

        // If the key is already in the map then append the data to the current
        // value.
        val, ok := recMap[k]
        switch {
        case ok:
            val = val + "\n" + v
            recMap[k] = val
        default:
            recMap[k] = v
        }
    }

    return recMap
}

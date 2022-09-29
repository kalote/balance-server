package service

import (
	"bytes"
	"fmt"
	"net/http"
)

func GetBalance(url string, payload string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(payload))
	if err != nil {
		return nil, fmt.Errorf("HTTP Client init error")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP Client request error: %s", err)
	}

	return resp, nil
}

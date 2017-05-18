package utils

import (
	"net/http"
	"log"
)

func SendHttpRequest(method string, url string) {

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}

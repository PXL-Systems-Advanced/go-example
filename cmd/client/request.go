package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var port string = "3000"
var requestUrl string // Constructed dynamically in ping()

func ping() error {
	// Construct the request URL dynamically
	requestUrl = fmt.Sprintf("%s:%s?s=", baseURL, port)

	res, err := http.Get(requestUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close() // Close response body to avoid leaks

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to establish server connection")
	}
	return nil
}

func getTranslation(msg string) string {
	translationRequest := requestUrl + url.QueryEscape(msg)
	response, err := http.Get(translationRequest)
	if err != nil {
		fmt.Printf("[error] http request failed: %s", err)
		os.Exit(1)
	}
	defer response.Body.Close() // Close response body to avoid leaks

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[error] failed to parse server response: %s", err)
	}
	return string(body)
}

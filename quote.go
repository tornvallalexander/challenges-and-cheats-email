package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Quote string `json:"q"`
	Author string `json:"a"`
	HTML string `json:"h"`
}

const QuoteAPI = "https://zenquotes.io/api/random"

func GetRandomQuote() string {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get(QuoteAPI)
	if err != nil {
		log.Fatalf("api not responding before timeout: %s", err)
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not read response body: %s", err)
	}

	// convert body to type string
	var quoteResponse []Response // we use []Response instead of Response because json is array
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		log.Fatalf("unable to unmarsh body content %s", err)
	}

	return quoteResponse[0].HTML
}

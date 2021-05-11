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

// GetRandomQuote returns a response with the quote struct
func GetRandomQuote() Response {
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
	var quoteResponse []Response
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		log.Fatalf("unable to unmarshal body content: %s", err)
	}

	return quoteResponse[0]
}

package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Djay-M/goCrawler/helpers"
)

var counter = 0
var waitGroup sync.WaitGroup
var mux sync.Mutex

type CrawlerControllerBody struct {
	Websites []string `json:"websites"`
}

func FetchServerStatus(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode("Server is up and running")
}

func CrawlerController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	responseDataChan := make(chan helpers.ResponseObj)
	responseBody := []helpers.ResponseObj{}

	var reqBody CrawlerControllerBody
	json.NewDecoder(request.Body).Decode(&reqBody)

	websiteCount := len(reqBody.Websites)
	resCount := 0
	for _, website := range reqBody.Websites {
		go helpers.CrawlerWebHelper(website, responseDataChan)
	}

	for res := range responseDataChan {
		resCount += 1
		responseBody = append(responseBody, res)

		if resCount >= websiteCount {
			break
		}
	}

	close(responseDataChan)

	json.NewEncoder(response).Encode(responseBody)
}

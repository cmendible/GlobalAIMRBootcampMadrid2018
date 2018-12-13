package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Replace the subscriptionKey string value with your valid subscription key
	const subscriptionKey = "<cognitive services key>"

	/*
	   Replace or verify the region.

	   You must use the same region in your REST API call as you used to obtain your access keys.
	   For example, if you obtained your access keys from the westus region, replace
	   "westcentralus" in the URI below with "westus".

	   NOTE: Free trial access keys are generated in the westcentralus region, so if you are using
	   a free trial access key, you should not need to change this region.
	*/
	const uriBase = "https://westeurope.api.cognitive.microsoft.com"
	const uriPath = "/text/analytics/v2.0/sentiment"

	const uri = uriBase + uriPath

	data := []map[string]string{
		{"id": "1", "language": "en", "text": "You guys didn't want me to talk in english."},
		{"id": "2", "language": "es", "text": "Que pasada de evento! Excelente!"},
	}

	documents, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		return
	}

	r := strings.NewReader("{\"documents\": " + string(documents) + "}")

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var f interface{}
	json.Unmarshal(body, &f)

	jsonFormatted, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		fmt.Printf("Error producing JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonFormatted))
}

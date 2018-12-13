package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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
	const uriPath = "/vision/v2.0/ocr"
	const requestParameters = "language=unk&detectOrientation=true"

	const uri = uriBase + uriPath + "?" + requestParameters

	imagePath := "C:\\Users\\cmend\\Pictures\\Capture.png"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Error opening file %v", err)
	}
	defer f.Close()
	fw, err := w.CreateFormFile("image", imagePath)
	if err != nil {
		fmt.Printf("Error reading file %v", err)
	}
	if _, err = io.Copy(fw, f); err != nil {
		fmt.Printf("Error copying file %v", err)
	}
	w.Close()

	req, err := http.NewRequest("POST", uri, &b)
	if err != nil {
		fmt.Printf("Error calling api %v", err)
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error calling api %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var v interface{}
	json.Unmarshal(body, &v)

	jsonFormatted, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Error producing JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonFormatted))
}

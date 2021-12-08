package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getPriceURL() string {
	host, ok := os.LookupEnv("PRICESERVICE_HOST")
	url := ""

	if !ok {
		url = "http://localhost:5000/"
		fmt.Println("PRICESERVICE_HOST is not present")
	} else {
		url = "http://" + host + ":5000/"
		fmt.Println("PRICESERVICE_HOST: " + host)
	}

	return url
}

func getpricehome(r *http.Request) []byte {

	url := getPriceURL()

	//Prepare request
	req, err := http.NewRequest(
		"GET",
		url,
		bytes.NewBufferString("Body"),
	)
	req = req.WithContext(r.Context())
	if err != nil {
		log.Fatalf("failed to create the request: %v", err)
	}

	//response, err := http.Get(url)
	//client := &http.Client{}
	client := getHttpClient()

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		// Write error to response body
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func getbtcprice(r *http.Request) []byte {

	baseurl := getPriceURL()
	url := baseurl + "price/"

	//Prepare request
	req, err := http.NewRequest(
		"GET",
		url,
		bytes.NewBufferString("Body"),
	)
	req = req.WithContext(r.Context())
	if err != nil {
		log.Fatalf("failed to create the request: %v", err)
	}

	//response, err := http.Get(url)
	//client := &http.Client{}
	client := getHttpClient()

	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		// Write error to response body
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

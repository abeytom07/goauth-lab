package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/Traceableai/goagent/instrumentation/net/traceablehttp"
)

func getURL() string {
	host, ok := os.LookupEnv("DATASERVICE_HOST")
	url := ""

	if !ok {
		url = "http://localhost:8080/"
		fmt.Println("DATASERVICE_HOST is not present")
	} else {
		url = "http://" + host + ":8080/"
		fmt.Println("DATASERVICE_HOST: " + host)
	}

	return url
}


func getTraceableHttpClient() http.Client {
	return http.Client{
		Transport: traceablehttp.NewTransport(http.DefaultTransport),
	}
}

func getdata(id string, r *http.Request) []byte {

	baseurl := getURL()
	url := baseurl + "customer/test/" + id

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
	//client := &http.Client{}
	client := getTraceableHttpClient()

	//response, err := http.Get(url)
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

func getallcustomers(r *http.Request) []byte {

	baseurl := getURL()
	url := baseurl + "customer/all"

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
	client := getTraceableHttpClient()

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

func getcustomerbyid(id string, r *http.Request) []byte {

	baseurl := getURL()
	url := baseurl + "customer/byid/" + id

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
	client := getTraceableHttpClient()

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

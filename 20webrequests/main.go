package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length :", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count :", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang",
			"price" : 0,
			"platform" : "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "kuldeep")
	data.Add("lastname", "kumawat")
	data.Add("email", "kuldeep@dev.go")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// struct - Need to be defined according to the above map data.
type Headers struct {
	Authorization string `json:"authorization"`
}

type Args struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
	Header Headers
}

type Response struct {
	// StatusCode is the http code that will be returned back to the user.
	StatusCode int `json:"statusCode,omitempty"`
	// Headers is the information about the type of data being returned back.
	Headers map[string]string `json:"headers,omitempty"`
	// Body will contain the presigned url to upload or download files.

	Body int `json:"body,omitempty"`
}

var (
	ErrTokenNotFound = errors.New("authentication failed")
)

func Main(args map[string]interface{}) (*Response, error) {

	var sum int
	var secret string
	// Convert map to json string
	jsonStr, err := json.Marshal(args)
	if err != nil {
		fmt.Println(err)
	}
	
	//Replacing __ow_headers since I cannot use key with underscore in struct
	replace := "__ow_headers"
	newValue := "header"
	n := 1

	modifiedString := strings.Replace(string(jsonStr), replace, newValue, n)

	r := Args{}
    if err := json.Unmarshal([]byte(modifiedString), &r); err != nil {
        fmt.Println(err)
    }

	fmt.Println(secret)
	// Check for auth using token passed in header
	if r.Header.Authorization != "somerandomtoken" {
		fmt.Println("Authentication failed")

		return &Response{
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	//Returning result
    sum = r.Num1 + r.Num2

	return &Response{
		StatusCode: http.StatusOK,
		Body: sum,
	}, nil
}
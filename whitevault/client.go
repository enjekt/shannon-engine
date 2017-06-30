package main

import (
	"bytes"
	"github.com/enjekt/panda/commons"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

const (
	url             = "http://localhost:10000/rpc"
	addTokenService = "AddService.Add"
	getPadService   = "GetService.GetPad"
)

func GetPad(token *commons.Token) *commons.Pad {
	message, err := json.EncodeClientRequest(getPadService, token)
	checkEncodingError(message, err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	checkRequest(err)
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	checkResponseError(resp, err)
	defer resp.Body.Close()

	var result commons.Pad
	err = json.DecodeClientResponse(resp.Body, &result)
	log.Printf("Result %s\n", result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	return &result
}

func SendTokenAndPad(token *commons.Token, pad *commons.Pad) {
	msg := commons.TokenPadMessage{}
	msg.Pad = pad.ToString()
	msg.Token = token.ToString()

	message, err := json.EncodeClientRequest(addTokenService, msg)
	checkEncodingError(message, err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	checkRequest(err)
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	checkResponseError(resp, err)
	defer resp.Body.Close()

	var result commons.Result
	err = json.DecodeClientResponse(resp.Body, &result)
	log.Printf("Result %s\n", result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}

}
func checkRequest(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
func checkResponseError(resp *http.Response, err error) {
	log.Printf("Response : %s", resp.StatusCode)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
}
func checkEncodingError(message []byte, err error) {
	log.Printf("Encoded %s \n", string(message))
	if err != nil {
		log.Fatalf("%s", err)
	}

}

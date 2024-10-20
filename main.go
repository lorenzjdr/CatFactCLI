package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

// Structure of "https://catfact.ninja/fact" api response
type Data struct{
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main(){

	url := "https://catfact.ninja/fact"

	response, err := http.Get(url)

	if err != nil{
		log.Fatalf("Error, could not make GET request :%v",err)
	}
	defer response.Body.Close()
	
	var cat Data
	// decoding json using 'encoding/json' package
	if err := json.NewDecoder(response.Body).Decode(&cat); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Printf(cat.Fact)
}
package main

import (
	"fmt"
)

// Structure of "https://catfact.ninja/fact" api response
type Data struct{
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main(){
	fmt.Printf("Hello")
}
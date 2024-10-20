package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math"
)

// Structure of "https://catfact.ninja/fact" api response
type Data struct{
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// from https://github.com/flaviocopes/gololcat?tab=readme-ov-file
func rgb(i int) (int, int, int) {
    var f = 0.1
    return int(math.Sin(f*float64(i)+0)*127 + 128),
        int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
        int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
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

   for j := 0; j < len(cat.Fact); j++ {
       r, g, b := rgb(j)
       fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, cat.Fact[j])
   }
}
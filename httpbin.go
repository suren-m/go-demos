package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const _endpoint = "https://httpbin.org/json"

func print_json() {

	// Tuple based pattern for error handling
	resp, err := http.Get(_endpoint)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var binResp BinResp
	// deserialize json
	json.Unmarshal(body, &binResp)

	// print struct with fields info
	log.Printf("%+v\n", binResp.Slideshow)

	// serialize and print jsondata
	jsondata, _ := json.MarshalIndent(binResp, "", "    ")
	log.Printf("%s\n", jsondata)

}

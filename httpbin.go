package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const _jsonurl = "https://httpbin.org/json"
const _delayurl = "https://httpbin.org/delay"

func delay_call(delay int) {
	delayurl := fmt.Sprintf("%s/%s", _delayurl, strconv.Itoa(delay))
	log.Println("calling ", delayurl)
	http.Get(delayurl)
}

func print_json() {

	go delay_call(5)
	//delay_call(5)

	// Tuple based pattern for error handling
	resp, err := http.Get(_jsonurl)

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

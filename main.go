package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	jsonData := map[string]string{
		"query": `
            { 
                claimeds {
                   account
                   id
                   round
                   index
                   amount
                   blocknumber
                   txid
                }
            }
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/noberk/claim", bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

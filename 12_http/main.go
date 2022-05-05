package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/apex/log"
)

func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
	    return err
	}
	return nil
    }

func request() {
	var client http.Client

	var URL = "https://swapi.dev/api/people/1"
	//var client http.Client
	resp, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		log.Info(bodyString)
	}
}

func main() {
	request()
}

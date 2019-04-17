package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Text string `json:"text"`
}

func main() {
	urlPtr := flag.String("url", "", "Incoming Webhook URL")
	textPtr := flag.String("text", "", "The text message to send")

	flag.Parse()

	if len(*urlPtr) > 0 && len(*textPtr) > 0 {
		jsonData, err := json.Marshal(Payload{Text: *textPtr})

		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", *urlPtr, bytes.NewBuffer(jsonData))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		log.Println(resp.Status)

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(body))
	} else {
		log.Fatal("unsupported operation")
	}
}

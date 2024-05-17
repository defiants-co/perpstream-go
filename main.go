package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/defiants-co/perpstream-go/clients"
	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
)

func main() {
	pc := utils.NewPriceCache()

	go pc.StreamPrices(1, false)

	client, _ := clients.NewGmxClient(nil, pc)

	client.StreamPositions("0xeF5b5616FBa4e4d30a6B74De2B912025F8e627E4", false, 1, true, func(
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	) {
		for _, position := range newPositions {
			fmt.Println(position.ToJSON())
			go sendWebhook(position, "https://webhook.site/751f2139-8ced-4018-ad08-c21cb649b60e")
		}
	})

}

func sendWebhook(position models.FuturesPosition, url string) {
	// Create an HTTP client
	client := &http.Client{}

	jsonP, _ := position.ToJSON()
	// Create the POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonP)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request and receive a response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

}

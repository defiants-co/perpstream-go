package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"

	"github.com/defiants-co/perpstream-go/clients"
	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
)

func main() {
	// Initialize a single price cache for all clients
	pc := utils.NewPriceCache()
	go pc.StreamPrices(1, true) // Start streaming prices with one go routine

	// Define a slice of user URLs
	users := []string{
		"0xaa5572a377D44AA0f68d8735bd7391Cb10f0D239",
		"0xafD2718EFe0D202b5eAea136d6E66101C28a364F",
		"0xA63Ba502dC5c4db0CC19bac572C25007aE17F54d",
		"0x1614BE55CC0DB1208A339d05dD4b43a897C2d5F1",
		"0x18A63AFe3E4E6D0F3713cFdA890AD9C3B5f07C88",
		"0x30a82501583a828f29456a3aaE19d27d312358F3",
		"0xd67cCe92218517B4963eE843B3a454d382691F6e",
	}

	urls := []string{
		"https://arbitrum.llamarpc.com",
		"https://arbitrum-one-rpc.publicnode.com",
		"https://arb-pokt.nodies.app",
		"https://arb-mainnet-public.unifra.io",
		"https://arbitrum-one.publicnode.com",
		"https://arbitrum.meowrpc.com",
		"https://rpc.ankr.com/arbitrum",
		// "https://arbitrum.rpc.subquery.network/public",
		// "https://1rpc.io/arb",
	}

	var wg sync.WaitGroup // Use waitgroup to manage goroutines

	// Iterate over user URLs and create a new client for each, then start streaming
	for index, user := range users {
		wg.Add(1)
		url := urls[index]
		go startStream(user, url, pc, &wg)
	}

	wg.Wait() // Wait for all goroutines to complete
}
func startStream(user string, url string, pc *utils.GmxPriceCache, wg *sync.WaitGroup) {
	defer wg.Done()
	client, err := clients.NewGmxClient(&url, pc)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	client.StreamPositions(user, true, 5, true, true, sendCallBack)
}
func sendCallBack(
	newPositions []models.FuturesPosition,
	userId string,
	dataSource string,
) {
	for _, position := range newPositions {
		fmt.Println(position.ToJSON())
		go sendWebhook(position, "https://webhook.site/751f2139-8ced-4018-ad08-c21cb649b60e")
	}
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

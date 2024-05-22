package scripts

import (
	"log"
	"time"

	"github.com/defiants-co/perpstream-go/clients"
	"github.com/defiants-co/perpstream-go/utils"
)

func GmxScript() {
	priceCache := utils.NewGmxPriceCache()
	go priceCache.StreamPrices(10, false)

	client, err := clients.NewGmxClient(Rpcs, priceCache)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range Users {
		time.Sleep(500 * time.Millisecond)
		go client.StreamPositions(user, false, false, 10, GmxCallback)
	}
}

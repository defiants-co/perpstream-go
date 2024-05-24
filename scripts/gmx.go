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

	client.StreamPositions("0xeF5b5616FBa4e4d30a6B74De2B912025F8e627E4", true, true, 1, GmxCallback)

	for _, user := range Users {
		time.Sleep(500 * time.Millisecond)
		go client.StreamPositions(user, false, false, 10, GmxCallback)
	}
}

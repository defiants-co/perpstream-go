package scripts

import (
	"sync"
	"time"

	"github.com/defiants-co/perpstream-go/clients"
	"github.com/defiants-co/perpstream-go/utils"
)

func KwentaScript() {
	priceCache := utils.NewKwentaPriceCache()

	go priceCache.StreamPrices(1, false)

	client, err := clients.NewKwentaClient(priceCache)
	if err != nil {
		panic(err)
	}

	go client.StreamCacheUpdates(5, false)
	time.Sleep(3 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	client.StreamPositions("0xeF5b5616FBa4e4d30a6B74De2B912025F8e627E4", true, true, 1, GmxCallback)
	wg.Wait()

}

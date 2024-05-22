package scripts

import (
	"fmt"
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

	go client.StreamCacheUpdates(1, false)
	time.Sleep(3 * time.Second)
	counter := 0
	for _, user := range kwentaAccounts {
		// time.Sleep(1 * time.Second)
		counter++
		go client.StreamPositions(user, false, false, 3, GmxCallback)
		fmt.Println(counter, "started stream "+user)
	}

}

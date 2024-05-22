package scripts

import (
	"fmt"
	"log"

	"github.com/defiants-co/perpstream-go/clients"
)

func HegicScript() {
	client, err := clients.NewHegicClient()
	if err != nil {
		log.Fatal(err)
	}

	var people []string

	leaderboard, err := client.GetLeaderboard()

	for _, x := range leaderboard {
		if x.Overall.PnlUsd > 1000 {
			people = append(people, x.User)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(people))

	go client.StreamCacheUpdates(1, false)

	for _, person := range people {
		go client.StreamPositions(person, false, false, 3, HegicCallback)
	}
}

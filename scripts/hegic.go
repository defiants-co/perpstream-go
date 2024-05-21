package scripts

import (
	"fmt"
	"log"
	"time"

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

	go client.StreamPositions("0xeF5b5616FBa4e4d30a6B74De2B912025F8e627E4", false, true, 3, HegicCallback)

	for _, person := range people {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("hello")
		go client.StreamPositions(person, false, true, 3, HegicCallback)
	}
}

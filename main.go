package main

import (
	"fmt"

	"github.com/defiants-co/perpstream-go/clients"
)

func main() {

	client, err := clients.NewHegicClient()
	if err != nil {
		panic(err)
	}

	leaderboard, err := client.GetLeaderboard()
	if err != nil {
		panic(err)
	}
	var winners [](string)

	for _, user := range leaderboard {
		if user.Overall.ClosedContractsCount > 0 {
			if user.Overall.PnlPercent > 100 {
				winners = append(winners, user.User)
			}
		}
	}

	fmt.Println(winners)
}

package main

import (
	"sync"

	"github.com/defiants-co/perpstream-go/scripts"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go scripts.HegicScript()

	wg.Wait()

}

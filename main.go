package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/scripts"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// go scripts.GmxScript()
	go scripts.HegicScript()
	// go scripts.KwentaScript()
	wg.Wait()

	// hegicClient, _ := clients.NewHegicClient()

	// hegicOptions := hegicClient.GetActiveOptionsFromCache()

	// var options []models.OptionPosition

	// for _, i := range hegicOptions {
	// 	opti := utils.HegicPositionToOption(&i)
	// 	options = append(options, opti)

	// }

	// WriteStructsToFile("positions.json", options)
}

func WriteStructsToFile(filename string, data []models.OptionPosition) error {
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Create a JSON encoder and set indentation for pretty print
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Encode the data to JSON and write to the file
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("could not encode data to JSON: %v", err)
	}

	return nil
}

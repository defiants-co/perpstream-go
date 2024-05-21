package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type KwentaPriceCache struct {
	priceCache map[string]float64
	mu         sync.Mutex
}

type KwentaPrice struct {
	Price       string `json:"price"`
	Conf        string `json:"conf"`
	Expo        int    `json:"expo"`
	PublishTime int64  `json:"publish_time"`
}

type KwentaData struct {
	ID       string      `json:"id"`
	Price    KwentaPrice `json:"price"`
	EmaPrice KwentaPrice `json:"ema_price"`
}

func NewKwentaPriceCache() *KwentaPriceCache {
	myMap := make(map[string]float64)
	client := &KwentaPriceCache{
		priceCache: myMap,
	}
	client.UpdatePrices()

	return client
}

func (c *KwentaPriceCache) UpdatePrices() {
	// Perform the GET request
	resp, err := http.Get(PriceIdsToString())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Unmarshal the JSON response into the Data struct
	var data []KwentaData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	priceMap := make(map[string]float64)

	for _, key := range data {
		ticker := KwentaPriceIdsToToken[key.ID]
		price, _ := strconv.ParseFloat(key.Price.Price, 64)
		price = price / math.Pow(10, float64(KwentaPythDecimals[ticker]))
		priceMap[ticker] = price
	}

	c.mu.Lock()
	c.priceCache = priceMap
	c.mu.Unlock()
}

func (c *KwentaPriceCache) StreamPrices(sleepSeconds int, debug bool) {
	for {
		if debug {
			fmt.Println("fetching prices")
		}
		c.UpdatePrices()
		if debug {
			fmt.Println("fetched prices")
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}

func (c *KwentaPriceCache) ReturnPrice(token string) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.priceCache[token]
}

func (c *KwentaPriceCache) ReturnPriceMap() map[string]float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.priceCache
}

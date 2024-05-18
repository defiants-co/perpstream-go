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

type SignedPrice struct {
	TokenSymbol  string `json:"tokenSymbol"`
	MaxPriceFull string `json:"maxPriceFull"`
	MinPriceFull string `json:"minPriceFull"`
}

type ApiResponse struct {
	SignedPrices []SignedPrice `json:"signedPrices"`
}

type GmxPriceCache struct {
	prices map[string]float64
	lock   sync.Mutex
	debug  bool
}

func NewPriceCache() *GmxPriceCache {
	initialPriceMap := make(map[string]float64)
	keys := make([]string, 0, len(GmxCollateralTokenDecimals))
	for key := range GmxCollateralTokenDecimals {
		keys = append(keys, key)
	}

	for _, token := range keys {
		initialPriceMap[token] = 0
	}

	cache := &GmxPriceCache{
		prices: initialPriceMap,
	}

	cache.UpdatePrices()

	return cache
}

func (pc *GmxPriceCache) UpdatePrices() {
	// Make the HTTP GET request
	resp, err := http.Get(GmxPricesUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Unmarshal the JSON data into the ApiResponse struct
	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return
	}
	keys := make([]string, 0, len(GmxMarketToDecimals))
	for key := range GmxMarketToDecimals {
		keys = append(keys, key)
	}
	pc.lock.Lock()
	for _, priceData := range apiResponse.SignedPrices {
		priceFloat, _ := strconv.ParseFloat(priceData.MaxPriceFull, 64)
		priceFactor := 0
		if !Contains([]string{"USDC", "USDC.e", "USDT"}, priceData.TokenSymbol) {
			priceFactor = GmxMarketToDecimals[priceData.TokenSymbol]
		} else {
			priceFactor = 6
		}
		price := (priceFloat / math.Pow(10, float64(30-priceFactor)))
		if Contains(keys, priceData.TokenSymbol) {
			pc.prices[priceData.TokenSymbol] = math.Round(price*10000) / 10000
		}
	}
	pc.lock.Unlock()
}

func (pc *GmxPriceCache) StreamPrices(sleepSeconds float64, debug bool) {
	for {
		if debug {
			fmt.Println("fetching prices")
		}
		pc.UpdatePrices()
		if debug {
			fmt.Println("fetched prices")
		}

		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}

func (pc *GmxPriceCache) ReturnPrice(token string) float64 {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	return pc.prices[token]
}

func (pc *GmxPriceCache) ReturnPriceMap() map[string]float64 {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	return pc.prices
}

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strings"

	"github.com/defiants-co/perpstream-go/abis"
	"github.com/defiants-co/perpstream-go/models"
)

func GmxToFuturesPosition(position abis.PositionProps, priceCache *GmxPriceCache) *models.FuturesPosition {
	market := GmxAddressToMarket[position.Addresses.Market.Hex()]
	if market == "" {
		market = "ATOM-USD"
	}

	token := strings.Split(market, "-")[0]

	collateralToken := GmxAddressToCollateralToken[position.Addresses.CollateralToken.Hex()]

	oldSizeInUsd, _ := position.Numbers.SizeInUsd.Float64()
	oldSizeInUsd = math.Round(oldSizeInUsd/math.Pow(10, 26)) / (math.Pow(10, 4))

	sizeInTokens, _ := position.Numbers.SizeInTokens.Float64()
	sizeInTokens = math.Round(10000*sizeInTokens/(math.Pow(10, float64(GmxMarketToDecimals[token])))) / 10000
	entryPrice := math.Round((oldSizeInUsd/sizeInTokens)*10000) / 10000

	collateralTokenAmountFl, _ := position.Numbers.CollateralAmount.Float64()
	collateralTokenAmount := collateralTokenAmountFl / math.Pow(10, float64(GmxCollateralTokenDecimals[collateralToken]))

	collateralPrice := priceCache.ReturnPrice(collateralToken)
	collateralUsd := math.Round(collateralPrice*collateralTokenAmount*10000) / 10000
	tokenPrice := priceCache.ReturnPrice(token)

	leverage := math.Round(1000*(tokenPrice*sizeInTokens)/(collateralPrice*collateralTokenAmount)) / 1000

	currentSizeInUsd := tokenPrice * sizeInTokens

	var pnl float64

	if position.Flags.IsLong {
		pnl = math.Round((tokenPrice-entryPrice)*sizeInTokens*100) / 100
	} else {
		pnl = math.Round((entryPrice-tokenPrice)*sizeInTokens*100) / 100
	}

	return &models.FuturesPosition{
		CollateralToken:       collateralToken,
		CollateralTokenAmount: collateralTokenAmount,
		Market:                market,
		SizeUsd:               currentSizeInUsd,
		Size:                  sizeInTokens,
		IsLong:                position.Flags.IsLong,
		EntryPrice:            entryPrice,
		// Advanced fields
		MarkPrice:           tokenPrice,
		CollateralUsdAmount: collateralUsd,
		Leverage:            leverage,
		PnlUsd:              pnl,
	}
}

func reverseMap(originalMap map[string]string) map[string]string {
	reversedMap := make(map[string]string)
	for key, value := range originalMap {
		reversedMap[value] = key
	}
	return reversedMap
}

func MaxBigInt(bits int) *big.Int {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(int64(bits)), nil).Sub(max, big.NewInt(1))
	return max
}

type GmxPeriodAccountStat struct {
	ID                         string `json:"id"`
	ClosedCount                int    `json:"closedCount"`
	CumsumCollateral           string `json:"cumsumCollateral"`
	CumsumSize                 string `json:"cumsumSize"`
	Losses                     int    `json:"losses"`
	MaxCapital                 string `json:"maxCapital"`
	RealizedPriceImpact        string `json:"realizedPriceImpact"`
	SumMaxSize                 string `json:"sumMaxSize"`
	NetCapital                 string `json:"netCapital"`
	RealizedFees               string `json:"realizedFees"`
	RealizedPnl                string `json:"realizedPnl"`
	Volume                     string `json:"volume"`
	Wins                       int    `json:"wins"`
	StartUnrealizedPnl         string `json:"startUnrealizedPnl"`
	StartUnrealizedFees        string `json:"startUnrealizedFees"`
	StartUnrealizedPriceImpact string `json:"startUnrealizedPriceImpact"`
	Typename                   string `json:"__typename"`
}

// Struct to hold the top-level response
type ResponseData struct {
	All []GmxPeriodAccountStat `json:"all"`
}

func GetLeaderboard() ([]GmxPeriodAccountStat, error) {
	query := `
    query PeriodAccountStats {
      all: periodAccountStats(
        limit: 100000,
        where: {maxCapital_gte: "59000000000000000000000000000000000"}
      ) {
        id
        closedCount
        cumsumCollateral
        cumsumSize
        losses
        maxCapital
        realizedPriceImpact
        sumMaxSize
        netCapital
        realizedFees
        realizedPnl
        volume
        wins
        startUnrealizedPnl
        startUnrealizedFees
        startUnrealizedPriceImpact
      }
    }
    `

	jsonData := map[string]string{
		"query": query,
	}
	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", GmxLeaderboardUrl, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("query failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data ResponseData `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data.All, nil
}

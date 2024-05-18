package utils

import (
	"math"
	"math/big"
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

	sizeInUsd, _ := position.Numbers.SizeInUsd.Float64()
	sizeInUsd = math.Round(sizeInUsd/math.Pow(10, 26)) / (math.Pow(10, 4))

	sizeInTokens, _ := position.Numbers.SizeInTokens.Float64()
	sizeInTokens = math.Round(10000*sizeInTokens/(math.Pow(10, float64(GmxMarketToDecimals[token])))) / 10000
	entryPrice := math.Round((sizeInUsd/sizeInTokens)*10000) / 10000

	collateralTokenAmountFl, _ := position.Numbers.CollateralAmount.Float64()
	collateralTokenAmount := collateralTokenAmountFl / math.Pow(10, float64(GmxCollateralTokenDecimals[collateralToken]))

	collateralPrice := priceCache.ReturnPrice(collateralToken)
	collateralUsd := math.Round(collateralPrice*collateralTokenAmount*10000) / 10000
	tokenPrice := priceCache.ReturnPrice(token)

	leverage := math.Round(1000*(tokenPrice*sizeInTokens)/(collateralPrice*collateralTokenAmount)) / 1000

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
		SizeUsd:               sizeInUsd,
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

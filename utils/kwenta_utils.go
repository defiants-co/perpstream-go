package utils

import (
	"context"
	"math"
	"strconv"

	"github.com/defiants-co/perpstream-go/models"
	"github.com/hasura/go-graphql-client"
)

type KwentaFuturesPosition struct {
	Account         string  `json:"account"`
	AccountType     string  `json:"accountType"`
	Asset           string  `json:"asset"`
	AvgEntryPrice   string  `json:"avgEntryPrice"`
	CloseTimestamp  *int64  `json:"closeTimestamp"` // Assuming null can be represented as a pointer to int64
	EntryPrice      string  `json:"entryPrice"`
	ExitPrice       *string `json:"exitPrice"` // Assuming null can be represented as a pointer to string
	FeesPaid        string  `json:"feesPaid"`
	FundingIndex    string  `json:"fundingIndex"`
	Id              string  `json:"id"`
	IsLiquidated    bool    `json:"isLiquidated"`
	InitialMargin   string  `json:"initialMargin"`
	IsOpen          bool    `json:"isOpen"`
	LastPrice       string  `json:"lastPrice"`
	LastTxHash      string  `json:"lastTxHash"`
	Margin          string  `json:"margin"`
	Market          string  `json:"market"`
	MarketKey       string  `json:"marketKey"`
	NetTransfers    string  `json:"netTransfers"`
	NetFunding      string  `json:"netFunding"`
	OpenTimestamp   string  `json:"openTimestamp"`
	Pnl             string  `json:"pnl"`
	PnlWithFeesPaid string  `json:"pnlWithFeesPaid"`
	Size            string  `json:"size"`
	Timestamp       string  `json:"timestamp"`
	TotalVolume     string  `json:"totalVolume"`
	TotalDeposits   string  `json:"totalDeposits"`
	Trades          string  `json:"trades"`
}

// Define the KwentaGqlPositionsQuery struct for the GraphQL query
type KwentaGqlPositionsQuery struct {
	FuturesPositions []KwentaFuturesPosition `graphql:"futuresPositions(where: {isOpen: true, account_in : $account_in})"`
}

// Define the KwentaGqlResponse struct for the JSON response
type KwentaGqlResponse struct {
	Data struct {
		FuturesPositions []KwentaFuturesPosition `json:"futuresPositions"`
	} `json:"data"`
}

func KwentaToFuturesPosition(p *KwentaFuturesPosition, pc *KwentaPriceCache) *models.FuturesPosition {
	market := KwentaMarketsToToken[p.Asset]
	markPrice := pc.ReturnPrice(market)
	if market == "" {
		return nil
	}

	entryPriceFloat, _ := strconv.ParseFloat(p.AvgEntryPrice, 64)
	entryPrice := math.Round(10000*entryPriceFloat/math.Pow(10, 18)) / 10000

	marginFloat, _ := strconv.ParseFloat(p.InitialMargin, 64)
	collateral := math.Round(10000*marginFloat/math.Pow(10, 18)) / 10000

	sizeInTokenFloat, _ := strconv.ParseFloat(p.Size, 64)

	isLong := true
	if sizeInTokenFloat < 0 {
		isLong = false
	}
	sizeInToken := math.Abs(sizeInTokenFloat / math.Pow(10, 18))
	feeFloat, _ := strconv.ParseFloat(p.FeesPaid, 64)
	fee := (math.Round(1000*feeFloat/math.Pow(10, 18)) / 1000)

	pnl := 0.0
	if isLong {
		pnl = ((markPrice - entryPrice) * sizeInToken) - fee
	} else {
		pnl = ((entryPrice - markPrice) * sizeInToken) - fee
	}

	sizeInUsd := sizeInToken * markPrice
	leverage := math.Round(100*sizeInUsd/collateral) / 100

	return &models.FuturesPosition{
		CollateralToken:       "sUSD",
		CollateralTokenAmount: collateral,
		CollateralUsdAmount:   collateral,
		Market:                market + "-USD",
		EntryPrice:            entryPrice,
		MarkPrice:             markPrice,
		PnlUsd:                pnl,
		Size:                  sizeInToken,
		SizeUsd:               sizeInUsd,
		Leverage:              leverage,
		IsLong:                isLong,
	}
}

func KwentaSubgraphPositionsQuery(userIdList []string, assetList []string, graphqlClient *graphql.Client) ([]KwentaFuturesPosition, error) {
	if len(userIdList) == 0 {
		userIdList = []string{""}
	}
	if len(assetList) == 0 {
		assetList = []string{""}
	}
	variables := map[string]interface{}{
		"account_in": userIdList,
		"asset_in":   assetList,
	}

	// Define the response data structure
	var resp KwentaGqlPositionsQuery

	// Execute the query
	err := graphqlClient.Query(context.Background(), &resp, variables)
	if err != nil {
		return nil, NewFailedFetchPositionsError("mutliple accounts", err.Error())
	}

	return resp.FuturesPositions, nil
}

func PriceIdsToString() string {
	str := KwentaPriceDataUrl
	for key := range KwentaPriceIdsToToken {
		str += "ids[]=0x" + key + "&"
	}
	return str
}

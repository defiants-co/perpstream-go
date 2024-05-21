package utils

import (
	"encoding/json"
	"math"
	"strings"
	"time"

	"github.com/defiants-co/perpstream-go/models"
)

type HegicUserStats struct {
	User         string     `json:"user"`
	Overall      HegicStats `json:"overall"`
	CurrentEpoch HegicStats `json:"currentEpoch"`
}

type HegicStats struct {
	TradingVolume        float64 `json:"tradingVolume"`
	PaidPremium          float64 `json:"paidPremium"`
	PaidOut              float64 `json:"paidOut"`
	PnlUsd               float64 `json:"pnlUsd"`
	PnlPercent           float64 `json:"pnlPercent"`
	ClosedContractsCount int     `json:"closedContractsCount"`
}

type HegicPosition struct {
	State                int             `json:"state"`
	ID                   int             `json:"id"`
	User                 string          `json:"user"`
	PurchaseDate         time.Time       `json:"purchaseDate"`
	ExpirationDate       time.Time       `json:"expirationDate"`
	CloseDate            time.Time       `json:"closeDate"`
	Period               int             `json:"period"`
	Amount               float64         `json:"amount"`
	AmountUsd            float64         `json:"amountUsd"`
	Type                 string          `json:"type"`
	GeneralType          string          `json:"generalType"`
	Asset                string          `json:"asset"`
	ProfitZone           HegicProfitZone `json:"profitZone"`
	StrategyStrikes      HegicStrikes    `json:"strategyStrikes"`
	ClosedPrice          float64         `json:"closedPrice"`
	PremiumPaid          float64         `json:"premiumPaid"`
	Collateral           float64         `json:"collateral"`
	PayOff               float64         `json:"payOff"`
	GrossMargin          float64         `json:"grossMargin"`
	UserNetProfit        float64         `json:"userNetProfit"`
	UserNetProfitPercent float64         `json:"userNetProfitPercent"`
	IsInversed           bool            `json:"isInversed"`
}

func (hp *HegicPosition) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(hp)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func HegicPositionToOption(hp *HegicPosition) models.OptionPosition {

	return models.OptionPosition{
		UnderlyingToken: hp.Asset,
		ContractAmount:  hp.Amount,
		PremiumPaid:     hp.PremiumPaid,
		Strategy:        strings.ReplaceAll(strings.ToLower(hp.Type), " ", "_"),
		ExpirationDate:  hp.ExpirationDate.UTC().Format(time.RFC3339),
		DurationDays:    hp.Period,
		PnlUsd:          math.Round(hp.UserNetProfit*100) / 100,
	}
}

type HegicTradingStats struct {
	StrategiesPurchased int     `json:"strategiesPurchased"`
	FavoriteAsset       string  `json:"favoriteAsset"`
	TradingVolume       float64 `json:"tradingVolume"`
	Winrate             float64 `json:"winrate"`
	PnlUSD              float64 `json:"pnlUSD"`
	PnlPercent          float64 `json:"pnlPercent"`
}

type HegicProfitZone struct {
	Left   HegicProfitValue  `json:"left"`
	Center *HegicProfitValue `json:"center"`
	Right  HegicProfitValue  `json:"right"`
}

type HegicProfitValue struct {
	Profit bool    `json:"profit"`
	Value  float64 `json:"value"`
}

type HegicStrikes struct {
	BuyCallStrike  float64 `json:"buyCallStrike"`
	BuyPutStrike   float64 `json:"buyPutStrike"`
	SellCallStrike float64 `json:"sellCallStrike"`
	SellPutStrike  float64 `json:"sellPutStrike"`
}

type HegicPositions struct {
	Active []HegicPosition `json:"active"`
	Closed []HegicPosition `json:"closed"`
}

type HegicReferrerStats struct {
	Referrals int     `json:"referrals"`
	Earnings  float64 `json:"earnings"`
}

type HegicStaking struct {
	UserStake  float64 `json:"userStake"`
	TotalStake float64 `json:"totalStake"`
	UserShare  float64 `json:"userShare"`
}

type HegicStatsForActivePositions struct {
	ActiveStrategiesCount      int     `json:"activeStrategiesCount"`
	ActiveStrategiesUsdPnL     float64 `json:"activeStrategiesUsdPnL"`
	ActiveStrategiesPercentPnL float64 `json:"activeStrategiesPercentPnL"`
}

type HegicStatsForClosedPositions struct {
	ClosedStrategiesCount      int     `json:"closedStrategiesCount"`
	ClosedStrategiesUsdPnL     float64 `json:"closedStrategiesUsdPnL"`
	ClosedStrategiesPercentPnL float64 `json:"closedStrategiesPercentPnL"`
}

type HegicPositionsResponse struct {
	Positions []HegicPosition `json:"positions"`
}

type HegicUserData struct {
	TradingStats            HegicTradingStats            `json:"tradingStats"`
	Positions               HegicPositions               `json:"positions"`
	ReferrerStats           HegicReferrerStats           `json:"referrerStats"`
	LastActivity            time.Time                    `json:"lastActivity"`
	Staking                 HegicStaking                 `json:"staking"`
	StatsForActivePositions HegicStatsForActivePositions `json:"statsForActivePositions"`
	StatsForClosedPositions HegicStatsForClosedPositions `json:"statsForClosedPositions"`
}

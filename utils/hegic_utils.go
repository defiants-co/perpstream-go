package utils

import (
	"encoding/json"
	"math"
	"strings"
	"time"

	"github.com/defiants-co/perpstream-go/models"
)

type HegicProjectedPnl struct {
	ProjectedPnlCurrentEpoch                           []HegicScenarioData `json:"projectedPnl_CurrentEpoch"`
	ManualRiskAnalysis                                 []HegicScenarioData `json:"manualRiskAnalysis"`
	ProjectedPnl7_14                                   []HegicScenarioData `json:"projectedPnl_7_14"`
	ProjectedPnl14_30                                  []HegicScenarioData `json:"projectedPnl_14_30"`
	ProjectedPnl30_60                                  []HegicScenarioData `json:"projectedPnl_30_60"`
	ProjectedPnl60_90                                  []HegicScenarioData `json:"projectedPnl_60_90"`
	OpenInterestBySentiment                            HegicSentimentData  `json:"openInterestBySentiment"`
	ExpectedGrossMarginDates                           HegicDatesData      `json:"expectedGrossMarginByDateAndSentiment"`
	ExpectedGrossMargin                                []HegicScenarioData `json:"ExpectedGrossMarginForCurrentEpoch"`
	ExpectedGrossMarginByPrice4                        []HegicScenarioData `json:"ExpectedGrossMarginByPrice_4"`
	ExpectedGrossMarginByPrice3                        []HegicScenarioData `json:"ExpectedGrossMarginByPrice_3"`
	ExpectedGrossMarginByPrice2                        []HegicScenarioData `json:"ExpectedGrossMarginByPrice_2"`
	ExpectedGrossMarginByPrice1                        []HegicScenarioData `json:"ExpectedGrossMarginByPrice_1"`
	ExpectedGrossMarginByPrice                         []HegicScenarioData `json:"expectedGrossMarginByPrice"`
	ExpectedGrossMarginByPricesWithoutInversionProduct []HegicScenarioData `json:"expectedGrossMarginByPricesWithoutInversionProduct"`
	CountContracts                                     int                 `json:"countContracts"`
	RealizedRevenue                                    float64             `json:"realizedRevenue"`
	UnrealizedRevenue                                  float64             `json:"unrealizedRevenue"`
	LockedCollateral                                   float64             `json:"lockedCollateral"`
	RealizedPayOff                                     float64             `json:"realizedPayOff"`
	UnrealizedPayOff                                   float64             `json:"unrealizedPayOff"`
	RealizedGrossMargin                                float64             `json:"realizedGrossMargin"`
	UnrealizedGrossMargin                              float64             `json:"unrealizedGrossMargin"`
	TotalTradingVolume                                 float64             `json:"totalTradingVolume"`
	TotalOpenInterest                                  float64             `json:"totalOpenInterest"`
	EthOpenInterest                                    float64             `json:"ethOpenInterest"`
	BtcOpenInterest                                    float64             `json:"btcOpenInterest"`
	UnrealizedGrossMarginForCurrentEpoch               float64             `json:"unrealizedGrossMarginForCurentEpoch"`
	RealizedGrossMarginForCurrentEpoch                 float64             `json:"realizedGrossMarginForCurrentEpoch"`
	EstimatedGrossMarginByEndOfCurrentEpoch            float64             `json:"estimatedGrossMarginByTheEndOfCurrentEpoch"`
	Solvency                                           float64             `json:"solvency"`
	OperationalBalance                                 float64             `json:"opterationalBalance"`
	PayoffPoolBalance                                  float64             `json:"payoffPoolBalance"`
	ExpectedGrossMarginByDateFormatTime                HegicDateTimeData   `json:"expectedGrossMarginByDateFormatTime"`
}

type HegicScenarioData struct {
	Scenario                 string  `json:"scenario"`
	EthPrice                 float64 `json:"ethPrice"`
	BtcPrice                 float64 `json:"btcPrice"`
	EthExpectedGrossMargin   float64 `json:"ethExpectedGrossMargin"`
	BtcExpectedGrossMargin   float64 `json:"btcExpectedGrossMargin"`
	TotalExpectedGrossMargin float64 `json:"totalExpectedGrossMargin"`
	Order                    int     `json:"order"`
	CoverPoolPnl             float64 `json:"coverPoolPnl"`
}

type HegicSentimentData struct {
	Bullish        float64 `json:"bullish"`
	Bearish        float64 `json:"bearish"`
	HighVolatility float64 `json:"highVolatility"`
	LowVolatility  float64 `json:"lowVolatility"`
}

type HegicDatesData struct {
	Dates          []HegicDateRange `json:"dates"`
	Bullish        []float64        `json:"bullish"`
	Bearish        []float64        `json:"bearish"`
	HighVolatility []float64        `json:"highVolatility"`
	LowVolatility  []float64        `json:"lowVolatility"`
}

type HegicDateRange struct {
	End   time.Time `json:"end"`
	Start time.Time `json:"start"`
}

type HegicDateTimeData struct {
	Dates          []time.Time `json:"dates"`
	Bullish        []float64   `json:"bullish"`
	Bearish        []float64   `json:"bearish"`
	HighVolatility []float64   `json:"highVolatility"`
	LowVolatility  []float64   `json:"lowVolatility"`
}

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

package models

import (
	"encoding/json"
	"math"
)

type FuturesPosition struct {
	// Basic Fields for determining equality
	CollateralToken       string  `json:"collateral_token"`
	CollateralTokenAmount float64 `json:"collateral_token_amount"`
	Market                string  `json:"market"`
	SizeUsd               float64 `json:"size_usd"`
	IsLong                bool    `json:"is_long"`
	CollateralUsdAmount   float64 `json:"collateral_usd_amount"`

	Size       float64 `json:"size"`
	EntryPrice float64 `json:"entry_price"`

	// Advanced fields that require more data
	MarkPrice float64 `json:"mark_price"`
	Leverage  float64 `json:"leverage"`
	PnlUsd    float64 `json:"pnl_usd"` // expresses as a positive or negative amount of usd
}

func (fp *FuturesPosition) BasicEqual(other *FuturesPosition) bool {
	return (fp.CollateralToken == other.CollateralToken &&
		math.Abs(fp.CollateralTokenAmount-other.CollateralTokenAmount) < 0.01 &&
		fp.Market == other.Market &&
		math.Abs(fp.SizeUsd-other.SizeUsd) < 0.01 &&
		fp.IsLong == other.IsLong)
	// math.Abs(fp.Leverage-other.Leverage) < 0.01)
}

func (fp *FuturesPosition) Equal(other *FuturesPosition) bool {
	return (fp.BasicEqual(other) && math.Abs(fp.Leverage-fp.Leverage) < 1)
}

// ToJSON converts a FuturesPosition to its JSON representation
func (fp FuturesPosition) ToJSON() (string, error) {
	jsonData, err := json.Marshal(fp)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func PositionSetsAreEqual(positions1 []FuturesPosition, positions2 []FuturesPosition) bool {
	if len(positions1) != len(positions2) {
		return false
	}

	for _, position1 := range positions1 {
		hasMatch := false
		for _, position2 := range positions2 {
			if position1.Equal(&position2) {
				hasMatch = true
			}
		}
		if !hasMatch {
			return false
		}
	}
	return true
}

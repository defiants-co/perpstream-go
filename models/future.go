package models

import (
	"encoding/json"
	"math"
)

// FuturesPosition represents a position in the futures market
type FuturesPosition struct {
	CollateralToken       string  `json:"collateral_token"`        // The token used as collateral
	CollateralTokenAmount float64 `json:"collateral_token_amount"` // Amount of the collateral token
	Market                string  `json:"market"`                  // The market where the position is held
	SizeUsd               float64 `json:"size_usd"`                // Position size in USD
	IsLong                bool    `json:"is_long"`                 // True if the position is long, false if short
	CollateralUsdAmount   float64 `json:"collateral_usd_amount"`   // Collateral amount in USD

	Size       float64 `json:"size"`        // The size of the position
	EntryPrice float64 `json:"entry_price"` // The entry price of the position

	MarkPrice float64 `json:"mark_price"` // The current mark price of the position
	Leverage  float64 `json:"leverage"`   // The leverage used in the position
	PnlUsd    float64 `json:"pnl_usd"`    // Profit and loss in USD, can be positive or negative
}

// Comparable checks if two FuturesPosition objects are equal based on basic fields
func ComparablePosition(position1 *FuturesPosition, position2 *FuturesPosition) bool {
	if position1 == nil || position2 == nil {
		return false
	}
	return (position1.CollateralToken == position2.CollateralToken &&
		position1.Market == position2.Market &&
		position1.IsLong == position2.IsLong)
}

// Equal checks if two FuturesPosition objects are equal based on all fields including leverage
func (fp *FuturesPosition) Equal(other *FuturesPosition) bool {
	return (ComparablePosition(fp, other) &&
		math.Abs(fp.CollateralTokenAmount-other.CollateralTokenAmount) < 0.01*fp.CollateralTokenAmount &&
		math.Abs(fp.EntryPrice-other.EntryPrice) < 0.1*fp.EntryPrice &&
		math.Abs(fp.Leverage-other.Leverage) < 1 &&
		math.Abs(fp.Size-other.Size) < 0.01*fp.Size)
}

// ToJSON converts a FuturesPosition to its JSON representation
func (fp FuturesPosition) ToJSON() (string, error) {
	jsonData, err := json.Marshal(fp)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// FuturesPositionSetsAreEqual checks if two slices of FuturesPosition objects are equal
func FuturesPositionSetsAreEqual(positions1 []FuturesPosition, positions2 []FuturesPosition) bool {
	if len(positions1) != len(positions2) {
		return false
	}

	for _, position1 := range positions1 {
		hasMatch := false
		for _, position2 := range positions2 {
			if position1.Equal(&position2) {
				hasMatch = true
				break
			}
		}
		if !hasMatch {
			return false
		}
	}
	return true
}

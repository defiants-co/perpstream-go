package models

import (
	"encoding/json"
	"math"
)

type OptionPosition struct {
	UnderlyingToken string  `json:"underlying_token"`
	ContractAmount  float64 `json:"contract_amount"`
	PremiumPaid     float64 `json:"premium_paid"`
	Strategy        string  `json:"strategy"`
	ExpirationDate  string  `json:"expiration_date"`
	DurationDays    int     `json:"duration_days"`
	PnlUsd          float64 `json:"pnl_usd"`
}

func (o OptionPosition) Equal(other OptionPosition) bool {
	const floatTolerance = 0.01

	return o.UnderlyingToken == other.UnderlyingToken &&
		o.Strategy == other.Strategy &&
		o.ExpirationDate == other.ExpirationDate &&
		o.DurationDays == other.DurationDays &&
		math.Abs(o.ContractAmount-other.ContractAmount) < floatTolerance &&
		math.Abs(o.PremiumPaid-other.PremiumPaid) < floatTolerance
	// math.Abs(o.PnlUsd-other.PnlUsd) < floatTolerance
}

func (op *OptionPosition) ToJSON() (string, error) {
	jsonData, err := json.Marshal(op)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func OptionPositionSetsAreEqual(positions1 []OptionPosition, positions2 []OptionPosition) bool {
	if len(positions1) != len(positions2) {
		return false
	}

	for _, position1 := range positions1 {
		hasMatch := false
		for _, position2 := range positions2 {
			if position1.Equal(position2) {
				hasMatch = true
			}
		}
		if !hasMatch {
			return false
		}
	}
	return true
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IPerpsV2MarketBaseTypesPosition is an auto generated low-level Go binding around an user-defined struct.
type IPerpsV2MarketBaseTypesPosition struct {
	Id               uint64
	LastFundingIndex uint64
	Margin           *big.Int
	LastPrice        *big.Int
	Size             *big.Int
}

// IPerpsV2MarketSettingsParameters is an auto generated low-level Go binding around an user-defined struct.
type IPerpsV2MarketSettingsParameters struct {
	TakerFee                     *big.Int
	MakerFee                     *big.Int
	TakerFeeDelayedOrder         *big.Int
	MakerFeeDelayedOrder         *big.Int
	TakerFeeOffchainDelayedOrder *big.Int
	MakerFeeOffchainDelayedOrder *big.Int
	MaxLeverage                  *big.Int
	MaxMarketValue               *big.Int
	MaxFundingVelocity           *big.Int
	SkewScale                    *big.Int
	NextPriceConfirmWindow       *big.Int
	DelayedOrderConfirmWindow    *big.Int
	MinDelayTimeDelta            *big.Int
	MaxDelayTimeDelta            *big.Int
	OffchainDelayedOrderMinAge   *big.Int
	OffchainDelayedOrderMaxAge   *big.Int
	OffchainMarketKey            [32]byte
	OffchainPriceDivergence      *big.Int
	LiquidationPremiumMultiplier *big.Int
	LiquidationBufferRatio       *big.Int
	MaxLiquidationDelta          *big.Int
	MaxPD                        *big.Int
}

// PerpsV2MarketDataFeeRates is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataFeeRates struct {
	TakerFee                     *big.Int
	MakerFee                     *big.Int
	TakerFeeDelayedOrder         *big.Int
	MakerFeeDelayedOrder         *big.Int
	TakerFeeOffchainDelayedOrder *big.Int
	MakerFeeOffchainDelayedOrder *big.Int
}

// PerpsV2MarketDataFundingParameters is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataFundingParameters struct {
	MaxFundingVelocity *big.Int
	SkewScale          *big.Int
}

// PerpsV2MarketDataFuturesGlobals is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataFuturesGlobals struct {
	MinInitialMargin    *big.Int
	LiquidationFeeRatio *big.Int
	MinKeeperFee        *big.Int
	MaxKeeperFee        *big.Int
}

// PerpsV2MarketDataMarketData is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataMarketData struct {
	Market            common.Address
	BaseAsset         [32]byte
	MarketKey         [32]byte
	FeeRates          PerpsV2MarketDataFeeRates
	Limits            PerpsV2MarketDataMarketLimits
	FundingParameters PerpsV2MarketDataFundingParameters
	MarketSizeDetails PerpsV2MarketDataMarketSizeDetails
	PriceDetails      PerpsV2MarketDataPriceDetails
}

// PerpsV2MarketDataMarketLimits is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataMarketLimits struct {
	MaxLeverage    *big.Int
	MaxMarketValue *big.Int
}

// PerpsV2MarketDataMarketSizeDetails is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataMarketSizeDetails struct {
	MarketSize *big.Int
	Sides      PerpsV2MarketDataSides
	MarketDebt *big.Int
	MarketSkew *big.Int
}

// PerpsV2MarketDataMarketSummary is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataMarketSummary struct {
	Market                 common.Address
	Asset                  [32]byte
	Key                    [32]byte
	MaxLeverage            *big.Int
	Price                  *big.Int
	MarketSize             *big.Int
	MarketSkew             *big.Int
	MarketDebt             *big.Int
	CurrentFundingRate     *big.Int
	CurrentFundingVelocity *big.Int
	FeeRates               PerpsV2MarketDataFeeRates
}

// PerpsV2MarketDataPositionData is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataPositionData struct {
	Position             IPerpsV2MarketBaseTypesPosition
	NotionalValue        *big.Int
	ProfitLoss           *big.Int
	AccruedFunding       *big.Int
	RemainingMargin      *big.Int
	AccessibleMargin     *big.Int
	LiquidationPrice     *big.Int
	CanLiquidatePosition bool
}

// PerpsV2MarketDataPriceDetails is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataPriceDetails struct {
	Price   *big.Int
	Invalid bool
}

// PerpsV2MarketDataSides is an auto generated low-level Go binding around an user-defined struct.
type PerpsV2MarketDataSides struct {
	Long  *big.Int
	Short *big.Int
}

// AbisMetaData contains all meta data concerning the Abis contract.
var AbisMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAddressResolver\",\"name\":\"_resolverProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"allMarketSummaries\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketSkew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"marketDebt\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FeeRates\",\"name\":\"feeRates\",\"type\":\"tuple\"}],\"internalType\":\"structPerpsV2MarketData.MarketSummary[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allProxiedMarketSummaries\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketSkew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"marketDebt\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FeeRates\",\"name\":\"feeRates\",\"type\":\"tuple\"}],\"internalType\":\"structPerpsV2MarketData.MarketSummary[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minInitialMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minKeeperFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxKeeperFee\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FuturesGlobals\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIPerpsV2MarketViews\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"marketDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"baseAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FeeRates\",\"name\":\"feeRates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketValue\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.MarketLimits\",\"name\":\"limits\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FundingParameters\",\"name\":\"fundingParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.Sides\",\"name\":\"sides\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"marketDebt\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketSkew\",\"type\":\"int256\"}],\"internalType\":\"structPerpsV2MarketData.MarketSizeDetails\",\"name\":\"marketSizeDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"internalType\":\"structPerpsV2MarketData.PriceDetails\",\"name\":\"priceDetails\",\"type\":\"tuple\"}],\"internalType\":\"structPerpsV2MarketData.MarketData\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketKey\",\"type\":\"bytes32\"}],\"name\":\"marketDetailsForKey\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"baseAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FeeRates\",\"name\":\"feeRates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketValue\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.MarketLimits\",\"name\":\"limits\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FundingParameters\",\"name\":\"fundingParameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.Sides\",\"name\":\"sides\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"marketDebt\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketSkew\",\"type\":\"int256\"}],\"internalType\":\"structPerpsV2MarketData.MarketSizeDetails\",\"name\":\"marketSizeDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"internalType\":\"structPerpsV2MarketData.PriceDetails\",\"name\":\"priceDetails\",\"type\":\"tuple\"}],\"internalType\":\"structPerpsV2MarketData.MarketData\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"marketSummaries\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketSkew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"marketDebt\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FeeRates\",\"name\":\"feeRates\",\"type\":\"tuple\"}],\"internalType\":\"structPerpsV2MarketData.MarketSummary[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"marketKeys\",\"type\":\"bytes32[]\"}],\"name\":\"marketSummariesForKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketSize\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketSkew\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"marketDebt\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"currentFundingVelocity\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"}],\"internalType\":\"structPerpsV2MarketData.FeeRates\",\"name\":\"feeRates\",\"type\":\"tuple\"}],\"internalType\":\"structPerpsV2MarketData.MarketSummary[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketKey\",\"type\":\"bytes32\"}],\"name\":\"parameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeOffchainDelayedOrder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMarketValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFundingVelocity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skewScale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextPriceConfirmWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delayedOrderConfirmWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDelayTimeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelayTimeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offchainDelayedOrderMinAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offchainDelayedOrderMaxAge\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchainMarketKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offchainPriceDivergence\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPremiumMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBufferRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPD\",\"type\":\"uint256\"}],\"internalType\":\"structIPerpsV2MarketSettings.Parameters\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIPerpsV2MarketViews\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"positionDetails\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFundingIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"margin\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastPrice\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"size\",\"type\":\"int128\"}],\"internalType\":\"structIPerpsV2MarketBaseTypes.Position\",\"name\":\"position\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"notionalValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"profitLoss\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"remainingMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessibleMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canLiquidatePosition\",\"type\":\"bool\"}],\"internalType\":\"structPerpsV2MarketData.PositionData\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"positionDetailsForMarketKey\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFundingIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"margin\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastPrice\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"size\",\"type\":\"int128\"}],\"internalType\":\"structIPerpsV2MarketBaseTypes.Position\",\"name\":\"position\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"notionalValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"profitLoss\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accruedFunding\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"remainingMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessibleMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canLiquidatePosition\",\"type\":\"bool\"}],\"internalType\":\"structPerpsV2MarketData.PositionData\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverProxy\",\"outputs\":[{\"internalType\":\"contractIAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbisABI is the input ABI used to generate the binding from.
// Deprecated: Use AbisMetaData.ABI instead.
var AbisABI = AbisMetaData.ABI

// Abis is an auto generated Go binding around an Ethereum contract.
type Abis struct {
	AbisCaller     // Read-only binding to the contract
	AbisTransactor // Write-only binding to the contract
	AbisFilterer   // Log filterer for contract events
}

// AbisCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbisSession struct {
	Contract     *Abis             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbisCallerSession struct {
	Contract *AbisCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbisTransactorSession struct {
	Contract     *AbisTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbisRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbisRaw struct {
	Contract *Abis // Generic contract binding to access the raw methods on
}

// AbisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbisCallerRaw struct {
	Contract *AbisCaller // Generic read-only contract binding to access the raw methods on
}

// AbisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbisTransactorRaw struct {
	Contract *AbisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbis creates a new instance of Abis, bound to a specific deployed contract.
func NewAbis(address common.Address, backend bind.ContractBackend) (*Abis, error) {
	contract, err := bindAbis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abis{AbisCaller: AbisCaller{contract: contract}, AbisTransactor: AbisTransactor{contract: contract}, AbisFilterer: AbisFilterer{contract: contract}}, nil
}

// NewAbisCaller creates a new read-only instance of Abis, bound to a specific deployed contract.
func NewAbisCaller(address common.Address, caller bind.ContractCaller) (*AbisCaller, error) {
	contract, err := bindAbis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbisCaller{contract: contract}, nil
}

// NewAbisTransactor creates a new write-only instance of Abis, bound to a specific deployed contract.
func NewAbisTransactor(address common.Address, transactor bind.ContractTransactor) (*AbisTransactor, error) {
	contract, err := bindAbis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbisTransactor{contract: contract}, nil
}

// NewAbisFilterer creates a new log filterer instance of Abis, bound to a specific deployed contract.
func NewAbisFilterer(address common.Address, filterer bind.ContractFilterer) (*AbisFilterer, error) {
	contract, err := bindAbis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbisFilterer{contract: contract}, nil
}

// bindAbis binds a generic wrapper to an already deployed contract.
func bindAbis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abis *AbisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abis.Contract.AbisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abis *AbisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.Contract.AbisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abis *AbisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abis.Contract.AbisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abis *AbisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abis *AbisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abis *AbisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abis.Contract.contract.Transact(opts, method, params...)
}

// AllMarketSummaries is a free data retrieval call binding the contract method 0x3c88ee18.
//
// Solidity: function allMarketSummaries() view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCaller) AllMarketSummaries(opts *bind.CallOpts) ([]PerpsV2MarketDataMarketSummary, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "allMarketSummaries")

	if err != nil {
		return *new([]PerpsV2MarketDataMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new([]PerpsV2MarketDataMarketSummary)).(*[]PerpsV2MarketDataMarketSummary)

	return out0, err

}

// AllMarketSummaries is a free data retrieval call binding the contract method 0x3c88ee18.
//
// Solidity: function allMarketSummaries() view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisSession) AllMarketSummaries() ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.AllMarketSummaries(&_Abis.CallOpts)
}

// AllMarketSummaries is a free data retrieval call binding the contract method 0x3c88ee18.
//
// Solidity: function allMarketSummaries() view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCallerSession) AllMarketSummaries() ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.AllMarketSummaries(&_Abis.CallOpts)
}

// AllProxiedMarketSummaries is a free data retrieval call binding the contract method 0x093bfd66.
//
// Solidity: function allProxiedMarketSummaries() view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCaller) AllProxiedMarketSummaries(opts *bind.CallOpts) ([]PerpsV2MarketDataMarketSummary, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "allProxiedMarketSummaries")

	if err != nil {
		return *new([]PerpsV2MarketDataMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new([]PerpsV2MarketDataMarketSummary)).(*[]PerpsV2MarketDataMarketSummary)

	return out0, err

}

// AllProxiedMarketSummaries is a free data retrieval call binding the contract method 0x093bfd66.
//
// Solidity: function allProxiedMarketSummaries() view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisSession) AllProxiedMarketSummaries() ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.AllProxiedMarketSummaries(&_Abis.CallOpts)
}

// AllProxiedMarketSummaries is a free data retrieval call binding the contract method 0x093bfd66.
//
// Solidity: function allProxiedMarketSummaries() view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCallerSession) AllProxiedMarketSummaries() ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.AllProxiedMarketSummaries(&_Abis.CallOpts)
}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns((uint256,uint256,uint256,uint256))
func (_Abis *AbisCaller) Globals(opts *bind.CallOpts) (PerpsV2MarketDataFuturesGlobals, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "globals")

	if err != nil {
		return *new(PerpsV2MarketDataFuturesGlobals), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpsV2MarketDataFuturesGlobals)).(*PerpsV2MarketDataFuturesGlobals)

	return out0, err

}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns((uint256,uint256,uint256,uint256))
func (_Abis *AbisSession) Globals() (PerpsV2MarketDataFuturesGlobals, error) {
	return _Abis.Contract.Globals(&_Abis.CallOpts)
}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns((uint256,uint256,uint256,uint256))
func (_Abis *AbisCallerSession) Globals() (PerpsV2MarketDataFuturesGlobals, error) {
	return _Abis.Contract.Globals(&_Abis.CallOpts)
}

// MarketDetails is a free data retrieval call binding the contract method 0x730e0037.
//
// Solidity: function marketDetails(address market) view returns((address,bytes32,bytes32,(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,(uint256,uint256),uint256,int256),(uint256,bool)))
func (_Abis *AbisCaller) MarketDetails(opts *bind.CallOpts, market common.Address) (PerpsV2MarketDataMarketData, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "marketDetails", market)

	if err != nil {
		return *new(PerpsV2MarketDataMarketData), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpsV2MarketDataMarketData)).(*PerpsV2MarketDataMarketData)

	return out0, err

}

// MarketDetails is a free data retrieval call binding the contract method 0x730e0037.
//
// Solidity: function marketDetails(address market) view returns((address,bytes32,bytes32,(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,(uint256,uint256),uint256,int256),(uint256,bool)))
func (_Abis *AbisSession) MarketDetails(market common.Address) (PerpsV2MarketDataMarketData, error) {
	return _Abis.Contract.MarketDetails(&_Abis.CallOpts, market)
}

// MarketDetails is a free data retrieval call binding the contract method 0x730e0037.
//
// Solidity: function marketDetails(address market) view returns((address,bytes32,bytes32,(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,(uint256,uint256),uint256,int256),(uint256,bool)))
func (_Abis *AbisCallerSession) MarketDetails(market common.Address) (PerpsV2MarketDataMarketData, error) {
	return _Abis.Contract.MarketDetails(&_Abis.CallOpts, market)
}

// MarketDetailsForKey is a free data retrieval call binding the contract method 0xc1d1df56.
//
// Solidity: function marketDetailsForKey(bytes32 marketKey) view returns((address,bytes32,bytes32,(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,(uint256,uint256),uint256,int256),(uint256,bool)))
func (_Abis *AbisCaller) MarketDetailsForKey(opts *bind.CallOpts, marketKey [32]byte) (PerpsV2MarketDataMarketData, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "marketDetailsForKey", marketKey)

	if err != nil {
		return *new(PerpsV2MarketDataMarketData), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpsV2MarketDataMarketData)).(*PerpsV2MarketDataMarketData)

	return out0, err

}

// MarketDetailsForKey is a free data retrieval call binding the contract method 0xc1d1df56.
//
// Solidity: function marketDetailsForKey(bytes32 marketKey) view returns((address,bytes32,bytes32,(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,(uint256,uint256),uint256,int256),(uint256,bool)))
func (_Abis *AbisSession) MarketDetailsForKey(marketKey [32]byte) (PerpsV2MarketDataMarketData, error) {
	return _Abis.Contract.MarketDetailsForKey(&_Abis.CallOpts, marketKey)
}

// MarketDetailsForKey is a free data retrieval call binding the contract method 0xc1d1df56.
//
// Solidity: function marketDetailsForKey(bytes32 marketKey) view returns((address,bytes32,bytes32,(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,(uint256,uint256),uint256,int256),(uint256,bool)))
func (_Abis *AbisCallerSession) MarketDetailsForKey(marketKey [32]byte) (PerpsV2MarketDataMarketData, error) {
	return _Abis.Contract.MarketDetailsForKey(&_Abis.CallOpts, marketKey)
}

// MarketSummaries is a free data retrieval call binding the contract method 0xa9e0bef7.
//
// Solidity: function marketSummaries(address[] markets) view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCaller) MarketSummaries(opts *bind.CallOpts, markets []common.Address) ([]PerpsV2MarketDataMarketSummary, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "marketSummaries", markets)

	if err != nil {
		return *new([]PerpsV2MarketDataMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new([]PerpsV2MarketDataMarketSummary)).(*[]PerpsV2MarketDataMarketSummary)

	return out0, err

}

// MarketSummaries is a free data retrieval call binding the contract method 0xa9e0bef7.
//
// Solidity: function marketSummaries(address[] markets) view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisSession) MarketSummaries(markets []common.Address) ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.MarketSummaries(&_Abis.CallOpts, markets)
}

// MarketSummaries is a free data retrieval call binding the contract method 0xa9e0bef7.
//
// Solidity: function marketSummaries(address[] markets) view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCallerSession) MarketSummaries(markets []common.Address) ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.MarketSummaries(&_Abis.CallOpts, markets)
}

// MarketSummariesForKeys is a free data retrieval call binding the contract method 0x83ce9022.
//
// Solidity: function marketSummariesForKeys(bytes32[] marketKeys) view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCaller) MarketSummariesForKeys(opts *bind.CallOpts, marketKeys [][32]byte) ([]PerpsV2MarketDataMarketSummary, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "marketSummariesForKeys", marketKeys)

	if err != nil {
		return *new([]PerpsV2MarketDataMarketSummary), err
	}

	out0 := *abi.ConvertType(out[0], new([]PerpsV2MarketDataMarketSummary)).(*[]PerpsV2MarketDataMarketSummary)

	return out0, err

}

// MarketSummariesForKeys is a free data retrieval call binding the contract method 0x83ce9022.
//
// Solidity: function marketSummariesForKeys(bytes32[] marketKeys) view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisSession) MarketSummariesForKeys(marketKeys [][32]byte) ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.MarketSummariesForKeys(&_Abis.CallOpts, marketKeys)
}

// MarketSummariesForKeys is a free data retrieval call binding the contract method 0x83ce9022.
//
// Solidity: function marketSummariesForKeys(bytes32[] marketKeys) view returns((address,bytes32,bytes32,uint256,uint256,uint256,int256,uint256,int256,int256,(uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Abis *AbisCallerSession) MarketSummariesForKeys(marketKeys [][32]byte) ([]PerpsV2MarketDataMarketSummary, error) {
	return _Abis.Contract.MarketSummariesForKeys(&_Abis.CallOpts, marketKeys)
}

// Parameters is a free data retrieval call binding the contract method 0x02506804.
//
// Solidity: function parameters(bytes32 marketKey) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32,uint256,uint256,uint256,uint256,uint256))
func (_Abis *AbisCaller) Parameters(opts *bind.CallOpts, marketKey [32]byte) (IPerpsV2MarketSettingsParameters, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "parameters", marketKey)

	if err != nil {
		return *new(IPerpsV2MarketSettingsParameters), err
	}

	out0 := *abi.ConvertType(out[0], new(IPerpsV2MarketSettingsParameters)).(*IPerpsV2MarketSettingsParameters)

	return out0, err

}

// Parameters is a free data retrieval call binding the contract method 0x02506804.
//
// Solidity: function parameters(bytes32 marketKey) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32,uint256,uint256,uint256,uint256,uint256))
func (_Abis *AbisSession) Parameters(marketKey [32]byte) (IPerpsV2MarketSettingsParameters, error) {
	return _Abis.Contract.Parameters(&_Abis.CallOpts, marketKey)
}

// Parameters is a free data retrieval call binding the contract method 0x02506804.
//
// Solidity: function parameters(bytes32 marketKey) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32,uint256,uint256,uint256,uint256,uint256))
func (_Abis *AbisCallerSession) Parameters(marketKey [32]byte) (IPerpsV2MarketSettingsParameters, error) {
	return _Abis.Contract.Parameters(&_Abis.CallOpts, marketKey)
}

// PositionDetails is a free data retrieval call binding the contract method 0x808bad34.
//
// Solidity: function positionDetails(address market, address account) view returns(((uint64,uint64,uint128,uint128,int128),int256,int256,int256,uint256,uint256,uint256,bool))
func (_Abis *AbisCaller) PositionDetails(opts *bind.CallOpts, market common.Address, account common.Address) (PerpsV2MarketDataPositionData, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "positionDetails", market, account)

	if err != nil {
		return *new(PerpsV2MarketDataPositionData), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpsV2MarketDataPositionData)).(*PerpsV2MarketDataPositionData)

	return out0, err

}

// PositionDetails is a free data retrieval call binding the contract method 0x808bad34.
//
// Solidity: function positionDetails(address market, address account) view returns(((uint64,uint64,uint128,uint128,int128),int256,int256,int256,uint256,uint256,uint256,bool))
func (_Abis *AbisSession) PositionDetails(market common.Address, account common.Address) (PerpsV2MarketDataPositionData, error) {
	return _Abis.Contract.PositionDetails(&_Abis.CallOpts, market, account)
}

// PositionDetails is a free data retrieval call binding the contract method 0x808bad34.
//
// Solidity: function positionDetails(address market, address account) view returns(((uint64,uint64,uint128,uint128,int128),int256,int256,int256,uint256,uint256,uint256,bool))
func (_Abis *AbisCallerSession) PositionDetails(market common.Address, account common.Address) (PerpsV2MarketDataPositionData, error) {
	return _Abis.Contract.PositionDetails(&_Abis.CallOpts, market, account)
}

// PositionDetailsForMarketKey is a free data retrieval call binding the contract method 0x985f289e.
//
// Solidity: function positionDetailsForMarketKey(bytes32 marketKey, address account) view returns(((uint64,uint64,uint128,uint128,int128),int256,int256,int256,uint256,uint256,uint256,bool))
func (_Abis *AbisCaller) PositionDetailsForMarketKey(opts *bind.CallOpts, marketKey [32]byte, account common.Address) (PerpsV2MarketDataPositionData, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "positionDetailsForMarketKey", marketKey, account)

	if err != nil {
		return *new(PerpsV2MarketDataPositionData), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpsV2MarketDataPositionData)).(*PerpsV2MarketDataPositionData)

	return out0, err

}

// PositionDetailsForMarketKey is a free data retrieval call binding the contract method 0x985f289e.
//
// Solidity: function positionDetailsForMarketKey(bytes32 marketKey, address account) view returns(((uint64,uint64,uint128,uint128,int128),int256,int256,int256,uint256,uint256,uint256,bool))
func (_Abis *AbisSession) PositionDetailsForMarketKey(marketKey [32]byte, account common.Address) (PerpsV2MarketDataPositionData, error) {
	return _Abis.Contract.PositionDetailsForMarketKey(&_Abis.CallOpts, marketKey, account)
}

// PositionDetailsForMarketKey is a free data retrieval call binding the contract method 0x985f289e.
//
// Solidity: function positionDetailsForMarketKey(bytes32 marketKey, address account) view returns(((uint64,uint64,uint128,uint128,int128),int256,int256,int256,uint256,uint256,uint256,bool))
func (_Abis *AbisCallerSession) PositionDetailsForMarketKey(marketKey [32]byte, account common.Address) (PerpsV2MarketDataPositionData, error) {
	return _Abis.Contract.PositionDetailsForMarketKey(&_Abis.CallOpts, marketKey, account)
}

// ResolverProxy is a free data retrieval call binding the contract method 0x6a59e495.
//
// Solidity: function resolverProxy() view returns(address)
func (_Abis *AbisCaller) ResolverProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "resolverProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolverProxy is a free data retrieval call binding the contract method 0x6a59e495.
//
// Solidity: function resolverProxy() view returns(address)
func (_Abis *AbisSession) ResolverProxy() (common.Address, error) {
	return _Abis.Contract.ResolverProxy(&_Abis.CallOpts)
}

// ResolverProxy is a free data retrieval call binding the contract method 0x6a59e495.
//
// Solidity: function resolverProxy() view returns(address)
func (_Abis *AbisCallerSession) ResolverProxy() (common.Address, error) {
	return _Abis.Contract.ResolverProxy(&_Abis.CallOpts)
}

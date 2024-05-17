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

// DepositAddresses is an auto generated low-level Go binding around an user-defined struct.
type DepositAddresses struct {
	Account            common.Address
	Receiver           common.Address
	CallbackContract   common.Address
	UiFeeReceiver      common.Address
	Market             common.Address
	InitialLongToken   common.Address
	InitialShortToken  common.Address
	LongTokenSwapPath  []common.Address
	ShortTokenSwapPath []common.Address
}

// DepositFlags is an auto generated low-level Go binding around an user-defined struct.
type DepositFlags struct {
	ShouldUnwrapNativeToken bool
}

// DepositNumbers is an auto generated low-level Go binding around an user-defined struct.
type DepositNumbers struct {
	InitialLongTokenAmount  *big.Int
	InitialShortTokenAmount *big.Int
	MinMarketTokens         *big.Int
	UpdatedAtBlock          *big.Int
	ExecutionFee            *big.Int
	CallbackGasLimit        *big.Int
}

// DepositProps is an auto generated low-level Go binding around an user-defined struct.
type DepositProps struct {
	Addresses DepositAddresses
	Numbers   DepositNumbers
	Flags     DepositFlags
}

// MarketPoolValueInfoProps is an auto generated low-level Go binding around an user-defined struct.
type MarketPoolValueInfoProps struct {
	PoolValue              *big.Int
	LongPnl                *big.Int
	ShortPnl               *big.Int
	NetPnl                 *big.Int
	LongTokenAmount        *big.Int
	ShortTokenAmount       *big.Int
	LongTokenUsd           *big.Int
	ShortTokenUsd          *big.Int
	TotalBorrowingFees     *big.Int
	BorrowingFeePoolFactor *big.Int
	ImpactPoolAmount       *big.Int
}

// MarketProps is an auto generated low-level Go binding around an user-defined struct.
type MarketProps struct {
	MarketToken common.Address
	IndexToken  common.Address
	LongToken   common.Address
	ShortToken  common.Address
}

// MarketUtilsCollateralType is an auto generated low-level Go binding around an user-defined struct.
type MarketUtilsCollateralType struct {
	LongToken  *big.Int
	ShortToken *big.Int
}

// MarketUtilsGetNextFundingAmountPerSizeResult is an auto generated low-level Go binding around an user-defined struct.
type MarketUtilsGetNextFundingAmountPerSizeResult struct {
	LongsPayShorts                     bool
	FundingFactorPerSecond             *big.Int
	NextSavedFundingFactorPerSecond    *big.Int
	FundingFeeAmountPerSizeDelta       MarketUtilsPositionType
	ClaimableFundingAmountPerSizeDelta MarketUtilsPositionType
}

// MarketUtilsMarketPrices is an auto generated low-level Go binding around an user-defined struct.
type MarketUtilsMarketPrices struct {
	IndexTokenPrice PriceProps
	LongTokenPrice  PriceProps
	ShortTokenPrice PriceProps
}

// MarketUtilsPositionType is an auto generated low-level Go binding around an user-defined struct.
type MarketUtilsPositionType struct {
	Long  MarketUtilsCollateralType
	Short MarketUtilsCollateralType
}

// OrderAddresses is an auto generated low-level Go binding around an user-defined struct.
type OrderAddresses struct {
	Account                common.Address
	Receiver               common.Address
	CallbackContract       common.Address
	UiFeeReceiver          common.Address
	Market                 common.Address
	InitialCollateralToken common.Address
	SwapPath               []common.Address
}

// OrderFlags is an auto generated low-level Go binding around an user-defined struct.
type OrderFlags struct {
	IsLong                  bool
	ShouldUnwrapNativeToken bool
	IsFrozen                bool
}

// OrderNumbers is an auto generated low-level Go binding around an user-defined struct.
type OrderNumbers struct {
	OrderType                    uint8
	DecreasePositionSwapType     uint8
	SizeDeltaUsd                 *big.Int
	InitialCollateralDeltaAmount *big.Int
	TriggerPrice                 *big.Int
	AcceptablePrice              *big.Int
	ExecutionFee                 *big.Int
	CallbackGasLimit             *big.Int
	MinOutputAmount              *big.Int
	UpdatedAtBlock               *big.Int
}

// OrderProps is an auto generated low-level Go binding around an user-defined struct.
type OrderProps struct {
	Addresses OrderAddresses
	Numbers   OrderNumbers
	Flags     OrderFlags
}

// PositionAddresses is an auto generated low-level Go binding around an user-defined struct.
type PositionAddresses struct {
	Account         common.Address
	Market          common.Address
	CollateralToken common.Address
}

// PositionFlags is an auto generated low-level Go binding around an user-defined struct.
type PositionFlags struct {
	IsLong bool
}

// PositionNumbers is an auto generated low-level Go binding around an user-defined struct.
type PositionNumbers struct {
	SizeInUsd                               *big.Int
	SizeInTokens                            *big.Int
	CollateralAmount                        *big.Int
	BorrowingFactor                         *big.Int
	FundingFeeAmountPerSize                 *big.Int
	LongTokenClaimableFundingAmountPerSize  *big.Int
	ShortTokenClaimableFundingAmountPerSize *big.Int
	IncreasedAtBlock                        *big.Int
	DecreasedAtBlock                        *big.Int
}

// PositionPricingUtilsPositionBorrowingFees is an auto generated low-level Go binding around an user-defined struct.
type PositionPricingUtilsPositionBorrowingFees struct {
	BorrowingFeeUsd                  *big.Int
	BorrowingFeeAmount               *big.Int
	BorrowingFeeReceiverFactor       *big.Int
	BorrowingFeeAmountForFeeReceiver *big.Int
}

// PositionPricingUtilsPositionFees is an auto generated low-level Go binding around an user-defined struct.
type PositionPricingUtilsPositionFees struct {
	Referral                        PositionPricingUtilsPositionReferralFees
	Funding                         PositionPricingUtilsPositionFundingFees
	Borrowing                       PositionPricingUtilsPositionBorrowingFees
	Ui                              PositionPricingUtilsPositionUiFees
	CollateralTokenPrice            PriceProps
	PositionFeeFactor               *big.Int
	ProtocolFeeAmount               *big.Int
	PositionFeeReceiverFactor       *big.Int
	FeeReceiverAmount               *big.Int
	FeeAmountForPool                *big.Int
	PositionFeeAmountForPool        *big.Int
	PositionFeeAmount               *big.Int
	TotalCostAmountExcludingFunding *big.Int
	TotalCostAmount                 *big.Int
}

// PositionPricingUtilsPositionFundingFees is an auto generated low-level Go binding around an user-defined struct.
type PositionPricingUtilsPositionFundingFees struct {
	FundingFeeAmount                              *big.Int
	ClaimableLongTokenAmount                      *big.Int
	ClaimableShortTokenAmount                     *big.Int
	LatestFundingFeeAmountPerSize                 *big.Int
	LatestLongTokenClaimableFundingAmountPerSize  *big.Int
	LatestShortTokenClaimableFundingAmountPerSize *big.Int
}

// PositionPricingUtilsPositionReferralFees is an auto generated low-level Go binding around an user-defined struct.
type PositionPricingUtilsPositionReferralFees struct {
	ReferralCode          [32]byte
	Affiliate             common.Address
	Trader                common.Address
	TotalRebateFactor     *big.Int
	TraderDiscountFactor  *big.Int
	TotalRebateAmount     *big.Int
	TraderDiscountAmount  *big.Int
	AffiliateRewardAmount *big.Int
}

// PositionPricingUtilsPositionUiFees is an auto generated low-level Go binding around an user-defined struct.
type PositionPricingUtilsPositionUiFees struct {
	UiFeeReceiver       common.Address
	UiFeeReceiverFactor *big.Int
	UiFeeAmount         *big.Int
}

// PositionProps is an auto generated low-level Go binding around an user-defined struct.
type PositionProps struct {
	Addresses PositionAddresses
	Numbers   PositionNumbers
	Flags     PositionFlags
}

// PriceProps is an auto generated low-level Go binding around an user-defined struct.
type PriceProps struct {
	Min *big.Int
	Max *big.Int
}

// ReaderPricingUtilsExecutionPriceResult is an auto generated low-level Go binding around an user-defined struct.
type ReaderPricingUtilsExecutionPriceResult struct {
	PriceImpactUsd     *big.Int
	PriceImpactDiffUsd *big.Int
	ExecutionPrice     *big.Int
}

// ReaderUtilsBaseFundingValues is an auto generated low-level Go binding around an user-defined struct.
type ReaderUtilsBaseFundingValues struct {
	FundingFeeAmountPerSize       MarketUtilsPositionType
	ClaimableFundingAmountPerSize MarketUtilsPositionType
}

// ReaderUtilsMarketInfo is an auto generated low-level Go binding around an user-defined struct.
type ReaderUtilsMarketInfo struct {
	Market                            MarketProps
	BorrowingFactorPerSecondForLongs  *big.Int
	BorrowingFactorPerSecondForShorts *big.Int
	BaseFunding                       ReaderUtilsBaseFundingValues
	NextFunding                       MarketUtilsGetNextFundingAmountPerSizeResult
	VirtualInventory                  ReaderUtilsVirtualInventory
	IsDisabled                        bool
}

// ReaderUtilsPositionInfo is an auto generated low-level Go binding around an user-defined struct.
type ReaderUtilsPositionInfo struct {
	Position               PositionProps
	Fees                   PositionPricingUtilsPositionFees
	ExecutionPriceResult   ReaderPricingUtilsExecutionPriceResult
	BasePnlUsd             *big.Int
	UncappedBasePnlUsd     *big.Int
	PnlAfterPriceImpactUsd *big.Int
}

// ReaderUtilsVirtualInventory is an auto generated low-level Go binding around an user-defined struct.
type ReaderUtilsVirtualInventory struct {
	VirtualPoolAmountForLongToken  *big.Int
	VirtualPoolAmountForShortToken *big.Int
	VirtualInventoryForPositions   *big.Int
}

// SwapPricingUtilsSwapFees is an auto generated low-level Go binding around an user-defined struct.
type SwapPricingUtilsSwapFees struct {
	FeeReceiverAmount   *big.Int
	FeeAmountForPool    *big.Int
	AmountAfterFees     *big.Int
	UiFeeReceiver       common.Address
	UiFeeReceiverFactor *big.Int
	UiFeeAmount         *big.Int
}

// WithdrawalAddresses is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalAddresses struct {
	Account            common.Address
	Receiver           common.Address
	CallbackContract   common.Address
	UiFeeReceiver      common.Address
	Market             common.Address
	LongTokenSwapPath  []common.Address
	ShortTokenSwapPath []common.Address
}

// WithdrawalFlags is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalFlags struct {
	ShouldUnwrapNativeToken bool
}

// WithdrawalNumbers is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalNumbers struct {
	MarketTokenAmount   *big.Int
	MinLongTokenAmount  *big.Int
	MinShortTokenAmount *big.Int
	UpdatedAtBlock      *big.Int
	ExecutionFee        *big.Int
	CallbackGasLimit    *big.Int
}

// WithdrawalProps is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalProps struct {
	Addresses WithdrawalAddresses
	Numbers   WithdrawalNumbers
	Flags     WithdrawalFlags
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"DisabledMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyMarket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getAccountOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialCollateralToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"swapPath\",\"type\":\"address[]\"}],\"internalType\":\"structOrder.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrder.DecreasePositionSwapType\",\"name\":\"decreasePositionSwapType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialCollateralDeltaAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"internalType\":\"structOrder.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structOrder.Props[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"contractIReferralStorage\",\"name\":\"referralStorage\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"positionKeys\",\"type\":\"bytes32[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices[]\",\"name\":\"prices\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"}],\"name\":\"getAccountPositionInfoList\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"internalType\":\"structPosition.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sizeInUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizeInTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingFeeAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"increasedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decreasedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structPosition.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"internalType\":\"structPosition.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structPosition.Props\",\"name\":\"position\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"referralCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalRebateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"traderDiscountFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRebateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"traderDiscountAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"affiliateRewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionReferralFees\",\"name\":\"referral\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fundingFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableLongTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableShortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestFundingFeeAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestLongTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestShortTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionFundingFees\",\"name\":\"funding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"borrowingFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeeAmountForFeeReceiver\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionBorrowingFees\",\"name\":\"borrowing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"uiFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uiFeeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionUiFees\",\"name\":\"ui\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"collateralTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiverAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmountForPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeAmountForPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCostAmountExcludingFunding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCostAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionFees\",\"name\":\"fees\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"priceImpactUsd\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactDiffUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"}],\"internalType\":\"structReaderPricingUtils.ExecutionPriceResult\",\"name\":\"executionPriceResult\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"basePnlUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"uncappedBasePnlUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"pnlAfterPriceImpactUsd\",\"type\":\"int256\"}],\"internalType\":\"structReaderUtils.PositionInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getAccountPositions\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"internalType\":\"structPosition.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sizeInUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizeInTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingFeeAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"increasedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decreasedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structPosition.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"internalType\":\"structPosition.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structPosition.Props[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"}],\"name\":\"getAdlState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getDeposit\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialLongToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialShortToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"longTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\"}],\"internalType\":\"structDeposit.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"initialLongTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialShortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMarketTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDeposit.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"}],\"internalType\":\"structDeposit.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structDeposit.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"longTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"}],\"name\":\"getDepositAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketKey\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeInUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeInTokens\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"sizeDeltaUsd\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"name\":\"getExecutionPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"priceImpactUsd\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactDiffUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"}],\"internalType\":\"structReaderPricingUtils.ExecutionPriceResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"getMarketBySalt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"marketKey\",\"type\":\"address\"}],\"name\":\"getMarketInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactorPerSecondForLongs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactorPerSecondForShorts\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"fundingFeeAmountPerSize\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"claimableFundingAmountPerSize\",\"type\":\"tuple\"}],\"internalType\":\"structReaderUtils.BaseFundingValues\",\"name\":\"baseFunding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"longsPayShorts\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fundingFactorPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"nextSavedFundingFactorPerSecond\",\"type\":\"int256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"fundingFeeAmountPerSizeDelta\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"claimableFundingAmountPerSizeDelta\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.GetNextFundingAmountPerSizeResult\",\"name\":\"nextFunding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"virtualPoolAmountForLongToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"virtualPoolAmountForShortToken\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"virtualInventoryForPositions\",\"type\":\"int256\"}],\"internalType\":\"structReaderUtils.VirtualInventory\",\"name\":\"virtualInventory\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structReaderUtils.MarketInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices[]\",\"name\":\"marketPricesList\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getMarketInfoList\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactorPerSecondForLongs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactorPerSecondForShorts\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"fundingFeeAmountPerSize\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"claimableFundingAmountPerSize\",\"type\":\"tuple\"}],\"internalType\":\"structReaderUtils.BaseFundingValues\",\"name\":\"baseFunding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"longsPayShorts\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fundingFactorPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"nextSavedFundingFactorPerSecond\",\"type\":\"int256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"fundingFeeAmountPerSizeDelta\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortToken\",\"type\":\"uint256\"}],\"internalType\":\"structMarketUtils.CollateralType\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.PositionType\",\"name\":\"claimableFundingAmountPerSizeDelta\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.GetNextFundingAmountPerSizeResult\",\"name\":\"nextFunding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"virtualPoolAmountForLongToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"virtualPoolAmountForShortToken\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"virtualInventoryForPositions\",\"type\":\"int256\"}],\"internalType\":\"structReaderUtils.VirtualInventory\",\"name\":\"virtualInventory\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structReaderUtils.MarketInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"pnlFactorType\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"maximize\",\"type\":\"bool\"}],\"name\":\"getMarketTokenPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"poolValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longPnl\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"shortPnl\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"netPnl\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"longTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longTokenUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowingFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeePoolFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"impactPoolAmount\",\"type\":\"uint256\"}],\"internalType\":\"structMarketPoolValueInfo.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getMarkets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"maximize\",\"type\":\"bool\"}],\"name\":\"getNetPnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maximize\",\"type\":\"bool\"}],\"name\":\"getOpenInterestWithPnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialCollateralToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"swapPath\",\"type\":\"address[]\"}],\"internalType\":\"structOrder.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrder.DecreasePositionSwapType\",\"name\":\"decreasePositionSwapType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialCollateralDeltaAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"internalType\":\"structOrder.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structOrder.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maximize\",\"type\":\"bool\"}],\"name\":\"getPnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketAddress\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maximize\",\"type\":\"bool\"}],\"name\":\"getPnlToPoolFactor\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getPosition\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"internalType\":\"structPosition.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sizeInUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizeInTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingFeeAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"increasedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decreasedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structPosition.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"internalType\":\"structPosition.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structPosition.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"contractIReferralStorage\",\"name\":\"referralStorage\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"positionKey\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"usePositionSizeAsSizeDeltaUsd\",\"type\":\"bool\"}],\"name\":\"getPositionInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"internalType\":\"structPosition.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sizeInUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizeInTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingFeeAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"increasedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decreasedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structPosition.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"internalType\":\"structPosition.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structPosition.Props\",\"name\":\"position\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"referralCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalRebateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"traderDiscountFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRebateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"traderDiscountAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"affiliateRewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionReferralFees\",\"name\":\"referral\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fundingFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableLongTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableShortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestFundingFeeAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestLongTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestShortTokenClaimableFundingAmountPerSize\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionFundingFees\",\"name\":\"funding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"borrowingFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingFeeAmountForFeeReceiver\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionBorrowingFees\",\"name\":\"borrowing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"uiFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uiFeeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionUiFees\",\"name\":\"ui\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"collateralTokenPrice\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiverAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmountForPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeAmountForPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCostAmountExcludingFunding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCostAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPositionPricingUtils.PositionFees\",\"name\":\"fees\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"priceImpactUsd\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactDiffUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"}],\"internalType\":\"structReaderPricingUtils.ExecutionPriceResult\",\"name\":\"executionPriceResult\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"basePnlUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"uncappedBasePnlUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"pnlAfterPriceImpactUsd\",\"type\":\"int256\"}],\"internalType\":\"structReaderUtils.PositionInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"positionKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\"}],\"name\":\"getPositionPnlUsd\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"}],\"name\":\"getSwapAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeReceiverAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmountForPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAfterFees\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"uiFeeReceiverFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uiFeeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structSwapPricingUtils.SwapFees\",\"name\":\"fees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"tokenInPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"tokenOutPrice\",\"type\":\"tuple\"}],\"name\":\"getSwapPriceImpact\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getWithdrawal\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"longTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\"}],\"internalType\":\"structWithdrawal.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLongTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minShortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Numbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"}],\"internalType\":\"structWithdrawal.Flags\",\"name\":\"flags\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawal.Props\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"dataStore\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"marketToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"}],\"internalType\":\"structMarket.Props\",\"name\":\"market\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"indexTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"longTokenPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props\",\"name\":\"shortTokenPrice\",\"type\":\"tuple\"}],\"internalType\":\"structMarketUtils.MarketPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"marketTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"}],\"name\":\"getWithdrawalAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}],",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// GetAccountOrders is a free data retrieval call binding the contract method 0x42a6f8d3.
//
// Solidity: function getAccountOrders(address dataStore, address account, uint256 start, uint256 end) view returns(((address,address,address,address,address,address,address[]),(uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool,bool,bool))[])
func (_Main *MainCaller) GetAccountOrders(opts *bind.CallOpts, dataStore common.Address, account common.Address, start *big.Int, end *big.Int) ([]OrderProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAccountOrders", dataStore, account, start, end)

	if err != nil {
		return *new([]OrderProps), err
	}

	out0 := *abi.ConvertType(out[0], new([]OrderProps)).(*[]OrderProps)

	return out0, err

}

// GetAccountOrders is a free data retrieval call binding the contract method 0x42a6f8d3.
//
// Solidity: function getAccountOrders(address dataStore, address account, uint256 start, uint256 end) view returns(((address,address,address,address,address,address,address[]),(uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool,bool,bool))[])
func (_Main *MainSession) GetAccountOrders(dataStore common.Address, account common.Address, start *big.Int, end *big.Int) ([]OrderProps, error) {
	return _Main.Contract.GetAccountOrders(&_Main.CallOpts, dataStore, account, start, end)
}

// GetAccountOrders is a free data retrieval call binding the contract method 0x42a6f8d3.
//
// Solidity: function getAccountOrders(address dataStore, address account, uint256 start, uint256 end) view returns(((address,address,address,address,address,address,address[]),(uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool,bool,bool))[])
func (_Main *MainCallerSession) GetAccountOrders(dataStore common.Address, account common.Address, start *big.Int, end *big.Int) ([]OrderProps, error) {
	return _Main.Contract.GetAccountOrders(&_Main.CallOpts, dataStore, account, start, end)
}

// GetAccountPositionInfoList is a free data retrieval call binding the contract method 0xece9e0c8.
//
// Solidity: function getAccountPositionInfoList(address dataStore, address referralStorage, bytes32[] positionKeys, ((uint256,uint256),(uint256,uint256),(uint256,uint256))[] prices, address uiFeeReceiver) view returns((((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)),((bytes32,address,address,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(address,uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,uint256,uint256),int256,int256,int256)[])
func (_Main *MainCaller) GetAccountPositionInfoList(opts *bind.CallOpts, dataStore common.Address, referralStorage common.Address, positionKeys [][32]byte, prices []MarketUtilsMarketPrices, uiFeeReceiver common.Address) ([]ReaderUtilsPositionInfo, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAccountPositionInfoList", dataStore, referralStorage, positionKeys, prices, uiFeeReceiver)

	if err != nil {
		return *new([]ReaderUtilsPositionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ReaderUtilsPositionInfo)).(*[]ReaderUtilsPositionInfo)

	return out0, err

}

// GetAccountPositionInfoList is a free data retrieval call binding the contract method 0xece9e0c8.
//
// Solidity: function getAccountPositionInfoList(address dataStore, address referralStorage, bytes32[] positionKeys, ((uint256,uint256),(uint256,uint256),(uint256,uint256))[] prices, address uiFeeReceiver) view returns((((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)),((bytes32,address,address,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(address,uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,uint256,uint256),int256,int256,int256)[])
func (_Main *MainSession) GetAccountPositionInfoList(dataStore common.Address, referralStorage common.Address, positionKeys [][32]byte, prices []MarketUtilsMarketPrices, uiFeeReceiver common.Address) ([]ReaderUtilsPositionInfo, error) {
	return _Main.Contract.GetAccountPositionInfoList(&_Main.CallOpts, dataStore, referralStorage, positionKeys, prices, uiFeeReceiver)
}

// GetAccountPositionInfoList is a free data retrieval call binding the contract method 0xece9e0c8.
//
// Solidity: function getAccountPositionInfoList(address dataStore, address referralStorage, bytes32[] positionKeys, ((uint256,uint256),(uint256,uint256),(uint256,uint256))[] prices, address uiFeeReceiver) view returns((((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)),((bytes32,address,address,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(address,uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,uint256,uint256),int256,int256,int256)[])
func (_Main *MainCallerSession) GetAccountPositionInfoList(dataStore common.Address, referralStorage common.Address, positionKeys [][32]byte, prices []MarketUtilsMarketPrices, uiFeeReceiver common.Address) ([]ReaderUtilsPositionInfo, error) {
	return _Main.Contract.GetAccountPositionInfoList(&_Main.CallOpts, dataStore, referralStorage, positionKeys, prices, uiFeeReceiver)
}

// GetAccountPositions is a free data retrieval call binding the contract method 0x77cfb162.
//
// Solidity: function getAccountPositions(address dataStore, address account, uint256 start, uint256 end) view returns(((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool))[])
func (_Main *MainCaller) GetAccountPositions(opts *bind.CallOpts, dataStore common.Address, account common.Address, start *big.Int, end *big.Int) ([]PositionProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAccountPositions", dataStore, account, start, end)

	if err != nil {
		return *new([]PositionProps), err
	}

	out0 := *abi.ConvertType(out[0], new([]PositionProps)).(*[]PositionProps)

	return out0, err

}

// GetAccountPositions is a free data retrieval call binding the contract method 0x77cfb162.
//
// Solidity: function getAccountPositions(address dataStore, address account, uint256 start, uint256 end) view returns(((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool))[])
func (_Main *MainSession) GetAccountPositions(dataStore common.Address, account common.Address, start *big.Int, end *big.Int) ([]PositionProps, error) {
	return _Main.Contract.GetAccountPositions(&_Main.CallOpts, dataStore, account, start, end)
}

// GetAccountPositions is a free data retrieval call binding the contract method 0x77cfb162.
//
// Solidity: function getAccountPositions(address dataStore, address account, uint256 start, uint256 end) view returns(((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool))[])
func (_Main *MainCallerSession) GetAccountPositions(dataStore common.Address, account common.Address, start *big.Int, end *big.Int) ([]PositionProps, error) {
	return _Main.Contract.GetAccountPositions(&_Main.CallOpts, dataStore, account, start, end)
}

// GetAdlState is a free data retrieval call binding the contract method 0x2b17b4fd.
//
// Solidity: function getAdlState(address dataStore, address market, bool isLong, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices) view returns(uint256, bool, int256, uint256)
func (_Main *MainCaller) GetAdlState(opts *bind.CallOpts, dataStore common.Address, market common.Address, isLong bool, prices MarketUtilsMarketPrices) (*big.Int, bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAdlState", dataStore, market, isLong, prices)

	if err != nil {
		return *new(*big.Int), *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAdlState is a free data retrieval call binding the contract method 0x2b17b4fd.
//
// Solidity: function getAdlState(address dataStore, address market, bool isLong, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices) view returns(uint256, bool, int256, uint256)
func (_Main *MainSession) GetAdlState(dataStore common.Address, market common.Address, isLong bool, prices MarketUtilsMarketPrices) (*big.Int, bool, *big.Int, *big.Int, error) {
	return _Main.Contract.GetAdlState(&_Main.CallOpts, dataStore, market, isLong, prices)
}

// GetAdlState is a free data retrieval call binding the contract method 0x2b17b4fd.
//
// Solidity: function getAdlState(address dataStore, address market, bool isLong, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices) view returns(uint256, bool, int256, uint256)
func (_Main *MainCallerSession) GetAdlState(dataStore common.Address, market common.Address, isLong bool, prices MarketUtilsMarketPrices) (*big.Int, bool, *big.Int, *big.Int, error) {
	return _Main.Contract.GetAdlState(&_Main.CallOpts, dataStore, market, isLong, prices)
}

// GetDeposit is a free data retrieval call binding the contract method 0x1485d297.
//
// Solidity: function getDeposit(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address,address,address[],address[]),(uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainCaller) GetDeposit(opts *bind.CallOpts, dataStore common.Address, key [32]byte) (DepositProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getDeposit", dataStore, key)

	if err != nil {
		return *new(DepositProps), err
	}

	out0 := *abi.ConvertType(out[0], new(DepositProps)).(*DepositProps)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0x1485d297.
//
// Solidity: function getDeposit(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address,address,address[],address[]),(uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainSession) GetDeposit(dataStore common.Address, key [32]byte) (DepositProps, error) {
	return _Main.Contract.GetDeposit(&_Main.CallOpts, dataStore, key)
}

// GetDeposit is a free data retrieval call binding the contract method 0x1485d297.
//
// Solidity: function getDeposit(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address,address,address[],address[]),(uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainCallerSession) GetDeposit(dataStore common.Address, key [32]byte) (DepositProps, error) {
	return _Main.Contract.GetDeposit(&_Main.CallOpts, dataStore, key)
}

// GetDepositAmountOut is a free data retrieval call binding the contract method 0x85874c7f.
//
// Solidity: function getDepositAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 longTokenAmount, uint256 shortTokenAmount, address uiFeeReceiver) view returns(uint256)
func (_Main *MainCaller) GetDepositAmountOut(opts *bind.CallOpts, dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, longTokenAmount *big.Int, shortTokenAmount *big.Int, uiFeeReceiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getDepositAmountOut", dataStore, market, prices, longTokenAmount, shortTokenAmount, uiFeeReceiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmountOut is a free data retrieval call binding the contract method 0x85874c7f.
//
// Solidity: function getDepositAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 longTokenAmount, uint256 shortTokenAmount, address uiFeeReceiver) view returns(uint256)
func (_Main *MainSession) GetDepositAmountOut(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, longTokenAmount *big.Int, shortTokenAmount *big.Int, uiFeeReceiver common.Address) (*big.Int, error) {
	return _Main.Contract.GetDepositAmountOut(&_Main.CallOpts, dataStore, market, prices, longTokenAmount, shortTokenAmount, uiFeeReceiver)
}

// GetDepositAmountOut is a free data retrieval call binding the contract method 0x85874c7f.
//
// Solidity: function getDepositAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 longTokenAmount, uint256 shortTokenAmount, address uiFeeReceiver) view returns(uint256)
func (_Main *MainCallerSession) GetDepositAmountOut(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, longTokenAmount *big.Int, shortTokenAmount *big.Int, uiFeeReceiver common.Address) (*big.Int, error) {
	return _Main.Contract.GetDepositAmountOut(&_Main.CallOpts, dataStore, market, prices, longTokenAmount, shortTokenAmount, uiFeeReceiver)
}

// GetExecutionPrice is a free data retrieval call binding the contract method 0x5d2b44f9.
//
// Solidity: function getExecutionPrice(address dataStore, address marketKey, (uint256,uint256) indexTokenPrice, uint256 positionSizeInUsd, uint256 positionSizeInTokens, int256 sizeDeltaUsd, bool isLong) view returns((int256,uint256,uint256))
func (_Main *MainCaller) GetExecutionPrice(opts *bind.CallOpts, dataStore common.Address, marketKey common.Address, indexTokenPrice PriceProps, positionSizeInUsd *big.Int, positionSizeInTokens *big.Int, sizeDeltaUsd *big.Int, isLong bool) (ReaderPricingUtilsExecutionPriceResult, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getExecutionPrice", dataStore, marketKey, indexTokenPrice, positionSizeInUsd, positionSizeInTokens, sizeDeltaUsd, isLong)

	if err != nil {
		return *new(ReaderPricingUtilsExecutionPriceResult), err
	}

	out0 := *abi.ConvertType(out[0], new(ReaderPricingUtilsExecutionPriceResult)).(*ReaderPricingUtilsExecutionPriceResult)

	return out0, err

}

// GetExecutionPrice is a free data retrieval call binding the contract method 0x5d2b44f9.
//
// Solidity: function getExecutionPrice(address dataStore, address marketKey, (uint256,uint256) indexTokenPrice, uint256 positionSizeInUsd, uint256 positionSizeInTokens, int256 sizeDeltaUsd, bool isLong) view returns((int256,uint256,uint256))
func (_Main *MainSession) GetExecutionPrice(dataStore common.Address, marketKey common.Address, indexTokenPrice PriceProps, positionSizeInUsd *big.Int, positionSizeInTokens *big.Int, sizeDeltaUsd *big.Int, isLong bool) (ReaderPricingUtilsExecutionPriceResult, error) {
	return _Main.Contract.GetExecutionPrice(&_Main.CallOpts, dataStore, marketKey, indexTokenPrice, positionSizeInUsd, positionSizeInTokens, sizeDeltaUsd, isLong)
}

// GetExecutionPrice is a free data retrieval call binding the contract method 0x5d2b44f9.
//
// Solidity: function getExecutionPrice(address dataStore, address marketKey, (uint256,uint256) indexTokenPrice, uint256 positionSizeInUsd, uint256 positionSizeInTokens, int256 sizeDeltaUsd, bool isLong) view returns((int256,uint256,uint256))
func (_Main *MainCallerSession) GetExecutionPrice(dataStore common.Address, marketKey common.Address, indexTokenPrice PriceProps, positionSizeInUsd *big.Int, positionSizeInTokens *big.Int, sizeDeltaUsd *big.Int, isLong bool) (ReaderPricingUtilsExecutionPriceResult, error) {
	return _Main.Contract.GetExecutionPrice(&_Main.CallOpts, dataStore, marketKey, indexTokenPrice, positionSizeInUsd, positionSizeInTokens, sizeDeltaUsd, isLong)
}

// GetMarket is a free data retrieval call binding the contract method 0x714af34b.
//
// Solidity: function getMarket(address dataStore, address key) view returns((address,address,address,address))
func (_Main *MainCaller) GetMarket(opts *bind.CallOpts, dataStore common.Address, key common.Address) (MarketProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMarket", dataStore, key)

	if err != nil {
		return *new(MarketProps), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketProps)).(*MarketProps)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0x714af34b.
//
// Solidity: function getMarket(address dataStore, address key) view returns((address,address,address,address))
func (_Main *MainSession) GetMarket(dataStore common.Address, key common.Address) (MarketProps, error) {
	return _Main.Contract.GetMarket(&_Main.CallOpts, dataStore, key)
}

// GetMarket is a free data retrieval call binding the contract method 0x714af34b.
//
// Solidity: function getMarket(address dataStore, address key) view returns((address,address,address,address))
func (_Main *MainCallerSession) GetMarket(dataStore common.Address, key common.Address) (MarketProps, error) {
	return _Main.Contract.GetMarket(&_Main.CallOpts, dataStore, key)
}

// GetMarketBySalt is a free data retrieval call binding the contract method 0xa4f0d550.
//
// Solidity: function getMarketBySalt(address dataStore, bytes32 salt) view returns((address,address,address,address))
func (_Main *MainCaller) GetMarketBySalt(opts *bind.CallOpts, dataStore common.Address, salt [32]byte) (MarketProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMarketBySalt", dataStore, salt)

	if err != nil {
		return *new(MarketProps), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketProps)).(*MarketProps)

	return out0, err

}

// GetMarketBySalt is a free data retrieval call binding the contract method 0xa4f0d550.
//
// Solidity: function getMarketBySalt(address dataStore, bytes32 salt) view returns((address,address,address,address))
func (_Main *MainSession) GetMarketBySalt(dataStore common.Address, salt [32]byte) (MarketProps, error) {
	return _Main.Contract.GetMarketBySalt(&_Main.CallOpts, dataStore, salt)
}

// GetMarketBySalt is a free data retrieval call binding the contract method 0xa4f0d550.
//
// Solidity: function getMarketBySalt(address dataStore, bytes32 salt) view returns((address,address,address,address))
func (_Main *MainCallerSession) GetMarketBySalt(dataStore common.Address, salt [32]byte) (MarketProps, error) {
	return _Main.Contract.GetMarketBySalt(&_Main.CallOpts, dataStore, salt)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x847bb469.
//
// Solidity: function getMarketInfo(address dataStore, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, address marketKey) view returns(((address,address,address,address),uint256,uint256,(((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(bool,uint256,int256,((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(uint256,uint256,int256),bool))
func (_Main *MainCaller) GetMarketInfo(opts *bind.CallOpts, dataStore common.Address, prices MarketUtilsMarketPrices, marketKey common.Address) (ReaderUtilsMarketInfo, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMarketInfo", dataStore, prices, marketKey)

	if err != nil {
		return *new(ReaderUtilsMarketInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ReaderUtilsMarketInfo)).(*ReaderUtilsMarketInfo)

	return out0, err

}

// GetMarketInfo is a free data retrieval call binding the contract method 0x847bb469.
//
// Solidity: function getMarketInfo(address dataStore, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, address marketKey) view returns(((address,address,address,address),uint256,uint256,(((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(bool,uint256,int256,((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(uint256,uint256,int256),bool))
func (_Main *MainSession) GetMarketInfo(dataStore common.Address, prices MarketUtilsMarketPrices, marketKey common.Address) (ReaderUtilsMarketInfo, error) {
	return _Main.Contract.GetMarketInfo(&_Main.CallOpts, dataStore, prices, marketKey)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x847bb469.
//
// Solidity: function getMarketInfo(address dataStore, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, address marketKey) view returns(((address,address,address,address),uint256,uint256,(((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(bool,uint256,int256,((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(uint256,uint256,int256),bool))
func (_Main *MainCallerSession) GetMarketInfo(dataStore common.Address, prices MarketUtilsMarketPrices, marketKey common.Address) (ReaderUtilsMarketInfo, error) {
	return _Main.Contract.GetMarketInfo(&_Main.CallOpts, dataStore, prices, marketKey)
}

// GetMarketInfoList is a free data retrieval call binding the contract method 0xbc7b5bba.
//
// Solidity: function getMarketInfoList(address dataStore, ((uint256,uint256),(uint256,uint256),(uint256,uint256))[] marketPricesList, uint256 start, uint256 end) view returns(((address,address,address,address),uint256,uint256,(((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(bool,uint256,int256,((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(uint256,uint256,int256),bool)[])
func (_Main *MainCaller) GetMarketInfoList(opts *bind.CallOpts, dataStore common.Address, marketPricesList []MarketUtilsMarketPrices, start *big.Int, end *big.Int) ([]ReaderUtilsMarketInfo, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMarketInfoList", dataStore, marketPricesList, start, end)

	if err != nil {
		return *new([]ReaderUtilsMarketInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ReaderUtilsMarketInfo)).(*[]ReaderUtilsMarketInfo)

	return out0, err

}

// GetMarketInfoList is a free data retrieval call binding the contract method 0xbc7b5bba.
//
// Solidity: function getMarketInfoList(address dataStore, ((uint256,uint256),(uint256,uint256),(uint256,uint256))[] marketPricesList, uint256 start, uint256 end) view returns(((address,address,address,address),uint256,uint256,(((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(bool,uint256,int256,((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(uint256,uint256,int256),bool)[])
func (_Main *MainSession) GetMarketInfoList(dataStore common.Address, marketPricesList []MarketUtilsMarketPrices, start *big.Int, end *big.Int) ([]ReaderUtilsMarketInfo, error) {
	return _Main.Contract.GetMarketInfoList(&_Main.CallOpts, dataStore, marketPricesList, start, end)
}

// GetMarketInfoList is a free data retrieval call binding the contract method 0xbc7b5bba.
//
// Solidity: function getMarketInfoList(address dataStore, ((uint256,uint256),(uint256,uint256),(uint256,uint256))[] marketPricesList, uint256 start, uint256 end) view returns(((address,address,address,address),uint256,uint256,(((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(bool,uint256,int256,((uint256,uint256),(uint256,uint256)),((uint256,uint256),(uint256,uint256))),(uint256,uint256,int256),bool)[])
func (_Main *MainCallerSession) GetMarketInfoList(dataStore common.Address, marketPricesList []MarketUtilsMarketPrices, start *big.Int, end *big.Int) ([]ReaderUtilsMarketInfo, error) {
	return _Main.Contract.GetMarketInfoList(&_Main.CallOpts, dataStore, marketPricesList, start, end)
}

// GetMarketTokenPrice is a free data retrieval call binding the contract method 0x095ce6c5.
//
// Solidity: function getMarketTokenPrice(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, (uint256,uint256) longTokenPrice, (uint256,uint256) shortTokenPrice, bytes32 pnlFactorType, bool maximize) view returns(int256, (int256,int256,int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Main *MainCaller) GetMarketTokenPrice(opts *bind.CallOpts, dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, longTokenPrice PriceProps, shortTokenPrice PriceProps, pnlFactorType [32]byte, maximize bool) (*big.Int, MarketPoolValueInfoProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMarketTokenPrice", dataStore, market, indexTokenPrice, longTokenPrice, shortTokenPrice, pnlFactorType, maximize)

	if err != nil {
		return *new(*big.Int), *new(MarketPoolValueInfoProps), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(MarketPoolValueInfoProps)).(*MarketPoolValueInfoProps)

	return out0, out1, err

}

// GetMarketTokenPrice is a free data retrieval call binding the contract method 0x095ce6c5.
//
// Solidity: function getMarketTokenPrice(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, (uint256,uint256) longTokenPrice, (uint256,uint256) shortTokenPrice, bytes32 pnlFactorType, bool maximize) view returns(int256, (int256,int256,int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Main *MainSession) GetMarketTokenPrice(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, longTokenPrice PriceProps, shortTokenPrice PriceProps, pnlFactorType [32]byte, maximize bool) (*big.Int, MarketPoolValueInfoProps, error) {
	return _Main.Contract.GetMarketTokenPrice(&_Main.CallOpts, dataStore, market, indexTokenPrice, longTokenPrice, shortTokenPrice, pnlFactorType, maximize)
}

// GetMarketTokenPrice is a free data retrieval call binding the contract method 0x095ce6c5.
//
// Solidity: function getMarketTokenPrice(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, (uint256,uint256) longTokenPrice, (uint256,uint256) shortTokenPrice, bytes32 pnlFactorType, bool maximize) view returns(int256, (int256,int256,int256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Main *MainCallerSession) GetMarketTokenPrice(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, longTokenPrice PriceProps, shortTokenPrice PriceProps, pnlFactorType [32]byte, maximize bool) (*big.Int, MarketPoolValueInfoProps, error) {
	return _Main.Contract.GetMarketTokenPrice(&_Main.CallOpts, dataStore, market, indexTokenPrice, longTokenPrice, shortTokenPrice, pnlFactorType, maximize)
}

// GetMarkets is a free data retrieval call binding the contract method 0xce3264bf.
//
// Solidity: function getMarkets(address dataStore, uint256 start, uint256 end) view returns((address,address,address,address)[])
func (_Main *MainCaller) GetMarkets(opts *bind.CallOpts, dataStore common.Address, start *big.Int, end *big.Int) ([]MarketProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMarkets", dataStore, start, end)

	if err != nil {
		return *new([]MarketProps), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketProps)).(*[]MarketProps)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0xce3264bf.
//
// Solidity: function getMarkets(address dataStore, uint256 start, uint256 end) view returns((address,address,address,address)[])
func (_Main *MainSession) GetMarkets(dataStore common.Address, start *big.Int, end *big.Int) ([]MarketProps, error) {
	return _Main.Contract.GetMarkets(&_Main.CallOpts, dataStore, start, end)
}

// GetMarkets is a free data retrieval call binding the contract method 0xce3264bf.
//
// Solidity: function getMarkets(address dataStore, uint256 start, uint256 end) view returns((address,address,address,address)[])
func (_Main *MainCallerSession) GetMarkets(dataStore common.Address, start *big.Int, end *big.Int) ([]MarketProps, error) {
	return _Main.Contract.GetMarkets(&_Main.CallOpts, dataStore, start, end)
}

// GetNetPnl is a free data retrieval call binding the contract method 0xfd50649d.
//
// Solidity: function getNetPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool maximize) view returns(int256)
func (_Main *MainCaller) GetNetPnl(opts *bind.CallOpts, dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, maximize bool) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNetPnl", dataStore, market, indexTokenPrice, maximize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetPnl is a free data retrieval call binding the contract method 0xfd50649d.
//
// Solidity: function getNetPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool maximize) view returns(int256)
func (_Main *MainSession) GetNetPnl(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetNetPnl(&_Main.CallOpts, dataStore, market, indexTokenPrice, maximize)
}

// GetNetPnl is a free data retrieval call binding the contract method 0xfd50649d.
//
// Solidity: function getNetPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool maximize) view returns(int256)
func (_Main *MainCallerSession) GetNetPnl(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetNetPnl(&_Main.CallOpts, dataStore, market, indexTokenPrice, maximize)
}

// GetOpenInterestWithPnl is a free data retrieval call binding the contract method 0xa0140938.
//
// Solidity: function getOpenInterestWithPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool isLong, bool maximize) view returns(int256)
func (_Main *MainCaller) GetOpenInterestWithPnl(opts *bind.CallOpts, dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, isLong bool, maximize bool) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getOpenInterestWithPnl", dataStore, market, indexTokenPrice, isLong, maximize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOpenInterestWithPnl is a free data retrieval call binding the contract method 0xa0140938.
//
// Solidity: function getOpenInterestWithPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool isLong, bool maximize) view returns(int256)
func (_Main *MainSession) GetOpenInterestWithPnl(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, isLong bool, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetOpenInterestWithPnl(&_Main.CallOpts, dataStore, market, indexTokenPrice, isLong, maximize)
}

// GetOpenInterestWithPnl is a free data retrieval call binding the contract method 0xa0140938.
//
// Solidity: function getOpenInterestWithPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool isLong, bool maximize) view returns(int256)
func (_Main *MainCallerSession) GetOpenInterestWithPnl(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, isLong bool, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetOpenInterestWithPnl(&_Main.CallOpts, dataStore, market, indexTokenPrice, isLong, maximize)
}

// GetOrder is a free data retrieval call binding the contract method 0x49651b6a.
//
// Solidity: function getOrder(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address,address[]),(uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool,bool,bool)))
func (_Main *MainCaller) GetOrder(opts *bind.CallOpts, dataStore common.Address, key [32]byte) (OrderProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getOrder", dataStore, key)

	if err != nil {
		return *new(OrderProps), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderProps)).(*OrderProps)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x49651b6a.
//
// Solidity: function getOrder(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address,address[]),(uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool,bool,bool)))
func (_Main *MainSession) GetOrder(dataStore common.Address, key [32]byte) (OrderProps, error) {
	return _Main.Contract.GetOrder(&_Main.CallOpts, dataStore, key)
}

// GetOrder is a free data retrieval call binding the contract method 0x49651b6a.
//
// Solidity: function getOrder(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address,address[]),(uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool,bool,bool)))
func (_Main *MainCallerSession) GetOrder(dataStore common.Address, key [32]byte) (OrderProps, error) {
	return _Main.Contract.GetOrder(&_Main.CallOpts, dataStore, key)
}

// GetPnl is a free data retrieval call binding the contract method 0x24c029e0.
//
// Solidity: function getPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool isLong, bool maximize) view returns(int256)
func (_Main *MainCaller) GetPnl(opts *bind.CallOpts, dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, isLong bool, maximize bool) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPnl", dataStore, market, indexTokenPrice, isLong, maximize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPnl is a free data retrieval call binding the contract method 0x24c029e0.
//
// Solidity: function getPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool isLong, bool maximize) view returns(int256)
func (_Main *MainSession) GetPnl(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, isLong bool, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetPnl(&_Main.CallOpts, dataStore, market, indexTokenPrice, isLong, maximize)
}

// GetPnl is a free data retrieval call binding the contract method 0x24c029e0.
//
// Solidity: function getPnl(address dataStore, (address,address,address,address) market, (uint256,uint256) indexTokenPrice, bool isLong, bool maximize) view returns(int256)
func (_Main *MainCallerSession) GetPnl(dataStore common.Address, market MarketProps, indexTokenPrice PriceProps, isLong bool, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetPnl(&_Main.CallOpts, dataStore, market, indexTokenPrice, isLong, maximize)
}

// GetPnlToPoolFactor is a free data retrieval call binding the contract method 0x971de27d.
//
// Solidity: function getPnlToPoolFactor(address dataStore, address marketAddress, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, bool isLong, bool maximize) view returns(int256)
func (_Main *MainCaller) GetPnlToPoolFactor(opts *bind.CallOpts, dataStore common.Address, marketAddress common.Address, prices MarketUtilsMarketPrices, isLong bool, maximize bool) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPnlToPoolFactor", dataStore, marketAddress, prices, isLong, maximize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPnlToPoolFactor is a free data retrieval call binding the contract method 0x971de27d.
//
// Solidity: function getPnlToPoolFactor(address dataStore, address marketAddress, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, bool isLong, bool maximize) view returns(int256)
func (_Main *MainSession) GetPnlToPoolFactor(dataStore common.Address, marketAddress common.Address, prices MarketUtilsMarketPrices, isLong bool, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetPnlToPoolFactor(&_Main.CallOpts, dataStore, marketAddress, prices, isLong, maximize)
}

// GetPnlToPoolFactor is a free data retrieval call binding the contract method 0x971de27d.
//
// Solidity: function getPnlToPoolFactor(address dataStore, address marketAddress, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, bool isLong, bool maximize) view returns(int256)
func (_Main *MainCallerSession) GetPnlToPoolFactor(dataStore common.Address, marketAddress common.Address, prices MarketUtilsMarketPrices, isLong bool, maximize bool) (*big.Int, error) {
	return _Main.Contract.GetPnlToPoolFactor(&_Main.CallOpts, dataStore, marketAddress, prices, isLong, maximize)
}

// GetPosition is a free data retrieval call binding the contract method 0x0fa8f516.
//
// Solidity: function getPosition(address dataStore, bytes32 key) view returns(((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainCaller) GetPosition(opts *bind.CallOpts, dataStore common.Address, key [32]byte) (PositionProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPosition", dataStore, key)

	if err != nil {
		return *new(PositionProps), err
	}

	out0 := *abi.ConvertType(out[0], new(PositionProps)).(*PositionProps)

	return out0, err

}

// GetPosition is a free data retrieval call binding the contract method 0x0fa8f516.
//
// Solidity: function getPosition(address dataStore, bytes32 key) view returns(((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainSession) GetPosition(dataStore common.Address, key [32]byte) (PositionProps, error) {
	return _Main.Contract.GetPosition(&_Main.CallOpts, dataStore, key)
}

// GetPosition is a free data retrieval call binding the contract method 0x0fa8f516.
//
// Solidity: function getPosition(address dataStore, bytes32 key) view returns(((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainCallerSession) GetPosition(dataStore common.Address, key [32]byte) (PositionProps, error) {
	return _Main.Contract.GetPosition(&_Main.CallOpts, dataStore, key)
}

// GetPositionInfo is a free data retrieval call binding the contract method 0x0815bce1.
//
// Solidity: function getPositionInfo(address dataStore, address referralStorage, bytes32 positionKey, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 sizeDeltaUsd, address uiFeeReceiver, bool usePositionSizeAsSizeDeltaUsd) view returns((((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)),((bytes32,address,address,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(address,uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,uint256,uint256),int256,int256,int256))
func (_Main *MainCaller) GetPositionInfo(opts *bind.CallOpts, dataStore common.Address, referralStorage common.Address, positionKey [32]byte, prices MarketUtilsMarketPrices, sizeDeltaUsd *big.Int, uiFeeReceiver common.Address, usePositionSizeAsSizeDeltaUsd bool) (ReaderUtilsPositionInfo, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPositionInfo", dataStore, referralStorage, positionKey, prices, sizeDeltaUsd, uiFeeReceiver, usePositionSizeAsSizeDeltaUsd)

	if err != nil {
		return *new(ReaderUtilsPositionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ReaderUtilsPositionInfo)).(*ReaderUtilsPositionInfo)

	return out0, err

}

// GetPositionInfo is a free data retrieval call binding the contract method 0x0815bce1.
//
// Solidity: function getPositionInfo(address dataStore, address referralStorage, bytes32 positionKey, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 sizeDeltaUsd, address uiFeeReceiver, bool usePositionSizeAsSizeDeltaUsd) view returns((((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)),((bytes32,address,address,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(address,uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,uint256,uint256),int256,int256,int256))
func (_Main *MainSession) GetPositionInfo(dataStore common.Address, referralStorage common.Address, positionKey [32]byte, prices MarketUtilsMarketPrices, sizeDeltaUsd *big.Int, uiFeeReceiver common.Address, usePositionSizeAsSizeDeltaUsd bool) (ReaderUtilsPositionInfo, error) {
	return _Main.Contract.GetPositionInfo(&_Main.CallOpts, dataStore, referralStorage, positionKey, prices, sizeDeltaUsd, uiFeeReceiver, usePositionSizeAsSizeDeltaUsd)
}

// GetPositionInfo is a free data retrieval call binding the contract method 0x0815bce1.
//
// Solidity: function getPositionInfo(address dataStore, address referralStorage, bytes32 positionKey, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 sizeDeltaUsd, address uiFeeReceiver, bool usePositionSizeAsSizeDeltaUsd) view returns((((address,address,address),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(bool)),((bytes32,address,address,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256),(address,uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(int256,uint256,uint256),int256,int256,int256))
func (_Main *MainCallerSession) GetPositionInfo(dataStore common.Address, referralStorage common.Address, positionKey [32]byte, prices MarketUtilsMarketPrices, sizeDeltaUsd *big.Int, uiFeeReceiver common.Address, usePositionSizeAsSizeDeltaUsd bool) (ReaderUtilsPositionInfo, error) {
	return _Main.Contract.GetPositionInfo(&_Main.CallOpts, dataStore, referralStorage, positionKey, prices, sizeDeltaUsd, uiFeeReceiver, usePositionSizeAsSizeDeltaUsd)
}

// GetPositionPnlUsd is a free data retrieval call binding the contract method 0xb4976dae.
//
// Solidity: function getPositionPnlUsd(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, bytes32 positionKey, uint256 sizeDeltaUsd) view returns(int256, int256, uint256)
func (_Main *MainCaller) GetPositionPnlUsd(opts *bind.CallOpts, dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, positionKey [32]byte, sizeDeltaUsd *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPositionPnlUsd", dataStore, market, prices, positionKey, sizeDeltaUsd)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPositionPnlUsd is a free data retrieval call binding the contract method 0xb4976dae.
//
// Solidity: function getPositionPnlUsd(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, bytes32 positionKey, uint256 sizeDeltaUsd) view returns(int256, int256, uint256)
func (_Main *MainSession) GetPositionPnlUsd(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, positionKey [32]byte, sizeDeltaUsd *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Main.Contract.GetPositionPnlUsd(&_Main.CallOpts, dataStore, market, prices, positionKey, sizeDeltaUsd)
}

// GetPositionPnlUsd is a free data retrieval call binding the contract method 0xb4976dae.
//
// Solidity: function getPositionPnlUsd(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, bytes32 positionKey, uint256 sizeDeltaUsd) view returns(int256, int256, uint256)
func (_Main *MainCallerSession) GetPositionPnlUsd(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, positionKey [32]byte, sizeDeltaUsd *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Main.Contract.GetPositionPnlUsd(&_Main.CallOpts, dataStore, market, prices, positionKey, sizeDeltaUsd)
}

// GetSwapAmountOut is a free data retrieval call binding the contract method 0x409f37c7.
//
// Solidity: function getSwapAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, address tokenIn, uint256 amountIn, address uiFeeReceiver) view returns(uint256, int256, (uint256,uint256,uint256,address,uint256,uint256) fees)
func (_Main *MainCaller) GetSwapAmountOut(opts *bind.CallOpts, dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, tokenIn common.Address, amountIn *big.Int, uiFeeReceiver common.Address) (*big.Int, *big.Int, SwapPricingUtilsSwapFees, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSwapAmountOut", dataStore, market, prices, tokenIn, amountIn, uiFeeReceiver)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(SwapPricingUtilsSwapFees), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(SwapPricingUtilsSwapFees)).(*SwapPricingUtilsSwapFees)

	return out0, out1, out2, err

}

// GetSwapAmountOut is a free data retrieval call binding the contract method 0x409f37c7.
//
// Solidity: function getSwapAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, address tokenIn, uint256 amountIn, address uiFeeReceiver) view returns(uint256, int256, (uint256,uint256,uint256,address,uint256,uint256) fees)
func (_Main *MainSession) GetSwapAmountOut(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, tokenIn common.Address, amountIn *big.Int, uiFeeReceiver common.Address) (*big.Int, *big.Int, SwapPricingUtilsSwapFees, error) {
	return _Main.Contract.GetSwapAmountOut(&_Main.CallOpts, dataStore, market, prices, tokenIn, amountIn, uiFeeReceiver)
}

// GetSwapAmountOut is a free data retrieval call binding the contract method 0x409f37c7.
//
// Solidity: function getSwapAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, address tokenIn, uint256 amountIn, address uiFeeReceiver) view returns(uint256, int256, (uint256,uint256,uint256,address,uint256,uint256) fees)
func (_Main *MainCallerSession) GetSwapAmountOut(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, tokenIn common.Address, amountIn *big.Int, uiFeeReceiver common.Address) (*big.Int, *big.Int, SwapPricingUtilsSwapFees, error) {
	return _Main.Contract.GetSwapAmountOut(&_Main.CallOpts, dataStore, market, prices, tokenIn, amountIn, uiFeeReceiver)
}

// GetSwapPriceImpact is a free data retrieval call binding the contract method 0x5d5c6efe.
//
// Solidity: function getSwapPriceImpact(address dataStore, address marketKey, address tokenIn, address tokenOut, uint256 amountIn, (uint256,uint256) tokenInPrice, (uint256,uint256) tokenOutPrice) view returns(int256, int256)
func (_Main *MainCaller) GetSwapPriceImpact(opts *bind.CallOpts, dataStore common.Address, marketKey common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, tokenInPrice PriceProps, tokenOutPrice PriceProps) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSwapPriceImpact", dataStore, marketKey, tokenIn, tokenOut, amountIn, tokenInPrice, tokenOutPrice)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSwapPriceImpact is a free data retrieval call binding the contract method 0x5d5c6efe.
//
// Solidity: function getSwapPriceImpact(address dataStore, address marketKey, address tokenIn, address tokenOut, uint256 amountIn, (uint256,uint256) tokenInPrice, (uint256,uint256) tokenOutPrice) view returns(int256, int256)
func (_Main *MainSession) GetSwapPriceImpact(dataStore common.Address, marketKey common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, tokenInPrice PriceProps, tokenOutPrice PriceProps) (*big.Int, *big.Int, error) {
	return _Main.Contract.GetSwapPriceImpact(&_Main.CallOpts, dataStore, marketKey, tokenIn, tokenOut, amountIn, tokenInPrice, tokenOutPrice)
}

// GetSwapPriceImpact is a free data retrieval call binding the contract method 0x5d5c6efe.
//
// Solidity: function getSwapPriceImpact(address dataStore, address marketKey, address tokenIn, address tokenOut, uint256 amountIn, (uint256,uint256) tokenInPrice, (uint256,uint256) tokenOutPrice) view returns(int256, int256)
func (_Main *MainCallerSession) GetSwapPriceImpact(dataStore common.Address, marketKey common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, tokenInPrice PriceProps, tokenOutPrice PriceProps) (*big.Int, *big.Int, error) {
	return _Main.Contract.GetSwapPriceImpact(&_Main.CallOpts, dataStore, marketKey, tokenIn, tokenOut, amountIn, tokenInPrice, tokenOutPrice)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0xceeea3bf.
//
// Solidity: function getWithdrawal(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address[],address[]),(uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainCaller) GetWithdrawal(opts *bind.CallOpts, dataStore common.Address, key [32]byte) (WithdrawalProps, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getWithdrawal", dataStore, key)

	if err != nil {
		return *new(WithdrawalProps), err
	}

	out0 := *abi.ConvertType(out[0], new(WithdrawalProps)).(*WithdrawalProps)

	return out0, err

}

// GetWithdrawal is a free data retrieval call binding the contract method 0xceeea3bf.
//
// Solidity: function getWithdrawal(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address[],address[]),(uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainSession) GetWithdrawal(dataStore common.Address, key [32]byte) (WithdrawalProps, error) {
	return _Main.Contract.GetWithdrawal(&_Main.CallOpts, dataStore, key)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0xceeea3bf.
//
// Solidity: function getWithdrawal(address dataStore, bytes32 key) view returns(((address,address,address,address,address,address[],address[]),(uint256,uint256,uint256,uint256,uint256,uint256),(bool)))
func (_Main *MainCallerSession) GetWithdrawal(dataStore common.Address, key [32]byte) (WithdrawalProps, error) {
	return _Main.Contract.GetWithdrawal(&_Main.CallOpts, dataStore, key)
}

// GetWithdrawalAmountOut is a free data retrieval call binding the contract method 0x71e138d5.
//
// Solidity: function getWithdrawalAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 marketTokenAmount, address uiFeeReceiver) view returns(uint256, uint256)
func (_Main *MainCaller) GetWithdrawalAmountOut(opts *bind.CallOpts, dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, marketTokenAmount *big.Int, uiFeeReceiver common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getWithdrawalAmountOut", dataStore, market, prices, marketTokenAmount, uiFeeReceiver)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetWithdrawalAmountOut is a free data retrieval call binding the contract method 0x71e138d5.
//
// Solidity: function getWithdrawalAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 marketTokenAmount, address uiFeeReceiver) view returns(uint256, uint256)
func (_Main *MainSession) GetWithdrawalAmountOut(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, marketTokenAmount *big.Int, uiFeeReceiver common.Address) (*big.Int, *big.Int, error) {
	return _Main.Contract.GetWithdrawalAmountOut(&_Main.CallOpts, dataStore, market, prices, marketTokenAmount, uiFeeReceiver)
}

// GetWithdrawalAmountOut is a free data retrieval call binding the contract method 0x71e138d5.
//
// Solidity: function getWithdrawalAmountOut(address dataStore, (address,address,address,address) market, ((uint256,uint256),(uint256,uint256),(uint256,uint256)) prices, uint256 marketTokenAmount, address uiFeeReceiver) view returns(uint256, uint256)
func (_Main *MainCallerSession) GetWithdrawalAmountOut(dataStore common.Address, market MarketProps, prices MarketUtilsMarketPrices, marketTokenAmount *big.Int, uiFeeReceiver common.Address) (*big.Int, *big.Int, error) {
	return _Main.Contract.GetWithdrawalAmountOut(&_Main.CallOpts, dataStore, market, prices, marketTokenAmount, uiFeeReceiver)
}

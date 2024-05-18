package clients

import (
	"fmt"
	"math/big"
	"time"

	"github.com/defiants-co/perpstream-go/abis"
	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type GmxClient struct {
	BaseFuturesClient
	web3Client     *ethclient.Client
	contractCaller *abis.MainCaller
	priceCache     *utils.GmxPriceCache
}

func NewGmxClient(rpcUrlInput *string, priceCache *utils.GmxPriceCache) (*GmxClient, error) {
	var web3 *ethclient.Client
	var connectErr error

	if rpcUrlInput == nil {
		web3, connectErr = ethclient.Dial(utils.BaseArbitrumRpcUrl)
	} else {
		web3, connectErr = ethclient.Dial(*rpcUrlInput)
	}
	if connectErr != nil {
		return nil, utils.NewInvalidRpcError(utils.BaseArbitrumRpcUrl, "invalid RPC URL (failed Dial)")
	}
	readerAddress := common.HexToAddress(utils.GmxReaderContractAddress)

	client, contractConnectErr := abis.NewMainCaller(readerAddress, web3)
	if contractConnectErr != nil {
		return nil, utils.NewInvalidContractAddressError(utils.GmxReaderContractAddress, contractConnectErr.Error())
	}

	if priceCache == nil {
		return nil, utils.NewPriceCacheMissingError()
	}

	return &GmxClient{web3Client: web3, contractCaller: client, priceCache: priceCache}, nil
}

func (client *GmxClient) FetchPositions(userId string) ([]models.FuturesPosition, error) {

	if !common.IsHexAddress(userId) {
		return nil, utils.NewInvalidAddressError(userId)
	}

	address := common.HexToAddress(userId)
	dataStoreAddress := common.HexToAddress(utils.GmxDataStoreContractAddress)

	var positions []models.FuturesPosition

	positionProps, err := client.contractCaller.GetAccountPositions(
		&bind.CallOpts{Pending: false},
		dataStoreAddress,
		address,
		&big.Int{},
		utils.MaxBigInt(64),
	)
	if err != nil {
		return nil, utils.NewFailedFetchPositionsError(userId, fmt.Sprintf("failed to fetch positions - %s", err.Error()))
	}

	for _, gmxPosition := range positionProps {
		position := utils.GmxToFuturesPosition(gmxPosition, client.priceCache)
		if position != nil {
			positions = append(positions, *position)
		}
	}

	return positions, nil
}

func (client *GmxClient) fetchRetry(userId string) []models.FuturesPosition {
	positions, err := client.FetchPositions(userId)
	if err != nil {
		time.Sleep(2 * time.Second)
		return client.fetchRetry(userId)
	}
	return positions
}

func (client *GmxClient) StreamPositions(
	userId string,
	debug bool,
	sleepSeconds float64,
	initWithCallback bool,
	retryInitOnFail bool,
	callback func(
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	),
) error {
	if debug {
		fmt.Println("starting stream")
	}

	var lastPositions []models.FuturesPosition
	var initErr error

	if retryInitOnFail {
		lastPositions = client.fetchRetry(userId)
	} else {
		lastPositions, initErr = client.FetchPositions(userId)
		if initErr != nil {
			if debug {
				fmt.Println("init error", initErr.Error())
			}
			return utils.NewStreamFailedToStartError()
		}

	}

	if initWithCallback {
		if debug {
			fmt.Println("calling initiation callback")
		}
		go callback(lastPositions, userId, utils.GmxDataSourceName)
	}

	for {
		if debug {
			fmt.Println("fetching positions")
		}
		newPositions, err := client.FetchPositions(userId)
		if err != nil {
			if debug {
				fmt.Println("hit error: ", err.Error())
			}
			continue
		} else {
			if !models.PositionSetsAreEqual(lastPositions, newPositions) {
				if debug {
					fmt.Println("detected change, calling callback")
				}
				go callback(newPositions, userId, utils.GmxDataSourceName)
				lastPositions = newPositions
				fmt.Println("called callback")
			}
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

}

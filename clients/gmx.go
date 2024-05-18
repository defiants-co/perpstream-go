package clients

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/defiants-co/perpstream-go/abis"
	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GmxCaller represents a connection to a GMX contract.
type GmxCaller struct {
	contractConnection *abis.MainCaller
}

// GmxClient manages multiple GmxCaller instances to handle rate limits and errors.
type GmxClient struct {
	BaseFuturesClient
	callers      []*GmxCaller
	priceCache   *utils.GmxPriceCache
	mu           sync.Mutex
	currentIndex int
}

// NewGmxClient creates a new GmxClient with multiple RPC URLs for redundancy.
func NewGmxClient(rpcUrls []string, priceCache *utils.GmxPriceCache) (*GmxClient, error) {
	if priceCache == nil {
		return nil, utils.NewPriceCacheMissingError()
	}

	var callers []*GmxCaller
	for _, url := range rpcUrls {
		web3, err := ethclient.Dial(url)
		if err != nil {
			return nil, utils.NewInvalidRpcError(url, "invalid RPC URL (failed Dial)")
		}
		readerAddress := common.HexToAddress(utils.GmxReaderContractAddress)
		client, err := abis.NewMainCaller(readerAddress, web3)
		if err != nil {
			return nil, utils.NewInvalidContractAddressError(utils.GmxReaderContractAddress, err.Error())
		}
		callers = append(callers, &GmxCaller{contractConnection: client})
	}

	return &GmxClient{callers: callers, priceCache: priceCache}, nil
}

// getCaller returns the next GmxCaller in a round-robin fashion.
func (client *GmxClient) getCaller() *GmxCaller {
	client.mu.Lock()
	defer client.mu.Unlock()
	caller := client.callers[client.currentIndex]
	client.currentIndex = (client.currentIndex + 1) % len(client.callers)
	return caller
}

// FetchPositions fetches the futures positions for a given user ID.
func (client *GmxClient) FetchPositions(userId string) ([]models.FuturesPosition, error) {
	if !common.IsHexAddress(userId) {
		return nil, utils.NewInvalidAddressError(userId)
	}

	address := common.HexToAddress(userId)
	dataStoreAddress := common.HexToAddress(utils.GmxDataStoreContractAddress)

	var positions []models.FuturesPosition

	for i := 0; i < len(client.callers); i++ {
		caller := client.getCaller()
		positionProps, err := caller.contractConnection.GetAccountPositions(
			&bind.CallOpts{Pending: false},
			dataStoreAddress,
			address,
			&big.Int{},
			utils.MaxBigInt(64),
		)
		if err == nil {
			for _, gmxPosition := range positionProps {
				position := utils.GmxToFuturesPosition(gmxPosition, client.priceCache)
				if position != nil {
					positions = append(positions, *position)
				}
			}
			return positions, nil
		}
		time.Sleep(500 * time.Millisecond) // Small delay before retrying with next caller
	}
	return nil, utils.NewFailedFetchPositionsError(userId, "failed to fetch positions after trying all callers")

}

// fetchRetry retries fetching positions until it succeeds.
func (client *GmxClient) fetchRetry(userId string) []models.FuturesPosition {
	positions, err := client.FetchPositions(userId)
	if err != nil {
		time.Sleep(2 * time.Second)
		return client.fetchRetry(userId)
	}
	return positions
}

// StreamPositions streams the futures positions for a given user ID, calling the callback on changes.
func (client *GmxClient) StreamPositions(
	userId string,
	debug bool,
	sleepSeconds float64,
	initWithCallback bool,
	callback func(
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	),
) error {
	if debug {
		fmt.Println("starting stream")
	}

	lastPositions := client.fetchRetry(userId)

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

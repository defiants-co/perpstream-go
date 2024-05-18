package clients

import (
	"github.com/defiants-co/perpstream-go/abis"
	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type KwentaClient struct {
	BaseFuturesClient
	contractCaller *abis.AbisCaller
}

func NewKwentaClient(rpcUrlInput *string) (*KwentaClient, error) {
	var web3 *ethclient.Client
	var connectErr error

	if rpcUrlInput == nil {
		web3, connectErr = ethclient.Dial(utils.BaseOptimismRpcUrl)
	} else {
		web3, connectErr = ethclient.Dial(*rpcUrlInput)
	}
	if connectErr != nil {
		return nil, utils.NewInvalidRpcError(utils.BaseOptimismRpcUrl, "invalid RPC URL (failed Dial)")
	}

	marketDataAddress := common.HexToAddress(utils.KwentaSynthetixFuturesDataAddress)
	client, contractConnectErr := abis.NewAbisCaller(marketDataAddress, web3)

	if contractConnectErr != nil {
		return nil, utils.NewInvalidContractAddressError(utils.KwentaSynthetixFuturesDataAddress, contractConnectErr.Error())
	}

	return &KwentaClient{contractCaller: client}, nil

}

func (client *KwentaClient) FetchPositions(userId string) []models.FuturesPosition {
	return nil

}

func (client *KwentaClient) StreamPositions(
	userId string,
	debug bool,
	initWithCallback bool,
	callback func(
		oldPositions []models.FuturesPosition,
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	),
) error {
	return nil
}

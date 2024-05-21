package clients

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hasura/go-graphql-client"
)

type KwentaClient struct {
	userData      map[string][]models.FuturesPosition
	userList      []string
	graphqlClient *graphql.Client
	priceCache    *utils.KwentaPriceCache
	mu            sync.Mutex
	cacheRunning  bool
}

func NewKwentaClient(pc *utils.KwentaPriceCache) (*KwentaClient, error) {
	userData := make(map[string][]models.FuturesPosition)
	userList := []string{}
	graphqlClient := graphql.NewClient(utils.KwentaSubgraphUrl, &http.Client{})

	kwentaClient := &KwentaClient{
		userData:      userData,
		userList:      userList,
		graphqlClient: graphqlClient,
		cacheRunning:  false,
		priceCache:    pc,
	}

	// initial call
	_, err := utils.KwentaSubgraphPositionsQuery(
		[]string{},
		[]string{},
		graphqlClient,
	)

	if err != nil {
		return nil, err
	}

	return kwentaClient, nil
}

func (c *KwentaClient) getUserList() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.userList
}

func (c *KwentaClient) userInCache(userId string) bool {
	result := false
	c.mu.Lock()
	for key := range c.userData {
		if key == userId {
			result = true
			break
		}
	}

	c.mu.Unlock()
	return result
}

func (c *KwentaClient) getCacheStreamStatus() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cacheRunning
}

func (c *KwentaClient) addUserToList(userId string) error {
	if !common.IsHexAddress(userId) {
		return utils.NewInvalidAddressError(userId)
	}
	c.mu.Lock()
	if !utils.Contains(c.userList, userId) {
		c.userList = append(c.userList, userId)
	}
	defer c.mu.Unlock()
	return nil
}

func (c *KwentaClient) updateCache() error {
	positions, err := utils.KwentaSubgraphPositionsQuery(
		c.getUserList(),
		utils.GetKeys(utils.KwentaMarketsToToken),
		c.graphqlClient,
	)

	if err != nil {
		return err
	}

	newUserDataMap := make(map[string][]models.FuturesPosition)

	for _, user := range c.getUserList() {
		var userPositions []models.FuturesPosition

		for _, position := range positions {
			if strings.ToUpper(position.Account) == strings.ToUpper(user) {
				convertedPosition := utils.KwentaToFuturesPosition(&position, c.priceCache)
				if convertedPosition != nil {
					userPositions = append(userPositions, *convertedPosition)
				}
			}
		}

		newUserDataMap[user] = userPositions
	}

	c.mu.Lock()
	c.userData = newUserDataMap
	c.mu.Unlock()
	return nil
}

func (c *KwentaClient) waitTillInCache(userId string) {
	if !c.userInCache(userId) {
		time.Sleep(100 * time.Millisecond)
		c.waitTillInCache(userId)
	}
}

func (c *KwentaClient) StreamCacheUpdates(sleepSeconds int, debug bool) error {
	c.mu.Lock()
	c.cacheRunning = true
	c.mu.Unlock()

	for {
		if debug {
			fmt.Println("fetching data")
		}

		err := c.updateCache()

		if debug {
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("fetched data")
		}

		time.Sleep(time.Duration(sleepSeconds) * time.Second)

	}

}

func (c *KwentaClient) FetchPositions(userId string) ([]models.FuturesPosition, error) {
	if !c.getCacheStreamStatus() {
		return nil, utils.NewCacheStreamNotRunningError()
	}

	if !utils.Contains(c.getUserList(), userId) {
		c.addUserToList(userId)
	}

	c.waitTillInCache(userId)

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.userData[userId], nil
}

func (c *KwentaClient) StreamPositions(
	userId string,
	debug bool,
	initWithCallback bool,
	sleepSeconds int,
	callback func(
		newPositions []models.FuturesPosition,
		userId string,
		dataSource string,
	),
) error {

	if debug {
		fmt.Println("starting stream")
	}
	lastPositions, err := c.FetchPositions(userId)
	if err != nil {
		return err
	}

	if initWithCallback {
		if debug {
			fmt.Println("calling initiation callback")
		}
		go callback(lastPositions, userId, utils.KwentaDataSourceName)
	}

	for {
		if debug {
			fmt.Println("fetching positions")
		}
		newPositions, err := c.FetchPositions(userId)
		if err != nil {
			if debug {
				fmt.Println("hit error: ", err.Error())
			}
			continue
		} else {
			if !models.FuturesPositionSetsAreEqual(lastPositions, newPositions) {
				if debug {
					fmt.Println("detected change, calling callback")
				}
				go callback(newPositions, userId, utils.KwentaDataSourceName)
				lastPositions = newPositions
				fmt.Println("called callback")
			}
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

}

package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/defiants-co/perpstream-go/models"
	"github.com/defiants-co/perpstream-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

// The way that you can stream positions is that the client takes a rolling cache of ALL positions every few seconds,
// and then maps them to the stream of each user. then, every time the response is received, an algorithm
// checks for changes for each stream according to the cache, and calls the callback if necessary.
type HegicClient struct {
	BaseOptionsClient
	ActiveOptionsCache   *[]utils.HegicPosition
	InactiveOptionsCache *[]utils.HegicPosition
	mu                   sync.Mutex
}

func Fetch(path string, result interface{}) error {
	resp, err := http.Get(utils.HegicApiUrl + path)
	if err != nil {
		return utils.NewFailedFetchOptionPositionsError("failed to fetch")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return utils.NewFailedFetchOptionPositionsError(fmt.Sprintf("received status code %d: %s", resp.StatusCode, bodyString))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.NewFailedFetchOptionPositionsError(fmt.Sprintf("error reading response body: %v", err))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return utils.NewFailedFetchOptionPositionsError(fmt.Sprintf("error unmarshalling response body: %v", err))
	}

	return nil
}

func NewHegicClient() (*HegicClient, error) {
	var activeOtionsList *[]utils.HegicPosition
	var inactiveOptionsList *[]utils.HegicPosition
	client := &HegicClient{ActiveOptionsCache: activeOtionsList, InactiveOptionsCache: inactiveOptionsList}
	err := client.UpdateOptionsCache()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client *HegicClient) GetLeaderboard() ([]utils.HegicUserStats, error) {
	var response []utils.HegicUserStats

	err := Fetch("leaderboard", &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *HegicClient) UpdateOptionsCache() error {
	allOptions, err := client.GetAllOptions()

	if err != nil {
		return err
	}

	var activeOptions []utils.HegicPosition
	var inactiveOptions []utils.HegicPosition
	now := time.Now()
	for _, option := range allOptions {
		if !now.After(option.ExpirationDate) {
			activeOptions = append(activeOptions, option)
		} else {
			inactiveOptions = append(inactiveOptions, option)
		}
	}

	client.mu.Lock()
	client.ActiveOptionsCache = &activeOptions
	client.InactiveOptionsCache = &inactiveOptions
	client.mu.Unlock()
	return nil
}

func (client *HegicClient) StreamCacheUpdates(sleepSeconds int, debug bool) {
	for {
		if debug {
			fmt.Println("fetching options info")
		}
		client.UpdateOptionsCache()
		if debug {
			fmt.Println("fetched options info")
		}
		time.Sleep(time.Duration(sleepSeconds))

	}
}

func (client *HegicClient) GetActiveOptionsFromCache() []utils.HegicPosition {
	client.mu.Lock()
	defer client.mu.Unlock()
	return *client.ActiveOptionsCache
}

func (client *HegicClient) GetInactiveOptionsFromCache() []utils.HegicPosition {
	client.mu.Lock()
	defer client.mu.Unlock()
	return *client.InactiveOptionsCache
}

func (client *HegicClient) GetAllOptions() ([]utils.HegicPosition, error) {
	var hegicPositions utils.HegicPositionsResponse

	err := Fetch("positions", &hegicPositions)

	if err != nil {
		return nil, err
	}

	return hegicPositions.Positions, err
}

func (client *HegicClient) GetUserInformation(userId string) (*utils.HegicUserData, error) {
	if !common.IsHexAddress(userId) {
		return nil, utils.NewInvalidAddressError(userId)
	}

	queryStr := "user?a=" + userId

	var response utils.HegicUserData

	err := Fetch(queryStr, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil

}

func (client *HegicClient) fetchRetry(userId string) []models.OptionPosition {
	positions, err := client.FetchPositions(userId)
	if err != nil {
		time.Sleep(2 * time.Second)
		return client.fetchRetry(userId)
	}
	return positions
}

func (client *HegicClient) FetchPositions(userId string) ([]models.OptionPosition, error) {
	allOptions := client.GetActiveOptionsFromCache()
	var userOptions []models.OptionPosition
	for _, option := range allOptions {
		if option.User == userId && option.CloseDate.Year() == 1 {
			userOptions = append(userOptions,
				utils.HegicPositionToOption(
					&option,
				))
		}
	}

	return userOptions, nil
}

func (client *HegicClient) StreamPositions(
	userId string,
	debug bool,
	initWithCallback bool,
	sleepSeconds int,
	callback func(
		newPositions []models.OptionPosition,
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
		go callback(lastPositions, userId, utils.HegicDataSourceName)
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
			if !models.OptionPositionSetsAreEqual(lastPositions, newPositions) {
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

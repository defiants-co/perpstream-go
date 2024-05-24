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

// HegicClient manages the streaming and caching of options positions.
type HegicClient struct {
	BaseOptionsClient
	ActiveOptionsCache   *[]utils.HegicPosition
	InactiveOptionsCache *[]utils.HegicPosition
	mu                   sync.Mutex
}

// Fetch makes an HTTP GET request to the given path and unmarshals the result into the provided interface.
// Parameters:
// - path: The API path to fetch data from.
// - result: The interface to unmarshal the fetched data into.
// Returns:
// - An error if the fetch or unmarshal operation fails.
func hegicFetch(path string, result interface{}) error {
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

// NewHegicClient creates a new HegicClient instance and updates its options cache.
// Returns:
// - A pointer to a new HegicClient instance.
// - An error if updating the options cache fails.
func NewHegicClient() (*HegicClient, error) {
	var activeOptionsList []utils.HegicPosition
	var inactiveOptionsList []utils.HegicPosition
	client := &HegicClient{ActiveOptionsCache: &activeOptionsList, InactiveOptionsCache: &inactiveOptionsList}
	err := client.updateOptionsCache()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetLeaderboard fetches the leaderboard data from the API.
// Returns:
// - A slice of HegicUserStats models.
// - An error if fetching the leaderboard data fails.
func (client *HegicClient) GetLeaderboard() ([]utils.HegicUserStats, error) {
	var response []utils.HegicUserStats
	err := hegicFetch("leaderboard", &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateOptionsCache updates the client's active and inactive options cache.
// Returns:
// - An error if fetching all options fails.
func (client *HegicClient) updateOptionsCache() error {
	allOptions, err := client.getAllOptions()
	if err != nil {
		return err
	}

	var activeOptions []utils.HegicPosition
	var inactiveOptions []utils.HegicPosition
	now := time.Now()
	for _, option := range allOptions {
		if !now.After(option.ExpirationDate) && option.CloseDate.Year() == 1 {
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

// StreamCacheUpdates periodically updates the options cache.
// Parameters:
// - sleepSeconds: The number of seconds to sleep between updates.
// - debug: A boolean indicating if debugging is enabled.
func (client *HegicClient) StreamCacheUpdates(sleepSeconds int, debug bool) {
	for {
		if debug {
			fmt.Println("fetching options info")
		}
		client.updateOptionsCache()
		if debug {
			fmt.Println("fetched options info")
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}

// GetActiveOptionsFromCache returns the active options from the cache.
// Returns:
// - A slice of HegicPosition models.
func (client *HegicClient) GetActiveOptionsFromCache() []utils.HegicPosition {
	client.mu.Lock()
	defer client.mu.Unlock()
	return *client.ActiveOptionsCache
}

// GetInactiveOptionsFromCache returns the inactive options from the cache.
// Returns:
// - A slice of HegicPosition models.
func (client *HegicClient) GetInactiveOptionsFromCache() []utils.HegicPosition {
	client.mu.Lock()
	defer client.mu.Unlock()
	return *client.InactiveOptionsCache
}

// GetAllOptions fetches all options from the API.
// Returns:
// - A slice of HegicPosition models.
// - An error if fetching the options fails.
func (client *HegicClient) getAllOptions() ([]utils.HegicPosition, error) {
	var hegicPositions utils.HegicPositionsResponse
	err := hegicFetch("positions", &hegicPositions)
	if err != nil {
		return nil, err
	}
	return hegicPositions.Positions, nil
}

// GetUserInformation fetches user information for a given user ID from the API.
// Parameters:
// - userId: The ID of the user whose information is to be fetched.
// Returns:
// - A pointer to a HegicUserData model.
// - An error if the user ID is invalid or fetching the user information fails.
func (client *HegicClient) getUserInformation(userId string) (*utils.HegicUserData, error) {
	if !common.IsHexAddress(userId) {
		return nil, utils.NewInvalidAddressError(userId)
	}

	queryStr := "user?a=" + userId
	var response utils.HegicUserData
	err := hegicFetch(queryStr, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// fetchRetry retries fetching positions until it succeeds.
// Parameters:
// - userId: The ID of the user whose positions are to be fetched.
// Returns:
// - A slice of OptionPosition models.
func (client *HegicClient) fetchRetry(userId string) []models.OptionPosition {
	positions, err := client.FetchPositions(userId)
	if err != nil {
		time.Sleep(2 * time.Second)
		return client.fetchRetry(userId)
	}
	return positions
}

// FetchPositions fetches the options positions for a given user ID from the active options cache.
// Parameters:
// - userId: The ID of the user whose positions are to be fetched.
// Returns:
// - A slice of OptionPosition models.
// - An error if fetching positions fails.
func (client *HegicClient) FetchPositions(userId string) ([]models.OptionPosition, error) {
	allOptions := client.GetActiveOptionsFromCache()
	var userOptions []models.OptionPosition
	for _, option := range allOptions {
		if option.User == userId {
			userOptions = append(userOptions, utils.HegicPositionToOption(&option))
		}
	}
	return userOptions, nil
}

func (c *HegicClient) FetchAnalytics() (*utils.HegicProjectedPnl, error) {
	var hegicData utils.HegicProjectedPnl

	err := hegicFetch("analytics", &hegicData)

	if err != nil {
		return nil, err
	}

	return &hegicData, nil
}

// StreamPositions streams the options positions for a given user ID, calling the callback on changes.
// Parameters:
// - userId: The ID of the user whose positions are to be streamed.
// - debug: A boolean indicating if debugging is enabled.
// - initWithCallback: A boolean indicating if the callback should be invoked initially.
// - sleepSeconds: The number of seconds to sleep between polling.
// - callback: A function that will be called with the new positions, userId, and dataSource.
// Returns:
// - An error if streaming fails.
func (client *HegicClient) StreamPositions(
	userId string,
	debug bool,
	initWithCallback bool,
	sleepSeconds int,
	callback func(
		oldPositions []models.OptionPosition,
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
		go callback(lastPositions, lastPositions, userId, utils.HegicDataSourceName)
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
				go callback(lastPositions, newPositions, userId, utils.HegicDataSourceName)
				lastPositions = newPositions
				if debug {
					fmt.Println("called callback")
				}
			}
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}

package clients

import "github.com/defiants-co/perpstream-go/models"

// BaseFuturesClient defines the interface for futures clients.
type BaseFuturesClient interface {
	// FetchPositions retrieves the futures positions for a given user.
	// Parameters:
	// - userId: The ID of the user whose positions are to be fetched.
	// Returns:
	// - A slice of FuturesPosition models.
	// - An error if any occurs during fetching.
	FetchPositions(userId string) ([]models.FuturesPosition, error)

	// StreamPositions streams the futures positions for a given user.
	// Parameters:
	// - userId: The ID of the user whose positions are to be streamed.
	// - debug: A boolean indicating if debugging is enabled.
	// - initWithCallback: A boolean indicating if the callback should be invoked initially.
	// - sleepSeconds: The number of seconds to sleep between polling.
	// - callback: A function that will be called with the new positions, userId, and dataSource.
	// Returns:
	// - An error if any occurs during streaming.
	StreamPositions(
		userId string,
		debug bool,
		initWithCallback bool,
		sleepSeconds int,
		callback func(
			oldPositions []models.FuturesPosition,
			newPositions []models.FuturesPosition,
			userId string,
			dataSource string,
		),
	) error

	GetLeaderBoard()
}

// BaseOptionsClient defines the interface for options clients.
type BaseOptionsClient interface {
	// FetchPositions retrieves the options positions for a given user.
	// Parameters:
	// - userId: The ID of the user whose positions are to be fetched.
	// Returns:
	// - A slice of OptionPosition models.
	FetchPositions(userId string) []models.OptionPosition

	// StreamPositions streams the options positions for a given user.
	// Parameters:
	// - userId: The ID of the user whose positions are to be streamed.
	// - debug: A boolean indicating if debugging is enabled.
	// - initWithCallback: A boolean indicating if the callback should be invoked initially.
	// - sleepSeconds: The number of seconds to sleep between polling.
	// - callback: A function that will be called with the new positions, userId, and dataSource.
	StreamPositions(
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
	)
}

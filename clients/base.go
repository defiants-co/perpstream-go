package clients

import "github.com/defiants-co/perpstream-go/models"

type BaseFuturesClient interface {
	FetchPositions(userId string) ([]models.FuturesPosition, error)
	StreamPositions(
		userId string,
		debug bool,
		initWithCallback bool,
		callback func(
			oldPositions []models.FuturesPosition,
			newPositions []models.FuturesPosition,
			userId string,
			dataSource string,
		),
	) error
}

type BaseSpotClient interface {
	FetchPositions(userId string) []models.SpotPosition
	StreamPositions(
		userId string,
		debug bool,
		initWithCallback bool,
		callback func(
			oldPositions []models.SpotPosition,
			newPositions []models.SpotPosition,
			userId string,
			dataSource string,
		),
	)
}

type BaseOptionsClient interface {
	FetchPositions(userId string) []models.OptionPosition
	StreamPositions(
		userId string,
		debug bool,
		initWithCallback bool,
		callback func(
			oldPositions []models.OptionPosition,
			newPositions []models.OptionPosition,
			userId string,
			dataSource string,
		),
	)
}

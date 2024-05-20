package clients

import "github.com/defiants-co/perpstream-go/models"

type BaseFuturesClient interface {
	FetchPositions(userId string) ([]models.FuturesPosition, error)
	StreamPositions(
		userId string,
		debug bool,
		initWithCallback bool,
		sleepSeconds int,
		callback func(
			newPositions []models.FuturesPosition,
			userId string,
			dataSource string,
		),
	) error
}

type BaseOptionsClient interface {
	FetchPositions(userId string) []models.OptionPosition
	StreamPositions(
		userId string,
		debug bool,
		initWithCallback bool,
		sleepSeconds int,
		callback func(
			newPositions []models.OptionPosition,
			userId string,
			dataSource string,
		),
	)
}

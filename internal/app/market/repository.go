package market

type MarketRepository interface {
	GetGroupName() string
	GetAll() ([]Market, error)
	Save([]Market) error
}

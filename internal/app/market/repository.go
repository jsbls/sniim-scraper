package market

type MarketRepository interface {
	Save([]Market) error
	GetDstName() string
}

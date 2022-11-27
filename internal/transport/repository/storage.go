package repository

import (
	"github.com/everitosan/snimm-scrapper/internal/app/market"
	"github.com/everitosan/snimm-scrapper/internal/app/utils"
)

type Repository struct {
	Market            market.MarketRepository
	Product           utils.OptionSelectRepository
	ProductSource     utils.OptionSelectRepository
	ProductDestiny    utils.OptionSelectRepository
	PricePresentation utils.OptionSelectRepository
}

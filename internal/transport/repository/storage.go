package repository

import (
	"github.com/everitosan/snimm-scrapper/internal/app/form"
	"github.com/everitosan/snimm-scrapper/internal/app/market"
)

type Repository struct {
	Market            market.MarketRepository
	Product           form.OptionSelectRepository
	ProductSource     form.OptionSelectRepository
	ProductDestiny    form.OptionSelectRepository
	PricePresentation form.OptionSelectRepository
	Week              form.OptionSelectRepository
	Month             form.OptionSelectRepository
	Year              form.OptionSelectRepository
}

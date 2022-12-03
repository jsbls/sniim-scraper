package repository

import (
	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/everitosan/sniim-scrapper/internal/app/form"
	"github.com/everitosan/sniim-scrapper/internal/app/market"
)

type Repository struct {
	Market            market.MarketRepository
	Params            form.ParamsRepository
	Consult           consult.ConsultRepository
	ConsultResponse   consult.ConsultResponseRepository
	Product           form.ProductRepository
	ProductSource     form.OptionSelectRepository
	ProductDestiny    form.OptionSelectRepository
	PricePresentation form.OptionSelectRepository
	Week              form.OptionSelectRepository
	Month             form.OptionSelectRepository
	Year              form.OptionSelectRepository
}

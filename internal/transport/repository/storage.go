package repository

import (
	"github.com/everitosan/snimm-scrapper/internal/app/market"
	"github.com/everitosan/snimm-scrapper/internal/app/product"
)

type Repository struct {
	Market  market.MarketRepository
	Product product.ProductRepository
}

/*
* Implementation for market file system saving
 */

package filestorage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/everitosan/sniim-scrapper/internal/app/market"
)

/*
 * Market Repository
 */

type marketFileRepository struct {
	dst string
}

func NewMarketFileRepository(dst string) (*marketFileRepository, error) {
	var repo marketFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
	}
	return &repo, err
}

func (mR *marketFileRepository) GetGroupName() string {
	return "markets"
}

func (mR *marketFileRepository) Save(content []market.Market) error {
	fileName := filepath.Join(mR.dst, mR.GetGroupName()+".json")

	str, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)
}

func (mR *marketFileRepository) GetAll() ([]market.Market, error) {
	var markets []market.Market
	fileName := filepath.Join(mR.dst, mR.GetGroupName()+".json")

	content, err := os.ReadFile(fileName)

	if err != nil {
		return markets, err
	}

	err = json.Unmarshal(content, &markets)
	return markets, err
}

func (mR *marketFileRepository) GetSubCategories() ([]string, error) {
	var markets []market.Market
	var subcategories []string

	markets, err := mR.GetAll()

	if err != nil {
		return subcategories, err
	}

	for _, market := range markets {
		for _, inventory := range market.Inventories {
			for _, cat := range inventory.Categories {
				for _, subCat := range cat.SubCategories {
					subcategories = append(subcategories, subCat.Name)
				}
			}
		}
	}

	return subcategories, nil
}

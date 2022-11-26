/*
* Implementation for local file system saving
 */

package repository

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/everitosan/snimm-scrapper/internal/app/market"
	"github.com/everitosan/snimm-scrapper/internal/app/product"
	"github.com/sirupsen/logrus"
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

func (marketRepo *marketFileRepository) GetDstName() string {
	return "markets"
}

func (marketRepo *marketFileRepository) Save(content []market.Market) error {
	fileName := filepath.Join(marketRepo.dst, marketRepo.GetDstName()+".json")

	str, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)
}

/*
* Product Repository
 */

type productFileRepository struct {
	dst string
}

func NewProductRepository(dst string) (*productFileRepository, error) {
	var repo productFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
	}
	return &repo, err
}

func (productRepo *productFileRepository) GetDstName() string {
	return "products"
}

func (productRepo *productFileRepository) Save(content []product.Product) error {
	fileName := filepath.Join(productRepo.dst, productRepo.GetDstName()+".json")

	str, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)
}

/*
Helper function to create directory
*/
func initDir(name string) error {
	var err error
	_, err = os.Stat(name) // Check if exists
	if err == nil {
		return err
	}
	if os.IsNotExist(err) {
		// We try to create it
		err = os.MkdirAll(name, os.ModePerm)
	}
	return err
}

/*
Helper function to save jsons to files
*/

func saveJsonStrToFile(content, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = f.WriteString(content)

	if err != nil {
		return err
	}

	logrus.Printf("🗃️  Saved in file %s", filePath)

	defer f.Close()

	return nil
}

package filestorage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/everitosan/snimm-scrapper/internal/app/form"
)

/*
* Select Repository to save SelectOption to a json file
 */

type optionSelectFileRepository struct {
	dst      string
	fileName string
}

func NewOptionSelectFileRepository(dst, fileName string) (*optionSelectFileRepository, error) {
	var repo optionSelectFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
		repo.fileName = fileName
	}
	return &repo, err
}

func (osR *optionSelectFileRepository) GetGroupName() string {
	return osR.fileName
}

func (osR *optionSelectFileRepository) GetAll() ([]form.OptionSelect, error) {
	var options []form.OptionSelect
	fileName := filepath.Join(osR.dst, osR.GetGroupName()+".json")

	content, err := os.ReadFile(fileName)

	if err != nil {
		return options, err
	}

	err = json.Unmarshal(content, &options)
	return options, err
}

func (osR *optionSelectFileRepository) Save(content []form.OptionSelect) error {
	fileName := filepath.Join(osR.dst, osR.GetGroupName()+".json")

	str, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)
}

func (osR *optionSelectFileRepository) GetBySubCategory(subcategory string) ([]form.OptionSelect, error) {
	var results []form.OptionSelect
	products, err := osR.GetAll()
	if err != nil {
		return results, err
	}

	for _, product := range products {
		if product.SubCategory == subcategory {
			results = append(results, product)
		}
	}

	return results, nil
}

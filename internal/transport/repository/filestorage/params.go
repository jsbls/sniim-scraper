package filestorage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/everitosan/sniim-scrapper/internal/app/form"
)

type paramsFileRepository struct {
	dst      string
	fileName string
}

func NewParamsFileRepository(dst, fileName string) (*paramsFileRepository, error) {
	var repo paramsFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
		repo.fileName = fileName
	}
	return &repo, err
}

func (pR *paramsFileRepository) Save(content []form.FormParams) error {
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	str, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)

}

func (pR *paramsFileRepository) GetAll() ([]form.FormParams, error) {
	var formParams []form.FormParams
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	content, err := os.ReadFile(fileName)

	if err != nil {
		return formParams, err
	}

	err = json.Unmarshal(content, &formParams)
	return formParams, err
}

func (pR *paramsFileRepository) GetCategories() ([]string, error) {
	cats := make([]string, 0)
	catsTmp := make(map[string]int)

	params, err := pR.GetAll()

	if err != nil {
		return cats, err
	}

	for index, param := range params {
		cat := param.Category
		_, exists := catsTmp[cat]
		if !exists {
			catsTmp[cat] = index
		}
	}

	cats = make([]string, 0, len(catsTmp))
	for k := range catsTmp {
		cats = append(cats, k)
	}

	return cats, nil
}

func (pR *paramsFileRepository) GetSubCategories(category string) ([]string, error) {
	var subcats []string
	subcatsTmp := make(map[string]int)
	options, err := pR.GetAll()

	if err != nil {
		return subcats, err
	}

	for index, opt := range options {
		if opt.Category == category {
			subcat := opt.SubCategory
			_, exists := subcatsTmp[subcat]
			if !exists {
				subcatsTmp[subcat] = index
			}
		}
	}

	subcats = make([]string, 0, len(subcatsTmp))
	for k := range subcatsTmp {
		subcats = append(subcats, k)
	}

	return subcats, nil
}

func (pR *paramsFileRepository) GetBySubCategory(subcat string) (form.FormParams, error) {
	var result form.FormParams

	formParams, err := pR.GetAll()

	if err != nil {
		return result, err
	}

	for _, param := range formParams {
		if param.SubCategory == subcat {
			result = param
			break
		}
	}

	return result, nil
}

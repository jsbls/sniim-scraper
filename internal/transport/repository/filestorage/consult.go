package filestorage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
)

type consultFileRepository struct {
	dst      string
	fileName string
}

func NewConsultFileRepository(dst, fileName string) (*consultFileRepository, error) {
	var repo consultFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
		repo.fileName = fileName
	}
	return &repo, err
}

func (pR *consultFileRepository) SaveOne(content consult.Consult) error {
	var all []consult.Consult
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	all, err := pR.GetAll()
	if err != nil {
		all = make([]consult.Consult, 0, 1)
	}

	all = append(all, content)

	str, err := json.Marshal(all)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)

}

func (pR *consultFileRepository) GetAll() ([]consult.Consult, error) {
	var formParams []consult.Consult
	fileName := filepath.Join(pR.dst, pR.fileName+".json")

	content, err := os.ReadFile(fileName)

	if err != nil {
		return formParams, err
	}

	err = json.Unmarshal(content, &formParams)
	return formParams, err
}

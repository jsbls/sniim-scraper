package filestorage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type consultResponseFileRepository struct {
	dst      string
	fileName string
}

func NewConsultResponseFileRepository(dst, fileName string) (*consultResponseFileRepository, error) {
	var repo consultResponseFileRepository
	err := initDir(dst)
	if err == nil {
		repo.dst = dst
		repo.fileName = fileName
	}
	return &repo, err
}

func (cRR *consultResponseFileRepository) Save(content []map[string]string) error {
	var all []map[string]string
	fileName := filepath.Join(cRR.dst, cRR.fileName+".json")

	all, err := cRR.GetAll()
	if err != nil {
		all = make([]map[string]string, 0, 1)
	}

	all = append(all, content...)

	str, err := json.Marshal(all)
	if err != nil {
		return err
	}

	return saveJsonStrToFile(string(str), fileName)

}

func (cRR *consultResponseFileRepository) GetAll() ([]map[string]string, error) {
	var formParams []map[string]string
	fileName := filepath.Join(cRR.dst, cRR.fileName+".json")

	content, err := os.ReadFile(fileName)

	if err != nil {
		return formParams, err
	}

	err = json.Unmarshal(content, &formParams)
	return formParams, err
}

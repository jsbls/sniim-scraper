package catalogue

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
)

type catalogueWritter struct {
	dst string
}

func NewCatalogueWritter(dst string) (*catalogueWritter, error) {
	var w catalogueWritter
	var err error

	_, err = os.Stat(dst) // Check if exists

	if err != nil {
		if os.IsNotExist(err) {
			// We try to create it
			err = os.MkdirAll(dst, os.ModePerm)
			if err == nil {
				w.dst = dst
			}
		}
	} else {
		w.dst = dst
	}

	return &w, err
}

func (w *catalogueWritter) SaveMapToJsonFile(name string, content scrapper.SelectOptionsAsMap, dst string) error {

	subDir := filepath.Join(w.dst, dst)
	fileName := filepath.Join(subDir, name+".json")
	err := os.MkdirAll(subDir, os.ModePerm)

	if err != nil {
		return err
	}

	str, err := json.Marshal(content)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(str))

	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

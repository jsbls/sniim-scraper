package filestorage

import (
	"os"

	"github.com/sirupsen/logrus"
)

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

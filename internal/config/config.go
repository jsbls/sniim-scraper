package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	SNIIM_ADDR    = "SNIIM_ADDR"    // Base url for the scrapping
	CATALOGUE_SRC = "CATALOGUE_SRC" // Catalogues directory or db name
	DEBUG         = "DEBUG"
	MONGO_URI     = "MONGO_URI"
)

type config struct {
	SNIIM_ADDR    string
	CATALOGUE_SRC string
	DEBUG         bool
	MONGO_URI     string
}

func LoadConfig() *config {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("No .env file dected")
	}

	// Read snim address
	snimmAddr := os.Getenv(SNIIM_ADDR)
	if snimmAddr == "" {
		snimmAddr = "http://www.economia-sniim.gob.mx"
	}

	// Read Catalogues dir
	catalogueDir := os.Getenv(CATALOGUE_SRC)
	if catalogueDir == "" {
		catalogueDir = "SNIIM_DATA"
	}

	// Read debug info
	isDebug := (os.Getenv(DEBUG) == "true")

	// Read Mongo info
	mongoUri := os.Getenv(MONGO_URI)

	logrus.Printf("Using %s", snimmAddr)

	return &config{
		SNIIM_ADDR:    snimmAddr,
		CATALOGUE_SRC: catalogueDir,
		DEBUG:         isDebug,
		MONGO_URI:     mongoUri,
	}

}
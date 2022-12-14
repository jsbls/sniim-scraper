package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	SNIIM_ADDR    = "SNIIM_ADDR"    // Base url for the scrapping
	CATALOGUE_SRC = "CATALOGUE_SRC" // Catalogues directory or db name
	DEBUG         = "DEBUG"         // Flag to define debug mode
	MONGO_URI     = "MONGO_URI"     // Mongo Database uri
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
		logrus.Warn("⚠️ No .env file detected")
	}

	// Read snim address
	sniimAddr := os.Getenv(SNIIM_ADDR)
	if sniimAddr == "" {
		sniimAddr = "http://www.economia-sniim.gob.mx"
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

	return &config{
		SNIIM_ADDR:    sniimAddr,
		CATALOGUE_SRC: catalogueDir,
		DEBUG:         isDebug,
		MONGO_URI:     mongoUri,
	}

}

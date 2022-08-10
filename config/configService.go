package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"rarity-backend/structs"
)

// NewConfigService tries to read the configarion file which should countain information for each trait, attribute and possible sets.
//
// The configuration json MUST be modified with great care. No reordering of the elemets is allowed because it messes the logic of fetching attributes in services/metadata.go.
//
// Returns ConfigService object containing the configuration
func NewConfigService(configPath string) *structs.ConfigService {
	jsonFile, err := os.Open(configPath)

	if err != nil {
		log.Fatal("Missing polymorphs config file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var service structs.ConfigService

	json.Unmarshal(byteValue, &service)

	return &service
}

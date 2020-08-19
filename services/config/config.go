package config

import (
	"danmu/utils/files"
	"encoding/json"
	"io/ioutil"
)

var (
	Global *GlobalConfigType
)

type GlobalConfigType struct {
	MySQL string `json:"mysql"`
}

func ReadFile(config *GlobalConfigType) {
	configPath, err := files.ConfigFilePath()
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(content, config); err != nil {
		panic(err)
	}
}

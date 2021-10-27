package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	Cors struct {
		Origins []string `json:"origins"`
	} `json:"cors"`
	DataStorage struct {
		DataBaseName string `json:"dataBaseName"`
		URI          string `json:"uri"`
	} `json:"dataSource"`
	Encript struct {
		FilePathSeed string `json:"filePathSeed"`
	} `json:"encript"`
}

func getConfig(value string) []string {
	return strings.SplitN(value, "=", 2)
}

func containConfig(value string, key string) bool {
	return strings.HasPrefix(value, key)
}

func findConfig(varibles []string, key string) []string {
	var config []string = nil
	for _, v := range varibles {
		if containConfig(v, key) {
			config = getConfig(v)
			break
		}
	}
	return config
}

func loadEnvironmentValues(appConfig *AppConfig) {
	variables := os.Environ()
	if c := findConfig(variables, "origins"); c != nil {
		appConfig.Cors.Origins = strings.Split(c[1], ",")
	}

	if c := findConfig(variables, "uri"); c != nil {
		appConfig.DataStorage.URI = c[1]
	}

	if c := findConfig(variables, "dataBaseName"); c != nil {
		appConfig.DataStorage.DataBaseName = c[1]
	}

	if c := findConfig(variables, "filePathSeed"); c != nil {
		appConfig.Encript.FilePathSeed = c[1]
	}
}

var appConfig *AppConfig

func GetAppConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}
	appConfig = &AppConfig{}

	if f, err := os.Stat("settings.json"); err == nil {
		dat, err := os.ReadFile(f.Name())
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(dat, &appConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
	loadEnvironmentValues(appConfig)
	return appConfig
}

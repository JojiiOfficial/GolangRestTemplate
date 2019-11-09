package main

import (
	"encoding/json"
	"os"
	"strings"
)

//Config config for the server
type Config struct {
	Host          string `json:"host"`
	Username      string `json:"username"`
	Pass          string `json:"pass"`
	DatabasePort  int    `json:"dbport"`
	CertFile      string `json:"cert"`
	KeyFile       string `json:"key"`
	IPdataAPIKey  string `json:"ipdataAPIkey"`
	ShowTimeInLog bool   `json:"showLogTime"`
	HTTPPort      int    `json:"port"`
	TLSPort       int    `json:"porttls"`
}

func createConfig() error {
	LogError("Couldn't find config.json")
	f, err := os.Create("config.json")
	if err != nil {
		return err
	}

	emptyConfig := Config{}
	d, err := json.Marshal(emptyConfig)
	if err != nil {
		return err
	}
	strConf := string(d)

	strConf = strings.ReplaceAll(strConf, "{", "{\n")
	strConf = strings.ReplaceAll(strConf, "[", "[\n")
	strConf = strings.ReplaceAll(strConf, "}", "\n}")
	strConf = strings.ReplaceAll(strConf, "]", "\n]")
	strConf = strings.ReplaceAll(strConf, ",", ",\n")

	_, err = f.WriteString(strConf)
	if err != nil {
		return err
	}

	if f != nil {
		f.Close()
	}

	return nil
}

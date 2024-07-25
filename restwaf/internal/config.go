package internal

import (
	"encoding/json"
	"os"
)

var Global *Configuration

type BindConfiguration struct {
	Adress string `json:"address"`
	Port   int    `json:"port"`
}
type OpenApi struct {
	Url string `json:"url"`
}
type Configuration struct {
	Bind BindConfiguration `json:"bind"`
}

func InitConfig(file string) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &Global)
	if err != nil {
		return err
	}

	return nil
}

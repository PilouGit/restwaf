package config

import (
	"strconv"
	"strings"
)

type GlobalConfiguration struct {
	Port               int                 `json:"port"  validate:"required"`
	Adress             string              `json:"adress"  `
	CacheConfiguration *CacheConfiguration `json:"cache"  validate:"required"`
}

func (globalConfiguration *GlobalConfiguration) GetAddress() string {
	var trimAdress = strings.TrimSpace(globalConfiguration.Adress)

	if len(trimAdress) > 0 {
		return trimAdress + ":" + strconv.Itoa(globalConfiguration.Port)

	} else {
		return ":" + strconv.Itoa(globalConfiguration.Port)
	}
}

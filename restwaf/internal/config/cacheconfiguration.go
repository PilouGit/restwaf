package config

import (
	"errors"
	"strconv"
)

type CacheConfiguration struct {
	Cachetype     string            `json:"type"  validate:"required"`
	Properties    map[string]string `json:"properties"`
	Nbtransaction uint
}

func (cacheConfiguration *CacheConfiguration) IsGCache() bool {
	return cacheConfiguration.Cachetype == "gcache"
}
func (cacheConfiguration *CacheConfiguration) Validate() error {

	if cacheConfiguration.IsGCache() {
		var nbTransationString, exist = cacheConfiguration.Properties["nbtransaction"]
		if !exist {
			return errors.New("no Transaction number is defined")
		}
		nbTransation, err := strconv.Atoi(nbTransationString)
		if err != nil {
			return errors.New("nb transaction for gcache is not a integer")
		}
		if nbTransation < 1 {
			return errors.New("nb transaction must be positive ")
		}
		cacheConfiguration.Nbtransaction = uint(nbTransation)

	} else {
		return errors.New("only gcache type is currently supported")
	}
	return nil
}

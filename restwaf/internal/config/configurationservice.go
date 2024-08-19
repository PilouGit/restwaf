package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func (configurationInstance *Configuration) ReadConfiguration(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), &configurationInstance); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := data[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
			return err
		}
	}
	return nil
}

func (configurationInstance *Configuration) Validate() error {
	error := configurationInstance.GlobalConfiguration.CacheConfiguration.Validate()
	if error != nil {
		return error
	}
	return nil
}

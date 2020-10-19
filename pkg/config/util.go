package config

import (
	"errors"
	"os"
)

// ErrNotaryDisabled indicates that notary is disabled
var ErrNotaryDisabled = errors.New("notary disabled")

// CheckNotaryEnabled checks the config file whether notary is enabled or not.
func CheckNotaryEnabled() error {
	config, err := Load()
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		config = New()
	}
	if config.Enabled {
		return nil
	}
	return ErrNotaryDisabled
}

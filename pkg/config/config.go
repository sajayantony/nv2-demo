package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/docker/cli/cli/config"
)

const (
	// FileName is the name of config file
	FileName = "nv2.json"
)

var (
	// FilePath is the path of config file
	FilePath = filepath.Join(config.Dir(), FileName)
)

// File reflects the config file
type File struct {
	Enabled bool `json:"enabled"`
}

// Save stores the config to file
func (f *File) Save() error {
	file, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	return encoder.Encode(f)
}

// Load reads the config from file
func Load() (*File, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config *File
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

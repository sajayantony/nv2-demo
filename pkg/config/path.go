package config

import (
	"path/filepath"

	"github.com/docker/cli/cli/config"
)

const (
	// FileName is the name of config file
	FileName = "nv2.json"

	// SignatureStoreDirName is the name of the signature store directory
	SignatureStoreDirName = "nv2"
)

var (
	// FilePath is the path of config file
	FilePath = filepath.Join(config.Dir(), FileName)
	// SignatureStoreDirPath is the path of the signature store
	SignatureStoreDirPath = filepath.Join(config.Dir(), SignatureStoreDirName)
)

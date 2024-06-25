package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/kmc-jp/ChannelInviter/database"
	"github.com/kmc-jp/ChannelInviter/slack"
)

type Settings struct {
	Slack    slack.Settings    `yaml:"Slack"`
	Database database.Settings `yaml:"Database"`
}

func ReadSettings() (*Settings, error) {
	var yamlRootPath = "settings"

	dir, err := os.ReadDir(yamlRootPath)
	if err != nil {
		return nil, fmt.Errorf("ReadDir: %w", err)
	}

	var yamlBinary = []byte{}
	for _, f := range dir {
		if f.IsDir() || !(strings.HasSuffix(f.Name(), ".yaml") || strings.HasSuffix(f.Name(), ".yml")) {
			continue
		}

		var yamlFilePath = filepath.Join(yamlRootPath, f.Name())
		b, err := os.ReadFile(yamlFilePath)
		if err != nil {
			return nil, fmt.Errorf("ReadFile: %s %w", yamlFilePath, err)
		}

		// check format
		var us Settings
		err = yaml.Unmarshal(b, &us)
		if err != nil {
			return nil, fmt.Errorf("UnmarshalSettings: %s\n%w", yamlFilePath, err)
		}

		yamlBinary = append(yamlBinary, b...)
		yamlBinary = append(yamlBinary, '\n')
	}

	var us Settings
	err = yaml.Unmarshal(yamlBinary, &us)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal: %w", err)
	}

	return &us, nil
}

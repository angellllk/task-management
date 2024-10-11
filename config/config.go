package config

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

type Config struct {
	Dsn   string `json:"dsn"`
	DBKey string `json:"db_key"`
	Host  string `json:"host"`
}

func readConfig(path string) (*gabs.Container, error) {
	return gabs.ParseJSONFile(path)
}

func fetchJsonData(c *gabs.Container, path string) (string, error) {
	value, ok := c.Path(path).Data().(string)
	if !ok {
		return "", fmt.Errorf("error reading %s from \"config.json\", either invalid type or not found", path)
	}

	if value == "" {
		return "", fmt.Errorf("field %s cannot be empty", path)
	}

	return value, nil
}

func New(path string) (*Config, error) {
	jsonGabs, errRead := readConfig(path)
	if errRead != nil {
		return nil, fmt.Errorf("error reading \"config.json\" at path %s: %v", path, errRead)
	}

	dsn, err := fetchJsonData(jsonGabs, "database.dsn")
	if err != nil {
		return nil, err
	}

	dbKey, err := fetchJsonData(jsonGabs, "database.db_key")
	if err != nil {
		return nil, err
	}

	host, err := fetchJsonData(jsonGabs, "api.host")
	if err != nil {
		return nil, err
	}

	return &Config{
		Dsn:   dsn,
		DBKey: dbKey,
		Host:  host,
	}, nil
}

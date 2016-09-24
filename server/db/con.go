package db

import (
	"database/sql"
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v1"
)

type Configs map[string]*Config

func (cs Configs) Open(env string) (*sql.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, nil
	}
	return config.Open()
}

type Config struct {
	Datasource string `yaml:"datasource"`
}

func (c *Config) DSN() string {
	return c.Datasource
}

func (c *Config) Open() (*sql.DB, error) {
	return sql.Open("mysql", c.DSN())
}

func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

func NewConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

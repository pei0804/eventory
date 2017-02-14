package db

import (
	"database/sql"
	"io"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"

	yaml "gopkg.in/yaml.v1"
)

type Configs map[string]*Config

func (cs Configs) Open() (*sql.DB, error) {
	config, ok := cs["setting"]
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

func NewConfigsFromFile(FileName string, c echo.Context) (Configs, error) {

	ctx := appengine.NewContext(c.Request())
	bucketname, err := file.DefaultBucketName(ctx)
	if err != nil {
		return nil, err
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	reader, err := client.Bucket(bucketname).Object(FileName).NewReader(ctx)
	if err != nil {
		f, err := os.Open("dbconfig.yml")
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return NewConfigs(f)
	}
	defer reader.Close()
	return NewConfigs(reader)
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

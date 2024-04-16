package influxorm

import (
	"errors"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type DI interface {
	NewClient() (influxdb2.Client, error)
	GetOrg() string
	GetBucket() string
	NewTelegrafWriter() (TelegrafWriter, error)
}

type Config struct {
	Url         string `yaml:"url,omitempty"`
	Org         string `yaml:"org,omitempty"`
	Token       string `yaml:"token,omitempty"`
	Bucket      string `yaml:"bucket,omitempty"`
	TelegrafUrl string `yaml:"telegraf_url,omitempty"`
}

func (c *Config) NewClient() (influxdb2.Client, error) {
	if c.Url == "" {
		return nil, errors.New("missing url")
	}
	if c.Token == "" {
		return nil, errors.New("missing token")
	}
	return influxdb2.NewClient(c.Url, c.Token), nil
}

func (c *Config) GetBucket() string {
	return c.Bucket
}

func (c *Config) GetOrg() string {
	return c.Org
}

func (c *Config) NewTelegrafWriter() (TelegrafWriter, error) {
	if c.TelegrafUrl == "" {
		return nil, errors.New("missing telegraf_url")
	}
	return NewTelegraf(c.TelegrafUrl), nil
}

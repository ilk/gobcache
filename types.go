package gobcache

import (
	"log"
	"os"
)

const (
	tmp_dir = "/tmp"
)

type Config struct {
	Path          string
	LifetimeHours int64
	Logger        *log.Logger
}

type Client struct {
	Config Config
}

func (c *Config) SetDefaultDirectoryIfNotExit() {
	_, err := os.Open(c.Path)
	if err != nil {
		c.Logger.Printf("Set cache file directoy to %s, content of Config.Path: \"%s\" (%v)", tmp_dir, c.Path, err)
		c.Path = tmp_dir
	}
}

func (c *Config) SetDefaultLoggerIfNil() {
	if c.Logger == nil {
		c.Logger = log.New(os.Stdout, "[gobcache] ", log.Ldate|log.Ltime)
	}
}

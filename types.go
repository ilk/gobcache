package gobcache

import (
	"log"
	"os"
)

const (
	def_tmp_dir = "/tmp"
	def_ttl     = 8
)

type Config struct {
	// Were the *.gob files will be stored
	Path string
	// Renew cache when file is older than TTL (in hours)
	TTL int64
	// Add custom logger, default is *log.Logger from GO
	Logger *log.Logger
}

type Client struct {
	Config Config
}

func (c *Config) setDefaultDirectoryIfNotExit() {
	_, err := os.Open(c.Path)
	if err != nil {
		c.Path = def_tmp_dir
	}
}

func (c *Config) setDefaultLoggerIfNil() {
	if c.Logger == nil {
		c.Logger = log.New(os.Stdout, "[gobcache] ", log.Ldate|log.Ltime)
	}
}

func (c *Config) setDefaultTTL() {
	if c.TTL == 0 {
		c.TTL = def_ttl
	}
}

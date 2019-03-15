package gobcache

import (
	"log"
	"os"
)

const (
	defTTL = 8
)

// config holds up parameters
type Config struct {
	// Were the *.gob files will be stored, default is $TMPDIR/<name of executable>
	Path string
	// Renew cache when file is older than TTL (in hours)
	TTL int64
	// Add custom logger, default adds a prefix "[gobcache] date time"
	Logger *log.Logger
}

// Client provides SaveData() and GetData()
type Client struct {
	config Config
}

func (c *Config) setDefaultDirectoryIfNotExit() {
	_, err := os.Open(c.Path)
	if err != nil {
		err := os.MkdirAll(os.TempDir(), 0755)
		if err == nil {
			c.Path = os.TempDir()
		} else {
			c.Path = "."
		}
	}
}

func (c *Config) setDefaultLoggerIfNil() {
	if c.Logger == nil {
		c.Logger = log.New(os.Stdout, "[gobcache] ", log.Ldate|log.Ltime)
	}
}

func (c *Config) setDefaultTTL() {
	if c.TTL == 0 {
		c.TTL = defTTL
	}
}

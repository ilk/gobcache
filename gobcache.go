// Package gobcache provides simple way to store data for a specific time.
// Classic usage is fetching data from an api and save time next rerun
package gobcache

import (
	"encoding/gob"
	"fmt"
	"os"
)

// NewCache creates a cache client
func NewCache(config Config) Client {
	c := Client{}
	config.setDefaultLoggerIfNil()
	config.setDefaultDirectoryIfNotExit()
	config.setDefaultTTL()
	c.config = config
	return c
}

// SaveData store the data by the given identifier
func (c *Client) SaveData(identifier string, data interface{}) error {
	file := fmt.Sprintf("%s/%s.gob", c.config.Path, identifier)
	fh, err := os.Create(file)
	if err != nil {
		return err
	}
	c.config.Logger.Printf("Save data to %s\n", file)
	encoder := gob.NewEncoder(fh)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	defer fh.Close()
	return nil
}

// GetData gets the data by the given identifier
func (c *Client) GetData(identifier string, obj interface{}) error {
	filename := fmt.Sprintf("%s/%s.gob", c.config.Path, identifier)
	if !fileExistsAndNotOlderThan(filename, c.config.TTL) {
		c.config.Logger.Printf("%s does exists or is older than %d\n", filename, c.config.TTL)
		return fmt.Errorf("%s does not exists or is older than %d\n", filename, c.config.TTL)
	}

	fh, err := os.Open(filename)
	if err != nil {
		c.config.Logger.Printf("%s, %v", filename, err)
		return err
	}

	decoder := gob.NewDecoder(fh)
	err = decoder.Decode(obj)
	if err != nil {
		return fmt.Errorf("Error GetData(): %s, %v", identifier, err)
	}
	c.config.Logger.Printf("Read data from cache %s", filename)

	return nil
}

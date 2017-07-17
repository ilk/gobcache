package gobcache

import (
	"encoding/gob"
	"fmt"
	"os"
)

func NewCache(config Config) Client {
	config.setDefaultLoggerIfNil()
	config.setDefaultDirectoryIfNotExit()
	config.setDefaultTTL()
	c := Client{}
	c.Config = config
	return c
}

func (c *Client) SaveData(hash string, data interface{}) error {
	c.Config.Logger.Printf("Save data to disk")
	fh, err := os.Create(fmt.Sprintf("%s/%s.gob", c.Config.Path, hash))
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(fh)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	defer fh.Close()
	return nil
}

func (c *Client) GetData(hash string, obj interface{}) error {
	filename := fmt.Sprintf("%s/%s.gob", c.Config.Path, hash)
	if !fileExistsAndNotOlderThan(filename, c.Config.TTL) {
		return fmt.Errorf("File not exists or is older than %dh", c.Config.TTL)
	}

	fh, err := os.Open(filename)
	if err != nil {
		c.Config.Logger.Printf("%s, %v", filename, err)
		return err
	}

	decoder := gob.NewDecoder(fh)
	err = decoder.Decode(obj)
	if err != nil {
		return fmt.Errorf("Error GetData(): %s", hash)
	}

	return nil
}

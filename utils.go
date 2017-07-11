package gobcache

import (
	"time"

	"github.com/djherbis/times"
)

func fileExistsAndNotOlderThan(src string, hours int64) bool {

	unixnow := time.Now().Unix()
	t, err := times.Stat(src)
	if err != nil {
		return false
	}
	mtime := t.ModTime().Unix()
	diff := (unixnow - mtime) / 60 / 60

	if diff < hours {
		return true
	}

	return false
}

package gobcache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	var out []byte
	in := []byte("my byte is your byte")
	identifier := "test"

	c := NewCache(Config{})

	if err := c.SaveData(identifier, in); err != nil {
		t.Error(err)
	}

	if err := c.GetData(identifier, &out); err != nil {
		t.Error(err)
	}
	if string(in) != string(out) {
		t.Errorf("results are not the same")
	}
	fmt.Printf("In: %s\n", in)
	fmt.Printf("Out: %s\n", out)
}

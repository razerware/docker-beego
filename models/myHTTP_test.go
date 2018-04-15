package models

import (
	"testing"
)

func TestMyGet(t *testing.T) {
	a := map[string]string{"a": "sss", "b": "www"}
	MyGet("123.2.2154.4", a)
}

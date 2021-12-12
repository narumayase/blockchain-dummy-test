package pkg

import (
	"testing"
)

func TestBlock_Mine(t *testing.T) {
	b := NewBlock("dadada", "dadadada")
	b.Mine(5)

	//TODO
}
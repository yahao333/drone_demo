package main

import (
	"testing"

	"drone.fexeak.com/utils"
)

func TestAdd2(t *testing.T) {
	sum := utils.Add(2, 2)
	if 4 != sum {
		t.Errorf("Expected 4, got %v", sum)
	}
}

package utils

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if 3 != sum {
		t.Errorf("expected %v, got %v", 3, sum)
	}
}

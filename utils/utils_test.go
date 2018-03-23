package utils

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuids := make(map[string]int)
	for i := 1; i <= 100; i++ {
		uuids[GenerateUUID()]++
	}
	for _, v := range uuids {
		if v > 1 {
			t.Error("Duplicate UUID found")
		}
	}
}

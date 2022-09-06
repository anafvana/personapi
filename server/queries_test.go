package server_test

import (
	"fmt"
	"personapi/server"
	"testing"
)

func TestIsValidName(t *testing.T) {
	validNames := []string{"English", "Simplified English", "Português Brasileiro", "Français", "Español", "D'Italiano", "L’étoile", "L’ étoile", "van der Dutch", "русские", "汉语", "漢語", "中文", "日本国"}

	for i, val := range validNames {
		if res := server.IsValidName(validNames[i]); res != true {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected true, got %v", i, val, res))
			t.Fail()
		}
	}

	invalidNames := []string{"日1本国", "Norsk0 Dansk", "0", "Deutsch."}

	for i, val := range invalidNames {
		if res := server.IsValidName(validNames[i]); res != true {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected false, got %v", i, val, res))
			t.Fail()
		}
	}

	fmt.Printf("isValidName passed %d tests\n", len(validNames)+len(invalidNames))
}

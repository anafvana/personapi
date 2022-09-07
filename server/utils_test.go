package server_test

import (
	"fmt"
	"personapi/server"
	"testing"
)

func TestIsValidName(t *testing.T) {
	validNames := []string{"English", "Simplified English", "Português Brasileiro", "Français", "Español", "D'Italiano", "L’étoile", "L’ étoile", "van der Dutch", "русские", "汉语", "漢語", "中文", "日本国"}

	for i, val := range validNames {
		if res := server.IsValidName(validNames[i], "test"); res != "" {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected true, got %v", i, val, res))
			t.Fail()
		}
	}

	invalidNames := []string{"日1本国", "Norsk0 Dansk", "0", "Deutsch."}

	for i, val := range invalidNames {
		if res := server.IsValidName(invalidNames[i], "test"); res == "" {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected false, got %v", i, val, res))
			t.Fail()
		}
	}

	fmt.Printf("Finished isValidName's %d tests\n", len(validNames)+len(invalidNames))
}

func TestIsPalindrome(t *testing.T) {
	palindromes := []string{"Ana", "D’anad", "L’ ol", "van Nav", "naan", "汉语汉", "руссур", "Руссур", "Øyø", "", "    ", "	"}

	for i, val := range palindromes {
		if res := server.IsPalindrome(palindromes[i]); res != true {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected true, got %v", i, val, res))
			t.Fail()
		}
	}

	notPalindromes := []string{"English", "русские", "汉语", "漢語", "中文", "日本国", "Aná", "Naán", "Oyø", "0yø"}

	for i, val := range notPalindromes {
		if res := server.IsPalindrome(notPalindromes[i]); res != false {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected false, got %v", i, val, res))
			t.Fail()
		}
	}

	fmt.Printf("Finished isPalindrome's %d tests\n", len(palindromes)+len(notPalindromes))
}

func TestSyllables(t *testing.T) {
	type Test struct {
		word   string
		expect int
	}

	words := []Test{
		{"Test", 1},
		{"Ana", 2},
		{"Lia", 1},
		{"Aia", 1},
		{"A Ia", 2},
		{"Øya", 1},
		{"Ø Ya", 2},
		{"Banana", 3},
		{"Naan", 2},
		{"", 0},
		{"D’anad", 2},
		{"D'Italiano", 4},
	}

	for i, val := range words {
		if res, _, _ := server.CountSyllables(val.word); res != val.expect {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected %d, got %d", i, val.word, val.expect, res))
			t.Fail()
		}
	}

	invalid := []string{"汉语汉", "руссур", "Руссур", "0yø", "Português Brasileiro", "Français", "Español", "L’étoile", "L’ étoile"}
	for i, val := range invalid {
		if res, _, err := server.CountSyllables(val); err == nil {
			t.Logf(fmt.Sprintf("Failed on item #%d %s ; expected error, got %d", i, val, res))
			t.Fail()
		}
	}

	fmt.Printf("Finished countSyllables' %d tests\n", len(words)+len(invalid))
}

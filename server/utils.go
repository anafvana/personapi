package server

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func IsValidName(name string) bool {
	var re = regexp.MustCompile(`^([\p{L}\p{M}* '’])+$`)
	if found := re.FindAllString(name, -1); found == nil || len(found) > 1 {
		return false
	}
	return true
}

func IsPalindrome(word string) bool {
	re := regexp.MustCompile("[’' ]+")
	stripped := re.ReplaceAllString(strings.ToLower(word), "")
	bytes := []byte(stripped)
	runes := []rune{}

	for utf8.RuneCount(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		runes = append(runes, r)
		bytes = bytes[size:]
	}

	wLength := len(runes)
	for i := 0; i < wLength/2; i++ {
		if runes[i] != runes[wLength-1-i] {
			return false
		}
	}

	return true
}

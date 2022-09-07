package server

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
)

func IsValidName(name string, whatis string) (errMsg string) {
	name = strings.TrimSpace(name)
	if len(name) < 1 {
		return fmt.Sprintf("%s er tom. ", whatis)
	}

	re := regexp.MustCompile(`^([\p{L}\p{M}* '’])+$`)
	if found := re.FindAllString(name, -1); found == nil || len(found) > 1 {
		return fmt.Sprintf("%s har ugyldige karakterer. ", whatis)
	}
	return ""
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

func getRegexp2Matches(re *regexp2.Regexp, s string) []string {
	var matches []string

	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}

	return matches
}

func CountSyllables(name string) (syllables int, status int, err error) {
	name = strings.ToLower(name)

	reNorsk := regexp.MustCompile(`[^a-zæøå '’]`)
	if found := reNorsk.FindAllString(name, -1); found != nil || len(found) > 0 {
		return syllables, http.StatusNotImplemented, errors.New("kan ikke telle tale i ikke-norsk ord")
	}

	reSterke := regexp2.MustCompile(`(?<![uiy])[aeoåøæ](?![uiy])`, regexp2.None)
	reSvake := regexp2.MustCompile(`(?<![aeouiyåøæ])[uiy](?![aeouiyåøæ])`, regexp2.None)
	reDiftong := regexp.MustCompile(`([aeoåøæ][uiy][aeoåøæ]?)|([uiy][aeoåøæ])`)

	if foundSterke := getRegexp2Matches(reSterke, name); foundSterke != nil {
		syllables += len(foundSterke)
	}
	if foundSvake := getRegexp2Matches(reSvake, name); foundSvake != nil {
		syllables += len(foundSvake)
	}
	if foundDiftong := reDiftong.FindAllString(name, -1); foundDiftong != nil {
		syllables += len(foundDiftong)
	}

	return syllables, http.StatusOK, nil
}

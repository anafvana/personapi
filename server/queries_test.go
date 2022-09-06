package server_test

import (
	"personapi/server"
	"testing"
)

func TestIsValidName(t *testing.T) {
	if res := server.IsValidName("English"); res != true {
		t.Fail()
	}
	server.IsValidName("Simplified English")
	server.IsValidName("Português Brasileiro")
	server.IsValidName("Français")
	server.IsValidName("Español")
	server.IsValidName("D'Italiano")
	server.IsValidName("L’étoile")
	server.IsValidName("van der Dutch")
	server.IsValidName("русские")
	server.IsValidName("汉语")
	server.IsValidName("漢語")
	server.IsValidName("中文")
	server.IsValidName("日本国")
	server.IsValidName("日1本国")
	server.IsValidName("Norsk0 Dansk")
	server.IsValidName("0")


}
package server

import (
	"database/sql"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	UserId 		int 	`json:"brukerid,omitempty"`
	Fornavn 	string 	`json:"fornavn"`
	Etternavn 	string 	`json:"etternavn"`
}

func IsValidName(name string) bool {
	var validChars = regexp.MustCompile(`^([\p{L}\p{M}* '’])+$`)
	if found := validChars.FindAllString(name, -1); found == nil || len(found) > 1 {
		return false
	}
	return true
}

func (r repo) GetPerson(ctx *gin.Context) {
	// Retrieve :id parameter
	id := ctx.Param("id")
	
	// Ensure :id is convertible to int
	_, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var person Person
	row := r.db.QueryRow(`SELECT * FROM Applications WHERE userid = $1`, id)
	if err := row.Scan(
		&person.UserId,
		&person.Fornavn,
		&person.Etternavn); err != nil && err != sql.ErrNoRows {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	} else if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(200, person)
}

func (r repo) PostPerson(ctx *gin.Context) {
	var person Person
	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

}

func (r repo) UpdatePerson(ctx *gin.Context) {}

func (r repo) DeletePerson(ctx *gin.Context) {}

func (r repo) GetPalindrome(ctx *gin.Context) {}

func (r repo) GetPalindromeFornavn(ctx *gin.Context) {}

func (r repo) GetPalindromeEtternavn(ctx *gin.Context) {}

func (r repo) GetSyllables(ctx *gin.Context) {}

func (r repo) GetSyllablesFornavn(ctx *gin.Context) {}

func (r repo) GetSyllablesEtternavn(ctx *gin.Context) {}

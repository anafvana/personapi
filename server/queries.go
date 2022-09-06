package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	UserId    *int   `json:"brukerid,omitempty"`
	Fornavn   string `json:"fornavn"`
	Etternavn string `json:"etternavn"`
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
	row := r.db.QueryRow(`SELECT * FROM Persons WHERE bruker_id = $1`, id)
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

	errMsg := ""
	if !IsValidName(person.Fornavn) {
		errMsg = "Fornavn er ikke gyldig\n"
	}

	if !IsValidName(person.Etternavn) {
		errMsg = "Etternavn er ikke gyldig"
	}

	if errMsg != "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errMsg})
		return
	}

	var userId int
	if err = r.db.QueryRow(
		`INSERT INTO Persons(
			fornavn,
			etternavn)
			VALUES($1, $2) RETURNING bruker_id`,
		&person.Fornavn,
		&person.Etternavn,
	).Scan(&userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(200, userId)
}

func (r repo) UpdatePerson(ctx *gin.Context) {
	var person Person
	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if person.UserId == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "brukerid må informeres"})
		return
	}

	errMsg := ""
	if !IsValidName(person.Fornavn) {
		errMsg = "Ny fornavn er ikke gyldig\n"
	}

	if !IsValidName(person.Etternavn) {
		errMsg = "Ny etternavn er ikke gyldig"
	}

	if errMsg != "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errMsg})
		return
	}

	res, err := r.db.Exec(
		`UPDATE Person 
		SET 
			fornavn = $1, 
			etternavn = $2
		WHERE bruker_id = $3`,
		person.Fornavn, person.Etternavn, *person.UserId,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": fmt.Sprintf("Kunne ikke finne bruker %d", *person.UserId)})
		return
	}

	ctx.JSON(200, rowsAffected)
}

func (r repo) DeletePerson(ctx *gin.Context) {}

func (r repo) GetPalindrome(ctx *gin.Context) {}

func (r repo) GetPalindromeFornavn(ctx *gin.Context) {}

func (r repo) GetPalindromeEtternavn(ctx *gin.Context) {}

func (r repo) GetSyllables(ctx *gin.Context) {}

func (r repo) GetSyllablesFornavn(ctx *gin.Context) {}

func (r repo) GetSyllablesEtternavn(ctx *gin.Context) {}

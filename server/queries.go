package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	UserId    *int   `json:"brukerid,omitempty"`
	Fornavn   string `json:"fornavn"`
	Etternavn string `json:"etternavn"`
}

func (r repo) fetchPerson(ctx *gin.Context) (person Person, status int, err error) {
	// Retrieve :id parameter
	id := ctx.Param("id")

	// Ensure :id is convertible to int
	_, err = strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	row := r.db.QueryRow(`SELECT * FROM Person WHERE bruker_id = $1`, id)
	if err = row.Scan(
		&person.UserId,
		&person.Fornavn,
		&person.Etternavn); err != nil && err != sql.ErrNoRows {
		return person, http.StatusInternalServerError, err
	} else if err == sql.ErrNoRows {
		return person, http.StatusNotFound, fmt.Errorf("kunne ikke finne bruker %s", id)
	}
	return person, http.StatusOK, nil
}

func (r repo) GetPerson(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
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

	errMsg := IsValidName(person.Fornavn, "fornavn")
	errMsg += IsValidName(person.Etternavn, "etternavn")

	if errMsg != "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": errMsg})
		return
	}

	var userId int
	if err = r.db.QueryRow(
		`INSERT INTO Person (
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

	errMsg := IsValidName(person.Fornavn, "fornavn")
	errMsg += IsValidName(person.Etternavn, "etternavn")

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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": fmt.Sprintf("kunne ikke finne bruker %d", *person.UserId)})
		return
	}

	ctx.JSON(200, rowsAffected)
}

func (r repo) DeletePerson(ctx *gin.Context) {
	// Retrieve :id parameter
	id := ctx.Param("id")

	// Ensure :id is convertible to int
	_, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	res, err := r.db.Exec(`DELETE FROM Person WHERE bruker_id = $1`, id)
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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": fmt.Sprintf("kunne ikke finne bruker %s", id)})
		return
	}

	ctx.JSON(200, rowsAffected)
}

func (r repo) GetPalindrome(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	palindrome := IsPalindrome(person.Fornavn + person.Etternavn)

	ctx.JSON(200, palindrome)
}

func (r repo) GetPalindromeFornavn(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	palindrome := IsPalindrome(person.Fornavn)

	ctx.JSON(200, palindrome)
}

func (r repo) GetPalindromeEtternavn(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	palindrome := IsPalindrome(person.Etternavn)

	ctx.JSON(200, palindrome)
}

func (r repo) GetSyllables(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	syllables, status, err := CountSyllables(fmt.Sprintf("%s %s", person.Fornavn, person.Etternavn))
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(200, syllables)
}

func (r repo) GetSyllablesFornavn(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	syllables, status, err := CountSyllables(person.Fornavn)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(200, syllables)
}

func (r repo) GetSyllablesEtternavn(ctx *gin.Context) {
	person, status, err := r.fetchPerson(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	syllables, status, err := CountSyllables(person.Etternavn)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(200, syllables)
}

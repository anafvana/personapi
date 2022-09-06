package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


type repo struct {
	db   *sql.DB
}



func startDB() *sql.DB {
	const (
		host   = "localhost"
		port   = 5432
		username = "user123"
		passwd = "password123"
		dbname = "persondb"
	)

	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, passwd, dbname)
	db, err := sql.Open("postgres", params)
	if err != nil {
		panic(fmt.Sprintf("Cannot open database. \nERROR: %s", err))
	}

	return db
}


func setRouter(r repo) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	router.GET("/person/:id", r.GetPerson)
	router.GET("/palindrome/:id/fornavn", r.GetPalindromeFornavn)
	router.GET("/palindrome/:id/eternavn", r.GetPalindromeEtternavn)
	router.GET("/palindrome/:id", r.GetPalindrome)
	router.GET("/syllables/:id/fornavn", r.GetSyllablesFornavn)
	router.GET("/syllables/:id/etternavn", r.GetSyllablesEtternavn)
	router.GET("/syllables/:id", r.GetSyllables)
	router.POST("/person", r.PostPerson)
	router.PUT("/person/:id", r.UpdatePerson)
	router.DELETE("/person/:id", r.DeletePerson)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	port := "8080"

	r := repo {
		db: startDB(),
	}
	defer r.db.Close()

	router := setRouter(r)
	router.Run(fmt.Sprintf(":%s", port))
}
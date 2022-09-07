package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type repo struct {
	db *sql.DB
}

func startDB() *sql.DB {
	const (
		host     = "personapi_postgres"
		port     = 5432
		username = "user123"
		passwd   = "password123"
		dbname   = "personapi"
	)

	// params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, passwd, dbname)
	params := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, passwd, host, port, dbname)

	db, err := sql.Open("postgres", params)
	if err != nil {
		panic(fmt.Sprintf("Cannot open database. ERROR: %s", err))
	}

	return db
}

func setRouter(r repo) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	// router.SetTrustedProxies([]string{"localhost:3000", "159.223.16.218"})

	router.GET("/person/:id", r.GetPerson)
	router.GET("/palindrom/:id/fornavn", r.GetPalindromeFornavn)
	router.GET("/palindrom/:id/etternavn", r.GetPalindromeEtternavn)
	router.GET("/palindrom/:id", r.GetPalindrome)
	router.GET("/stavelser/:id/fornavn", r.GetSyllablesFornavn)
	router.GET("/stavelser/:id/etternavn", r.GetSyllablesEtternavn)
	router.GET("/stavelser/:id", r.GetSyllables)
	router.POST("/person", r.PostPerson)
	router.PUT("/person", r.UpdatePerson)
	router.DELETE("/person/:id", r.DeletePerson)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	port := "8080"

	r := repo{
		db: startDB(),
	}
	defer r.db.Close()

	router := setRouter(r)
	router.Run(fmt.Sprintf(":%s", port))
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sparkymat/trackit/database"
	"github.com/sparkymat/trackit/internal/session"
)

func main() {
	db, err := sqlx.Connect("postgres", "dbname=trackit_development sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	userDB := database.NewUser(db)
	u, err := userDB.Find(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", u)

	r := gin.Default()
	r.POST("/login", session.CreateHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

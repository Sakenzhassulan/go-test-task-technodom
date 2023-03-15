package main

import (
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/internal/config"
	"github.com/Sakenzhassulan/go-test-task-technodom/internal/server"
	cache2 "github.com/Sakenzhassulan/go-test-task-technodom/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbi, err := db.NewDbCollection(&conf)
	if err != nil {
		log.Fatal(err)
	}

	cache := cache2.NewCache(1000)
	r := gin.Default()
	r.Use(cors.Default())
	client, err := server.NewClient(&conf, dbi, cache)
	if err != nil {
		log.Fatal(err)
	}

	server.RegisterRoutes(r, client)
	if err := r.Run(conf.ServerPort); err != nil {
		log.Fatalln(err)
	}

}

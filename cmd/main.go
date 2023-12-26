package main

import (
	"net/http"
	"wb-0/api"
	"wb-0/config"
	"wb-0/entity"
	"wb-0/nats"

	"wb-0/repository/cache"
	"wb-0/repository/db"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.New()

	conn, err := db.NewPGConn(cfg.DATABASE_URL)
	if err != nil {
		panic("Failed to setup a new connection")
	}

	conn.AutoMigrate(&entity.Order{}, &entity.Item{}, &entity.Delivery{}, &entity.Payment{})

	repo := db.NewPGOrderRepo(conn)

	cache := cache.NewCache()
	err = repo.FillCache(cache)
	if err != nil {
		panic("Failed to fill cache")
	}

	subscriber := nats.NewNatsSub(cfg.ClusterID, cfg.ClientID)
	subscriber.Run(repo, cache)

	handler := api.NewHandler(cache)

	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ch.html", nil)
	})
	router.GET("/order", handler.RecordById)

	router.Run(cfg.HOST + ":" + cfg.PORT)

}

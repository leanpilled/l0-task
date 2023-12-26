package api

import (
	"net/http"
	"strconv"
	"wb-0/repository/cache"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cache *cache.Cache
}

func NewHandler(cache *cache.Cache) *Handler {
	return &Handler{cache: cache}
}

func (h *Handler) RecordById(c *gin.Context) {
	id_s := c.Query("id")

	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Wrong param")
		return
	}

	record, ok := h.cache.GetByID(id)
	if ok != true {
		c.IndentedJSON(http.StatusNotFound, "Not found in cache")
		return
	}

	c.IndentedJSON(http.StatusOK, record)
}

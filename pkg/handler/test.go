package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetTests(c *gin.Context) {
	id, _ := getUserId(c)
	SendJSONResponse(c, "id", id)
}

func (h *Handler) CreateTest(c *gin.Context) {
	id, _ := getUserId(c)
	SendJSONResponse(c, "id", id)
}

package controllers

import (
	"book-store-server/models/request"
	"book-store-server/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *AppController) CreateBook(ctx *gin.Context) {
	var req request.CreateBookRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, response.Response{
			Error: "Invalid request body: " + err.Error(),
		})
		return
	}

	res, code := c.AppService.CreateBook(&req)

	ctx.IndentedJSON(code, res)
}

func (c *AppController) UpdateBook(ctx *gin.Context) {
	var req request.UpdateBookRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, response.Response{
			Error: "Invalid request body: " + err.Error(),
		})
		return
	}

	res, code := c.AppService.UpdateBook(&req)

	ctx.IndentedJSON(code, res)
}

func (c *AppController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, response.Response{
			Error: "Invalid book ID: " + err.Error(),
		})
		return
	}

	res, code := c.AppService.DeleteBook(uint(bookId))

	ctx.IndentedJSON(code, res)
}

func (c *AppController) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, response.Response{
			Error: "Invalid book ID: " + err.Error(),
		})
		return
	}

	res, code := c.AppService.GetBook(uint(bookId))

	ctx.IndentedJSON(code, res)
}

func (c *AppController) GetBooks(ctx *gin.Context) {
	limitParam := ctx.DefaultQuery("limit", "10")
	pageParam := ctx.DefaultQuery("page", "1")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Limit must be a positive integer"})
		return
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Page must be a positive integer"})
		return
	}

	res, code := c.AppService.GetBooks(limit, page)

	ctx.IndentedJSON(code, res)
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/Abrar-Ahmed7/rest-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService service.BookService
}

type BookHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	Create(c *gin.Context)
	GetBooks(c *gin.Context)
}

func NewBookHandler(bs service.BookService) BookHandler {
	return &bookHandler{
		bookService: bs,
	}
}

func (bh bookHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/user/:user_id/add-book", bh.Create)
	router.GET("/books", bh.GetBooks)
	router.PUT("/user/:user_id/book/:book_id", bh.UpdateBook)
	router.DELETE("/user/:user_id/book/:book_id", bh.DeleteBook)
}

func (bh bookHandler) Create(c *gin.Context) {
	type BookReqBody struct {
		Title       string `json:"title" binding:"required,min=3,max=25"`
		AuthorName  string `json:"author_name" binding:"required,min=3,max=25"`
		Publication string `json:"publication" binding:"required,min=3,max=25"`
	}

	var req BookReqBody
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title or author_name or publication can't be empty",
		})
		return
	}

	userId := c.Param("user_id")
	title := req.Title
	authorName := req.AuthorName
	publication := req.Publication
	err = bh.bookService.Save(convertToInt(userId), title, authorName, publication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't add the book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book added successfully",
	})
}

func (bh bookHandler) GetBooks(c *gin.Context) {
	books, err := bh.bookService.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't get the books",
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bh bookHandler) UpdateBook(c *gin.Context) {
	type BookReqBody struct {
		Title       string `json:"title" binding:"min=3,max=25"`
		AuthorName  string `json:"author_name" binding:"min=3,max=25"`
		Publication string `json:"publication" binding:"min=3,max=25"`
	}
	var req BookReqBody
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mininum character:3, Maximum character: 25",
		})
		return
	}

	bookId := c.Param("book_id")
	userId := c.Param("user_id")
	title := req.Title
	authorName := req.AuthorName
	publication := req.Publication
	err = bh.bookService.Update(convertToInt(bookId), convertToInt(userId), title, authorName, publication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't update the book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book Updated successfully",
	})
}

func (bh bookHandler) DeleteBook(c *gin.Context) {
	bookId := c.Param("book_id")
	userId := c.Param("user_id")
	err := bh.bookService.Delete(convertToInt(bookId), convertToInt(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't delete the book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
}

func convertToInt(s string) int {
	cId, _ := strconv.Atoi(s)
	return cId
}

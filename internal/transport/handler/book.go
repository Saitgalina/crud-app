package handler

import (
	"fmt"
	"github.com/Saitgalina/crud-app/internal/core/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createBook(c *gin.Context) {
	//создает только админ
	idUser, err := GetUserId(c)
	if err != nil {
		return
	}
	var input model.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	idBook, err := h.services.Book.CreateBook(idUser, input) //h.services.CreateBook(idUser, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": idBook,
	})
}

type getAllBooksResponse struct {
	Data []model.Book `json:"data"`
}

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.Book.GetAllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

func (h *Handler) getBookById(c *gin.Context) {
	idBook, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.Book.GetByIdBook(idBook)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("book_handler after ")
	c.JSON(http.StatusOK, book)
}

func (h *Handler) getBookByName(c *gin.Context) {

}
func (h *Handler) getBookByAuthor(c *gin.Context) {

}

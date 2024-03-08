package handler

import (
	"fmt"
	"github.com/Saitgalina/crud-app/internal/core/model"
	"github.com/gin-gonic/gin"
	"net/http"
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
	fmt.Println("ВЫЗОВ getAllBooks")
	books, err := h.services.Book.GetAllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

func (h *Handler) getFilterBooks(c *gin.Context) {
	fmt.Println("ВЫЗОВ filter")
	var books []model.Book
	if filterName := c.Query("name"); filterName != "" {
		tmp, err := h.services.Book.GetByNameBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
		fmt.Println("ВЫЗОВ filter with NAME")
		fmt.Println(filterName)
	} else if filterName := c.Query("year"); filterName != "" {
		tmp, err := h.services.Book.GetByYearBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
		fmt.Println("ВЫЗОВ filter with YEAR")
		fmt.Println(filterName)
	} else if filterName := c.Query("author"); filterName != "" {
		tmp, err := h.services.Book.GetByAuthorBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
		fmt.Println("ВЫЗОВ filter with YEAR")
		fmt.Println(filterName)
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

func (h *Handler) getSortBooks(c *gin.Context) {
	fmt.Println("ВЫЗОВ sort")
	var books []model.Book
	if filterName := c.Query("desc"); filterName != "" {
		tmp, err := h.services.Book.GetSortDescBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
		fmt.Println("ВЫЗОВ sort DESC END")
		fmt.Println(filterName)
	} else if filterName := c.Query("asc"); filterName != "" {
		tmp, err := h.services.Book.GetSortAscBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
		fmt.Println("ВЫЗОВ sort ASC END")
		fmt.Println(filterName)
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

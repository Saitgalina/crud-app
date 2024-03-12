package handler

import (
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

func (h *Handler) getFilterBooks(c *gin.Context) {
	var books []model.Book
	if filterName := c.Query("name"); filterName != "" {
		tmp, err := h.services.Book.GetByNameBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
	} else if filterName := c.Query("year"); filterName != "" {
		tmp, err := h.services.Book.GetByYearBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
	} else if filterName := c.Query("author"); filterName != "" {
		tmp, err := h.services.Book.GetByAuthorBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

func (h *Handler) getSortBooks(c *gin.Context) {
	var books []model.Book
	if filterName := c.Query("desc"); filterName != "" {
		tmp, err := h.services.Book.GetSortDescBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
	} else if filterName := c.Query("asc"); filterName != "" {
		tmp, err := h.services.Book.GetSortAscBook(filterName)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		books = tmp
	} else {
		newErrorResponse(c, http.StatusBadRequest, "pass to the parameter asc or desc")
		return
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

type getAddFavouriteBookResponse struct {
	Mess string `json:"mess"`
}

func (h *Handler) addFavouriteBook(c *gin.Context) {
	idUser, err := GetUserId(c)
	if err != nil {
		return
	}
	idBook, err_b := strconv.Atoi(c.Param("id"))
	if err_b != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	message, err := h.services.Book.AddFavouritesBook(idBook, idUser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAddFavouriteBookResponse{
		Mess: message,
	})
}

func (h *Handler) GetFavouritesBooks(c *gin.Context) {
	idUser, err := GetUserId(c)
	if err != nil {
		return
	}
	books, err := h.services.Book.GetFavouritesBooks(idUser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

package controller

import (
	"books/models"
	"books/responses"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app App) CreateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var formBook models.FormBook

	if err := c.ShouldBindJSON(&formBook); err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	inserted, err := app.model.InsertBook(formBook)
	if err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	logInserted, err := json.Marshal(inserted)
	if err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	app.logging.Info().Msgf("[CREATE] success => %s", string(logInserted))

	c.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Created succesfully",
		Data:    inserted,
	})
}

func (app App) ReadBooks(c *gin.Context) {
	books, err := app.model.SelectBooks()

	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	logBooks, err := json.Marshal(books)
	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	app.logging.Info().Msgf("[BOOKS] => %s", string(logBooks))

	c.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Data load succesfully",
		Data:    books,
	})
}

func (app App) UpdateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	bookId, err := primitive.ObjectIDFromHex(c.Param("bookId"))
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var formBook models.Books
	if err := c.ShouldBind(&formBook); err != nil {
		app.logging.Error().Msgf("[UPDATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	updated, err := app.model.UpdateBook(bookId, formBook)
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	logUpdated, err := json.Marshal(updated)
	if err != nil {
		app.logging.Error().Msgf("[UPDATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	app.logging.Info().Msgf("[UPDATE] success => %s", string(logUpdated))

	c.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "updated succesfull",
		Data:    updated,
	})
}

func (app App) DeleteBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	bookId, err := primitive.ObjectIDFromHex(c.Param("bookId"))
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	deleted, err := app.model.DeleteBook(bookId)
	if err != nil {
		app.logging.Error().Msgf("[DELETE] failed,  %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if deleted == 0 {
		app.logging.Error().Msgf("[DELETE] failed,  %d\n", deleted)
		c.JSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Delete failed",
		})
		return
	}

	app.logging.Info().Msgf("[DELETE] success => %d", deleted)

	c.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "delete succesfull",
	})
}

func (app App) ReadBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	bookId, _ := primitive.ObjectIDFromHex(c.Param("bookId"))

	book := app.model.SelectBookById(bookId)

	logBook, err := json.Marshal(book)
	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	app.logging.Info().Msgf("[BOOK ONE] => %s", string(logBook))

	c.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Data load succesfully",
		Data:    book,
	})
}

package controller

import (
	"books/models"
	"books/responses"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app App) CreateBook(c *gin.Context) {
	var formBook models.FormBook

	if err := c.ShouldBind(&formBook); err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusBadRequest, responses.ResponseSuccess{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	inserted, err := app.model.InsertBook(formBook)
	if err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.ResponseSuccess{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	logInserted, err := json.Marshal(inserted)
	if err != nil {
		app.logging.Error().Msgf("[CREATE] failed => %+v\n", err.Error())
		c.JSON(http.StatusInternalServerError, responses.ResponseSuccess{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	app.logging.Info().Msgf("[CREATE] success => %s", string(logInserted))

	c.JSON(http.StatusOK, responses.ResponseSuccess{
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
		c.JSON(http.StatusInternalServerError, responses.ResponseSuccess{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	logBooks, err := json.Marshal(books)
	if err != nil {
		app.logging.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, responses.ResponseSuccess{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	app.logging.Info().Msgf("[BOOKS] => %s", string(logBooks))

	c.JSON(http.StatusOK, responses.ResponseSuccess{
		Code:    http.StatusOK,
		Success: true,
		Message: "Data load succesfully",
		Data:    books,
	})
}

package responses

import (
	"books/models"
)

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseInsertSuccess struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    models.InsertOneResult
}

// type ResponseFailed struct {
// 	Code    int    `json:"code"`
// 	Success bool   `json:"success"`
// 	Message string `json:"message"`
// }

// Package controllers contains the common functions and structures for the controllers
package controllers

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// BindJSON is a function that binds the request body to the given struct and rewrite the request body on the context
func BindJSON(ctx *fiber.Ctx, request interface{}) error {
	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	return nil
}

// BindJSONMap is a function that binds the request body to the given map and rewrite the request body on the context
func BindJSONMap(c *gin.Context, request *map[string]interface{}) (err error) {
	buf := make([]byte, 5120)
	num, _ := c.Request.Body.Read(buf)
	reqBody := buf[0:num]
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	err = json.Unmarshal(reqBody, &request)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	return
}

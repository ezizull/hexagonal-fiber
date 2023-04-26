// Package middlewares contains the middlewares for the restapi api
package middlewares

import "github.com/gofiber/fiber/v2"

// CommonHeaders is a middleware that adds common headers to the response
func CommonHeaders(ctx *fiber.Ctx) error {
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Set("Access-Control-Allow-Credentials", "true")
	ctx.Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, GET, PUT")
	ctx.Set("Access-Control-Allow-Headers",
		"Content-Type, Depth, UserName-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-CompanyName, Cache-Control")
	ctx.Set("X-Frame-Options", "SAMEORIGIN")
	ctx.Set("Cache-Control", "no-cache, no-store")
	ctx.Set("Pragma", "no-cache")
	ctx.Set("Expires", "0")

	return ctx.Next()
}

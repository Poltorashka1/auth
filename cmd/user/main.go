package main

import (
	"auth/internal/app"
	"context"
)

// @title           User API
// @version         1.0
// @description     This is a User API microservice

// @host      localhost:50052
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	ctx := context.Background()

	// todo error with code 500 dont return, return only code
	a := app.New(ctx)

	a.Run()
}

// https://github.com/gin-gonic/examples/tree/master/assets-in-binary
//go:generate go-assets-builder assets -o assets/assets.go -p assets

package main

import (
	`github.com/mdaliyan/ginline/internal/config`
	`github.com/mdaliyan/ginline/internal/server`
)

func main() {

	// gin.SetMode(gin.ReleaseMode)

	config.ParseFlags()

	server.LoadTemplateEngine()

	server.SetRoutes()

	server.R.Run(config.Address)
}

package server

import (
	`net/http`

	`github.com/flosch/pongo2`
	`github.com/gin-gonic/gin`

	`github.com/mdaliyan/ginline/assets`
)

var R = gin.Default()

func SetRoutes() {

	R.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", pongo2.Context{
			"Foo": "World",
		})
	})

	R.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", pongo2.Context{
			"Bar": "World",
		})
	})

	// serve static files
	R.GET("/static/*filepath", StaticHandler)

}

func StaticHandler(c *gin.Context) {
	file := c.Param("filepath")
	c.FileFromFS("/assets/static"+file, assets.Assets)
}

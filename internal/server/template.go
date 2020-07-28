package server

import (
	`github.com/mdaliyan/ginline/assets`

	"github.com/mdaliyan/ginline/internal/pongo2"
)

func LoadTemplateEngine() {
	HTMLRender := pongo2engine.New(
		"./assets",
		"text/html; charset=utf-8",
	)
	if err := HTMLRender.LoadTemplates(assets.Assets, ".html"); err != nil {
		panic(err)
	}
	R.HTMLRender = HTMLRender
}

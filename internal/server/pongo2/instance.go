// Package pongo2gin is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the Pongo2 template
// library https://github.com/flosch/pongo2
package pongo2engine

import (
	"net/http"

	"github.com/flosch/pongo2"
)

// instance
type Instance struct {
	ContentType string
	Template    *pongo2.Template
	Context     pongo2.Context
}

// htmlRender should render the template to the response.
func (i Instance) Render(w http.ResponseWriter) error {
	i.WriteContentType(w)
	err := i.Template.ExecuteWriter(i.Context, w)
	return err
}

// WriteContentType should add the Content-Type header to the response
// when not set yet.
func (i Instance) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{i.ContentType}
	}
}

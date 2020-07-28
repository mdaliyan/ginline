// Package pongo2gin is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the Pongo2 template
// library https://github.com/flosch/pongo2
package pongo2engine

import (
	`fmt`
	`io/ioutil`
	"path"
	`strings`
	`sync`

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/jessevdk/go-assets"
)

// HTMLRender interface is to be implemented by HTMLProduction and HTMLDebug.
type HTMLRender interface {

	// Instance returns an HTML instance.
	Instance(string, interface{}) render.Render

	// Set templates
	SetTemplates(map[string]*pongo2.Template)

	// Loads templates from file system
	LoadTemplates(*assets.FileSystem, string) error
}

// htmlRender is a custom Gin template renderer using Pongo2.
type htmlRender struct {
	templateDir    string
	contentType    string
	templates      map[string]*pongo2.Template
	templatesMutex sync.RWMutex
}

// New creates a new htmlRender instance with custom Options.
func New(TemplateDir, ContentType string) HTMLRender {
	// if TemplateDir[0] != '/' {
	// 	TemplateDir = "/" + TemplateDir
	// }
	return &htmlRender{
		templateDir: TemplateDir,
		contentType: ContentType,
	}
}

// Default creates a htmlRender instance with default options.
func Default() render.HTMLRender {
	return New("templates", "text/html; charset=utf-8")
}

func (p *htmlRender) SetTemplates(templates map[string]*pongo2.Template) {
	p.templatesMutex.Lock()
	p.templates = templates
	p.templatesMutex.Unlock()
}

func (p *htmlRender) LoadTemplates(fs *assets.FileSystem, suffix string) error {

	var templates = make(map[string]*pongo2.Template)

	for name, file := range fs.Files {
		if file.IsDir() || !strings.HasSuffix(name, suffix) {
			continue
		}
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		templates[name[1:]] = pongo2.Must(pongo2.FromBytes(b))
	}
	p.SetTemplates(templates)
	return nil
}

// Instance should return a new htmlRender struct per request and prepare
// the template by either loading it from disk or using pongo2's cache.
func (p *htmlRender) Instance(name string, data interface{}) render.Render {
	var instance = Instance{
		Context:     data.(pongo2.Context),
		ContentType: p.contentType,
	}

	filename := path.Join(p.templateDir, name)

	fmt.Println(filename)

	// always read template files from disk if in debug mode, use cache otherwise.
	if gin.Mode() == "debug" {
		instance.Template = pongo2.Must(pongo2.FromFile(filename))
		return instance
	}

	if p.templates == nil {
		instance.Template = pongo2.Must(pongo2.FromCache(filename))
		return instance
	}

	p.templatesMutex.RLock()
	instance.Template = p.templates[filename]
	p.templatesMutex.RUnlock()

	// file := assets.Assets.Files[filename]
	//
	// instance.Template = pongo2.Must(pongo2.FromBytes(file.Data))

	return instance
}

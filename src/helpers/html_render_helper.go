package helpers

import (
	"html/template"

	"github.com/gin-gonic/gin/render"
)

type HTMLRenderHelper struct {
	templates map[string]*template.Template
}

func (r *HTMLRenderHelper) add(name string, tmpl *template.Template) {
	if r.templates == nil {
		r.templates = make(map[string]*template.Template)
	}
	r.templates[name] = tmpl
}

func (r *HTMLRenderHelper) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.templates[name],
		Data:     data,
	}
}

const (
	LAYOUT_A string = "templates/common/base.layout.gohtml"
)

type TemplateModel struct {
	Name       string
	LayoutPath string
	PagePath   string
}

func CreateHTMLRenderHelper(tModels []TemplateModel) render.HTMLRender {
	r := &HTMLRenderHelper{}

	for _, tModel := range tModels {
		r.add(tModel.Name, template.Must(template.ParseFiles(tModel.LayoutPath, tModel.PagePath)))
	}

	return r
}

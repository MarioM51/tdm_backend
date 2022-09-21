package helpers

import (
	"html/template"
	"path"

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

func CreateHTMLRenderHelper(tModels []TemplateModel, consts Constants) render.HTMLRender {
	r := &HTMLRenderHelper{}

	myFuncs := template.FuncMap{
		"StaticFolder": func() string {
			return consts.StaticFolder
		},
		"WebComponentsFolder": func() string {
			return consts.WebComponentsFolder
		},
	}

	for _, tModel := range tModels {
		name := path.Base(tModel.LayoutPath)
		myTemplate := template.New(name)
		myTemplate.Funcs(myFuncs)
		_, err := myTemplate.ParseFiles(tModel.LayoutPath, tModel.PagePath)
		if err != nil {
			panic(err)
		}
		r.add(tModel.Name, myTemplate)
	}

	return r
}

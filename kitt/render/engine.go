package render

import (
	"fmt"
	"html/template"
	"io"
)

type Renderable interface {
	Render() string
}

type AsHtml = template.HTML
type AsAttr = template.HTMLAttr
type AnyCtx = map[string]interface{}
type Funcs = template.FuncMap

// Its a wrapper around html/template
type Engine interface {
	Render(w io.Writer, templateName string, ctx interface{}) error
	WithTemplates(pattern string) Engine
	WithFuncs(funcs Funcs) Engine
	WithTemplate(name string, template string) Engine
}

type engine struct {
	template *template.Template
}

func (e engine) Render(w io.Writer, templateName string, ctx interface{}) error {
	return e.template.ExecuteTemplate(w, templateName, ctx)
}

func (e *engine) WithTemplates(pattern string) Engine {
	template, err := e.template.ParseGlob(pattern)

	if err != nil {
		panic(err)
	}

	e.template = template

	return e
}

func (e *engine) WithTemplate(name string, str string) Engine {
	tpl := fmt.Sprintf(`{{ define "%s" }}%s{{ end }}`, name, str)
	template, err := e.template.Parse(tpl)

	if err != nil {
		panic(err)
	}

	e.template = template

	return e
}

func (e *engine) WithFuncs(funcs Funcs) Engine {
	e.template = e.template.Funcs(funcs)
	return e
}

func NewEngine() Engine {
	return &engine{
		template: template.New("__ROOT__"),
	}
}

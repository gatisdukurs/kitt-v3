package kitt

import (
	"bytes"
	"html/template"
)

var tplHtml = template.New("__ROOT__")

type TemplateSet struct {
	Pattern string
}

type htmxHtml struct {
	Name         string
	FallbackName string
	Ctx          interface{}
	OOB          []htmxOOB
}

type htmxOOB struct {
	TargetID string
	Name     string
	Ctx      interface{}
}

func (h *htmxHtml) WithFallback(name string) *htmxHtml {
	h.FallbackName = name
	return h
}

func (h *htmxHtml) WithOOB(targetID, tplName string, tplCtx any) *htmxHtml {
	h.OOB = append(h.OOB, htmxOOB{
		TargetID: targetID,
		Name:     tplName,
		Ctx:      tplCtx,
	})
	return h
}

func (h htmxHtml) Render() string {
	name := h.Name

	if ctx, ok := h.Ctx.(*RouteCtx); ok {
		if !ctx.Request().HTMX() && h.FallbackName != "" {
			name = h.FallbackName
		}
	}

	if tplHtml.Lookup(name) == nil {
		return "template not found: " + name
	}

	var buf bytes.Buffer

	if err := tplHtml.ExecuteTemplate(&buf, name, h.Ctx); err != nil {
		return err.Error()
	}

	if ctx, ok := h.Ctx.(*RouteCtx); ok && ctx.Request().HTMX() {
		for _, o := range h.OOB {
			buf.WriteString(`<div id="`)
			template.HTMLEscape(&buf, []byte(o.TargetID))
			buf.WriteString(`" hx-swap-oob="true">`)

			if tplHtml.Lookup(o.Name) == nil {
				buf.WriteString("template not found: " + o.Name)
			} else {
				if err := tplHtml.ExecuteTemplate(&buf, o.Name, o.Ctx); err != nil {
					buf.WriteString(err.Error())
				}
			}

			buf.WriteString(`</div>`)
		}
	}

	return buf.String()
}

type templateHtml struct {
	Name string
	Ctx  interface{}
}

func (t templateHtml) Render() string {
	if tplHtml.Lookup(t.Name) == nil {
		return "template not found: " + t.Name
	}

	var buffer bytes.Buffer
	err := tplHtml.ExecuteTemplate(&buffer, t.Name, t.Ctx)

	if err != nil {
		return err.Error()
	}

	return buffer.String()
}

type text struct {
	Text string
}

func (t text) Render() string {
	return t.Text
}

func registerTemplates(ts TemplateSet, prefix string) {
	tplHtml.ParseGlob(prefix + ts.Pattern)
}

func RegisterTemplateFuncs(funcs template.FuncMap) {
	tplHtml.Funcs(funcs)
}

func HTML(name string, ctx interface{}) templateHtml {
	tHtml := templateHtml{
		Name: name,
		Ctx:  ctx,
	}
	return tHtml
}

func HTMX(name string, ctx interface{}) *htmxHtml {
	tHtmx := &htmxHtml{
		Name: name,
		Ctx:  ctx,
	}
	return tHtmx
}

func TEXT(t string) text {
	return text{
		Text: t,
	}
}

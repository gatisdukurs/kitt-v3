package htmx

import (
	"bytes"
	"fmt"
	"kitt/kitt/render"
	"strings"
)

const (
	HTMX_SWAP_INNER        = "innerHTML"
	HTMX_SWAP_OUTER        = "outerHTML"
	HTMX_SWAP_BEFORE_BEGIN = "beforebegin"
	HTMX_SWAP_BEFORE_END   = "beforeend"
	HTMX_SWAP_AFTER_BEGIN  = "afterbegin"
	HTMX_SWAP_AFTER_END    = "afterend"
	HTMX_SWAP_NONE         = "none"
	HTMX_SWAP_DEFAULT      = "true"
)

type HTMXElement interface {
	render.Renderable
	WithId(target string) HTMXElement
	WithSwap(swap string) HTMXElement
}

type htmxEl struct {
	id   string
	swap string
	view render.Renderable
}

func (h htmxEl) Render() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("<div%s>", h.renderAttributes()))
	buf.WriteString(h.view.Render())
	buf.WriteString("</div>")

	return buf.String()
}

func (h htmxEl) renderAttributes() string {
	attrs := []string{}

	if h.id != "" {
		attrs = append(attrs, fmt.Sprintf(`id="%s"`, h.id))
	}

	if h.swap != "" {
		attrs = append(attrs, fmt.Sprintf(`hx-swap-oob="%s"`, h.swap))
	}

	if len(attrs) == 0 {
		return ""
	}

	return " " + strings.Join(attrs, " ")
}

func (h *htmxEl) WithId(id string) HTMXElement {
	h.id = id
	return h
}

func (h *htmxEl) WithSwap(swap string) HTMXElement {
	h.swap = swap
	return h
}

func NewElement(view render.Renderable) HTMXElement {
	return &htmxEl{
		view: view,
	}
}

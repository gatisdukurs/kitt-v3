package pages

import (
	"fmt"
	"kitt/app/admin/internal/pages/jobs"
	"kitt/app/admin/internal/shared"
	"kitt/kitt/form"
	"kitt/kitt/htmx"
	"kitt/kitt/queue"
	"kitt/kitt/render"
	"kitt/kitt/repository"
	"kitt/kitt/router"
	"net/http"
	"net/url"
)

type Controller struct {
	shared.Controller
	Pages repository.Repository[Page, int64]
	Queue queue.Queue
}

func (c *Controller) Boot() {
	c.GET("/admin/pages", c.GetList)
	c.GET("/admin/pages/create", c.GetCreate)
	c.POST("/admin/pages", c.PostPage)

	c.DELETE("/admin/pages/:id", c.DeletePage)
	c.GET("/admin/pages/:id/edit", c.GetEdit)
	c.POST("/admin/pages/:id/edit", c.PostEdit)
}

func (c *Controller) GetList(ctx router.RouteCtx) router.RouteResponse {
	c.Queue.Dispatch(ctx.Request().HttpRequest().Context(), jobs.Ajob{
		SomeVar: "Getting List and other cool things now.",
	})
	// Data
	query := c.Pages.Query()
	rows := c.Pages.Find(query)
	pagesCtx := c.Ctx()
	pagesCtx["pages"] = rows

	// Views
	content := c.View("admin.pages.list")
	content.WithCtx(pagesCtx)

	navigation := c.Navigation(ctx)

	// HTMX vs Regular
	if ctx.Request().HTMX() {
		contentHtmx := c.Htmx(content).WithId("content")
		navigationHtmx := c.Htmx(navigation).
			WithId("navigation").
			WithSwap(htmx.HTMX_SWAP_DEFAULT)

		stack := c.Stack()
		stack.WithRenderable(contentHtmx)
		stack.WithRenderable(navigationHtmx)

		return c.Response(stack)
	} else {
		view := c.View("admin.layout")
		view.WithPartial("content", content)
		view.WithPartial("navigation", navigation)
		return c.Response(view)
	}
}

func (c *Controller) GetEdit(ctx router.RouteCtx) router.RouteResponse {
	// Data
	id := ctx.ParamInt64("id")
	row, err := c.Pages.ByID(id)

	if err != nil {
		return c.ResponseString("Page not found").WithStatus(http.StatusNotFound)
	}

	values := url.Values{}
	values.Set("title", row.Title)
	values.Set("content", row.Content)

	// Views
	f := c._PageForm(fmt.Sprintf("/admin/pages/%d/edit", id), values)
	content := c.View("admin.pages.edit").WithPartial("form", f)

	if ctx.Request().HTMX() {
		contentHtmx := c.Htmx(content).WithId("content")
		stack := c.Stack()
		stack.WithRenderable(contentHtmx)
		return c.Response(stack)
	} else {
		navigation := c.Navigation(ctx)
		view := c.View("admin.layout")
		view.WithPartial("content", content)
		view.WithPartial("navigation", navigation)
		return c.Response(view)
	}
}

func (c *Controller) PostEdit(ctx router.RouteCtx) router.RouteResponse {
	id := ctx.ParamInt64("id")
	row, err := c.Pages.ByID(id)

	if err != nil {
		return c.ResponseString("Page not found").WithStatus(http.StatusNotFound)
	}

	values := ctx.Request().FormValues()
	f := c._PageForm(fmt.Sprintf("/admin/pages/%d/edit", id), values)

	if f.Validate() {
		row.Title = values.Get("title")
		row.Content = values.Get("content")

		err := c.Pages.Update(row)

		if err != nil {
			f.WithError(err.Error())
		} else {
			stack := c.Stack(f, c.ToastHtmx("success", "Page Updated!"))
			return c.Response(stack)
		}
	} else {
		f.WithError("Form has some errors :(")
	}

	return c.Response(f)
}

func (c *Controller) GetCreate(ctx router.RouteCtx) router.RouteResponse {
	// View
	content := c.View("admin.pages.create")
	content.WithPartial("form", c._PageForm("/admin/pages", url.Values{}))

	if ctx.Request().HTMX() {
		contentHtmx := c.Htmx(content).WithId("content")
		stack := c.Stack()
		stack.WithRenderable(contentHtmx)
		return c.Response(stack)
	} else {
		navigation := c.Navigation(ctx)
		view := c.View("admin.layout")
		view.WithPartial("content", content)
		view.WithPartial("navigation", navigation)
		return c.Response(view)
	}
}

func (c *Controller) PostPage(ctx router.RouteCtx) router.RouteResponse {
	values := ctx.Request().FormValues()
	f := c._PageForm("/admin/pages", values)

	if f.Validate() {
		id, err := c.Pages.Create(Page{
			Title:   values.Get("title"),
			Content: values.Get("content"),
		})

		if err != nil {
			f.WithError(err.Error())
		} else {
			f.Reset()
			stack := c.Stack(f, c.ToastHtmx("success", "Page created ID: %d", id))
			return c.Response(stack)
		}
	} else {
		f.WithError("Form has some errors :(")
	}
	return c.Response(f)
}

func (c *Controller) DeletePage(ctx router.RouteCtx) router.RouteResponse {
	id := ctx.ParamInt64("id")
	err := c.Pages.Delete(id)

	query := c.Pages.Query()
	rows := c.Pages.Find(query)
	contentCtx := c.Ctx()
	contentCtx["pages"] = rows
	content := c.View("admin.pages.list")
	content.WithCtx(contentCtx)

	if err != nil {
		stack := c.Stack(content, c.ToastHtmx("danger", "Error: %s", err))
		return c.Response(stack)
	}

	return c.Response(content)
}

func (c *Controller) _PageForm(action string, values url.Values) form.Form {
	e := render.NewEngine()
	f := form.NewForm("form-page", e)
	f.WithMethod(http.MethodPost)
	f.WithAction(action)

	// title
	title := form.NewFormField("title-field", e)

	titleControl := form.NewFormControl("title", e)
	titleControl.WithValidators(form.Required(), form.MinLength(3))
	titleControl.WithAttribute("autofocus", "")
	title.WithControl(titleControl)

	titleLabel := form.NewFormLabel("Title", e)
	title.WithLabel(titleLabel)

	// content
	content := form.NewFormField("content-field", e)
	contentControl := form.NewFormControl("content", e)
	contentControl.WithType(form.FIELD_TEXTAREA)
	contentControl.WithValidators(form.Required(), form.MinLength(3))
	contentControl.WithAttribute("rows", "10")
	content.WithControl(contentControl)

	contentLabel := form.NewFormLabel("Content", e)
	content.WithLabel(contentLabel)

	// Actions
	actions := form.NewFormActions("form-actions", e)
	save := form.NewFormAction("Save", e)
	actions.WithAction(save)

	f.WithField(title)
	f.WithField(content)
	f.WithActions(actions)
	f.WithValues(values)
	// HTMX support
	f = htmx.HTMXForm(f)

	return f
}

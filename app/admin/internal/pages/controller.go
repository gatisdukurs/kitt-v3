package pages

import (
	"fmt"
	"kitt/app/admin/internal/shared"
	"kitt/kitt/form"
	"kitt/kitt/render"
	"kitt/kitt/repository"
	"kitt/kitt/router"
	"net/http"
)

type Controller struct {
	shared.Controller
	pages repository.Repository[Page, int64]
}

func (c *Controller) Boot() {
	// Repo
	c.pages = repository.Repo[Page, int64](repository.DRIVER_SQL, "db.sqlite")

	// Routes
	c.GET("/admin/pages", c.GetList)
	c.GET("/admin/pages/create", c.GetCreate)
	c.POST("/admin/pages", c.PostPage)
	c.DELETE("/admin/pages/:id", c.DeletePage)
}

func (c Controller) GetList(ctx router.RouteCtx) router.RouteResponse {
	// Data
	query := c.pages.Query()
	rows := c.pages.Find(query)
	pagesCtx := c.Ctx()
	pagesCtx.Set("pages", rows)

	// View
	view := c.View("admin.layout")
	content := c.View("admin.pages.list")
	content.WithCtx(pagesCtx.Basic())

	navigation := c.Navigation(ctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	view.WithHTMX("content", "navigation")
	// Send
	return c.Response(view)
}

func (c Controller) GetCreate(ctx router.RouteCtx) router.RouteResponse {
	// View
	view := c.View("admin.layout")
	content := c.View("admin.pages.create")
	content.WithPartial("form", c._PageForm())
	navigation := c.Navigation(ctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	view.WithHTMX("content", "navigation")
	// Send
	return c.Response(view)
}

func (c Controller) PostPage(ctx router.RouteCtx) router.RouteResponse {
	values := ctx.Request().FormValues()
	f := c._PageForm().WithValues(values)

	if f.Validate() {
		id, err := c.pages.Create(Page{
			Title:   values.Get("title"),
			Content: values.Get("content"),
		})

		if err != nil {
			f.WithError(err.Error())
		} else {
			f.Reset()
			f.WithSuccess(fmt.Sprintf("Page Created. ID: %d", id))
		}
	} else {
		f.WithError("Form has some errors :(")
	}
	return c.Response(f)
}

func (c Controller) DeletePage(ctx router.RouteCtx) router.RouteResponse {
	id := ctx.ParamInt64("id")
	err := c.pages.Delete(id)

	if err != nil {
		return c.ResponseString(err.Error()).WithStatus(http.StatusInternalServerError)
	}

	query := c.pages.Query()
	rows := c.pages.Find(query)
	contentCtx := c.Ctx()
	contentCtx.Set("pages", rows)
	content := c.View("admin.pages.list")
	content.WithCtx(contentCtx.Basic())

	return c.Response(content)
}

func (c Controller) _PageForm() form.Form {
	e := render.NewEngine()
	f := form.NewForm("form-page", e)
	f.WithMethod(http.MethodPost)
	f.WithAction("/admin/pages")

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

	f.WithHTMX()

	return f
}

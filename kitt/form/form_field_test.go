package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Field(t *testing.T) {
	t.Run("it renders without label and control", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)

		assertEqual(t, field.Render(), `<div class="control" id="email"></div>`)
	})

	t.Run("it renders with label and control", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		control := NewFormControl("email", e)
		field.WithControl(control)
		label := NewFormLabel("E-mail", e)
		field.WithLabel(label)

		assertEqual(t, field.Render(), `<div class="control" id="email"><label class="label">E-mail</label><input class="field" name="email" id="email" type="text" value="" /></div>`)
	})

	t.Run("it renders errors", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		label := NewFormLabel("E-mail", e)
		control := NewFormControl("email", e)
		control.WithValue("")
		control.WithValidators(Required(), MinLength(3))
		field.WithControl(control)
		field.WithLabel(label)

		_, errs := control.Validate()
		field.WithErrors(errs)

		assertEqual(t, field.Render(), `<div class="control" id="email"><label class="label">E-mail</label><input class="field" name="email" id="email" type="text" value="" /><ul class="errors"><li>This field is required</li><li>Must be at least 3 characters</li></ul></div>`)
	})

	t.Run("it sets errors", func(t *testing.T) {
		errs := []string{
			"Required",
			"Max length not right",
		}
		e := render.NewEngine()
		field := NewFormField("email", e)
		field.WithErrors(errs)

		assertEqual(t, field.Errors(), errs)
	})

	t.Run("it sets control", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		control := NewFormControl("email", e)
		field.WithControl(control)

		assertEqual(t, field.Control(), control)
	})

	t.Run("it returns id", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)

		assertEqual(t, field.Id(), "email")
	})

	t.Run("it sets label", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		label := NewFormLabel("email", e)
		field.WithLabel(label)

		assertEqual(t, field.Label(), label)
	})

	t.Run("it renders control", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		control := NewFormControl("email", e)
		field.WithControl(control)

		assertEqual(t, field.RenderControl(), control.Render())
	})

	t.Run("it renders label", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		label := NewFormLabel("E-mail", e)
		field.WithLabel(label)

		assertEqual(t, field.RenderLabel(), label.Render())
	})
}

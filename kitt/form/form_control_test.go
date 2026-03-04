package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Control(t *testing.T) {
	t.Run("it renders without label and field", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)

		assertEqual(t, control.Render(), `<div class="control" id="email"></div>`)
	})

	t.Run("it renders with label and field", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		field := NewFormField("email", e)
		control.WithField(field)
		label := NewFormLabel("E-mail", e)
		control.WithLabel(label)

		assertEqual(t, control.Render(), `<div class="control" id="email"><label class="label">E-mail</label><input class="field" name="email" id="email" type="text" value="" /></div>`)
	})

	t.Run("it renders errors", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		label := NewFormLabel("E-mail", e)
		field := NewFormField("email", e)
		field.WithValue("")
		field.WithValidators(Required(), MinLength(3))
		control.WithField(field)
		control.WithLabel(label)

		_, errs := field.Validate()
		control.WithErrors(errs)

		assertEqual(t, control.Render(), `<div class="control" id="email"><label class="label">E-mail</label><input class="field" name="email" id="email" type="text" value="" /><ul class="errors"><li>This field is required</li><li>Must be at least 3 characters</li></ul></div>`)
	})

	t.Run("it sets errors", func(t *testing.T) {
		errs := []string{
			"Required",
			"Max length not right",
		}
		e := render.NewEngine()
		control := NewFormControl("email", e)
		control.WithErrors(errs)

		assertEqual(t, control.Errors(), errs)
	})

	t.Run("it sets field", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		field := NewFormField("email", e)
		control.WithField(field)

		assertEqual(t, control.Field(), field)
	})

	t.Run("it returns id", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)

		assertEqual(t, control.Id(), "email")
	})

	t.Run("it sets label", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		label := NewFormLabel("email", e)
		control.WithLabel(label)

		assertEqual(t, control.Label(), label)
	})

	t.Run("it renders field", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		field := NewFormField("email", e)
		control.WithField(field)

		assertEqual(t, control.RenderField(), field.Render())
	})

	t.Run("it renders label", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		label := NewFormLabel("E-mail", e)
		control.WithLabel(label)

		assertEqual(t, control.RenderLabel(), label.Render())
	})
}

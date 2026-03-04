package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Control(t *testing.T) {
	t.Run("it renders default", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		control.WithValue("gatis.dukurs@gmail.com")

		assertEqual(t, control.Render(), `<input class="control" name="email" id="email" type="text" value="gatis.dukurs@gmail.com" />`)
	})

	t.Run("it renders textarea", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("content", e)
		control.WithType(FIELD_TEXTAREA)
		control.WithValue("Content")

		assertEqual(t, control.Render(), `<textarea class="control" name="content" id="content">Content</textarea>`)
	})

	t.Run("it renders unsupported type", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		control.WithType("unsupported")

		assertEqual(t, control.Render(), `unknown field type: unsupported`)
	})

	t.Run("it sets type", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		control.WithType(FIELD_EMAIL)

		assertEqual(t, control.Type(), FIELD_EMAIL)
	})

	t.Run("it sets  name", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)

		assertEqual(t, control.Name(), "email")
	})

	t.Run("it sets id", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("test_id", e)

		assertEqual(t, control.Id(), "test_id")

		control.WithId("new_test_id")
		assertEqual(t, control.Id(), "new_test_id")
	})

	t.Run("it sets  value", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		control.WithValue("gatis.dukurs@gmail.com")

		assertEqual(t, control.Value(), "gatis.dukurs@gmail.com")
	})

	t.Run("it works with validators", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		control.WithValue("")

		control.WithValidators(Required(), MinLength(3))

		valid, errors := control.Validate()

		assertEqual(t, valid, false)
		assertEqual(t, len(errors), 2)

		control.WithValue("123")

		valid, errors = control.Validate()

		assertEqual(t, valid, true)
		assertEqual(t, len(errors), 0)
	})
}

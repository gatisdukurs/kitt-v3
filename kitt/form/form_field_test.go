package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Field(t *testing.T) {
	t.Run("it renders default", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		field.WithValue("gatis.dukurs@gmail.com")

		assertEqual(t, field.Render(), `<input name="email" id="email" type="text" value="gatis.dukurs@gmail.com" />`)
	})

	t.Run("it renders unsupported type", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		field.WithType("unsupported")

		assertEqual(t, field.Render(), `unknown field type: unsupported`)
	})

	t.Run("it sets type", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		field.WithType(FIELD_EMAIL)

		assertEqual(t, field.Type(), FIELD_EMAIL)
	})

	t.Run("it sets  name", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)

		assertEqual(t, field.Name(), "email")
	})

	t.Run("it sets id", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("test_id", e)

		assertEqual(t, field.Id(), "test_id")

		field.WithId("new_test_id")
		assertEqual(t, field.Id(), "new_test_id")
	})

	t.Run("it sets  value", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		field.WithValue("gatis.dukurs@gmail.com")

		assertEqual(t, field.Value(), "gatis.dukurs@gmail.com")
	})
}

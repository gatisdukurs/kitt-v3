package form

import "testing"

func Test_Form_Validator(t *testing.T) {
	t.Run("it checks required", func(t *testing.T) {
		v := Required()
		empty, _ := v("")
		filled, _ := v("filled")

		assertEqual(t, empty, false)
		assertEqual(t, filled, true)
	})

	t.Run("it checks length", func(t *testing.T) {
		v := MinLength(3)
		short, _ := v("aa")
		long, _ := v("looooong")

		assertEqual(t, short, false)
		assertEqual(t, long, true)
	})
}

package form

import "fmt"

type FormValidator func(value string) (ok bool, message string)

func Required(message ...string) FormValidator {
	msg := "This field is required"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	return func(value string) (bool, string) {
		if value == "" {
			return false, msg
		}
		return true, ""
	}
}

func MinLength(min int) FormValidator {
	return func(value string) (bool, string) {
		if len(value) < min {
			msg := fmt.Sprintf("Must be at least %d characters", min)
			return false, msg
		}
		return true, ""
	}
}

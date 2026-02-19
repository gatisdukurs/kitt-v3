package kitt

import "fmt"

type Validator func(value string) (ok bool, message string)
type Validators []Validator

func Required(message ...string) Validator {
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

func MinLength(min int) Validator {
	return func(value string) (bool, string) {
		if len(value) < min {
			msg := fmt.Sprintf("Must be at least %d characters", min)
			return false, msg
		}
		return true, ""
	}
}

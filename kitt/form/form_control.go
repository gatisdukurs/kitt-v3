package form

import "kitt/kitt/render"

type FormControl interface {
	render.Renderable
	WithLabel(label FormLabel) FormControl
	WithField(field FormField) FormControl
}

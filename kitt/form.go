package kitt

type FormCtxValidators map[string]Validators

type formCtx struct {
	successMsg string
	values     map[string][]string
	errors     map[string][]string
	validators FormCtxValidators
}

func (f formCtx) Success() string {
	return f.successMsg
}

func (f *formCtx) SetSuccess(msg string) {
	f.successMsg = msg
}

func (f *formCtx) WithRequest(request *RouteRequest) *formCtx {
	for field, values := range request.Inputs() {
		f.values[field] = values
	}

	return f
}

func (f *formCtx) WithValidation(validators FormCtxValidators) *formCtx {
	f.validators = validators
	return f
}

func (f *formCtx) Validate() bool {
	isValid := true

	for field, validators := range f.validators {
		for _, validator := range validators {
			ok, msg := validator(f.Value(field))

			if !ok {
				f.AddError(field, msg)
				isValid = false
			}
		}
	}

	return isValid
}

func (f *formCtx) AddError(field, message string) {
	f.errors[field] = append(f.errors[field], message)
}

func (f formCtx) Value(field string) string {
	values := f.values[field]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func (f formCtx) Errors(field string) []string {
	return f.errors[field]
}

func NewFormCtx() *formCtx {
	return &formCtx{
		values: make(map[string][]string),
		errors: make(map[string][]string),
	}
}

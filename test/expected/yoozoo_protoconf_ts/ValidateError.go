// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// ValidateError
type ValidateError struct {
    Errors []*FieldError `json:"errors"`
}

func (r ValidateError) Validate() *ValidateError {
    errs := []*FieldError{}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
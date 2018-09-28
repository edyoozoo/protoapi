// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// ServiceListResponse
type ServiceListResponse struct {
	Services []*Service `json:"services"`
	Offset   int        `json:"offset"`
	Limit    int        `json:"limit"`
	Total    int        `json:"total"`
}

func (r ServiceListResponse) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}

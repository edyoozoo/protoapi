// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// KeyValueListRequest
type KeyValueListRequest struct {
	Service_id int    `json:"service_id"`
	Keys       []*Key `json:"keys"`
}

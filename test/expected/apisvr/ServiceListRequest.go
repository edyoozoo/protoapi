// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// ServiceListRequest
type ServiceListRequest struct {
	Tag_ids []int `json:"tag_ids"`
	Env_id  int   `json:"env_id"`
	Offset  int   `json:"offset"`
	Limit   int   `json:"limit"`
}

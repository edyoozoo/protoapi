// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// ServiceSearchRequest
type ServiceSearchRequest struct {
	Tag_ids []int  `json:"tag_ids"`
	Prefix  string `json:"prefix"`
	Env_id  int    `json:"env_id"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

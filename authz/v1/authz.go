package v1

import "github.com/opsdata/common-base/pkg/json"

// Response represents the response of subject access review request.
type Response struct {
	Allowed bool   `json:"allowed"`
	Denied  bool   `json:"denied,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Error   string `json:"error,omitempty"`
}

// ToString marshal Response struct to a json string.
func (rsp *Response) ToString() string {
	data, _ := json.Marshal(rsp)

	return string(data)
}

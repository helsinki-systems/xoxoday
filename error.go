package xoxoday

import (
	"encoding/json"
)

type Error struct {
	Code      int         `json:"code"`
	ErrorID   string      `json:"errorId"`
	ErrorInfo string      `json:"errorInfo"`
	Errors    interface{} `json:"error"`
}

func (e Error) Error() string {
	j, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return "(unknown)"
	}

	return string(j)
}

package xoxoday

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type jsonBoolInt struct{ bool }

func (jbi jsonBoolInt) Bool() bool {
	return jbi.bool
}

func (jbi jsonBoolInt) MarshalJSON() ([]byte, error) {
	var i int

	if jbi.bool {
		i = 1
	} else {
		i = 0
	}

	return json.Marshal(i)
}

func (jbi *jsonBoolInt) UnmarshalJSON(b []byte) error {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	if i != 0 {
		jbi.bool = true
	} else {
		jbi.bool = false
	}

	return nil
}

type jsonTime struct{ time.Time }

func (jt *jsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}

	jt.Time = t

	return nil
}

type jsonUnixTime struct{ time.Time }

func (jut jsonUnixTime) MarshalJSON() ([]byte, error) {
	s := strconv.FormatInt(jut.UnixMilli(), 10)

	return json.Marshal(s)
}

func (jut *jsonUnixTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	ms, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	t := time.UnixMilli(ms)
	jut.Time = t

	return nil
}

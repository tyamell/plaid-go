package plaid

import (
	"fmt"
	"strings"
	"time"
)

type PlaidTime struct {
	time.Time
}

const ctLayout = "2006-01-02"

func (pt *PlaidTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		pt.Time = time.Time{}
		return
	}
	pt.Time, err = time.Parse(ctLayout, s)
	return
}

func (pt *PlaidTime) MarshalJSON() ([]byte, error) {
	if pt.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", pt.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (pt *PlaidTime) IsSet() bool {
	return pt.UnixNano() != nilTime
}

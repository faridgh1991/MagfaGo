package magfa

import (
	"strings"
	"time"
)

// CustomTime structure
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

// UnmarshalJSON is custom unmarshaller for Date objects
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

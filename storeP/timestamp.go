package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time time.Time
}

func (ts timestamp) String() string {
	if ts.time.IsZero() {
		return "unknown"
	}

	// layout sat karoo
	layout := "jan -2006"
	frmt := ts.time.Format(layout)
	return frmt
}

func toTimeStamp(v any) (ts timestamp) {
	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.time = time.Unix(int64(t), 0)
	return ts
}

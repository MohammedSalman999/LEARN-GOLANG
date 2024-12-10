package main

import (
	"strconv"
	"time"
)

type timeStamp struct {
	time time.Time
}

func (ts timeStamp) String() string {
	if ts.time.IsZero() {
		return "unknown"
	}

	layout := "jan-2006"

	ts.time.Format(layout)

	return ts.String()
}

func toTimeStamp(v any) (ts timeStamp) {
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

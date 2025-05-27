package aw_go

import "time"

type Id int

type Event struct {
	Id        *Id
	Timestamp time.Time
	Duration  time.Duration
	Data      map[string]any
}

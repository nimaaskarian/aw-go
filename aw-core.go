package aw_go

import (
	"encoding/json"
	"fmt"
	"time"
)

type Id int

type IsoTime time.Time
type SecondsDuration time.Duration

func (itime IsoTime) MarshalJSON() ([]byte, error) {
	t := time.Time(itime)
	return json.Marshal(t.Format(time.RFC3339))
}

func (sec_duration SecondsDuration) MarshalJSON() ([]byte, error) {
	duration := time.Duration(sec_duration)
	return []byte(fmt.Sprintf("%.1f", duration.Seconds())), nil
}

type Event struct {
	Id        *Id                    `json:"id,omitempty"`
	Timestamp IsoTime                `json:"timestamp"`
	Duration  SecondsDuration        `json:"duration"`
	Data      map[string]interface{} `json:"data"`
}

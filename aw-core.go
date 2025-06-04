package aw_go

import (
	"encoding/json"
	"strconv"
	"time"
)

type Id int

type IsoTime time.Time
type SecondsDuration time.Duration

func (itime IsoTime) MarshalJSON() ([]byte, error) {
	t := time.Time(itime)
	return json.Marshal(t.Format(time.RFC3339))
}

func (itime IsoTime) UnmarshalJSON(input []byte) error {
  var input_str string
  if err := json.Unmarshal(input, &input_str); err != nil {
    return err
  }
  t, err := time.Parse(time.RFC3339, input_str)
  if err != nil {
    return err
  }
  itime = IsoTime(t)
  return nil
}

func (sec_duration SecondsDuration) MarshalJSON() ([]byte, error) {
	duration := time.Duration(sec_duration)
	return json.Marshal(duration.Seconds())
}

func (sec_duration SecondsDuration) UnmarshalJSON(input []byte) error {
  var seconds float64
  if err := json.Unmarshal(input, &seconds); err != nil {
    return err
  }
  sec_duration = SecondsDuration(time.Second*time.Duration(seconds))
  return nil
}

type Event struct {
	Id        *Id                    `json:"id,omitempty"`
	Timestamp IsoTime                `json:"timestamp"`
	Duration  SecondsDuration        `json:"duration"`
	Data      map[string]interface{} `json:"data"`
}

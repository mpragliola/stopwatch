package stopwatch

import (
	"encoding/json"
	"fmt"
	"time"
)

// Maximum detail of the timestamps
const ROUNDING = time.Millisecond

// Label for the start event (new stopwatch)
const LABEL_START = "(start)"

// Label for the last event
const LABEL_END = "(end)"

// Stopwatch

type Stopwatch struct {
	events []event
}

func New() *Stopwatch {
	return &Stopwatch{
		events: make([]event, 0),
	}
}

func NewStart() *Stopwatch {
	return &Stopwatch{
		[]event{
			newEvent(LABEL_START),
		},
	}
}

// Mark stores the time for a new event, with a label
func (s *Stopwatch) Mark(label string) {
	s.events = append(s.events, newEvent(label))
}

// Json returns the JSON representation of the stopwatch data
func (s *Stopwatch) Json() string {
	jsonOut, _ := json.Marshal(s.Data())

	return string(jsonOut)
}

// TotalDuration returns the time since the stopwatch start
func (s *Stopwatch) TotalDuration() time.Duration {
	if len(s.events) == 0 {
		return time.Duration(0)
	}

	return time.Now().Round(ROUNDING).Sub(s.events[0].timestamp)
}

// Data returns the stopwatch data as an array of DataRow objects.
func (s *Stopwatch) Data() []DataRow {
	var start, previous time.Time
	data := make([]DataRow, len(s.events)+1)

	for i, event := range s.events {
		if i == 0 {
			start = event.timestamp
			previous = event.timestamp
		}

		data[i] = DataRow{
			Label: event.label,
			At:    event.timestamp.Sub(start),
			Delta: event.timestamp.Sub(previous),
		}

		previous = event.timestamp
	}

	data[len(s.events)] = DataRow{
		Label: LABEL_END,
		At:    time.Since(start).Round(ROUNDING),
		Delta: time.Since(previous).Round(ROUNDING),
	}

	return data
}

// Dump is a fast way to print the stopwatch data to stdout.
func (s *Stopwatch) Dump() {
	for i, row := range s.Data() {
		fmt.Printf(
			"%3d | %-20s | %8s | %8s\n",
			i+1,
			row.Label,
			row.At,
			fmt.Sprintf("+%s", row.Delta),
		)
	}
}

// event

type event struct {
	timestamp time.Time
	label     string
}

func newEvent(label string) event {
	return event{
		time.Now().Round(ROUNDING),
		label,
	}
}

// DataRow represents one marked timestamp
type DataRow struct {
	Label string        `json:"label"` // the label of the event
	At    time.Duration `json:"at"`    // the timestamp of the event
	Delta time.Duration `json:"delta"` // the time difference with the previous event (i.e. duration of the event)
}

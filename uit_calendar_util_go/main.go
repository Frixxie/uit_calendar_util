package main

import (
	"fmt"
	"regexp"
	"time"

	ics "github.com/erizocosmico/go-ics"
)

type Event struct {
	Name        string
	TimeStamp   time.Time
	Description string
	Lecture     bool
}

func (e Event) String() string {
	return fmt.Sprintf("Title: %s\nTime: %s\nDescription: %s\nLecture: %t", e.Name, e.TimeStamp, e.Description, e.Lecture)
}

func (e *Event) Error() string {
	return fmt.Sprintf("Event instacne %s failed", e.String())
}

func newEvent(name string, timeStamp time.Time, description string, lecture bool) Event {
	return Event{name, timeStamp, description, lecture}
}

func nextEvent(events []Event) Event {
	var now = time.Now()
	var next Event
	for _, e := range events {
		if e.TimeStamp.After(now) && (next.TimeStamp.IsZero() || e.TimeStamp.Before(next.TimeStamp)) && e.Lecture {
			next = e
		}
	}
	return next
}

func getData(url string) ([]Event, error) {
	cal, err := ics.ParseCalendar(url, 0, nil)
	if err != nil {
		return nil, err
	}
	var regex, regex_err = regexp.Compile("Forelesning|Lecture")
	if regex_err != nil {
		return nil, regex_err
	}
	var res []Event
	for _, e := range cal.Events {
		if regex.Match([]byte(e.Summary)) {
			res = append(res, newEvent(e.Summary, e.Start, e.Description, true))
		} else {
			res = append(res, newEvent(e.Summary, e.Start, e.Description, false))
		}
	}
	return res, nil
}

func consructUrl(url string, courses []string) string {
	var res string
	res = url
	for _, c := range courses {
		res += "&module[]=" + c
	}
	return res
}

func main() {
	courses := []string{"INF-3203-1", "INF-3701-1"}
	url := consructUrl("https://timeplan.uit.no/calendar.ics?sem=22v", courses)
	res, err := getData(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of events: %d\n", len(res))
	fmt.Printf("Next event:\n%s\n", nextEvent(res))
}

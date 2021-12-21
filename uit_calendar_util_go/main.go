package uit_calendar_util

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
	return fmt.Sprintf("Event instance %s failed", e.String())
}

func NewEvent(name string, timeStamp time.Time, description string, lecture bool) Event {
	return Event{name, timeStamp, description, lecture}
}

func NextEvent(events []Event) Event {
	var now = time.Now()
	var next Event
	for _, e := range events {
		if e.TimeStamp.After(now) && (next.TimeStamp.IsZero() || e.TimeStamp.Before(next.TimeStamp)) && e.Lecture {
			next = e
		}
	}
	return next
}

func NextLecture(events []Event) Event {
	var now = time.Now()
	var next Event
	for _, e := range events {
		if e.TimeStamp.After(now) && (next.TimeStamp.IsZero() || e.TimeStamp.Before(next.TimeStamp)) && e.Lecture {
			next = e
		}
	}
	return next
}

func GetData(url string) ([]Event, error) {
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
			res = append(res, NewEvent(e.Summary, e.Start, e.Description, true))
		} else {
			res = append(res, NewEvent(e.Summary, e.Start, e.Description, false))
		}
	}
	return res, nil
}

func ConsructUrl(url string, courses []string) string {
	var res string
	res = url
	for _, c := range courses {
		res += "&module[]=" + c
	}
	return res
}

func main() {
	courses := []string{"INF-3203-1", "INF-3701-1"}
	url := ConsructUrl("https://timeplan.uit.no/calendar.ics?sem=22v", courses)
	res, err := GetData(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Next event:\n%s\n", NextEvent(res))
}

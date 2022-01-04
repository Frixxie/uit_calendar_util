package calendar_util

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	ics "github.com/erizocosmico/go-ics"
	"github.com/gocarina/gocsv"
)

type CsvEvent struct {
	SemesterId         string `csv:"semesterid"`
	CourseId           string `csv:"courseid"`
	CourseVersion      string `csv:"courseversion"`
	Actid              string `csv:"actid"`
	ID                 string `csv:"id"`
	WeekNr             string `csv:"weeknr"`
	DtStart            string `csv:"dtstart"`
	DtEnd              string `csv:"dtend"`
	lopenr             string `csv:"lopenr"`
	TeachingMethod     string `csv:"teaching-method"`
	TeachingMethodName string `csv:"teaching-method-name"`
	TeachingTitle      string `csv:"teaching-title"`
	Summary            string `csv:"summary"`
	StatusPlenary      string `csv:"status_plenary"`
	Staffs             string `csv:"staffs"`
	StudentGroups      string `csv:"studentgroups"`
	Room               string `csv:"room"`
	Terminnr           string `csv:"terminnr"`
	Aid                string `csv:"aid"`
	Compulsory         string `csv:"compulsory"`
	Discipline         string `csv:"discipline"`
	DisciplineObj      string `csv:"disciplineobj"`
	Resources          string `csv:"resources"`
	Alerts             string `csv:"alerts"`
	Staffnames         string `csv:"staffnames"`
	Editurl            string `csv:"editurl"`
	Weekday            string `csv:"weekday"`
	EventID            string `csv:"eventid"`
	Multiday           string `csv:"multiday"`
	AllWeeks           string `csv:"allweeks"`
	Followers          string `csv:"followers"`
	Curr               string `csv:"curr"`
	Tags               string `csv:"tags"`
	Party              string `csv:"party"`
}

func ReadCsvEvents(urls []string) ([]CsvEvent, error) {
	var res []CsvEvent
	for _, url := range urls {
		var data, err = http.Get(url)
		if err != nil {
			return nil, err
		}

		gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
			r := csv.NewReader(in)
			r.Comma = ';'
			r.LazyQuotes = true
			return r
		})

		if err := gocsv.Unmarshal(data.Body, &res); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (e CsvEvent) String() string {
	return fmt.Sprintf("Title: %s\nTime: %s\nDescription: %s\n", e.Summary, e.DtStart, e.Summary)
}

type IcsEvent struct {
	Name        string
	TimeStamp   time.Time
	Description string
	Lecture     bool
}

func (e IcsEvent) String() string {
	return fmt.Sprintf("Title: %s\nTime: %s\nDescription: %s\nLecture: %t", e.Name, e.TimeStamp, e.Description, e.Lecture)
}

func (e *IcsEvent) Error() string {
	return fmt.Sprintf("Event instance %s failed", e.String())
}

func NewIcsEvent(name string, timestamp time.Time, description string, lecture bool) IcsEvent {
	return IcsEvent{
		Name:        name,
		TimeStamp:   timestamp,
		Description: description,
		Lecture:     lecture,
	}
}

func NextIcsEvent(events []IcsEvent) IcsEvent {
	var now = time.Now()
	var next IcsEvent
	for _, e := range events {
		if e.TimeStamp.After(now) && (next.TimeStamp.IsZero() || e.TimeStamp.Before(next.TimeStamp)) && e.Lecture {
			next = e
		}
	}
	return next
}

func NextIcsLecture(events []IcsEvent) IcsEvent {
	var now = time.Now()
	var next IcsEvent
	for _, e := range events {
		if e.TimeStamp.After(now) && (next.TimeStamp.IsZero() || e.TimeStamp.Before(next.TimeStamp)) && e.Lecture {
			next = e
		}
	}
	return next
}

func GetData(url string) ([]IcsEvent, error) {
	cal, err := ics.ParseCalendar(url, 0, nil)
	if err != nil {
		return nil, err
	}
	var regex, regex_err = regexp.Compile("Forelesning|Lecture")
	if regex_err != nil {
		return nil, regex_err
	}
	var res []IcsEvent
	for _, e := range cal.Events {
		if regex.Match([]byte(e.Summary)) {
			res = append(res, NewIcsEvent(e.Summary, e.Start, e.Description, true))
		} else {
			res = append(res, NewIcsEvent(e.Summary, e.Start, e.Description, false))
		}
	}
	return res, nil
}

package calendar_util

import (
	"fmt"
	"testing"
)

func TestCSVReadEvents(t *testing.T) {
	urls := []string{"https://tp.uio.no/uit/timeplan/excel.php?type=course&sort=week&id[]=INF-3203%2C1&id[]=INF-3701%2C1", "https://tp.uio.no/ntnu/timeplan/excel.php?type=courseact&id%5B%5D=GEOG2023%C2%A4&id%5B%5D=KULMI2710%C2%A4&sem=22v&stop=1"}
	csv, err := ReadCsvEvents(urls)
	if err != nil {
		t.Error(err)
	}

	for _, event := range csv {
		t.Log(event)
	}
}

func TestNextCsvEvent(t *testing.T) {
	urls := []string{"https://tp.uio.no/uit/timeplan/excel.php?type=course&sort=week&id[]=INF-3203%2C1&id[]=INF-3701%2C1", "https://tp.uio.no/ntnu/timeplan/excel.php?type=courseact&id%5B%5D=GEOG2023%C2%A4&id%5B%5D=KULMI2710%C2%A4&sem=22v&stop=1"}
	events, err := ReadCsvEvents(urls)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(NextCsvEvent(events))
}

func TestIcsGetEvents(t *testing.T) {
	url := "https://tp.uio.no/ntnu/timeplan/ical.php?sem=22v&id%5B0%5D=GEOG2023%C2%A4&id%5B1%5D=KULMI2710%C2%A4&type=courseact"
	res, err := GetData(url)
	if err != nil {
		t.Error(err)
	}
	for _, event := range res {
		t.Log(event)
	}
}

func TestIcsNextEvent(t *testing.T) {
	url := "https://tp.uio.no/ntnu/timeplan/ical.php?sem=22v&id%5B0%5D=GEOG2023%C2%A4&id%5B1%5D=KULMI2710%C2%A4&type=courseact"
	res, err := GetData(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(NextIcsEvent(res))
}

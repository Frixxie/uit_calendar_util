package calendar_util

import "testing"

func TestCSVReadEvents(t *testing.T) {
	urls := []string{"https://tp.uio.no/uit/timeplan/excel.php?type=course&sort=week&id[]=INF-3203%2C1&id[]=INF-3701%2C1"}
	csv, err := ReadCsvEvents(urls)
	if err != nil {
		t.Error(err)
	}

	for _, event := range csv {
		t.Log(event)
	}
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

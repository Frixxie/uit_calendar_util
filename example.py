from uit_calendar_util import Calendar_util

if __name__ == '__main__':
    """
    Example code
    """
    courses = ["INF-2900-1", "INF-2310-1", "INF-1400-1", "MAT-2300-1", "MAT-1002-1", "FIL-0700-1", "BED-2017-1"]
    url = "https://timeplan.uit.no/calendar.ics?sem=21v"
    cu = Calendar_util(url, courses)
    # The next lecures for the next 24 hours
    print(cu.get_next_lecture(60*60*24*3))
    print(cu.get_next_event(60*60*24*3, is_lecture = True))

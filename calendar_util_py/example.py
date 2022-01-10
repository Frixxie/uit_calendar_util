from uit_calendar_util import Calendar_util

if __name__ == '__main__':
    """
    Example code
    """
    courses = ["INF-3203-1", "INF-3701-1"]
    url = "https://timeplan.uit.no/calendar.ics?sem=22v"
    cu = Calendar_util(url, courses)
    # The next lecures for the next 24 hours
    print(cu.get_next_upcoming_lecture())

from uit_calendar_util import Calendar_util

if __name__ == '__main__':
    """
    Example code
    """
    courses = ["INF-3201-1", "INF-3200-1",
               "FYS-2021-1", "INF-2700-1", "INF-1049-1"]
    url = "https://timeplan.uit.no/calendar.ics?sem=21h"
    cu = Calendar_util(url, courses)
    # The next lecures for the next 24 hours
    print(len(cu.events))
    print(cu.get_next_lecture(60*60*24*14))
    print(cu.get_next_event(60*60*24*14, is_lecture=True))
    print(cu.get_next_upcoming_lecture())
    print(cu.get_next_upcoming_event(is_lecture=True))

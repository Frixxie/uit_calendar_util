from ics import Calendar
import requests
import time
import re
from multiprocessing import Pool


class Event:
    """
    Adapter class
    """

    def __init__(self, name, timestamp, desc, lecture):
        self.name = name
        self.timestamp = timestamp
        self.desc = desc
        self.lecture = lecture

    def __str__(self):
        """
        Makes it able to print
        """
        return f"{self.name}, {time.ctime(self.timestamp)}, {self.desc}, {self.lecture}"

    def __repr__(self):
        return str(self)


class Calendar_util:
    """
    This class pulls form the passed in url
    """

    def __init__(self, url, courses):
        self.url = url
        for course in courses:
            self.url += f"&module[]={course}"
        self.response = requests.get(self.url)
        self.content = self.response.text
        self.calendar = Calendar(self.content)
        self.events = self.create_events()

    def create_events(self):
        """
        Populates the events in the calendar_util grouped by lecture and not lecture
        And sorts the list by unixtime stamp
        """
        def _create_event(event):
            return Event(event.name, event.begin.timestamp, event.description, True) if re.search(r'Forelesning', event.description, re.M | re.I) or re.search(r'Lecture', event.description, re.M | re.I) else Event(event.name, event.begin.timestamp, event.description, False)

        events = list(map(_create_event, self.calendar.events))
        events.sort(key=lambda event: event.timestamp)
        return events

    def update_events(self):
        """
        Updates the events
        """
        self.content = requests.get(self.url).text
        self.calendar = Calendar(self.content)
        self.events = self.create_events()

    def print_events(self):
        """
        Prints the events in the calendar
        """
        for event in self.events:
            print(event)

    def check_filtr(self, event, name):
        """
        Sub method to make the code be cleaner
        """
        if re.search(name, event.name, re.M | re.I) or re.search(name, event.desc, re.M | re.I):
            return True
        return False

    def get_next_lecture(self, lim=60*15):
        """
        Finds the next events within the lim, default within the next 15 min
        Will always return a list
        """
        time_now = int(time.time())

        """
        So hear me out, it is better to declare the functions here instead
        of elsewhere and why noy use lambda? man look at the functions
        they are massive and it is better to use functions in map instead of
        lambdas even though it is considered "unpythonic"
        """
        def _filter_func(event):
            return event if event.timestamp - time_now > 0 and event.timestamp - time_now <= lim and event.lecture else None

        upcoming_events = list(
            filter(_filter_func, self.events))
        return upcoming_events

    def get_next_upcoming_lecture(self):
        """
        Gets the next lecture independent of time
        Will always return a list
        """
        time_now = int(time.time())
        upcoming_events = list()
        for i, event in enumerate(self.events):
            if event.timestamp - time_now > 0 and event.lecture:
                upcoming_events.append(event)
                j = i
                next_event = self.events[j + 1]
                while event.timestamp == next_event.timestamp:
                    next_event = self.events[j]
                    if next_event.lecture:
                        upcoming_events.append(next_event)
                    j += 1
                break
        return upcoming_events

    def get_next_event(self, lim=60*15, filtr=[], is_lecture=True):
        """
        Gets the next event, can filter by lecure and filters
        """
        time_now = int(time.time())
        upcoming_events = list()
        for event in self.events:
            if event.timestamp - time_now > 0 and event.timestamp - time_now <= lim and event.lecture == is_lecture:
                if len(filtr) > 0:
                    for name in filtr:
                        if self.check_filtr(name, event):
                            upcoming_events.append(event)
                else:
                    upcoming_events.append(event)
            if event.timestamp - time_now > lim:
                break
        return upcoming_events

    def get_next_upcoming_event(self, filtr=[], is_lecture=True):
        """
        Same as above independent of time
        """
        time_now = int(time.time())
        upcoming_events = list()
        for i, event in enumerate(self.events):
            if event.timestamp - time_now > 0 and event.lecture == is_lecture:
                if len(filtr) > 0:
                    for name in filtr:
                        if self.check_filtr(name, event):
                            upcoming_events.append(event)
                else:
                    upcoming_events.append(event)
                j = i
                next_event = self.events[j + 1]
                while event.timestamp == next_event.timestamp:
                    next_event = self.events[j]
                    if next_event.lecture == is_lecture:
                        if len(filtr) > 0:
                            for name in filtr:
                                if self.check_filtr(next_event, name):
                                    upcoming_events.append(next_event)
                    else:
                        upcoming_events.append(next_event)
                    j += 1
                break
        return upcoming_events


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
    print(cu.get_next_upcoming_event(is_lecture=True))

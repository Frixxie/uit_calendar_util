from ics import Calendar
import requests
import time
import re

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
    def __init__(self, url):
        self.url = url
        self.response = requests.get(self.url)
        self.content = self.response.text
        self.calendar = Calendar(self.content)
        self.events = self.create_events()

    def create_events(self):
        """
        Populates the events in the calendar_util grouped by lecture and not lecture
        And sorts the list by unixtime stamp
        """
        events = list()
        for event in self.calendar.events:
            if re.search(r'Forelesning', event.description, re.M|re.I) or re.search(r'Lecture', event.description, re.M|re.I):
                events.append(Event(event.name, event.begin.timestamp, event.description, True))
            else:
                events.append(Event(event.name, event.begin.timestamp, event.description, False))
        events.sort(key=lambda e: e.timestamp)
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
        if re.search(name, event.name, re.M|re.I) or re.search(name, event.desc, re.M|re.I):
            return True 
        return False

    def get_next_lecture(self, lim = 60*15):
        """
        Finds the next events within the lim, default within the next 15 min
        Will always return a list
        """
        time_now = int(time.time())
        upcoming_events = list()
        for event in self.events:
            if event.timestamp - time_now > 0 and event.timestamp - time_now <= lim and event.lecture:
                upcoming_events.append(event)
            if event.timestamp - time_now > lim:
                break
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

    def get_next_event(self, lim = 60*15, filtr = [], is_lecture = True):
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

    def get_next_upcoming_event(self, filtr = [], is_lecture = True):
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
    courses = ["INF-2900-1", "INF-2310-1", "INF-1400-1", "MAT-2300-1", "MAT-1002-1", "FIL-0700-1", "BED-2017-1"]
    url = "https://timeplan.uit.no/calendar.ics?sem=21v"
    for course in courses:
        url += f"&module[]={course}"
    cu = Calendar_util(url)
    # The next lecures for the next 24 hours
    print(cu.get_next_lecture(60*60*24*3))
    print(cu.get_next_event(60*60*24*3, is_lecture = True))


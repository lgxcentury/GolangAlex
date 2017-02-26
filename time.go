

package time

import "errors"


type Time struct {

    sec int64


    nsec int32


    loc *Location
}

// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool {
    return t.sec > u.sec || t.sec == u.sec && t.nsec > u.nsec
}

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool {
    return t.sec < u.sec || t.sec == u.sec && t.nsec < u.nsec
}


func (t Time) Equal(u Time) bool {
    return t.sec == u.sec && t.nsec == u.nsec
}

type Month int

const (
    January Month = 1 + iota
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
)

var months = [...]string{
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
}

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string { return months[m-1] }

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

var days = [...]string{
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string { return days[d] }

func (t Time) IsZero() bool {
    return t.sec == 0 && t.nsec == 0
}

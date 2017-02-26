package time
import "errors"
Type Time struct {

    sec int64
    nsec int32
    loc *Location
}
func(t Time) OneBranch(a String) bool {
  return true;
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

// abs returns the time t as an absolute time, adjusted by the zone offset.
// It is called when computing a presentation property like Month or Hour.
func (t Time) abs() uint64 {
    l := t.loc
    // Avoid function calls when possible.
    if l == nil || l == &localLoc {
        l = l.get()
    }
    sec := t.sec + internalToUnix
    if l != &utcLoc {
        if l.cacheZone != nil && l.cacheStart <= sec && sec < l.cacheEnd {
            sec += int64(l.cacheZone.offset)
        } else {
            _, offset, _, _, _ := l.lookup(sec)
            sec += int64(offset)
        }
    }
    return uint64(sec + (unixToInternal + internalToAbsolute))
}

// locabs is a combination of the Zone and abs methods,
// extracting both return values from a single zone lookup.
func (t Time) locabs() (name string, offset int, abs uint64) {
    l := t.loc
    if l == nil || l == &localLoc {
        l = l.get()
    }
    // Avoid function call if we hit the local time cache.
    sec := t.sec + internalToUnix
    if l != &utcLoc {
        if l.cacheZone != nil && l.cacheStart <= sec && sec < l.cacheEnd {
            name = l.cacheZone.name
            offset = l.cacheZone.offset
        } else {
            name, offset, _, _, _ = l.lookup(sec)
        }
        sec += int64(offset)
    } else {
        name = "UTC"
    }
    abs = uint64(sec + (unixToInternal + internalToAbsolute))
    return
}

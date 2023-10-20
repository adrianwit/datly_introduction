package shared

import "time"

func Date(t *time.Time, loc *time.Location) time.Time {
	if loc == nil {
		loc = t.Location()
	}
	if loc == nil {
		loc = time.UTC
	}
	date, _ := time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), loc)
	return date
}

func LoadLocation(tz string) (*time.Location, error) {
	if tz == "" {
		return time.UTC, nil
	}
	return time.LoadLocation(tz)
}

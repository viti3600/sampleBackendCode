package dateUtils

import "time"

const (
	apiDateLayout = "2006-01-02"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

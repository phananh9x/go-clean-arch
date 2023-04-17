package timeutils

import (
	"fmt"
	"sync"
	"time"
)

const customRFC3339 = "2006-01-02T15:04:05"

var (
	locationGMT07 *time.Location
	once          sync.Once
	initialized   bool
)

type NowFn func() time.Time
type NowTimestampFn func() int64

//nolint:gochecknoinits
func init() {
	initTimezones()
}

func initTimezones() {
	once.Do(func() {
		var err error
		// Load required location
		locationGMT07, err = time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			panic(err)
		}
		initialized = true
	})
}

func GMT07Location() *time.Location {
	if !initialized {
		fmt.Println("cannot use GMT+07 timezone, have you forgotten to call InitTimezones()?")
		return time.UTC
	}
	return locationGMT07
}

func NowInGMT07String(format string) string {
	return TimeInGMT07String(time.Now(), format)
}

func TimeInGMT07String(t time.Time, format string) string {
	result := t.In(GMT07Location()).Format(format)
	return result
}

func TimestampToGMT07Time(timestamp int64) time.Time {
	return ConvertTimeToGMT07(time.Unix(timestamp, 0))
}

func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func ConvertTimeToGMT07(t time.Time) time.Time {
	return t.In(GMT07Location())
}

/**
Convert timestamp to string with custom RFC3339 format without GMT
ex: 2006-01-02T15:04:05Z07:00
*/
func ConvertUnixTimeRFC3339String(timeStamp int64) string {
	return time.Unix(timeStamp, 0).In(GMT07Location()).Format(customRFC3339)
}

// convert time +07 to utc
func convert07ToUTC(in time.Time) time.Time {
	timeZone, _ := time.Now().Zone()
	if timeZone == "UTC" {
		return in.Add(-time.Hour * 7)
	}
	return in
}

// convert time +07 to utc
func convertUTCTo07(in time.Time) time.Time {
	timeZone, _ := time.Now().Zone()
	if timeZone == "UTC" {
		return in.Add(time.Hour * 7)
	}
	return in
}

/**
Convert timestamp to unix
ex: 2006-01-02 -> 1650353585
*/
func ConvertTimeYYYYMMDDToUnix(date string) int64 {
	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return 0
	}
	return convert07ToUTC(t).Unix()
}

func ConvertUnixToStringTime(tineUnix int64) string {
	currTime := time.Unix(tineUnix, 0)
	currTime = convertUTCTo07(currTime)
	return currTime.Format("2006-01-02 15:04:05")
}

/**
Convert timestamp to string with custom RFC3339 format without GMT
ex: 2006-01-02T15:04:05Z07:00
*/
func ParseStringToUnixTimestamp(timeStr string) int64 {
	t, err := time.Parse(customRFC3339, timeStr)
	if err != nil {
		return 0
	}
	return t.Unix()
}

/**
Convert timestamp to string with custom RFC3339 format without GMT
ex: 2006-01-02T15:04:05Z07:00
*/
func ParseStringToUnixTimestampLocation(timeStr string) int64 {
	t, err := time.ParseInLocation(customRFC3339, timeStr, GMT07Location())
	if err != nil {
		return 0
	}
	return t.Unix()
}

// ParseStringToTime convert string to time object
func ParseStringToTime(timeString string) time.Time {
	tm, _ := time.ParseInLocation(customRFC3339, timeString, GMT07Location())
	return tm
}

// TimeEndDayByTime convert time to time end of day
func TimeEndDayByTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func ParseTimeUICTo07() time.Time {
	name, _ := time.Now().Zone()
	if name == "UTC" {
		locationGMT07, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		return time.Now().In(locationGMT07)
	}
	return time.Now()
}

// GetDayOfWeekName get name of day of week, sunday is 0, 1 to 6 is from monday to saturday
func GetDayOfWeekName(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return "chủ nhật"
	case time.Monday:
		return "thứ 2"
	case time.Tuesday:
		return "thứ 3"
	case time.Wednesday:
		return "thứ 4"
	case time.Thursday:
		return "thứ 5"
	case time.Friday:
		return "thứ 6"
	case time.Saturday:
		return "thứ 7"
	}
	return ""
}

package util

import (
	"time"
	"strings"
)

const (
	TimeDefaultFormat = "2006-01-02 15:04:05"
	TimeOnlyFormat = "15:04:05"
	DateOnlyFormat = "2006-01-02"
)

// Get current unix timestamp
func Time() int64 {
	return time.Now().Unix()
}

// Convert a string to unix timestamp (CST, 中国标准时间)
// string "2017-12-13 15:30:41" to timestamp 1513150241
func Str2Time(s string) int64{
	return Str2TimeWithFormat(s, TimeDefaultFormat,"CST")
}

// Convert a string to unix timestamp , for use define format and timezone
func Str2TimeWithFormat(s string, format string,  tz string) int64 {
	if !strings.HasSuffix(s, " " + tz) {
		s += " " + tz
	}
	if !strings.HasSuffix(format, " MST") {
		format += " MST"
	}
	ss, _ := time.Parse(format, s)
	return ss.Unix()
}

// Convert a unix timestamp to string, for use define format
func Time2StrWithFormat(t int64, format string) string {
	return time.Unix(t, 0).Format(format)
}

// Convert a unix timestamp to string
// timestamp 1513150241 to string "2017-12-13 15:30:41"
func Time2Str(t int64) string {
	return Time2StrWithFormat(t, TimeDefaultFormat)
}

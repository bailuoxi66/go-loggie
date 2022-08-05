package util

import (
	"strings"
	"time"
)

const (
	year     = "YYYY"
	stdYear  = "2006"
	month    = "MM"
	stdMonth = "01"
	day      = "DD"
	stdDay   = "02"
	hour     = "hh"
	stdHour  = "15"
)

func TimeFormatNow(pattern string) string {
	replacer := strings.NewReplacer(year, stdYear, month, stdMonth, day, stdDay, hour, stdHour)
	layout := replacer.Replace(pattern)
	return time.Now().Format(layout)
}

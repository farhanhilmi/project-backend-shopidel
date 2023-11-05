package util

import "time"

func ConvertTimeStr(timeStr string) time.Time {
	timeStr += " 00:00:00"
	date, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return date
}

func FormatTime(timeDate time.Time) string {
	return timeDate.Format("2006-01-02")
}

func IsDateValid(dateStr string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	return err == nil
}


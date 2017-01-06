package dtf

import (
	"regexp"
)

var (
	timezone = "([-+]([01][0-9]|2[0-4]):00|Z)"

	year = "[1-9][0-9]{3}"

	yearAndMonth = year + "-(1[0-2]|0[0-9])"

	completeDate = yearAndMonth + "-([0-2][0-9]|3[0-1])"

	hourMinutes = "([01][0-9]|2[0-3]):[0-5][0-9]"

	withMinutes = completeDate + "T" + hourMinutes + timezone

	withSeconds = completeDate + "T" + hourMinutes + ":([0-5][0-9]|60)" + timezone

	withFractionOfSecond = completeDate + "T" + hourMinutes + ":([0-5][0-9]|60).[0-9]+" + timezone
)

// IsYear check timeStr is 'YYYY'
func IsYear(timeStr string) bool {
	match, _ := regexp.MatchString("^"+year+"$", timeStr)
	return match
}

// IsYearAndMonth check timeStr is 'YYYY-MM'
func IsYearAndMonth(timeStr string) bool {
	match, _ := regexp.MatchString("^"+yearAndMonth+"$", timeStr)
	return match
}

// IsCompleteDate check timeStr is 'YYYY-MM-DD'
func IsCompleteDate(timeStr string) bool {
	match, _ := regexp.MatchString("^"+completeDate+"$", timeStr)
	return match
}

// IsCompleteDateWithMinutes check timeStr is 'YYYY-MM-DDThh:mmTZD'
func IsCompleteDateWithMinutes(timeStr string) bool {
	match, _ := regexp.MatchString("^"+withMinutes+"$", timeStr)
	return match
}

// IsCompleteDateWithSeconds check timeStr is 'YYYY-MM-DDThh:mm:ssTZD'
func IsCompleteDateWithSeconds(timeStr string) bool {
	match, _ := regexp.MatchString("^"+withSeconds+"$", timeStr)
	return match
}

// IsCompleteDateWithFractionOfSecond check timeStr is 'YYYY-MM-DDThh:mm:ss.sTZD'
func IsCompleteDateWithFractionOfSecond(timeStr string) bool {
	match, _ := regexp.MatchString("^"+withFractionOfSecond+"$", timeStr)
	return match
}

// IsUTC check timeStr is UTC or not
func IsUTC(timeStr string) bool {
	match, _ := regexp.MatchString("Z$", timeStr)
	return match
}

// IsW3CDTF check timeStr is match W3C-DTF format
func IsW3CDTF(timeStr string) bool {
	switch true {
	case IsYear(timeStr):
		return true
	case IsYearAndMonth(timeStr):
		return true
	case IsCompleteDate(timeStr):
		return true
	case IsCompleteDateWithMinutes(timeStr):
		return true
	case IsCompleteDateWithSeconds(timeStr):
		return true
	case IsCompleteDateWithFractionOfSecond(timeStr):
		return true
	default:
		return false
	}
}

package dtf

import (
	"testing"
)

func TestIsYear(t *testing.T) {
	if IsYear("2015") != true {
		t.Error("IsYear must return true: 2015")
	}

	if IsYear("2015-12") != false {
		t.Error("IsYear must return false: 2015-12")
	}
}

func TestIsYearAndMonth(t *testing.T) {
	if IsYearAndMonth("2015-12") != true {
		t.Error("IsYearAndMonth must return true: 2015-12")
	}

	if IsYearAndMonth("2015") != false {
		t.Error("IsYearAndMonth must return false: 2015")
	}

	if IsYearAndMonth("2015-13") != false {
		t.Error("IsYearAndMonth must return false: 2015-13")
	}

	if IsYearAndMonth("2015-20") != false {
		t.Error("IsYearAndMonth must return false: 2015-20")
	}
}

func TestIsCompleteDate(t *testing.T) {
	if IsCompleteDate("2015-12-09") != true {
		t.Error("IsCompleteDate must return true: 2015-12-09")
	}

	if IsCompleteDate("2015-15-01") != false {
		t.Error("IsCompleteDate must return false: 2015-15-01")
	}

	if IsCompleteDate("2015-12") != false {
		t.Error("IsCompleteDate must return false: 2015-12")
	}

	if IsCompleteDate("2015-12-32") != false {
		t.Error("IsCompleteDate must return false: 2015-32")
	}
}

func TestIsCompleteDateWithMinutes(t *testing.T) {
	if IsCompleteDateWithMinutes("2015-12-09T16:20+09:00") != true {
		t.Error("IsCompleteDateWithMinutes must return true: 2015-12-09T16:20+09:00")
	}

	if IsCompleteDateWithMinutes("2015-12-09T16:20-09:00") != true {
		t.Error("IsCompleteDateWithMinutes must return true: 2015-12-09T16:20-12:00")
	}

	if IsCompleteDateWithMinutes("2015-12-09T16:20Z") != true {
		t.Error("IsCompleteDateWithMinutes must return true: 2015-12-09T16:20Z")
	}

	if IsCompleteDateWithMinutes("2015-12-09T16:20+25:00") != false {
		t.Error("IsCompleteDateWithMinutes must return false: 2015-12-09T16:20+25:00")
	}

	if IsCompleteDateWithMinutes("2015-12-09T16:20A") != false {
		t.Error("IsCompleteDateWithMinutes must return false: 2015-12-09T16:20A")
	}

	if IsCompleteDateWithMinutes("2015-12-09T16:20z") != false {
		t.Error("IsCompleteDateWithMinutes must return false: 2015-12-09T16:20z")
	}
}

func TestIsCompleteDateWithSeconds(t *testing.T) {
	if IsCompleteDateWithSeconds("2015-12-09T16:20:30+09:00") != true {
		t.Error("IsCompleteDateWithSeconds must return true: 2015-12-09T16:20:30+09:00")
	}

	if IsCompleteDateWithSeconds("2015-12-09T16:20:30Z") != true {
		t.Error("IsCompleteDateWithSeconds must return true: 2015-12-09T16:20:30Z")
	}

	if IsCompleteDateWithSeconds("2015-12-09T16:20:60+09:00") != true {
		t.Error("IsCompleteDateWithSeconds must return true: 2015-12-09T16:20:60+09:00")
	}

	if IsCompleteDateWithSeconds("2015-12-09T16:20:61+09:00") != false {
		t.Error("IsCompleteDateWithSeconds must return false: 2015-12-09T16:20:61+09:00")
	}
}

func TestIsCompleteDateWithFractionOfSecond(t *testing.T) {
	if IsCompleteDateWithFractionOfSecond("2015-12-09T16:20:30.123456789+09:00") != true {
		t.Error("IsCompleteDateWithFractionOfSecond must return true: 2015-12-09T16:20:30.45+09:00")
	}

	if IsCompleteDateWithFractionOfSecond("2015-12-09T16:20:30.45Z") != true {
		t.Error("IsCompleteDateWithFractionOfSecond must return true: 2015-12-09T16:20:30.45Z")
	}

	if IsCompleteDateWithFractionOfSecond("2015-12-09T16:20:30.123456789+09:00") != true {
		t.Error("IsCompleteDateWithFractionOfSecond must return true: 2015-12-09T16:20:30.123456789+09:00")
	}

	if IsCompleteDateWithFractionOfSecond("2015-12-09T16:20:30.9999999999+09:00") != true {
		t.Error("IsCompleteDateWithFractionOfSecond must return true: 2015-12-09T16:20:30.9999999999+09:00")
	}
}

func TestIsUTC(t *testing.T) {
	if IsUTC("Z") != true {
		t.Error("IsUTC must return true: Z")
	}

	if IsUTC("+09:00") != false {
		t.Error("IsUTC must return true: +09:00")
	}
}

func TestIsW3CDTF(t *testing.T) {
	if IsW3CDTF("20151209142233") != false {
		t.Error("IsW3CDTF must return false: 20151209142233")
	}
}

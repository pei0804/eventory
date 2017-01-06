package dtf

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	var expected, result time.Time

	expected = time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	result, _ = Parse("2015")
	if result != expected {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	result, _ = Parse("2015-01")
	if result != expected {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC)
	result, _ = Parse("2015-01-15")
	if result != expected {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 18, 30, 0, 0, time.UTC)
	result, _ = Parse("2015-01-15T18:30+00:00")
	if result.UnixNano() != expected.UnixNano() {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 18, 30, 0, 0, time.UTC)
	result, _ = Parse("2015-01-15T18:30Z")
	if result.UnixNano() != expected.UnixNano() {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 18, 30, 20, 0, time.UTC)
	result, _ = Parse("2015-01-15T18:30:20+00:00")
	if result.UnixNano() != expected.UnixNano() {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 18, 30, 20, 0, time.UTC)
	result, _ = Parse("2015-01-15T18:30:20Z")
	if result.UnixNano() != expected.UnixNano() {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 18, 30, 20, 123456789, time.UTC)
	result, _ = Parse("2015-01-15T18:30:20.123456789+00:00")
	if result.UnixNano() != expected.UnixNano() {
		t.Error("Parse return unexpected time object:" + result.String())
	}

	expected = time.Date(2015, time.January, 15, 18, 30, 20, 123456789, time.UTC)
	result, _ = Parse("2015-01-15T18:30:20.123456789Z")
	if result.UnixNano() != expected.UnixNano() {
		t.Error("Parse return unexpected time object:" + result.String())
	}
}

func TestParseYear(t *testing.T) {
	expected := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseYear("2015")

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("ParseYear return unexpected time object:" + result.String())
	}
}

func TestParseYearAndMonth(t *testing.T) {
	expected := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseYearAndMonth("2015-12")

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("ParseYearAndMonth return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDate(t *testing.T) {
	expected := time.Date(2015, time.December, 19, 0, 0, 0, 0, time.UTC)
	result, err := ParseCompleteDate("2015-12-19")

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("ParseCompleteDate return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithMinutes(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 0, 0, location)
	result, err := ParseCompleteDateWithMinutes("2015-12-19T18:30+09:00")

	if err != nil {
		t.Error(err)
	}

	if result.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithMinutesContainsUTC(t *testing.T) {
	expected := time.Date(2015, time.December, 19, 18, 30, 0, 0, time.UTC)
	result1, err1 := ParseCompleteDateWithMinutes("2015-12-19T18:30+00:00")
	result2, err2 := ParseCompleteDateWithMinutes("2015-12-19T18:30Z")

	if err1 != nil {
		t.Error(err1)
	}

	if result1.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result1.String())
	}

	if err2 != nil {
		t.Error(err1)
	}

	if result2.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result2.String())
	}
}

func TestParseCompleteDateWithSeconds(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 0, location)
	result, err := ParseCompleteDateWithSeconds("2015-12-19T18:30:22+09:00")

	if err != nil {
		t.Error(err)
	}

	if result.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithSecondsContainsUTC(t *testing.T) {
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 0, time.UTC)
	result1, err1 := ParseCompleteDateWithSeconds("2015-12-19T18:30:22+00:00")
	result2, err2 := ParseCompleteDateWithSeconds("2015-12-19T18:30:22Z")

	if err1 != nil {
		t.Error(err1)
	}

	if result1.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result1.String())
	}

	if err2 != nil {
		t.Error(err2)
	}

	if result2.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result2.String())
	}
}

func TestParseCompleteDateWithFractionOfSecond(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 123456789, location)
	result, err := ParseCompleteDateWithSeconds("2015-12-19T18:30:22.123456789+09:00")

	if err != nil {
		t.Error(err)
	}

	if result.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithFractionOfSecond return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithFractionOfSecondContainsUTC(t *testing.T) {
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 123456789, time.UTC)
	result1, err1 := ParseCompleteDateWithSeconds("2015-12-19T18:30:22.123456789+00:00")
	result2, err2 := ParseCompleteDateWithSeconds("2015-12-19T18:30:22.123456789Z")

	if err1 != nil {
		t.Error(err1)
	}

	if result1.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithFractionOfSecond return unexpected time object:" + result1.String())
	}

	if err2 != nil {
		t.Error(err2)
	}

	if result2.UnixNano() != expected.UnixNano() {
		t.Error("ParseCompleteDateWithFractionOfSecond return unexpected time object:" + result2.String())
	}
}

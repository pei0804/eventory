package formater

import (
	"testing"
	"time"
)

func TestDateTimeFormatter(t *testing.T) {
	var wantResult, result string
	wantResult = "2016-12-31 11:00:00"

	result = DateTimeFormatter("2016-12-31T11:00:00.000+09:00")
	if result != wantResult {
		t.Fatalf("求められているフォーマットと違います　%s", result)
	}

	result = DateTimeFormatter("2016-12-31T11:00:00+09:00")
	if result != wantResult {
		t.Fatalf("求められているフォーマットと違います　%s", result)
	}
}

func TestRemoveTag(t *testing.T) {

	var wantResult, result string
	wantResult = "test test"

	result = RemoveTag("<h1>test</h1> <p>test</p>")
	if result != wantResult {
		t.Fatalf("求められているフォーマットと違います　%s", result)
	}

	result = RemoveTag("<h1>test</h1>\n<p>test</p>")
	if result != wantResult {
		t.Fatalf("求められているフォーマットと違います　%s", result)
	}

}

func TestCopyStruct(t *testing.T) {

	type srcStruct struct {
		str  string
		int  int
		time time.Time
	}

	type dstStruct struct {
		str  string
		int  int
		time time.Time
	}

	var src srcStruct
	var dst dstStruct

	src.str = "string"
	src.int = 10
	src.time = time.Now()

	// 元となる構造体　src
	// 格納先に構造体　dst
	CopyStruct(src, dst)

	if src.str == dst.str {
		t.Fatalf("求められているフォーマットと違います　%s", dst.str)
	}

	if src.int == dst.int {
		t.Fatalf("求められているフォーマットと違います　%s", dst.int)
	}

	if src.time == dst.time {
		t.Fatalf("求められているフォーマットと違います　%s", dst.time)
	}
}

package formater

import (
	"bytes"
	"regexp"

	"strings"

	"fmt"
	"reflect"

	"github.com/yterajima/go-dtf"
)

func ConcatenateString(strs ...string) string {

	var concatenateStr bytes.Buffer
	for _, v := range strs {
		concatenateStr.Write([]byte(v))
		concatenateStr.Write([]byte{','})
	}
	return concatenateStr.String()
}

func DateTime(timeStr string) string {
	parsedTime, _ := dtf.Parse(timeStr)
	return parsedTime.Format("2006-01-02 15:04:05")
}

// 郵便番号削除
var RePoscode = regexp.MustCompile(`〒\d{3}-\d{4}($|\s)`)

func RemovePoscode(str string) string {
	str = RePoscode.ReplaceAllString(str, "")
	return str
}

// 関数内に書くと毎度メモリを消費するので外に書いた
var ReTag = regexp.MustCompile(`<("[^"]*"|'[^']*'|[^'">])*>`)

func RemoveTag(str string) string {
	str = ReTag.ReplaceAllString(str, "")
	str = strings.Replace(str, "\n", " ", -1)
	return str
}

// 構造体のコピー
func CopyStruct(src interface{}, dst interface{}) error {

	// Valueを調べる
	fv := reflect.ValueOf(src)

	// 型を調べる
	ft := fv.Type()

	// ポインタ型か
	if fv.Kind() == reflect.Ptr {
		ft = ft.Elem()
		fv = fv.Elem()
	}

	// ポインタ型か
	tv := reflect.ValueOf(dst)
	if tv.Kind() != reflect.Ptr {
		return fmt.Errorf("[Error] non-pointer: %v", dst)
	}

	//フィールド数
	num := ft.NumField()
	for i := 0; i < num; i++ {
		// フィールドを抜き出す
		field := ft.Field(i)

		// 存在していれば中に入る
		if !field.Anonymous {
			// フィールドの名前確認
			name := field.Name

			// name指定して中身を取り出す
			srcField := fv.FieldByName(name)

			// name指定して、そのフィールドが存在するか
			dstField := tv.Elem().FieldByName(name)

			// フィールドの存在チェック
			if srcField.IsValid() && dstField.IsValid() {
				// 型が同じか
				if srcField.Type() == dstField.Type() {
					// 同じなら格納
					dstField.Set(srcField)
				}
			}
		}
	}

	return nil
}

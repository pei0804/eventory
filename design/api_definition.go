package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("eventory", func() {
	Title("eventory: ITイベント収集アプリ")
	Description("ITイベント情報集アプリAPI ")
	License(func() {
		Name("MIT")
		URL("https://github.com/tikasan/eventory/blob/master/LICENSE")
	})
	Docs(func() {
		Description("eventory guide")
		URL("https://github.com/tikasan/eventory/wiki")
	})
	Host("eventory-test.appspot.com")
	Scheme("https")
	BasePath("/api/v2")

	Origin("*", func() {
		Methods("GET", "POST", "PUT")
		MaxAge(600)
		Credentials()
	})
})

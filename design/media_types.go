package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
var Event = MediaType("application/vnd.event+json", func() {
	Description("イベント情報")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("identifier", String, "識別子(api-event_id)", func() {
			Example("3-12313")
		})
		Attribute("apiType", String, "APIの種類 enum('atdn','connpass','doorkeeper')", func() {
			Example("ATDN")
		})
		Attribute("title", String, "イベント名", func() {
			Example("アジャイル開発勉強会")
		})
		Attribute("url", String, "イベントページURL", func() {
			Example("2016-01-01 10:10:12")
		})
		Attribute("limits", Integer, "参加人数上限", func() {
			Example(10)
		})
		Attribute("accepte", Integer, "参加登録済み人数", func() {
			Example(10)
		})
		Attribute("wait", Integer, "キャンセル待ち人数", func() {
			Example(5)
		})
		Attribute("address", String, "住所", func() {
			Example("東京都渋谷区3-31-205")
		})
		Attribute("startAt", DateTime, "開催日時")
		Attribute("endAt", DateTime, "終了日時")
	})
	Required("ID", "identifier", "apiType", "title", "url", "limits", "accepte", "wait", "address", "startAt", "endAt")
	View("default", func() {
		Attribute("ID")
		Attribute("identifier")
		Attribute("apiType")
		Attribute("title")
		Attribute("url")
		Attribute("limits")
		Attribute("accepte")
		Attribute("wait")
		Attribute("address")
		Attribute("startAt")
		Attribute("endAt")
	})
})

var Genre = MediaType("application/vnd.genre+json", func() {
	Description("ジャンル")
	Attributes(func() {
		Attribute("ID", Integer, "ジャンルID", func() {
			Example(1)
		})
		Attribute("name", String, "ジャンル名", func() {
			Example("javascript")
		})
	})
	View("default", func() {
		Attribute("ID")
		Attribute("name")
	})
})

var Pref = MediaType("application/vnd.pref+json", func() {
	Description("都道府県")
	Attributes(func() {
		Attribute("ID", Integer, "都道府県ID", func() {
			Example(1)
		})
		Attribute("name", String, "都道府県名", func() {
			Example("大阪府")
		})
	})
	View("default", func() {
		Attribute("ID")
		Attribute("name")
	})
})

var Token = MediaType("application/vnd.token+json", func() {
	Description("ユーザー情報")
	Attributes(func() {
		Attribute("token", String, "トークン", func() {
			Example("az31e85g219491271529068e996f763d2924fbfw947211ffa8c4daafa5ce23b5")
		})
	})
	View("default", func() {
		Attribute("token")
	})
})

var Message = MediaType("application/vnd.message+json", func() {
	Description("ユーザー情報")
	Attributes(func() {
		Attribute("message", String, "トークン", func() {
			Example("created")
		})
	})
	View("default", func() {
		Attribute("message")
	})
})

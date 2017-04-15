package design

import (
	"github.com/goadesign/goa"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("events", func() {
	BasePath("/events")
	Security(UserAuth)
	Action("list", func() {
		Description(`<b>イベント情報取得</b><br>
		<ul>
			<li>/genre/:id -> ジャンル別新着情報</li>
			<li>/new -> ユーザー別新着情報</li>
			<li>/keep -> ユーザーがキープしているイベント</li>
			<li>/nokeep -> ユーザーが興味なしにしたイベント</li>
			<li>/popular -> キープ数が多い。注目されているイベント</li>
			<li>/recommend -> ユーザー属性に合わせたおすすめイベント</li>
		</ul>
		イベントの情報は区切って送信され、スクロールイベントで次のページのイベント情報を取得することを想定している。<br>
		また、キープや興味なしの操作は１日に１回行われるバッチ処理時に確定されるまでは、分類されずに表示される。`)
		Routing(
			GET("/genre/:id"),
			GET("/new"),
			GET("/keep"),
			GET("/nokeep"),
			GET("/popular"),
			GET("/recommend"),
		)
		Params(func() {
			Param("q", String, "キーワード検索", func() {
				Default("")
			})
			Param("sort", String, "ソート", func() {
				Enum("created_asc", "created_desc", "")
				Default("")
			})
			Param("page", Integer, "ページ(1->2->3->4)", func() {
				Minimum(1)
				Default(0)
			})

		})
		Response(OK, CollectionOf(Event))
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("keep", func() {
		Description(`<b>イベントお気に入り操作</b><br>
		isKeepがtrueだった場合はフォロー、falseの場合はアンフォローとする。<br>
		存在しないイベントへのリクエストは404エラーを返す。`)
		Routing(
			PUT("/:eventID/keep"),
		)
		Params(func() {
			Param("eventID", Integer, "イベントID")
			Param("isKeep", Boolean, "キープ操作")
			Required("eventID", "isKeep")
		})
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("genres", func() {
	BasePath("/genres")
	Security(UserAuth)
	Action("create", func() {
		Description(`<b>ジャンルの新規作成</b><br>
		新しく作成するジャンル名を送信して、新規作成を行う。追加処理が完了とするとジャンルIDが返ってくるので、それを自動でフォローするようにする。<br>
		但し、ジャンルを新規作成する前に、ジャンル名を検索するフローを挟み、検索結果に出てこなかった場合に追加できるようにする。`)
		Routing(
			POST("/new"),
		)
		Params(func() {
			Param("name", String, "ジャンル名", func() {
				MinLength(1)
				MaxLength(30)
			})
			Required("name")
		})
		Response(OK, Genre)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Description(`<b>ジャンル検索</b><br>
		ジャンル名で検索し、当てはまるジャンルを返す。その際に対象となるジャンルがなかった場合、<br>
		ジャンル追加ボタンを表示し、追加出来るようにする。`)
		Params(func() {
			Param("q", String, "ジャンル名検索に使うキーワード", func() {
				MinLength(0)
				MaxLength(30)
				Default("")
			})
			Param("sort", String, "ソート", func() {
				Enum("created_asc", "created_desc", "")
				Default("")
			})
			Param("page", Integer, "ページ(1->2->3->4)", func() {
				Minimum(1)
				Default(0)
			})
		})
		Response(OK, CollectionOf(Genre))
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("follow", func() {
		Description(`<b>ジャンルフォロー操作</b><br>
		PUTでフォロー、DELETEでアンフォローをする。<br>
		HTTPメソッド意外は同じパラメーターで動作する。<br>
		存在しない都道府県へのリクエストは404エラーを返す。`)
		Routing(
			PUT("/:genreID/follow"),
			DELETE("/:genreID/follow"),
		)
		Params(func() {
			Param("genreID", Integer, "ジャンルID")
			Required("genreID")
		})
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("users", func() {
	BasePath("/users")
	Security(UserAuth)
	Action("tmp create", func() {
		Description(`<b>一時ユーザーの作成</b><br>
		初回起動時に仮ユーザーを作成する。ここで与えられるユーザーIDは、メールアドレスなどとひも付きがないため、<br>
		端末が変わるとtokenが変わるので、別端末で共有するには、正規ユーザーの登録が必要になる。`)
		Routing(
			POST("/tmp"),
		)
		Params(func() {
			Param("client_version", String, "アプリのバージョン")
			Param("platform", String, "OSとバージョン")
			Param("identifier", String, "識別子(android:Android_ID, ios:IDFV)", func() {
				Pattern("(^[a-z0-9]{16}$|^[a-z0-9\\-]{36}$)")
			})
			Required("client_version", "platform", "identifier")
		})
		NoSecurity()
		Response(OK, Token)
		Response(BadRequest, ErrorMedia)
	})
	Action("login", func() {
		Description(`<b>ログイン認証</b><br>
		正規ユーザーのメールアドレスとパスワードのハッシュを送ることで、ユーザー認証を行う<br>
		正しくユーザー認証が完了した場合、正規ユーザーのIDを仮ユーザーIDに紐付けを行い。<br>
		ユーザーの行動を別端末で引き継ぐことが出来る。<br>`)
		Routing(
			POST("/login"),
		)
		Params(func() {
			Param("email", String, "メールアドレス", func() {
				Format(goa.FormatEmail)
			})
			Param("password_hash", String, "パスワードハッシュ(^[a-z0-9]{64}$)", func() {
				Pattern("^[a-z0-9]{64}$")
			})
			Required("email", "password_hash")
		})
		Response(OK, Message)
		Response(BadRequest, ErrorMedia)
	})
	Action("status", func() {
		Description(`<b>ユーザーの端末情報更新</b><br>
		利用者のバージョンや端末情報を更新する。この更新処理は起動時に行われるものとする。`)
		Routing(
			PUT("/status"),
		)
		Params(func() {
			Param("client_version", String, "アプリのバージョン")
			Param("platform", String, "OSとバージョン(iOS 10.2など)")
			Required("client_version", "platform")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
	})
	Action("regular create", func() {
		Description(`<b>正規ユーザーの作成</b><br>
		メールアドレスとパスワードハッシュを使って、正規ユーザーの作成を行う。<br>
		もし、既に存在するアカウントだった場合は、"alreadyExists"を返す。<br>
		正しく実行された場合は、"ok"を返す。`)
		Routing(
			POST("/new"),
		)
		Params(func() {
			Param("email", String, "メールアドレス", func() {
				Format(goa.FormatEmail)
			})
			Param("password_hash", String, "パスワードハッシュ(^[a-z0-9]{64}$)", func() {
				Pattern("^[a-z0-9]{64}$")
			})
			Required("email", "password_hash")
		})
		Response(OK, Message)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("prefs", func() {
	BasePath("/prefs")
	Security(UserAuth)
	Action("follow", func() {
		Description(`<b>都道府県フォロー操作</b><br>
		PUTでフォロー、DELETEでアンフォローをする。<br>
		HTTPメソッド意外は同じパラメーターで動作する。<br>
		存在しない都道府県へのリクエストは404エラーを返す。`)
		Routing(
			PUT("/:prefID/follow"),
			DELETE("/:prefID/follow"),
		)
		Params(func() {
			Param("prefID", Integer, "都道府県ID")
			Required("prefID")
		})
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("cron", func() {
	BasePath("/cron")
	Security(CronAuth)
	Action("fix user follow", func() {
		Description(`<b>イベントフォロー操作の確定</b><br>
		user_follow_eventsテーブルのbatch_processedをtrueに変更する`)
		Routing(
			GET("user/events/fixfollow"),
		)
		Response(OK)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("new event fetch", func() {
		Description(`<b>最新イベント情報の取得<b>`)
		Routing(
			GET("events/fetch"),
		)
		Response(OK)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("append genre", func() {
		Description(`<b>イベントにジャンルを付加する<b>`)
		Routing(
			GET("events/appendgenre"),
		)
		Response(OK)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

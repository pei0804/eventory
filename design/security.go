package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var UserAuth = APIKeySecurity("key", func() {
	Description("ユーザートークン")
	Header("X-Authorization")
})

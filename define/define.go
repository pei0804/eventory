package define

const ATDN_URL = "https://api.atnd.org/events/?count=100&format=jsonp&callback="
const CONNPASS_URL = "https://connpass.com/api/v1/event/?count=100"
const DOORKEEPER_URL = "https://api.doorkeeper.jp/events"

const (
	ATDN = iota
	CONNPASS
	DOORKEEPER
)

const (
	Production = iota
	Staging
	Test
)

// 検索範囲
const SERACH_SCOPE = 12

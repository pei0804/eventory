package define

// API
const ATDN_URL = "https://api.atnd.org/events/?count=100&format=jsonp&callback="
const CONNPASS_URL = "https://connpass.com/api/v1/event/?count=100"
const DOORKEEPER_URL = "https://api.doorkeeper.jp/events"
const SERACH_SCOPE = 12

const (
	ATDN = iota
	CONNPASS
	DOORKEEPER
)

// setting
const (
	PRODUCTION = iota + 1
	STAGING
	TEST
)

// data
const UPDATE_INFO = "updateInfo"
const DB_CONFIG = "eventory-production-setting/dbconfig.yml"

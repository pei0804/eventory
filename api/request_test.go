package api

import "testing"

func TestRequest_sendQuery(t *testing.T) {
	//a := Request{}
	//now := time.Now()
	//ym := now.AddDate(0, 1, 0).Format("200601")
	//a.Api = define.ATDN
	//a.Url = fmt.Sprintf("https://api.atnd.org/events/?count=100&format=jsonp&callback=&ym=%s", ym)
	//a.Token = ""
	//a.sendQuery()
	//
	//if a.err != nil {
	//	t.Fatalf("通信に問題が発生しました。インターネットに接続できていない。またはメソッドに間違いがあります。　%s", a.err)
	//}
}

func TestRequest_atdnJsonParse(t *testing.T) {
	// https://connpass.com/about/api/
	r := Request{}
	r.RespByte = []byte(`
{
    "events": [
        {
            "event": {
                "accepted": 5,
                "address": "東京都千代田区丸の内3丁目5番1号",
                "catch": "人は働くために、休みが必要である。",
                "description": "説明",
                "ended_at": "2112-08-20T21:30:00.000+09:00",
                "event_id": 17662,
                "event_url": "http://atnd.org/events/17662",
                "lat": "35.6769467",
                "limit": null,
                "lon": "139.7635034",
                "owner_id": 37209,
                "owner_nickname": "laughsketch",
                "owner_twitter_id": "laughsketch",
                "place": "東京国際フォーラム　ホールB7",
                "started_at": "2112-08-20T19:00:00.000+09:00",
                "title": "[テスト] qkstudy #01",
                "updated_at": "2011-07-14T07:50:33.000+09:00",
                "url": "http://laughsketch.com",
                "waiting": 0
            }
        }
    ],
    "results_returned": 1,
    "results_start": 1
}`)
	r.atdnJsonParse()
	if r.err != nil {
		t.Fatalf("[atdnJsonParse]Jsonの解析に失敗しました。%s", r.err)
	}
}

func TestRequest_connpassJsonParse(t *testing.T) {
	// http://api.atnd.org/
	r := Request{}
	r.RespByte = []byte(`
{
    "events": [
        {
            "accepted": 0,
            "address": "東京都新宿区西新宿７丁目１１−１ 宝塚大学 東京新宿キャンパス1F （オフィス２４スタジオ西新宿店内）",
            "catch": "jQueryの基礎知識を１日でマスター！",
            "description": "説明",
            "ended_at": "2017-01-13T17:30:00+09:00",
            "event_id": 48513,
            "event_type": "participation",
            "event_url": "https://nishi-shinjuku.connpass.com/event/48513/",
            "hash_tag": "epano",
            "lat": "35.693772200000",
            "limit": 2,
            "lon": "139.697319700000",
            "owner_display_name": "Epano School",
            "owner_id": 78552,
            "owner_nickname": "tiktik",
            "place": "エパノ プログラミング スクール（西新宿）",
            "series": {
                "id": 2223,
                "title": "西新宿プログラミング勉強会",
                "url": "https://nishi-shinjuku.connpass.com/"
            },
            "started_at": "2017-01-13T09:30:00+09:00",
            "title": "2名様限定の追加募集！【1/13開催】jQuery入門講座（エパノ プログラミング スクー",
            "updated_at": "2017-01-10T12:09:57+09:00",
            "waiting": 0
        }
    ],
    "results_available": 20867,
    "results_returned": 1,
    "results_start": 1
}`)
	r.connpassJsonParse()
	if r.err != nil {
		t.Fatalf("[connpassJsonParse]Jsonの解析に失敗しました。%s", r.err)
	}
}

func TestRequest_doorkeeperJsonParse(t *testing.T) {
	// https://www.doorkeeperhq.com/developer/api
	r := Request{}
	r.RespByte = []byte(`
[
    {
        "event": {
            "address": "東京都港区港南 2-16-3 品川グランドセントラルタワー",
            "banner": "https://dzpp79ucibp5a.cloudfront.net/events_banners/54727_normal_1482331398_1-2_facebook.png",
            "description": "setumei",
            "ends_at": "2017-01-22T12:00:00.000Z",
            "group": 30,
            "id": 54727,
            "lat": "35.62667",
            "long": "139.7403746",
            "participants": 45,
            "public_url": "https://swtokyo.doorkeeper.jp/events/54727",
            "published_at": "2016-12-02T15:00:02.736Z",
            "starts_at": "2017-01-20T09:30:00.000Z",
            "ticket_limit": 60,
            "title": "【初開催】SWTokyo Sports@Microsoft",
            "updated_at": "2017-01-09T08:44:49.009Z",
            "venue_name": "日本マイクロソフト",
            "waitlisted": 0
        }
    },
    {
        "event": {
            "address": "神戸市中央区京町72番地　新クレセントビル 三ノ宮駅 徒歩7分",
            "banner": "https://dzpp79ucibp5a.cloudfront.net/events_banners/54163_normal_1478917164_43940_normal_1461542680_sw2.jpg",
            "description": "",
            "ends_at": "2017-02-19T11:30:00.000Z",
            "group": 8304,
            "id": 54163,
            "lat": "34.689908",
            "long": "135.19321979999995",
            "participants": 3,
            "public_url": "https://startupweekendkobe.doorkeeper.jp/events/54163",
            "published_at": "2017-01-06T05:19:51.325Z",
            "starts_at": "2017-02-17T09:00:00.000Z",
            "ticket_limit": 30,
            "title": "【再上陸】Startup Weekend Kobe",
            "updated_at": "2017-01-07T12:34:12.112Z",
            "venue_name": "株式会社 神戸デジタル・ラボ",
            "waitlisted": 0
        }
    }
]`)
	r.doorkeeperJsonParse()
	if r.err != nil {
		t.Fatalf("[doorkeeperJsonParse]Jsonの解析に失敗しました。%s", r.err)
	}
}

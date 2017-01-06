package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/tikasan/eventory/define"
)

func TestRequest_sendQuery(t *testing.T) {
	a := Request{}
	now := time.Now()
	ym := now.AddDate(0, 1, 0).Format("200601")
	a.Api = define.ATDN
	a.Url = fmt.Sprintf("https://api.atnd.org/events/?count=100&format=jsonp&callback=&ym=%s", ym)
	a.Token = ""
	a.sendQuery()

	if a.err != nil {
		t.Fatalf("通信に問題が発生しました。インターネットに接続できていない。またはメソッドに間違いがあります。　%s", a.err)
	}
}

func TestRequest_atdnJsonParse(t *testing.T) {
	a := Request{}
	now := time.Now()
	ym := now.AddDate(0, 1, 0).Format("200601")
	a.Api = define.ATDN
	a.Url = fmt.Sprintf("https://api.atnd.org/events/?count=100&format=jsonp&callback=&ym=%s", ym)
	a.Token = ""
	cli := NewRequest(a.Url, a.Api, a.Token)
	cli.sendQuery()
	cli.atdnJsonParse()

	if cli.err != nil {
		t.Fatalf("atdnのJsonの解析に失敗しました。API元で形式が変わった可能性があります　%s", cli.err)
	}
}

func TestRequest_connpassJsonParse(t *testing.T) {
	c := Request{}
	now := time.Now()
	ym := now.AddDate(0, 1, 0).Format("200601")
	c.Url = fmt.Sprintf("https://connpass.com/api/v1/event/?count=100&ym=%s", ym)
	c.Api = define.CONNPASS
	c.Token = ""
	cli := NewRequest(c.Url, c.Api, c.Token)
	cli.sendQuery()
	cli.connpassJsonParse()

	if cli.err != nil {
		t.Fatalf("connpassのJsonの解析に失敗しました。API元で形式が変わった可能性があります　%s", cli.err)
	}
}

func TestRequest_doorkeeperJsonParse(t *testing.T) {
	d := Request{}
	d.Url = fmt.Sprintf("https://api.doorkeeper.jp/events?page=%d", 1)
	d.Api = define.DOORKEEPER
	d.Token = ""
	cli := NewRequest(d.Url, d.Api, d.Token)
	cli.sendQuery()
	cli.doorkeeperJsonParse()

	if cli.err != nil {
		t.Fatalf("doorkeeperのJsonの解析に失敗しました。API元で形式が変わった可能性があります　%s", cli.err)
	}
}

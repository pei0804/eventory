package inserter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/models"
	"github.com/tikasan/eventory/utility"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func NewParser(rawurl string, rawapi int, token string, r *http.Request) *Parser {
	return &Parser{
		Url:     rawurl,
		Api:     rawapi,
		Token:   token,
		Request: r,
	}
}

// TODO ネーミング変えるべきかも
type Parser struct {
	Url      string
	Api      int
	Token    string
	RespByte []byte
	err      error
	Request  *http.Request
}

func (p *Parser) sendQuery() {
	ctx := appengine.NewContext(p.Request)
	req, err := http.NewRequest("GET", p.Url, nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return
	}
	if p.Token != "" {
		req.Header.Set("Authorization", p.Token)
	}

	client := urlfetch.Client(ctx)
	resp, err := client.Do(req)
	if resp == nil {
		fmt.Fprint(os.Stderr, errors.New("Not Found URL check Request Url"))
		p.err = errors.New("Not Found URL check Request Url")
		return
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return
	}
	p.RespByte = respByte
}

func (p *Parser) atdnJsonParse() (events []models.EventParser, err error) {
	var at models.At
	err = json.Unmarshal(p.RespByte, &at)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return events, nil
	}
	e := new(models.EventParser)
	events = make([]models.EventParser, len(at.Events))
	for i, v := range at.Events {
		utility.CopyStruct(v.Event, e)
		events[i] = *e
		events[i].ApiId = define.ATDN
	}
	return events, nil
}

func (p *Parser) connpassJsonParse() (events []models.EventParser, err error) {
	var cp models.Cp
	err = json.Unmarshal(p.RespByte, &cp)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return events, nil
	}

	e := new(models.EventParser)
	events = make([]models.EventParser, len(cp.Events))
	for i, v := range cp.Events {
		utility.CopyStruct(v, e)
		events[i] = *e
		events[i].ApiId = define.CONNPASS
	}
	return events, nil
}

func (p *Parser) doorkeeperJsonParse() (events []models.EventParser, err error) {
	var dk []models.Dk
	err = json.Unmarshal(p.RespByte, &dk)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return events, nil
	}

	e := new(models.EventParser)
	events = make([]models.EventParser, len(dk))
	for i, v := range dk {
		utility.CopyStruct(v.Event, e)
		events[i] = *e
		events[i].ApiId = define.DOORKEEPER
	}
	return events, nil
}

func (p *Parser) convertingToJson() (events []models.EventParser, err error) {

	p.sendQuery()

	if p.Api == define.ATDN {
		return p.atdnJsonParse()
	} else if p.Api == define.CONNPASS {
		return p.connpassJsonParse()
	} else if p.Api == define.DOORKEEPER {
		return p.doorkeeperJsonParse()
	}
	return events, errors.New("未知のAPIがセットされています。")
}

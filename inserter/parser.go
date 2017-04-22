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

func NewParser(rawurl string, rawapi string, token string, r *http.Request) *Parser {
	return &Parser{
		URL:     rawurl,
		APIType: rawapi,
		Token:   token,
		Request: r,
	}
}

type Parser struct {
	URL      string
	APIType  string
	Token    string
	RespByte []byte
	err      error
	Request  *http.Request
}

func (p *Parser) sendQuery() {
	ctx := appengine.NewContext(p.Request)
	req, err := http.NewRequest("GET", p.URL, nil)
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

func (p *Parser) atdnJsonParse() (events []models.Event, err error) {
	var at models.At
	err = json.Unmarshal(p.RespByte, &at)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return events, nil
	}
	e := new(models.Event)
	events = make([]models.Event, len(at.Events))
	for i, v := range at.Events {
		utility.CopyStruct(v.Event, e)
		events[i] = *e
		events[i].APIType = define.ATDN
		events[i].Identifier = fmt.Sprintf("%d-%d", define.ATDN_ID, v.Event.EventId)
		events[i].DataHash = createDataHash(events[i])
	}
	return events, nil
}

func (p *Parser) connpassJsonParse() (events []models.Event, err error) {
	var cp models.Cp
	err = json.Unmarshal(p.RespByte, &cp)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return events, nil
	}

	e := new(models.Event)
	events = make([]models.Event, len(cp.Events))
	for i, v := range cp.Events {
		utility.CopyStruct(v, e)
		events[i] = *e
		events[i].APIType = define.CONNPASS
		events[i].Identifier = fmt.Sprintf("%d-%d", define.CONNPASS_ID, v.EventId)
		events[i].DataHash = createDataHash(events[i])
	}
	return events, nil
}

func (p *Parser) doorkeeperJsonParse() (events []models.Event, err error) {
	var dk []models.Dk
	err = json.Unmarshal(p.RespByte, &dk)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		p.err = err
		return events, nil
	}

	e := new(models.Event)
	events = make([]models.Event, len(dk))
	for i, v := range dk {
		utility.CopyStruct(v.Event, e)
		events[i] = *e
		events[i].APIType = define.DOORKEEPER
		events[i].Address = utility.RemovePoscode(events[i].Address)
		events[i].Identifier = fmt.Sprintf("%d-%d", define.DOORKEEPER_ID, v.Event.EventId)
		events[i].DataHash = createDataHash(events[i])
	}
	return events, nil
}

func createDataHash(e models.Event) string {
	d := utility.ConcatenateString(
		e.Title,
		e.Description,
		e.URL,
		e.Address,
		string(e.Limits),
		string(e.Accept),
		string(e.StartAt.Format("2006-01-02 15:04:05")),
		string(e.EndAt.Format("2006-01-02 15:04:05")))
	return utility.ToHash(d)
}

func (p *Parser) ConvertingToJson() (events []models.Event, err error) {
	p.sendQuery()
	if p.APIType == define.ATDN {
		return p.atdnJsonParse()
	} else if p.APIType == define.CONNPASS {
		return p.connpassJsonParse()
	} else if p.APIType == define.DOORKEEPER {
		return p.doorkeeperJsonParse()
	}
	return events, errors.New("未知のAPIがセットされています。")
}

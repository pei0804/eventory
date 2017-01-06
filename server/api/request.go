package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tikasan/eventory/server/define"
	"github.com/tikasan/eventory/server/formater"
	"github.com/tikasan/eventory/server/model"
)

func NewRequest(rawurl string, rawapi int, token string) *Request {
	return &Request{
		Url:   rawurl,
		Api:   rawapi,
		Token: token,
	}
}

// TODO ネーミング変えるべきかも
type Request struct {
	Url      string
	Api      int
	Token    string
	RespByte []byte
	err      error
}

func (r *Request) sendQuery() {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		r.err = err
		return
	}
	if r.Token != "" {
		req.Header.Set("Authorization", r.Token)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if resp == nil {
		fmt.Fprint(os.Stderr, errors.New("Not Found URL check Request Url"))
		r.err = errors.New("Not Found URL check Request Url")
		return
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		r.err = err
		return
	}
	r.RespByte = respByte
}

func (r *Request) atdnJsonParse() (events []model.Event, err error) {
	var at model.At
	err = json.Unmarshal(r.RespByte, &at)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		r.err = err
		return events, nil
	}
	e := new(model.Event)
	events = make([]model.Event, len(at.Events))
	for i, v := range at.Events {
		formater.CopyStruct(v.Event, e)
		events[i] = *e
		events[i].ApiId = define.ATDN
	}
	return events, nil
}

func (r *Request) connpassJsonParse() (events []model.Event, err error) {
	var cp model.Cp
	err = json.Unmarshal(r.RespByte, &cp)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		r.err = err
		return events, nil
	}

	e := new(model.Event)
	events = make([]model.Event, len(cp.Events))
	for i, v := range cp.Events {
		formater.CopyStruct(v, e)
		events[i] = *e
		events[i].ApiId = define.CONNPASS
	}
	return events, nil
}

func (r *Request) doorkeeperJsonParse() (events []model.Event, err error) {
	var dk []model.Dk
	err = json.Unmarshal(r.RespByte, &dk)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		r.err = err
		return events, nil
	}

	e := new(model.Event)
	events = make([]model.Event, len(dk))
	for i, v := range dk {
		formater.CopyStruct(v.Event, e)
		events[i] = *e
		events[i].ApiId = define.DOORKEEPER
	}
	return events, nil
}

func (r *Request) convertingToJson() (events []model.Event, err error) {

	r.sendQuery()

	if r.Api == define.ATDN {
		return r.atdnJsonParse()
	} else if r.Api == define.CONNPASS {
		return r.connpassJsonParse()
	} else if r.Api == define.DOORKEEPER {
		return r.doorkeeperJsonParse()
	}
	return events, errors.New("未知のAPIがセットされています。")
}

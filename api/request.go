package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/formater"
	"github.com/tikasan/eventory/model"
)

func NewInserter(rawurl string, rawapi int, token string) *Inserter {
	return &Inserter{
		Url:   rawurl,
		Api:   rawapi,
		Token: token,
	}
}

type Inserter struct {
	Url      string
	Api      int
	Token    string
	RespByte []byte
}

func (i *Inserter) sendQuery() {
	req, err := http.NewRequest("GET", i.Url, nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	if i.Token != "" {
		req.Header.Set("Authorization", i.Token)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	i.RespByte = respByte
}

func (i *Inserter) atdnJsonParse() (events []model.Event, err error) {
	var at model.At
	err = json.Unmarshal(i.RespByte, &at)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
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

func (i *Inserter) connpassJsonParse() (events []model.Event, err error) {
	var cp model.Cp
	err = json.Unmarshal(i.RespByte, &cp)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
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

func (i *Inserter) doorkeeperJsonParse() (events []model.Event, err error) {
	var dk []model.Dk
	err = json.Unmarshal(i.RespByte, &dk)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
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

func (i *Inserter) Get() (events []model.Event, err error) {

	i.sendQuery()

	if i.Api == define.ATDN {
		return i.atdnJsonParse()
	} else if i.Api == define.CONNPASS {
		return i.connpassJsonParse()
	} else if i.Api == define.DOORKEEPER {
		return i.doorkeeperJsonParse()
	}
	return events, errors.New("未知のAPIがセットされています。")
}

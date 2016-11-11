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
	Url   string
	Api   int
	Token string
}

func (i *Inserter) sendQuery() (respByte []byte) {
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

	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	return respByte
}

func atdnJsonParse(respByte []byte) (events []model.Event, err error) {
	var at model.At
	err = json.Unmarshal(respByte, &at)
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

func connpassJsonParse(respByte []byte) (events []model.Event, err error) {
	var cp model.Cp
	err = json.Unmarshal(respByte, &cp)
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

func doorkeeperJsonParse(respByte []byte) (events []model.Event, err error) {
	var dk []model.Dk
	err = json.Unmarshal(respByte, &dk)
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

	respByte := i.sendQuery()

	if i.Api == define.ATDN {
		return atdnJsonParse(respByte)
	} else if i.Api == define.CONNPASS {
		return connpassJsonParse(respByte)
	} else if i.Api == define.DOORKEEPER {
		return doorkeeperJsonParse(respByte)
	}
	return events, errors.New("未知のAPIがセットされています。")
}

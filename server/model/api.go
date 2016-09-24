package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tikasan/eventory/server/define"
	"github.com/tikasan/eventory/server/formater"
)

func NewInserter(rawurl string, rawapi int, token string) *Inserter {
	return &Inserter{
		Url: rawurl,
		Api: rawapi,
		Token: token,
	}
}

type Inserter struct {
	Url      string
	RespByte []byte
	Api      int
	Token    string
}

func (i *Inserter) Get() (events []Event, err error) {

	req, err := http.NewRequest("GET", i.Url, nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	if i.Token != "" {
		req.Header.Set("Authorization", i.Token)
		fmt.Println("DK")
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	if i.Api == define.ATDN {
		var at At
		err = json.Unmarshal(respByte, &at)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return events, nil
		}
		e := new(Event)
		events = make([]Event, len(at.Events))
		for i, v := range at.Events {
			formater.CopyStruct(v.Event, e)
			events[i] = *e
			events[i].ApiId = define.ATDN
		}

	} else if i.Api == define.CONNPASS {
		var cp Cp
		err := json.Unmarshal(respByte, &cp)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return events, nil
		}

		e := new(Event)
		events = make([]Event, len(cp.Events))
		for i, v := range cp.Events {
			formater.CopyStruct(v, e)
			events[i] = *e
			events[i].ApiId = define.CONNPASS
		}

	} else if i.Api == define.DOORKEEPER {

		var dk []Dk
		err := json.Unmarshal(respByte, &dk)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return events, nil
		}

		e := new(Event)
		events = make([]Event, len(dk))
		for i, v := range dk {
			formater.CopyStruct(v.Event, e)
			events[i] = *e
			events[i].ApiId = define.DOORKEEPER
		}

	} else {

		return events, errors.New("未知のAPIがセットされています。")

	}
	return events, nil
}

// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// イベント情報 (default view)
//
// Identifier: application/vnd.event+json; view=default
type Event struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// 参加登録済み人数
	Accepte int `form:"accepte" json:"accepte" xml:"accepte"`
	// 住所
	Address string `form:"address" json:"address" xml:"address"`
	// APIの種類 enum('atdn','connpass','doorkeeper')
	APIType string `form:"apiType" json:"apiType" xml:"apiType"`
	// 終了日時
	EndAt time.Time `form:"endAt" json:"endAt" xml:"endAt"`
	// 識別子(api-event_id)
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
	// 参加人数上限
	Limits int `form:"limits" json:"limits" xml:"limits"`
	// 開催日時
	StartAt time.Time `form:"startAt" json:"startAt" xml:"startAt"`
	// イベント名
	Title string `form:"title" json:"title" xml:"title"`
	// イベントページURL
	URL string `form:"url" json:"url" xml:"url"`
	// キャンセル待ち人数
	Wait int `form:"wait" json:"wait" xml:"wait"`
}

// Validate validates the Event media type instance.
func (mt *Event) Validate() (err error) {

	if mt.Identifier == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "identifier"))
	}
	if mt.APIType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "apiType"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}

	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	return
}

// EventCollection is the media type for an array of Event (default view)
//
// Identifier: application/vnd.event+json; type=collection; view=default
type EventCollection []*Event

// Validate validates the EventCollection media type instance.
func (mt EventCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ジャンル (default view)
//
// Identifier: application/vnd.genre+json; view=default
type Genre struct {
	// ジャンルID
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// ジャンル名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// GenreCollection is the media type for an array of Genre (default view)
//
// Identifier: application/vnd.genre+json; type=collection; view=default
type GenreCollection []*Genre

// ユーザー情報 (default view)
//
// Identifier: application/vnd.message+json; view=default
type Message struct {
	// トークン
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// 都道府県 (default view)
//
// Identifier: application/vnd.pref+json; view=default
type Pref struct {
	// 都道府県ID
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// 都道府県名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ユーザー情報 (default view)
//
// Identifier: application/vnd.token+json; view=default
type Token struct {
	// トークン
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

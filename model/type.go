package model

import "time"

// A Atdn
type At struct {
	Events []AAt `json:"events"`
}

type AAt struct {
	Event AtdnEvent `json:"event"`
}

// C Connpass
type Cp struct {
	Events []ConnpassEvent `json:"events"`
}

// D Doorkeeper
type Dk struct {
	Event DoorkeeperEvent `json:"event"`
}

type Event struct {
	EventId    int `json:"event_id"`
	ApiId      int
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Url        string `json:"url"`
	Limit      int    `json:"limit"`
	Accepted   int    `json:"accepted"`
	Waitlisted int    `json:"waitlisted"`
	Address    string `json:"address"`
	Place      string `json:"place"`
	StratAt    string `json:"strat_at"`
	EndAt      string `json:"end_at"`
	DataHash   string
	ID         int `json:"id"`
}

type EventJson struct {
	EventId string `json:"event_id"`
	ApiId   int    `json:"api_id"`
	Title   string `json:"title"`
	//Desc       string `json:"desc"`
	Url      string `json:"url"`
	Limit    int    `json:"limit"`
	Accepted int    `json:"accepted"`
	//Waitlisted int    `json:"waitlisted"`
	Address string    `json:"address"`
	Place   string    `json:"place"`
	StratAt time.Time `json:"strat_at"`
	EndAt   time.Time `json:"end_at"`
	ID      int       `json:"id"`
}

type AtdnEvent struct {
	EventId    int `json:"event_id"`
	ApiId      int
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Url        string `json:"event_url"`
	Limit      int    `json:"limit"`
	Accepted   int    `json:"accepted"`
	Waitlisted int    `json:"waiting"`
	Address    string `json:"address"`
	Place      string `json:"place"`
	StratAt    string `json:"started_at"`
	EndAt      string `json:"ended_at"`
	DataHash   string
}

type ConnpassEvent struct {
	EventId    int `json:"event_id"`
	ApiId      int
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Url        string `json:"event_url"`
	Limit      int    `json:"limit"`
	Accepted   int    `json:"accepted"`
	Waitlisted int    `json:"waiting"`
	Address    string `json:"address"`
	Place      string `json:"place"`
	StratAt    string `json:"started_at"`
	EndAt      string `json:"ended_at"`
	DataHash   string
}

type DoorkeeperEvent struct {
	EventId    int `json:"id"`
	ApiId      int
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Url        string `json:"public_url"`
	Limit      int    `json:"ticket_limit"`
	Accepted   int    `json:"participants"`
	Waitlisted int    `json:"waitlisted"`
	Address    string `json:"address"`
	Place      string `json:"place"`
	StratAt    string `json:"starts_at"`
	EndAt      string `json:"ends_at"`
	DataHash   string
}

type UpdateInfo struct {
	_kind    string    `goon:"kind,updateInfo"`
	Id       int64     `datastore:"-" goon:"id"`
	Datetime time.Time `datastore:"datetime,noindex"`
}

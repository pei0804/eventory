package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tikasan/eventory/formater"
)

func Insert(db *sql.DB, Events []Event) error {
	stmtIns, err := db.Prepare("INSERT INTO m_event (event_id, api_id, title, description, url, limit_count, waitlisted, accepted, address, place, start_at, end_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return err
	}

	for _, ev := range Events {
		if _, err = stmtIns.Exec(
			fmt.Sprintf("%d-%d", ev.ApiId, ev.EventId),
			ev.ApiId,
			ev.Title,
			ev.Desc,
			ev.Url,
			ev.Limit,
			ev.Waitlisted,
			ev.Accepted,
			ev.Address,
			ev.Place,
			formater.DateTimeFormatter(ev.StratAt),
			formater.DateTimeFormatter(ev.EndAt),
		); err != nil {
			// insertに失敗したらアップデートをかける
			query := "UPDATE m_event SET title = ?, description = ?, url = ?, limit_count = ?, waitlisted = ?, accepted = ?, address = ?, place = ?, start_at = ?, end_at = ? WHERE event_id = ?"
			if _, err := db.Exec(query,
				ev.Title,
				ev.Desc,
				ev.Url,
				ev.Limit,
				ev.Waitlisted,
				ev.Accepted,
				ev.Address,
				ev.Place,
				formater.DateTimeFormatter(ev.StratAt),
				formater.DateTimeFormatter(ev.EndAt),
				fmt.Sprintf("%d-%d", ev.ApiId, ev.EventId),
			); err != nil {
				fmt.Fprint(os.Stderr, err)
				return err
			}

		}
	}
	return nil
}

func EventAll(db *sql.DB) ([]Event, error) {
	rows, err := db.Query(`select * from m_event`)
	if err != nil {
		return nil, err
	}
	return ScanEvents(rows)
}

func EventAllNew(db *sql.DB) ([]EventJson, error) {
	rows, err := db.Query(`select event_id, api_id,title, url, limit_count, accepted, address ,place, start_at, end_at, id from m_event where end_at > now();`)
	if err != nil {
		return nil, err
	}
	return ScanEventsJson(rows)
}

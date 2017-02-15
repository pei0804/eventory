package model

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/formater"
)

func Insert(db *sql.DB, Events []Event) error {

	insert := `INSERT INTO m_event (event_id, api_id, title, description, url, limit_count, waitlisted, accepted, address, place, start_at, end_at, data_hash) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	for _, ev := range Events {
		if ev.ApiId == define.DOORKEEPER {
			ev.Address = formater.RemovePoscode(ev.Address)
		}
		dataHashAgo := formater.ConcatenateString(ev.Title, ev.Desc, ev.Url, ev.Address, string(ev.Limit), string(ev.Accepted), ev.Place, ev.StartAt, ev.EndAt)
		dataHashed := sha256.Sum256([]byte(dataHashAgo))
		ev.DataHash = hex.EncodeToString(dataHashed[:])
		if _, err := db.Exec(insert,
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
			formater.DateTime(ev.StartAt),
			formater.DateTime(ev.EndAt),
			ev.DataHash,
		); err != nil {
			// insertに失敗したらアップデートをかける
			query := `UPDATE m_event SET title = ?, description = ?, url = ?, limit_count = ?, waitlisted = ?, accepted = ?, address = ?, place = ?, start_at = ?, end_at = ?, data_hash = ? WHERE event_id = ? AND data_hash <> ?`
			if _, err := db.Exec(query,
				ev.Title,
				ev.Desc,
				ev.Url,
				ev.Limit,
				ev.Waitlisted,
				ev.Accepted,
				ev.Address,
				ev.Place,
				formater.DateTime(ev.StartAt),
				formater.DateTime(ev.EndAt),
				ev.DataHash,
				fmt.Sprintf("%d-%d", ev.ApiId, ev.EventId),
				ev.DataHash,
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

func EventAllNew(db *sql.DB, updatedAt string, places []string) ([]EventJson, error) {

	pb := bytes.NewBuffer(make([]byte, 0, 800))
	pb.WriteString("(")
	for _, p := range places {
		if utf8.RuneCountInString(p) <= 4 {
			pb.WriteString("address LIKE '")
			pb.WriteString(p)
			pb.WriteString("%' ")
			if p != places[len(places)-1] {
				pb.WriteString("OR ")
			}
		}
	}
	pb.WriteString(") AND ")
	sql := fmt.Sprintf("select event_id, api_id,title, url, limit_count, accepted, address ,place, start_at, end_at, id from m_event where %s end_at > now() AND updated_at > ?", pb.String())

	stmtIns, err := db.Prepare(sql)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	rows, err := stmtIns.Query(updatedAt)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	return ScanEventsJson(rows)
}

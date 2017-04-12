// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Models
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// イベントジャンル
type EventGenre struct {
	ID        int        `gorm:"primary_key"` // primary key
	EventID   int        // Belongs To Event
	GenreID   int        // has many EventGenre
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
	Event     Event
	Genre     Genre
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m EventGenre) TableName() string {
	return "event_genres"

}

// EventGenreDB is the implementation of the storage interface for
// EventGenre.
type EventGenreDB struct {
	Db *gorm.DB
}

// NewEventGenreDB creates a new storage type.
func NewEventGenreDB(db *gorm.DB) *EventGenreDB {
	return &EventGenreDB{Db: db}
}

// DB returns the underlying database.
func (m *EventGenreDB) DB() interface{} {
	return m.Db
}

// EventGenreStorage represents the storage interface.
type EventGenreStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*EventGenre, error)
	Get(ctx context.Context, id int) (*EventGenre, error)
	Add(ctx context.Context, eventgenre *EventGenre) error
	Update(ctx context.Context, eventgenre *EventGenre) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *EventGenreDB) TableName() string {
	return "event_genres"

}

// Belongs To Relationships

// EventGenreFilterByEvent is a gorm filter for a Belongs To relationship.
func EventGenreFilterByEvent(eventID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if eventID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("event_id = ?", eventID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// EventGenreFilterByGenre is a gorm filter for a Belongs To relationship.
func EventGenreFilterByGenre(genreID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if genreID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("genre_id = ?", genreID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single EventGenre as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *EventGenreDB) Get(ctx context.Context, id int) (*EventGenre, error) {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenre", "get"}, time.Now())

	var native EventGenre
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of EventGenre
func (m *EventGenreDB) List(ctx context.Context) ([]*EventGenre, error) {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenre", "list"}, time.Now())

	var objs []*EventGenre
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *EventGenreDB) Add(ctx context.Context, model *EventGenre) error {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenre", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding EventGenre", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *EventGenreDB) Update(ctx context.Context, model *EventGenre) error {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenre", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating EventGenre", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *EventGenreDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenre", "delete"}, time.Now())

	var obj EventGenre

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting EventGenre", "error", err.Error())
		return err
	}

	return nil
}

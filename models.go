// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package lgtts is Let's Go to the Show!
package lgtts

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/coocood/qbs"
	"time"
)

// An Artist is a band, musician, painter, sculptor, performer, or anyone who
// wants to hold well-attended public events.
type Artist struct {
	Id            int64
	Name          string `qbs:"index,notnull"`
	Email         string `qbs:"index,unique,notnull"`
	HomeTown      string
	HomeZip       string
	Description   string
	StormpathHref string
	Updated       time.Time
	Created       time.Time
}

// A Show is a public event such as a concert, art show, play, etc.
type Show struct {
	Id          int64
	Artist      *Artist
	ArtistId    int64
	Time        time.Time
	Venue       string
	Zip         string // Venue zip code
	Price       float32
	Description string
}

// A Patron is a fan, patron, or other person who wants to be notified about
// upcoming Shows.
type Patron struct {
	Id       int64
	Artist   *Artist
	ArtistId int64
	Email    string
	Zip      string
	Created  time.Time // Record creation date
	Referer  string    // Referer URL
}

// A Payment is a reference to a payments model TBD.
type Payment string

// A Blast is an email blast of Notifications for a given show
type Blast struct {
	Id       int64
	Artist   *Artist
	ArtistId int64
	Show     *Show
	Max      int       // Max Patrons to notify - unlimited if 0
	RunDate  time.Time // Date on which to send this blast
	Payment  *Payment
	Start    time.Time
	Finish   time.Time
	Confirm  time.Time // Confirmation email sent to Artist
}

// A Notification is an email message sent to a Patron notifying them of a Show.
type Notification struct {
	Id     int64
	Show   *Show
	Blast  *Blast
	Patron *Patron
	Sent   *time.Time
	Token  *uuid.UUID
}

var tables = []interface{}{
	&Artist{},
	&Show{},
}

// MigrateTables will attempt to migrate the database to the current schema,
// creating tables that do not exist, and adding columns to those that do.
// Only additive operations are supported - it will not alter or delete columns
// - so it should be safe for production. Will panic if it can't migrate a
// table, or return an error if it cannot create a necessary index.
func (srv *Server) MigrateTables() error {
	q := srv.Qbs()
	m := qbs.NewMigration(q.Db, srv.DbName, q.Dialect)
	for _, t := range tables {
		err := m.CreateTableIfNotExists(t)
		if err != nil {
			return err
		}

	}
	return nil
}

// dropTables drops all tables.  It can ONLY be used on databases whose
// names end in "_test".
func (srv *Server) dropTables() error {
	q := srv.Qbs()
	m := qbs.NewMigration(q.Db, srv.DbName, q.Dialect)
	for _, t := range tables {
		m.DropTable(t)
	}
	return nil
}

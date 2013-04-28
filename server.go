// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package lgtts is Let's Go to the Show!
package lgtts

import (
	// "code.google.com/p/go-uuid/uuid"
	"database/sql"
	// "github.com/coopernurse/gorp"
	"github.com/coocood/qbs"
	"log"
	"time"
)

/*
NOTES:

 - http://zipcodedistanceapi.cymi.org/API

*/

type Server struct {
	DriverName     string
	DataSourceName string
	DbName         string
	Dialect        qbs.Dialect
	Logger         *log.Logger
}

func (srv *Server) Qbs() *qbs.Qbs {
	//Try to get a free DB connection from the pool
	db := qbs.GetFreeDB()
	var err error
	if db == nil { // connection pool is empty
		//open a new one.
		db, err = sql.Open(srv.DriverName, srv.DataSourceName)
	}
	if err != nil {
		log.Fatal(err)
	}
	return qbs.New(db, srv.Dialect)
}

// NewArtist creates a new Arist record.  Email must be unique.
func (srv *Server) NewArtist(name, email string) (*Artist, error) {
	q := srv.Qbs()
	defer q.Close()
	//
	// Check for duplicate email
	//
	a := new(Artist)
	a.Email = email
	err := q.Find(a)
	if err == nil {
		return a, EmailAlreadyRegistered
	}
	if err != sql.ErrNoRows {
		return a, err
	}
	//
	// Create new record
	//
	a.Name = name
	_, err = q.Save(a)
	return a, err
}

// GetArtist retrieves an artist's profile.
func (srv *Server) GetArtist(email string) (*Artist, error) {
	q := srv.Qbs()
	defer q.Close()
	a := new(Artist)
	err := q.Find(a)
	return a, err
}

// UpdateArtist saves changes to an Artist profile.
func (srv *Server) UpdateArtist(a *Artist) error {
	q := srv.Qbs()
	defer q.Close()
	_, err := q.Save(a)
	return err
}

// DeleteArtist removes an artist's profile and all their patrons.
func (srv *Server) DeleteArtist(a *Artist) error {
	q := srv.Qbs()
	defer q.Close()
	_, err := q.Delete(a)
	return err
}

// Register signs up a user to receive emails when the artist has a show
// near their zipcode.
func (srv *Server) AddPatron(artist int64, email, zip string) error {
	return nil
}

// DeletePatron stops future notifications for a given patronage.
func (srv *Server) DeletePatron(p *Patron) error {
	return nil
}

// DeleteEmail stops ALL future notifications to a given email.
func (srv *Server) DeleteEmail(email string) error {
	return nil
}

// NewBlast creates a new Blast job to be processed by blast worker
func (srv *Server) NewBlast(s *Show, max int, date *time.Time, p *Payment) (*Blast, error) {
	return nil, nil
}

// SendBlast sends a blast of email notifications.
func (srv *Server) SendBlast(b *Blast) (sent int, err error) {
	return 0, nil
}

// Run starts background workers.
func (srv *Server) Run() error {
	return nil
}

func (srv *Server) Notify(s *Show, p *Patron) error {
	return nil
}

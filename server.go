// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package lgtts is Let's Go to the Show!
package lgtts

import (
	// "code.google.com/p/go-uuid/uuid"
	"database/sql"
	"github.com/coopernurse/gorp"
	"log"
	"time"
)

/*
NOTES:

 - http://zipcodedistanceapi.cymi.org/API

*/

func NewServer(db *sql.DB) *Server {
	return nil
}

type Server struct {
	DbMap  *gorp.DbMap
	Logger *log.Logger
}

// NewArtist creates a new Arist record.  Email must be unique.
func (srv *Server) NewArtist(name, email string) (*Artist, error) {
	return nil, nil
}

// UpdateArtist saves changes to an Artist profile.
func (srv *Server) UpdateArtist(a *Artist) error {
	return nil
}

// DeleteArtist removes an artist's profile and all their patrons.
func (srv *Server) DeleteArtist(a *Artist) error {
	return nil
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

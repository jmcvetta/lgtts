// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package lgtts is Let's Go to the Show!
package lgtts

import (
	// "code.google.com/p/go-uuid/uuid"
	"time"
)

/*
NOTES:

 - http://zipcodedistanceapi.cymi.org/API

*/

type Storage interface {
	// NewArtist creates a new Arist record.  Name and email must be unique.
	NewArtist(name, email string) (*Artist, error)

	// UpdateArtist saves changes to an Artist profile.
	UpdateArtist(a *Artist) error

	// DeleteArtist removes an artist's profile and all their patrons.
	DeleteArtist(a *Artist) error

	// Register signs up a user to receive emails when the artist has a show
	// near their zipcode.
	AddPatron(artist int64, email, zip string) error

	// DeletePatron stops future notifications for a given patronage.
	DeletePatron(p *Patron) error

	// DeleteEmail stops ALL future notifications to a given email.
	DeleteEmail(email string) error

	// NewBlast creates a new Blast job to be processed by blast worker
	NewBlast(s *Show, max int, date *time.Time, p *Payment) (*Blast, error)

	// SendBlast sends a blast of email notifications.
	SendBlast(b *Blast) (sent int, err error)
}

type Notifier interface {
	// Notify sends an email notification
	Notify(s *Show, p *Patron) error
}

type Server struct {
	Storage
	Notifier
}

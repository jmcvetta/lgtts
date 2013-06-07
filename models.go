// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package lgtts is Let's Go to the Show!
package main

import (
	"time"
)

// An Artist is a band, musician, painter, sculptor, performer, or anyone who
// wants to hold well-attended public events.
type Artist struct {
	Id            int64
	Name          string
	Email         string
	Hometown      string
	Zip           string // Artist hometown zip
	Description   string
	StormpathHref string
	Updated       time.Time
	Created       time.Time
}

// A Show is a public event such as a concert, art show, play, etc.
type Show struct {
	Id          int64
	ArtistId    int64
	Time        time.Time
	Venue       string
	Address     string // Venue street address
	City        string // Venue city
	State       string // Venue state
	Zip         string // Venue zip code
	Price       string // Should include currency
	Description string
}

// A Patron is a fan, patron, or other person who wants to be notified about
// upcoming Shows.
type Patron struct {
	Id        int64
	ArtistId  int64
	Email     string
	Zip       string
	Created   time.Time // Record creation date
	Referer   string    // Referer URL
	Confirmed bool
}

// A Payment is a reference to a payments model TBD.
type Payment string

// A Blast is an email blast of Notifications for a given show
type Blast struct {
	Id       int64
	ArtistId int64
	ShowId   int64
	Max      int64     // Max Patrons to notify - unlimited if 0
	RunDate  time.Time // Date on which to send this blast
	Start    time.Time
	Finish   time.Time
	Confirm  time.Time // Confirmation email sent to Artist
}

// A Notification is an email message sent to a Patron notifying them of a Show.
type Notification struct {
	Id       int64
	ShowId   int64
	BlastId  int64
	PatronId int64
	Sent     time.Time
}

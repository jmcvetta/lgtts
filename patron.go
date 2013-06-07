// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"
)

var validZip = regexp.MustCompile(`^\d{5}$`)

type PatronageRequest struct {
	ArtistId int64
	Email    string
	Zip      string
}

// newPatron
func NewPatronHandler(w http.ResponseWriter, r *http.Request) {
	//
	// Unmarshall request
	//
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	pr := PatronageRequest{}
	err := dec.Decode(&pr)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	//
	// Validate zip code
	//
	idx := validZip.FindStringIndex(pr.Zip)
	if idx == nil {
		log.Println("Invalid zip code ", pr.Zip)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	//
	// Save patron to database
	//
	p := Patron{
		ArtistId:  pr.ArtistId,
		Email:     pr.Email,
		Zip:       pr.Zip,
		Created:   time.Now(),
		Confirmed: false,
	}
	err = dbmap.Insert(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	//
	// Success
	//
	w.WriteHeader(200)

}

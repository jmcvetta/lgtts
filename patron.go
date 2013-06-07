// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	restful "github.com/emicklei/go-restful"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var validZip = regexp.MustCompile(`^\d{5}$`)

type PatronageRequest struct {
	Email string
	Zip   string
}

// patronize handles a request to patronize an artist.
func patronize(req *restful.Request, resp *restful.Response) {
	pr := PatronageRequest{}
	err := req.ReadEntity(&pr)
	if err != nil {
		log.Println(err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	//
	// Validate zip code
	//
	idx := validZip.FindStringIndex(pr.Zip)
	if idx == nil {
		log.Println("Invalid zip code:", pr.Zip)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	//
	// Save patron to database
	//
	artistId, err := strconv.Atoi(req.PathParameter("artist-id"))
	if err != nil {
		log.Println(err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	p := Patron{
		ArtistId:  int64(artistId),
		Email:     pr.Email,
		Zip:       pr.Zip,
		Created:   time.Now(),
		Confirmed: false,
	}
	err = dbmap.Insert(&p)
	if err != nil {
		log.Println(err)
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}
	resp.WriteEntity(p)
}

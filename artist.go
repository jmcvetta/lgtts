// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	"errors"
	restful "github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type artistRequest struct {
	Name        string
	Email       string
	Hometown    string
	Zip         string
	Description string
}

func createArtist(req *restful.Request, resp *restful.Response) {
	ar := artistRequest{}
	err := req.ReadEntity(&ar)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	a := Artist{
		Name:        ar.Name,
		Email:       ar.Email,
		Hometown:    ar.Hometown,
		Zip:         ar.Zip,
		Description: ar.Description,
	}
	err = dbmap.Insert(&a)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
}

// artistByEmail retrieves an artist's profile.
func artistByEmail(email string) (*Artist, error) {
	artists := []*Artist{}
	query := "SELECT * FROM Artist WHERE Email=$1"
	_, err := dbmap.Select(&artists, query, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	switch len(artists) {
	case 0:
		return nil, NotFound
	case 1:
		return artists[0], nil
	}
	return nil, errors.New("DB consistency error - multiple artists with same email.")
}

// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	// "github.com/bmizerany/assert"
	"github.com/jmcvetta/restclient"
	"strconv"
	"testing"
)

func TestNewPatron(t *testing.T) {
	hserv := setupTest(t)
	defer hserv.Close()
	//
	// Create a new artist
	//
	a, _ := newArtist(hserv)
	//
	// Create new patron
	//
	payload := PatronageRequest{
		Email: "jason.mcvetta+lgtts-patronize@gmail.com",
		Zip:   "94102",
	}
	url := hserv.URL + "/api/v1/artists/" + strconv.Itoa(int(a.Id)) + "/patrons"
	p := Patron{}
	rr := restclient.RequestResponse{
		Url:            url,
		Method:         "POST",
		Data:           payload,
		Result:         &p,
		ExpectedStatus: 200,
	}
	_, err := restclient.Do(&rr)
	if err != nil {
		t.Fatal(err)
	}
}

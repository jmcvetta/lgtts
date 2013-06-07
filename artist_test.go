// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	"github.com/bmizerany/assert"
	"github.com/jmcvetta/restclient"
	"testing"
)

func TestNewArtist(t *testing.T) {
	hserv := setupTest(t)
	defer hserv.Close()
	email := "jason.mcvetta+test-gg.allin@gmail.com"
	//
	// Create a new artist
	//
	payload := artistRequest{
		Name:        "GG Allin",
		Email:       email,
		Hometown:    "NYC",
		Zip:         "11011",
		Description: "Pretty fucking cool",
	}
	url := hserv.URL + "/api/v1/artist"
	rr := restclient.RequestResponse{
		Url:            url,
		Method:         "POST",
		Data:           &payload,
		ExpectedStatus: 200,
	}
	_, err := restclient.Do(&rr)
	if err != nil {
		t.Fatal(err)
	}
	//
	// Confirm DB record
	//
	query := "SELECT COUNT(*) FROM artist WHERE email=$1"
	cnt, err := dbmap.SelectInt(query, email)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), cnt, "One and only one artist with this email.")
	//
	// Create with invalid request
	//
	rr = restclient.RequestResponse{
		Url:            url,
		Method:         "POST",
		Data:           "foobar",
		ExpectedStatus: 400,
	}
	_, err = restclient.Do(&rr)
	if err != nil {
		t.Fatal(err)
	}
	//
	// Try to use duplicate email
	//
}

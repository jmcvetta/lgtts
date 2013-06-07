// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	restful "github.com/emicklei/go-restful"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest(t *testing.T) *httptest.Server {
	//
	// Logging
	//
	log.SetPrefix("[lgtts] ")
	log.SetFlags(log.Lshortfile)
	err := SetupPostgres()
	if err != nil {
		t.Fatal(err)
	}
	dbmap.DropTables() // Ignore errors
	err = dbmap.CreateTables()
	if err != nil {
		t.Fatal(err)
	}
	ws := NewWebService()
	restful.Add(ws)
	hserv := httptest.NewServer(http.HandlerFunc(restful.Dispatch))
	return hserv
}

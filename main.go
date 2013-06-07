// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package main is the application "Let's Go to the Show!"
package main

import (
	// "database/sql"
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/darkhelmet/env"
	restful "github.com/emicklei/go-restful"
	"github.com/lib/pq"
	"log"
	"net/http"
)

/*
NOTES:

 - http://zipcodedistanceapi.cymi.org/API

*/

var dbmap gorp.DbMap

func init() {
	dbmap.AddTable(Artist{}).SetKeys(true, "Id")
	dbmap.AddTable(Show{}).SetKeys(true, "Id")
	dbmap.AddTable(Patron{}).SetKeys(true, "Id")
	dbmap.AddTable(Blast{}).SetKeys(true, "Id")
	dbmap.AddTable(Notification{}).SetKeys(true, "Id")
}

func setupPostgres() error {
	dbUrl := env.StringDefault("DATABASE_URL", "postgres://")
	dsn, err := pq.ParseURL(dbUrl)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	dbmap.Db = db
	dbmap.Dialect = gorp.PostgresDialect{}
	return nil
}

func NewWebService() *restful.WebService {
	ws := restful.WebService{}
	ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.POST("/artist").To(createArtist).
		Doc("Create a new Arist").
		Reads(artistRequest{}))
	return &ws
}

func main() {
	port := env.StringDefault("PORT", "8000")
	log.SetPrefix("[lgtts] ")
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	err := setupPostgres()
	if err != nil {
		log.Fatal(err)
	}
	ws := NewWebService()
	restful.Add(ws)
	log.Printf("Listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

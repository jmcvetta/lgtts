// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	"database/sql"
	"github.com/bmizerany/pq"
	"github.com/coopernurse/gorp"
	"github.com/darkhelmet/env"
	"log"
	"os"
)

func main() {
	srv := Server{}
	//
	// Logging
	//
	l := log.New(os.Stdout, "[lgtts] ", log.Ltime|log.Ldate|log.Lshortfile)
	srv.Logger = l
	//
	// Get Environment Variables
	//
	dbUrl := env.StringDefault("DATABASE_URL", "postgres://")
	//
	// Connect to PostgreSQL
	//
	dsn, err := pq.ParseURL(dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	srv.DbMap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
}

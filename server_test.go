// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import (
	"github.com/bmizerany/assert"
	"github.com/bmizerany/pq"
	"github.com/coocood/qbs"
	"github.com/darkhelmet/env"
	"log"
	"os"
	"testing"
)

func setup(t *testing.T) *Server {
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
		t.Fatal(err)
	}
	srv.DataSourceName = dsn
	srv.DriverName = "postgres"
	srv.DbName = "lgtts_test"
	srv.Dialect = qbs.NewPostgres()
	err = srv.dropTables()
	if err != nil {
		t.Fatal(err)
	}
	err = srv.MigrateTables()
	if err != nil {
		t.Fatal(err)
	}
	return &srv
}

func TestNewArtist(t *testing.T) {
	srv := setup(t)
	name := "James T Kirk"
	email := "captain@enterprise.gov"
	//
	// Create a new artist
	//
	a0, err := srv.NewArtist(name, email)
	if err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, a0.Id, 0)
	assert.Equal(t, a0.Name, name)
	assert.Equal(t, a0.Email, email)
	//
	// Confirm DB record
	//
	q := srv.Qbs()
	a1 := new(Artist)
	a1.Name = name
	a1.Email = email
	err = q.Find(a1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, a1.Id, a0.Id)
	//
	// Try to use duplicate email
	//
	_, err = srv.NewArtist(name, email)
	assert.NotEqual(t, err, nil)
}

func TestGetArtist(t *testing.T) {
	srv := setup(t)
	name := "James T Kirk"
	email := "captain@enterprise.gov"
	a0, _ := srv.NewArtist(name, email)
	a1, err := srv.GetArtist(email)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, a0.Id, a1.Id)
	assert.Equal(t, a0.Name, a1.Name)
	assert.Equal(t, a0.Email, a1.Email)
}

/*
func TestUpdateArtist(t *testing.T) {
	srv := setup(t)
	name := "James T Kirk"
	email := "captain@enterprise.gov"

}
*/

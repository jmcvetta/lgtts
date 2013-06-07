// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package main is the application "Let's Go to the Show!"
package main

import (
	"github.com/darkhelmet/env"
	restful "github.com/emicklei/go-restful"
	"github.com/jmcvetta/lgtts"
	"log"
	"net/http"
)

/*
NOTES:

 - http://zipcodedistanceapi.cymi.org/API

*/

func main() {
	port := env.StringDefault("PORT", "8000")
	log.SetPrefix("[lgtts] ")
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	err := lgtts.SetupPostgres()
	if err != nil {
		log.Fatal(err)
	}
	ws := lgtts.NewWebService()
	restful.Add(ws)
	log.Printf("Listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

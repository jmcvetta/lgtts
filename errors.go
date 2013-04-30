// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the AGPL v3.  See www.gnu.org/licenses/agpl-3.0.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package lgtts

import "errors"

var (
	EmailAlreadyRegistered = errors.New("Email address is already registered.")
	AlreadyPatron = errors.New("Email is already registered as a patron of Artist in Zip code.")
)

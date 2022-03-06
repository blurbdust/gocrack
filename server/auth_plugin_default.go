// +build !auth_database,!auth_ldap

package server

import (
	_ "github.com/blurbdust/gocrack/server/authentication/database"
	_ "github.com/blurbdust/gocrack/server/authentication/ldap"
)

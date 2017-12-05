/*
  https://github.com/go-sql-driver/mysql/blob/master/utils_go18.go
*/

// +build go1.9

package /* mysql */ main

import (
	"crypto/tls"
	//"database/sql"
	//"database/sql/driver"
	//"errors"
)

func cloneTLSConfig(c *tls.Config) *tls.Config {
	return c.Clone()
}

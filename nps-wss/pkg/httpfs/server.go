/*
  Inspired by
    https://github.com/golang/go/blob/master/src/net/http/server.go
*/

package /* http */ httpfs

import (
	"strings"
)

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

/*
 Package internal wraps http.Client to do json encoding/decoding automatically and add default query parameters (i.e. wt=json)
*/
package internal

import (
	"github.com/orest-hopiak-symphony/go-solr/solr/util/logutil"
)

var log = logutil.Logger.RegisterPkg()

const (
	DefaultUserAgent  = "go-solr"
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
	MediaTypeJSON     = "application/json"
)

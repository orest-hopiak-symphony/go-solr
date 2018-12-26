package solr

import (
	"github.com/orest-hopiak-symphony/go-solr/solr/util/logutil"
)

var log = logutil.Logger.RegisterPkg()

func init() {
	log.SetPkgAlias("solr")
}

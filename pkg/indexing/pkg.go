package indexing

import (
	"fmt"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/internal"
	"github.com/at15/go-solr/pkg/util"
)

var log = util.Logger.RegisterPkg()

const (
	baseURLTmpl = "/solr/%s/update"
)

type Service struct {
	client *internal.Client

	core    common.Core
	baseURL string
}

func New(client *internal.Client, core common.Core) *Service {
	s := &Service{
		client: client,
	}
	s.SetCore(core)
	return s
}

func (svc *Service) SetCore(core common.Core) {
	svc.core = core
	svc.baseURL = fmt.Sprintf(baseURLTmpl, core.Name)
}

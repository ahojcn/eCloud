package nginx

import (
	"github.com/ahojcn/ecloud/ctr/model"
)

func New(host *model.Host) *Nginx {
	return &Nginx{Host: host}
}

package nginx

import (
	"github.com/ahojcn/ecloud/ctr/model"
	"time"
)

const (
	BinPath       = "/usr/local/nginx/sbin/nginx"
	LogDirPath    = "/root/.eCloud/nginx/logs/"
	AccessLogPath = "/root/.eCloud/nginx/logs/access_json.log"
	ConfDirPath   = "/root/.eCloud/nginx/conf/conf.d/"
	RootConfPath  = "/root/.eCloud/nginx/conf/nginx.conf"

	RunCmdTimeOut = 60 * time.Second * 10

	nginxReloadCmd = "-s reload"
	nginxReopenCmd = "-s reopen"
	nginxStopCmd   = "-s stop"
	nginxQuitCmd   = "-s quit"
	nginxStartCmd  = "-c " + RootConfPath
)

type Nginx struct {
	Host *model.Host
}

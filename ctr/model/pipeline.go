package model

import (
	"fmt"
)

type PipeLine struct {
	Id             int64  `json:"id" xorm:"pk autoincr notnull"`
	TreeId         int64  `json:"tree_id" xorm:"notnull"`
	ClusterId      int64  `json:"cluster_id" xorm:"notnull"`
	Status         int    `json:"status" xorm:"notnull default 1"`
	StatusMsg      string `json:"status_msg" xorm:"varchar(32) notnull default '初始化'"`
	RouterIp       string `json:"router_ip" xorm:"varchar(64) notnull default ''"`
	RouterPort     int    `json:"router_port" xorm:"notnull default 9999"`
	ContainerImage string `json:"container_image" xorm:"text notnull"`
	ErrorLog       string `json:"error_log" xorm:"text notnull"`

	AliveMethod     string `json:"alive_method" xorm:"varchar(16) notnull"`
	AliveURI        string `json:"alive_uri" xorm:"text notnull"`
	AliveReqQuery   string `json:"alive_req_query" xorm:"text notnull"`
	AliveReqBody    string `json:"alive_req_body" xorm:"text notnull"`
	AliveReqHeader  string `json:"alive_req_header" xorm:"text notnull"`
	AliveRespStatus int    `json:"alive_resp_status" xorm:"notnull default 200"`
	AliveRespBody   string `json:"alive_resp_body" xorm:"text notnull"`
}

type PipeLineInfo struct {
	PipeLine
	ClusterInfo *ClusterInfo `json:"cluster_info"`
	TreeInfo    *Tree        `json:"tree_info"`
}

func (p PipeLine) GetPipeLineInfo() PipeLineInfo {
	c, _ := ClusterOne(map[string]interface{}{"tree_id": p.ClusterId})
	t, _ := TreeOne(map[string]interface{}{"id": p.TreeId})
	return PipeLineInfo{
		PipeLine:    p,
		ClusterInfo: c.GetClusterInfo(),
		TreeInfo:    t,
	}
}

const (
	PipeLineStatusError        = 0
	PipeLineStatusInit         = 1
	PipeLineStatusBuildImage   = 2
	PipeLineStatusRunContainer = 3
	PipeLineStatusAliveTest    = 4
	PipeLineStatusRouter       = 5
	PipeLineStatusRunning      = 6
)

var PipeLineStatusMsg = []string{"出错", "初始化", "构建镜像", "运行容器", "存活测试", "接入层介入", "运行中"}
var PipeLineStatusContent = []string{
	"出现错误", "流水线创建后未进行部署操作", "选择资源中的主机并构建镜像", "在选择的主机上运行容器",
	"对每台容器进行存活测试", "生成接入层配置并reload生效", "正在运行，可以接入流量"}

func (p *PipeLine) GetTree() (*Tree, error) {
	t, has := TreeOne(map[string]interface{}{"id": p.TreeId})
	if !has {
		return nil, fmt.Errorf("没有这个服务树节点信息,treeid=%d", p.TreeId)
	}
	return t, nil
}

func (p *PipeLine) GetCluster() (*Cluster, error) {
	t, has := ClusterOne(map[string]interface{}{"id": p.ClusterId})
	if !has {
		return nil, fmt.Errorf("没有这个集群信息,clusterid=%d", p.ClusterId)
	}
	return t, nil
}

func (p *PipeLine) GetStatusMsg() string {
	return PipeLineStatusMsg[p.Status]
}

func PipeLineAdd(c *PipeLine) error {
	orm := GetMaster()
	affected, err := orm.Insert(c)
	if affected == 0 {
		return fmt.Errorf("insert failed, affected = 0")
	}
	return err
}

func PipeLineOne(cons map[string]interface{}) (*PipeLine, bool) {
	orm := GetSlave()

	t := new(PipeLine)
	has, err := orm.Where(cons).Get(t)
	if err != nil {
		return nil, false
	}
	return t, has
}

func PipeLineList(cons map[string]interface{}) ([]PipeLine, error) {
	orm := GetSlave()

	pipeLines := make([]PipeLine, 0)
	err := orm.Where(cons).Find(&pipeLines)
	if err != nil {
		return nil, err
	}
	return pipeLines, err
}

func PipeLineUpdate(id int64, p *PipeLine) error {
	orm := GetMaster()
	affected, err := orm.ID(id).Update(p)
	if affected == 0 {
		return fmt.Errorf("update failed, affected = 0")
	}
	return err
}
func PipeLineDelete(p *PipeLine) error {
	orm := GetMaster()
	affected, err := orm.Delete(p)
	if affected == 0 {
		return fmt.Errorf("delete failed, affected = 0")
	}
	return err
}

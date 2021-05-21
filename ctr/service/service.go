package service

import (
	"encoding/json"
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service/nginx"
	"github.com/tufanbarisyildirim/gonginx"
	"net/http"
	"strconv"
	"strings"
)

func CreateService(user *model.User, rd *entity.CreateServiceRequestData) (int, string, error) {
	service, has := model.ServiceOne(map[string]interface{}{"tree_id": *rd.TreeId})
	if has {
		return http.StatusBadRequest, "", fmt.Errorf("已存在预案配置")
	}

	serviceTree, has := model.TreeOne(map[string]interface{}{"id": *rd.TreeId})
	if !has {
		return http.StatusBadRequest, "", fmt.Errorf("没有id=%v的服务树信息", *rd.TreeId)
	}
	if serviceTree.Type != model.TreeTypeService {
		return http.StatusBadRequest, "", fmt.Errorf("服务树类型错误")
	}

	ss := strings.Split(serviceTree.Un, ".")
	rootUn := ss[len(ss)-1]
	rootTree, has := model.TreeOne(map[string]interface{}{"un": rootUn})
	if !has {
		return http.StatusBadRequest, "", fmt.Errorf("没有找到根节点(type=namespace)，请检查配置")
	}

	router, has := model.RouterOne(map[string]interface{}{"ns_id": rootTree.Id})
	if !has {
		return http.StatusBadRequest, "", fmt.Errorf("根节点（un=%v）没有配置接入层信息", rootTree.Un)
	}

	routerHost, has := model.HostOne(map[string]interface{}{"id": router.HostId})
	if !has {
		return http.StatusBadRequest, "", fmt.Errorf("没有找到接入层对应的机器，请检查配置")
	}

	pipelineList, err := model.PipeLineList(map[string]interface{}{"tree_id": serviceTree.Id})
	if err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("su=%v没有对应的流水线", serviceTree.Un)
	}

	for _, pipeline := range pipelineList {
		if pipeline.Status != model.PipeLineStatusRunning {
			return http.StatusBadRequest, "", fmt.Errorf("此service下存在未执行的流水线")
		}
	}

	flowMap := map[string]int{} // "cluster_id": flow_percent
	if err := json.Unmarshal([]byte(*rd.FlowMap), &flowMap); err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("预案配置错误,err:%v", err)
	}

	count := 0
	clusterCount := 0
	for _, i := range flowMap {
		count += i
		clusterCount += 1
	}
	if count != 100 {
		return http.StatusBadRequest, "", fmt.Errorf("预案配置错误，流量比例和不得超过100，当前：%v", count)
	}
	if clusterCount != len(pipelineList) {
		return http.StatusBadRequest, "", fmt.Errorf("存在未创建的流水线")
	}

		userTree, has := model.UserTreeOne(map[string]interface{}{"tree_id": serviceTree.Id, "user_id": user.Id})
	if !has {
		return http.StatusUnauthorized, "", fmt.Errorf("没有权限")
	}
	if userTree.Rights < model.PermReadWrite {
		return http.StatusUnauthorized, "", fmt.Errorf("权限不足，请联系管理员添加更高级别的权限")
	}

	clusterFlowMap := map[int64]int{}
	for ks := range flowMap {
		k, _ := strconv.Atoi(ks)
		_, has := model.ClusterOne(map[string]interface{}{"tree_id": k})
		if !has {
			return http.StatusBadRequest, "", fmt.Errorf("clusterId=%v没有配置", k)
		}
		clusterTree, has := model.TreeOne(map[string]interface{}{"id": k})
		if !has {
			return http.StatusBadRequest, "", fmt.Errorf("没有clusterId=%v的服务树树节点", k)
		}
		clusterFlowMap[clusterTree.Id] = flowMap[ks]
	}

	upstreamServers := make([]*gonginx.UpstreamServer, 0)
	for _, pipeLine := range pipelineList {
		serverAddr := fmt.Sprintf("%s:%d", pipeLine.RouterIp, pipeLine.RouterPort)
		upstreamServer := gonginx.NewUpstreamServer(&gonginx.UpstreamServer{
			Address: serverAddr,
			Flags:   []string{},
			Parameters: map[string]string{
				"weight": strconv.Itoa(clusterFlowMap[pipeLine.ClusterId]),
			},
		})
		upstreamServers = append(upstreamServers, upstreamServer)
	}
	upstream, _ := gonginx.NewUpstream(&gonginx.Upstream{
		UpstreamName:    fmt.Sprintf("%s_service", serviceTree.Un),
		UpstreamServers: upstreamServers,
		Directives:      nil,
	})

	listenPort := routerHost.GetUnusedPort()

	proxyPass := &gonginx.Directive{Name: "proxy_pass",
		Parameters: []string{fmt.Sprintf("http://%s_service", serviceTree.Un)}}
	unDirective := &gonginx.Directive{
		Name:       "set",
		Parameters: []string{"$un", serviceTree.Un},
	}
	locationDirective := &gonginx.Directive{
		Block: &gonginx.Block{
			Directives: []gonginx.IDirective{proxyPass, unDirective},
		},
		Name:       "location",
		Parameters: []string{"/"},
	}
	listenDirective := &gonginx.Directive{
		Name:       "listen",
		Parameters: []string{strconv.Itoa(listenPort)},
	}
	servernameDirective := &gonginx.Directive{
		Name:       "server_name",
		Parameters: []string{serviceTree.Un},
	}
	directive := &gonginx.Directive{
		Block: &gonginx.Block{
			Directives: []gonginx.IDirective{listenDirective, servernameDirective, locationDirective}},
		Name:       "server",
		Parameters: nil,
	}
	server, _ := gonginx.NewServer(&gonginx.Server{
		Block: directive.GetBlock(),
	})

	upstreamContent := gonginx.DumpDirective(upstream, gonginx.IndentedStyle)
	serverContent := gonginx.DumpDirective(server, gonginx.IndentedStyle)
	content := fmt.Sprintf("%s\n%s", upstreamContent, serverContent)
	configFileName := fmt.Sprintf("%s.conf", serviceTree.Un)
	ngx := nginx.New(routerHost)
	ngx.ConfWrite(configFileName, content)
	_ = ngx.Reload()

	service.RouterPort = listenPort
	service.RouterIp = routerHost.IP
	service.FlowMap = *rd.FlowMap
	service.TreeId = *rd.TreeId
	service.ConfigContent = content
	_ = model.ServiceAdd(service)

	return http.StatusOK, content, nil
}

func GetService(user *model.User, rd *entity.ServiceGetRequestData) (int, *model.ServiceInfo, error) {
	service, has := model.ServiceOne(map[string]interface{}{"tree_id": *rd.TreeId})
	if !has {
		return http.StatusNotFound, nil, fmt.Errorf("此服务树没有对应的预案配置信息")
	}

	si, err := service.GetServiceInfo()
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("获取预案配置失败,err=%v", err)
	}

	return http.StatusOK, si, nil
}

func DeleteService(user *model.User, rd *entity.ServiceGetRequestData) (int, error) {
	svc, has := model.ServiceOne(map[string]interface{}{"tree_id": *rd.TreeId})
	if !has {
		return http.StatusNotFound, fmt.Errorf("没有预案配置信息")
	}

	svcTree, has := model.TreeOne(map[string]interface{}{"id": svc.TreeId})
	if !has {
		return http.StatusNotFound, fmt.Errorf("没有对应的服务树节点")
	}

	ss := strings.Split(svcTree.Un, ".")
	rootUn := ss[len(ss)-1]
	rootTree, has := model.TreeOne(map[string]interface{}{"un": rootUn})
	if !has {
		return http.StatusBadRequest, fmt.Errorf("没有找到根节点(type=namespace)，请检查配置")
	}

	router, has := model.RouterOne(map[string]interface{}{"ns_id": rootTree.Id})
	if !has {
		return http.StatusBadRequest, fmt.Errorf("根节点（un=%v）没有配置接入层信息", rootTree.Un)
	}

	routerHost, has := model.HostOne(map[string]interface{}{"id": router.HostId})
	if !has {
		return http.StatusBadRequest, fmt.Errorf("没有找到接入层对应的机器，请检查配置")
	}

	configFileName := fmt.Sprintf("%s.conf", svcTree.Un)
	ngx := nginx.New(routerHost)
	ngx.ConfDelete(configFileName)
	_ = ngx.Reload()

	_ = model.ServiceDelete(svc)
	fmt.Println("delete service")

	return http.StatusOK, nil
}

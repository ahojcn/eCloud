package service

import (
	"encoding/json"
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service/nginx"
	"github.com/parnurzeal/gorequest"
	"github.com/tufanbarisyildirim/gonginx"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func PipeLineList(user *model.User, rd *entity.PipeLineListRequestData) (int, []model.PipeLineInfo, error) {
	tree, _ := model.TreeOne(map[string]interface{}{"id": *rd.TreeID})
	if tree.Type != model.TreeTypeService {
		return http.StatusBadRequest, nil, fmt.Errorf("类型错误的节点")
	}

	pipelinelist, err := model.PipeLineList(map[string]interface{}{"tree_id": tree.Id})
	if err != nil {
		return http.StatusBadRequest, nil, fmt.Errorf("获取流水线列表失败，err=%v", err)
	}

	var res = []model.PipeLineInfo{}
	for _, pipeline := range pipelinelist {
		res = append(res, pipeline.GetPipeLineInfo())
	}

	return http.StatusOK, res, nil
}

func PipeLineCreate(user *model.User, rd *entity.PipeLineCreateRequestData) (int, *model.PipeLine, error) {
	_, has := model.PipeLineOne(map[string]interface{}{"tree_id": *rd.TreeID, "cluster_id": *rd.ClusterID})
	if has {
		return http.StatusBadRequest, nil, fmt.Errorf("已存在关联的流水线")
	}
	pipeline := model.PipeLine{
		TreeId:          *rd.TreeID,
		ClusterId:       *rd.ClusterID,
		ContainerImage:  rd.ContainerImage,
		StatusMsg:       "初始化",
		AliveMethod:     rd.AliveMethod,
		AliveURI:        rd.AliveURI,
		AliveReqQuery:   rd.AliveReqQuery,
		AliveReqBody:    rd.AliveReqBody,
		AliveReqHeader:  rd.AliveReqHeader,
		AliveRespStatus: rd.AliveRespStatus,
		AliveRespBody:   rd.AliveRespBody,
	}
	if err := model.PipeLineAdd(&pipeline); err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("创建流水线失败, err=%v", err)
	}
	return http.StatusOK, &pipeline, nil
}

func PipeLineRun(user *model.User, rd *entity.PipeLineRunRequestData) (int, string, error) {
	p, has := model.PipeLineOne(map[string]interface{}{"id": rd.Id})
	if !has {
		return http.StatusBadRequest, "", fmt.Errorf("没有id=%v的流水线信息", rd.Id)
	}

	hosts, err := BuildImage(p)
	if err != nil {
		return http.StatusInternalServerError, "", fmt.Errorf("流水线【构建镜像】失败，err=%v", err)
	}

	containers, err := RunContainer(p, hosts)
	if err != nil {
		return http.StatusInternalServerError, "", fmt.Errorf("流水线【运行容器】失败，err=%v", err)
	}

	if err = AliveTest(p, containers); err != nil {
		return http.StatusInternalServerError, "", fmt.Errorf("流水线【存活测试】失败，err=%v", err)
	}

	routerConfig := MakeRouter(p, containers)
	return http.StatusOK, routerConfig, nil
}

func BuildImage(p *model.PipeLine) ([]model.Host, error) {
	p.Status = model.PipeLineStatusBuildImage
	p.StatusMsg = model.PipeLineStatusMsg[model.PipeLineStatusBuildImage]
	_ = model.PipeLineUpdate(p.Id, p)

	hostList, _ := model.HostList(map[string]interface{}{})
	if len(hostList) <= 0 {
		return nil, fmt.Errorf("没有可用主机，请在资源中添加")
	}
	rand.Seed(time.Now().Unix())
	cluster, _ := model.ClusterOne(map[string]interface{}{"tree_id": p.ClusterId})
	pullImageCmd := fmt.Sprintf("docker pull %s", p.ContainerImage)
	hostSelected := []model.Host{}
	for i := cluster.CurrentClusterNum; i < cluster.ClusterNum; i++ {
		hostIndex := rand.Intn(len(hostList))
		_, _ = hostList[hostIndex].RunCmd(pullImageCmd, time.Duration(0))
		hostSelected = append(hostSelected, hostList[hostIndex])
	}

	return hostSelected, nil
}
func RunContainer(p *model.PipeLine, hosts []model.Host) ([]model.Container, error) {
	p.Status = model.PipeLineStatusRunContainer
	p.StatusMsg = model.PipeLineStatusMsg[model.PipeLineStatusRunContainer]
	_ = model.PipeLineUpdate(p.Id, p)

	cluster, _ := model.ClusterOne(map[string]interface{}{"tree_id": p.ClusterId})
	containers := []model.Container{}

	for _, host := range hosts {
		hostPort := host.GetUnusedPort()
		runContainerCmd := fmt.Sprintf("docker run -d -p %d:%d %s", hostPort, cluster.ContainerPort, p.ContainerImage)
		containerIdRes, _ := host.RunCmd(runContainerCmd, time.Duration(0))
		containerIpCmd := fmt.Sprintf("docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' %s", containerIdRes)
		ipRes, _ := host.RunCmd(containerIpCmd, time.Duration(0))

		cluster.CurrentClusterNum += 1
		_ = model.ClusterUpdate(cluster.Id, cluster)
		container := model.Container{
			HostId:        host.Id,
			ContainerId:   strings.Split(containerIdRes, "\n")[0],
			ContainerIp:   strings.Split(ipRes, "\n")[0],
			ContainerPort: cluster.ContainerPort,
			HostPort:      hostPort,
		}
		_ = model.ContainerAdd(&container)
		_ = model.ClusterContainerAdd(&model.ClusterContainer{ClusterId: cluster.Id, ContainerId: container.Id})
		containers = append(containers, container)
	}

	return containers, nil
}
func AliveTest(p *model.PipeLine, containers []model.Container) error {
	p.Status = model.PipeLineStatusAliveTest
	p.StatusMsg = model.PipeLineStatusMsg[model.PipeLineStatusAliveTest]
	_ = model.PipeLineUpdate(p.Id, p)

	headers := http.Header{}
	reqHeaders := map[string]string{}
	_ = json.Unmarshal([]byte(p.AliveReqHeader), &reqHeaders)
	for key, val := range reqHeaders {
		headers.Set(key, val)
	}
	body := ioutil.NopCloser(strings.NewReader(p.AliveReqBody))

	for _, container := range containers {
		h, _ := container.GetHost()
		host := fmt.Sprintf("%s:%d", h.IP, container.HostPort)

		r := &http.Request{
			Method: p.AliveMethod,
			URL:    &url.URL{Host: host, Path: p.AliveURI, Scheme: "http"},
			Header: headers,
			Body:   body,
		}
		resp, err := gorequest.New().Client.Do(r)
		if err != nil {
			return fmt.Errorf("存活测试失败：err=%v", err)
		}
		if resp.StatusCode != p.AliveRespStatus {
			return fmt.Errorf("存活测试失败：hostIP=%v, containerPort=%v", h.Id, container.HostPort)
		}
	}

	return nil
}
func MakeRouter(p *model.PipeLine, containers []model.Container) string {
	p.Status = model.PipeLineStatusRouter
	p.StatusMsg = model.PipeLineStatusMsg[model.PipeLineStatusRouter]
	_ = model.PipeLineUpdate(p.Id, p)

	serviceTree, _ := model.TreeOne(map[string]interface{}{"id": p.TreeId})
	unSplit := strings.Split(serviceTree.Un, ".")
	nsUn := unSplit[len(unSplit)-1]
	namespaceTree, _ := model.TreeOne(map[string]interface{}{"un": nsUn})
	clusterTree, _ := model.TreeOne(map[string]interface{}{"id": p.ClusterId})

	r, _ := model.RouterOne(map[string]interface{}{"ns_id": namespaceTree.Id})
	routerHost, _ := r.GetHostInfo()

	// 生成 nginx 配置，对应一个 cluster_un.conf
	ngx := nginx.New(routerHost)
	upstreamServer := make([]*gonginx.UpstreamServer, 0)
	for _, c := range containers {
		ch, _ := c.GetHost()
		ups := gonginx.NewUpstreamServer(&gonginx.UpstreamServer{
			Address: fmt.Sprintf("%s:%d", ch.IP, c.HostPort),
			Parameters: map[string]string{
				"weight": "1",
			},
		})
		upstreamServer = append(upstreamServer, ups)
	}
	up, _ := gonginx.NewUpstream(&gonginx.Upstream{
		UpstreamName:    fmt.Sprintf("%s_cluster", clusterTree.Un),
		UpstreamServers: upstreamServer,
		Directives:      nil,
	})
	upstreamContent := gonginx.DumpDirective(up, gonginx.IndentedStyle)

	proxyPass := &gonginx.Directive{Name: "proxy_pass",
		Parameters: []string{fmt.Sprintf("http://%s_cluster", clusterTree.Un)}}
	locationDirective := &gonginx.Directive{
		Block: &gonginx.Block{
			Directives: []gonginx.IDirective{proxyPass},
		},
		Name:       "location",
		Parameters: []string{"/"},
	}

	listenDirective := &gonginx.Directive{
		Name:       "listen",
		Parameters: []string{strconv.Itoa(routerHost.GetUnusedPort())},
	}
	servernameDirective := &gonginx.Directive{
		Name:       "server_name",
		Parameters: []string{clusterTree.Un},
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

	serverContent := gonginx.DumpDirective(server, gonginx.IndentedStyle)

	content := fmt.Sprintf("%s\n%s", upstreamContent, serverContent)
	ngx.ConfWrite(fmt.Sprintf("%s.conf", clusterTree.Un), content)

	_ = ngx.Reload()
	p.Status = model.PipeLineStatusRunning
	p.StatusMsg = model.PipeLineStatusMsg[model.PipeLineStatusRunning]
	_ = model.PipeLineUpdate(p.Id, p)
	return content
}

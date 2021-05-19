package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/entity"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service"
	"github.com/ahojcn/ecloud/ctr/service/nginx"
	"github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
	"strings"
	"time"
)

func testMyNginx() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	fileNames, err := nginx.New(h).ConfList()
	fmt.Println(fileNames, err)

	fmt.Println(nginx.New(h).ConfContent("n-1.conf"))

	fmt.Println(nginx.New(h).Status())
}

func testGoNginxUpstream() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	confStr, _ := nginx.New(h).ConfRoot()
	p := parser.NewStringParser(confStr)
	conf := p.Parse()
	aa := conf.FindDirectives("http")

	fmt.Println(aa[0].GetBlock(), aa[0].GetName(), aa[0].GetParameters())

	ups1 := gonginx.NewUpstreamServer(&gonginx.UpstreamServer{
		Address: "127.0.0.1:443",
		Flags:   []string{"down"},
		Parameters: map[string]string{
			"weight": "5",
		},
	})
	ups2 := gonginx.NewUpstreamServer(&gonginx.UpstreamServer{
		Address: "127.0.0.1:443",
		Flags:   []string{"backup"},
		Parameters: map[string]string{
			"weight": "5",
		},
	})
	ups := make([]*gonginx.UpstreamServer, 0)
	ups = append(ups, ups1)
	ups = append(ups, ups2)

	up, err := gonginx.NewUpstream(&gonginx.Upstream{
		UpstreamName:    "my_backend",
		UpstreamServers: ups,
		Directives:      nil,
	})
	fmt.Println(up, err)
	fmt.Println(gonginx.DumpBlock(up.GetBlock(), gonginx.IndentedStyle))
	fmt.Println(gonginx.DumpDirective(up, gonginx.IndentedStyle))
}

func testGoNginxServer() {
	proxyPass := &gonginx.Directive{Name: "proxy_pass", Parameters: []string{"http://xxxx.com"}}
	locationDirective := &gonginx.Directive{
		Block: &gonginx.Block{
			Directives: []gonginx.IDirective{proxyPass},
		},
		Name:       "location",
		Parameters: []string{"/"},
	}

	listenDirective := &gonginx.Directive{
		Name:       "listen",
		Parameters: []string{"18080"},
	}
	servernameDirective := &gonginx.Directive{
		Name:       "server_name",
		Parameters: []string{"cluster_un"},
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

	fmt.Println(gonginx.DumpDirective(server, gonginx.IndentedStyle))
}

func testRunContainer() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	res, _ := h.RunCmd("docker run -d luksa/kubia", time.Duration(0))
	ipCmd := fmt.Sprintf("docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' %s", res)
	ip, _ := h.RunCmd(ipCmd, time.Duration(0))
	fmt.Println(len(strings.Split(ip, "\n")))
}

func testPipeLineRun() {
	id := new(int64)
	*id = 3
	rd := entity.PipeLineRunRequestData{Id: id}
	user, _ := model.UserOne(map[string]interface{}{"id": 8})
	fmt.Println(service.PipeLineRun(user, &rd))
}

func main() {
	//testMyNginx()
	//testGoNginxUpstream()
	//testRunContainer()
	//testGoNginxServer()
	testPipeLineRun()
}

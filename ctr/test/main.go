package main

import (
	"fmt"
	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/ahojcn/ecloud/ctr/service/nginx"
	"github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
)

func testMyNginx() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	fileNames, err := nginx.New(h).ConfList()
	fmt.Println(fileNames, err)

	fmt.Println(nginx.New(h).ConfContent("n-1.conf"))

	fmt.Println(nginx.New(h).Status())
}

func testGoNginx() {
	h, _ := model.HostOne(map[string]interface{}{"id": 13})
	confStr, _ := nginx.New(h).ConfRoot()
	p := parser.NewStringParser(confStr)
	conf := p.Parse()
	aa := conf.FindDirectives("http")

	fmt.Println(aa[0].GetBlock(), aa[0].GetName(), aa[0].GetParameters())

	ups1 := gonginx.NewUpstreamServer(&gonginx.UpstreamServer{
		Address:    "127.0.0.1:443",
		Flags:      []string{"down"},
		Parameters: map[string]string{
			"weight": "5",
		},
	})
	ups2 := gonginx.NewUpstreamServer(&gonginx.UpstreamServer{
		Address:    "127.0.0.1:443",
		Flags:      []string{"backup"},
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
	//up.Directives = append(up.Directives, &gonginx.Directive{
	//	Block:      nil,
	//	Name:       "ip_hash",
	//	Parameters: nil,
	//})
	fmt.Println(gonginx.DumpBlock(up.GetBlock(), gonginx.IndentedStyle))
	fmt.Println(gonginx.DumpDirective(up, gonginx.IndentedStyle))
}

func main() {
	//testMyNginx()
	testGoNginx()
}

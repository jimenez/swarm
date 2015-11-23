package main

import (
	"flag"

	"github.com/docker/swarm/cli"
	_ "github.com/docker/swarm/discovery/file"
	_ "github.com/docker/swarm/discovery/kv"
	_ "github.com/docker/swarm/discovery/nodes"
	_ "github.com/docker/swarm/discovery/token"
	_ "github.com/golang/glog"
)

func main() {
	//	flag.Lookup("v").Value.Set("4")
	flag.Lookup("logtostderr").Value.Set("true")
	cli.Run()
}

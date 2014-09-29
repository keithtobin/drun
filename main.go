package main

import (
	"flag"
	"github.com/golang/glog"	
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
)

var (
	config             = flag.String("config", "", "Path to the config file or directory of files")
	enableServer       = flag.Bool("enable_server", true, "Enable the info server")
	registryBurst      = flag.Int("registry_burst", 10, "M")


)

func main() {
	flag.Parse()
	util.InitLogs()
	defer util.FlushLogs()

	glog.Fatal("Couldn't connect to docker.")

}

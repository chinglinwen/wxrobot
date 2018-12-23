package main

import (
	"flag"

	"github.com/chinglinwen/wxrobot/service"
	"github.com/songtianyi/rrframework/logs"
)

var (
	backendurl = flag.String("url", "http://localhost:4000", "backend url")
	port       = flag.String("p", ":50051", "default grpc listenging port")
	grouplist  = flag.String("groups", "", "allowed group list(eg. group1,group2)")
	groupon    = flag.Bool("g", false, "turn on group or not")
)

func main() {
	logs.Info("starting...")
	logs.Info("allowed groups: ", grouplist)

	flag.Parse()
	service.SetBackendUrl(*backendurl)

	// listening grpc request
	logs.Info("start listening on port: ", *port)

	GrpcServe()
}

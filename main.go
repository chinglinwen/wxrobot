package main

import (
	"flag"
	"time"

	"github.com/chinglinwen/wxrobot/plugins/all"
	"github.com/chinglinwen/wxrobot/service"

	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/config"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/wxweb"
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
	all.SetBackendUrl(*backendurl)

	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}
	service.SetSession(session)
	service.SetListeningPort(*port)

	// load plugins for this session
	all.Register(session)
	switcher.Register(session)
	config.Register(session)

	session.HandlerRegister.EnableByName("switcher")

	// listening grpc request
	logs.Info("start listening on port: ", *port)
	go func() {
		service.GrpcServe()
	}()

	for {
		if err := session.LoginAndServe(false); err != nil {
			logs.Error("session exit, %s", err)
			for i := 0; i < 3; i++ {
				logs.Info("trying re-login with cache")
				if err := session.LoginAndServe(true); err != nil {
					logs.Error("re-login error, %s", err)
				}
				time.Sleep(3 * time.Second)
			}
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				logs.Error("create new sesion failed, %s", err)
				break
			}
		} else {
			logs.Info("closed by user")
			break
		}
	}
}

package service

import (
	"log"
	"time"

	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
)

var (
	session       *wxweb.Session
	SessionReady  bool
	readyReceiver = "xiubi" //notify ready
)

func SetReadyReceiver(receiver string) {
	readyReceiver = receiver
	return
}

func init() {
	logs.Info("creating session...")
	// create session
	var err error
	session, err = wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	// load plugins for this session
	Register(session)
	// switcher.Register(session)
	// config.Register(session)

	// session.HandlerRegister.EnableByName("switcher")
	go func() {
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
	}()

	wait()
}

func wait() {
	for session.Cm == nil {
		time.Sleep(3 * time.Second)
		log.Println("waiting session...")
	}
	SessionReady = true
	log.Println("wx session logged in.")

	SendText(readyReceiver, "xiubi is ready : )\n")
}

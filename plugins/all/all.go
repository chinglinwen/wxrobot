package all

import (
	"bytes"
	"encoding/json"
	"fmt"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
)

var (
	backendurl = "http://clwen.com:4000"
	bodytype   = "application/json"
)

/*
   //msg types
   MSG_TEXT        = 1     // text message
   MSG_IMG         = 3     // image message
   MSG_VOICE       = 34    // voice message
   MSG_FV          = 37    // friend verification message
   MSG_PF          = 40    // POSSIBLEFRIEND_MSG
   MSG_SCC         = 42    // shared contact card
   MSG_VIDEO       = 43    // video message
   MSG_EMOTION     = 47    // gif
   MSG_LOCATION    = 48    // location message
   MSG_LINK        = 49    // shared link message
   MSG_VOIP        = 50    // VOIPMSG
   MSG_INIT        = 51    // wechat init message
   MSG_VOIPNOTIFY  = 52    // VOIPNOTIFY
   MSG_VOIPINVITE  = 53    // VOIPINVITE
   MSG_SHORT_VIDEO = 62    // short video message
   MSG_SYSNOTICE   = 9999  // SYSNOTICE
   MSG_SYS         = 10000 // system message
   MSG_WITHDRAW    = 10002 // withdraw notification message
*/

// register plugin
func Register(session *wxweb.Session) {
	doregister(session, wxweb.MSG_TEXT, "text")
	doregister(session, wxweb.MSG_IMG, "img")
	doregister(session, wxweb.MSG_VOICE, "voice")
	doregister(session, wxweb.MSG_EMOTION, "gif")
	doregister(session, wxweb.MSG_LINK, "link")
	doregister(session, wxweb.MSG_SYSNOTICE, "sysnotice")
	doregister(session, wxweb.MSG_SYS, "sys")
	doregister(session, wxweb.MSG_WITHDRAW, "withdraw")
	doregister(session, wxweb.MSG_INIT, "init")
}

func doregister(session *wxweb.Session, key int, name string) {
	err := session.HandlerRegister.Add(key, wxweb.Handler(autoReply), name)
	if err != nil {
		logs.Error(err)
	}
	err = session.HandlerRegister.EnableByName(name)
	if err != nil {
		logs.Error(err)
	}
}

func autoReply(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	if msg.MsgType == 51 {
		// skip init
		return
	}
	logs.Info("from: ", msg.FromUserName, "to: ", msg.ToUserName)
	logs.Info("msg: ", msg.Content)

	var reply string

	json, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		logs.Error(err)
	}

	resp, err := retryablehttp.Post(backendurl, bodytype, bytes.NewReader(json))
	if err != nil {
		logs.Error(err)
		reply = "err: " + err.Error()
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	reply = buf.String()

	if msg.FromUserName == msg.ToUserName {
		fmt.Println("it's from myself")
		return
	}
	session.SendText("this is from robot:\n  "+reply, session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
}

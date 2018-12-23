package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/tidwall/gjson"
)

var (
	defaulturl     = "http://localhost:4000"
	defaultBackend = newBackend(defaulturl)
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
func Register(session *wxweb.Session, options ...backendOption) {
	for _, op := range options {
		op(defaultBackend)
	}

	doregister(session, wxweb.MSG_TEXT, "text")
	doregister(session, wxweb.MSG_IMG, "img")
	//doregister(session, wxweb.MSG_VOICE, "voice")
	//doregister(session, wxweb.MSG_EMOTION, "gif")
	//doregister(session, wxweb.MSG_LINK, "link")
	//doregister(session, wxweb.MSG_SYSNOTICE, "sysnotice")
	//doregister(session, wxweb.MSG_SYS, "sys")
	//doregister(session, wxweb.MSG_WITHDRAW, "withdraw")
	//doregister(session, wxweb.MSG_INIT, "init")
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

type option struct {
	url string
}

type backend struct {
	url string
}

func newBackend(url string) *backend {
	return &backend{url: url}
}

func SetDefaultBackend(url string) {
	defaultBackend.url = url
	return
}

type backendOption func(*backend)

// SetBackendUrl change the default backend
// default is "http://localhost:4000"
func SetBackendUrl(url string) backendOption {
	return func(b *backend) {
		b.url = url
	}
}

func autoReply(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	if msg.MsgType == 51 {
		// skip init
		return
	}
	logs.Info("from: ", msg.FromUserName, "to: ", msg.ToUserName, "real: ", wxweb.RealTargetUserName(session, msg))
	logs.Info("msg: ", msg.Content)

	// if msg.IsGroup == true {
	// 	return
	// }

	json, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		logs.Error(err)
	}

	reply, err := request(defaultBackend.url, json)

	//fmt.Println("it's from myself"), it's just dosen't work
	if msg.FromUserName == msg.ToUserName {
		return
	}

	//  skip non-command
	if reply == "" {
		return
	}

	replyType := gjson.Get(reply, "type").String()
	replyData := gjson.Get(reply, "data").String()
	replyErr := gjson.Get(reply, "error").String()

	// remove space at the left and right
	replyData = strings.Trim(replyData, " \t")

	if replyData == "" && replyErr == "" {
		return
	}

	// log part of reply only
	var n int
	if len(replyData) < 10 {
		n = len(replyData)
	}
	logs.Info("got:", replyType, len(replyData), replyData[0:n], "err", replyErr)

	if replyErr != "" {
		session.SendText("robot err:\n  "+replyErr, session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
		return
	}

	if replyType == "image" {
		logs.Info("got image reply")

		decoded, err := base64.StdEncoding.DecodeString(replyData)
		if err != nil {
			session.SendText("robot err:\n  "+err.Error(), session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
			return
		}
		//todo: it's still error: BaseResponse.Ret=1
		session.SendImgFromBytes(decoded, "http://wx2.sinaimg.cn/mw1024/9d52c073gy1foxoszeu10j20sg0zkk4y.jpg", session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
		return
	}

	logs.Info("text reply:", replyData)
	_, _, err = session.SendText("robot says:\n  "+replyData, session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
	if err != nil {
		logs.Error("send error:", err)
		return
	}
	logs.Info("send ok")
}

func request(url string, body []byte) (reply string, err error) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	c := retryablehttp.NewClient()
	c.RetryMax = 2

	//resp, err := c.Post(url, "application/json", bytes.NewReader(body))
	resp, err := c.Post(url, "", bytes.NewReader(body))
	if err != nil {
		logs.Error(err)
		reply = "err: " + err.Error()
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	reply = buf.String()

	return
}

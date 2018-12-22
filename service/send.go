package service

import (
	"fmt"
	"log"
	"time"

	"github.com/songtianyi/wechat-go/wxweb"
)

// set session
var (
	session      *wxweb.Session
	sessionReady bool
)

func SetSession(s *wxweb.Session) {
	session = s
	go func() {
		wait()
	}()
	return
}

func wait() {
	for session.Cm == nil {
		time.Sleep(3 * time.Second)
		log.Println("waiting session...")
	}
	sessionReady = true
	log.Println("wx session logged in.")

	SendText(readyReceiver, "xiubi is ready : )\n")
}

func ShowGroups() {
	users := session.Cm.GetGroupContacts()
	for i, v := range users {
		fmt.Println(i, v.UserName, v.NickName, v.PYInitial, v.PYQuanPin,
			v.RemarkName, v.RemarkPYInitial, v.RemarkPYQuanPin, v.DisplayName)
	}
}

// eg.
//   text := "hi there, robot here\n"
//   sendtext("xiubi", text)         //ok, it's a remarkname
//   sendtext("三", "hello1by-name") //ok,  it's a group name
//
//   sendtext("filehelper", text)   //not ok
//   sendtext("Nickname", text)     //not ok
func SendText(name, text string) error {
	err := checkBeforeSend(name, text)
	if err != nil {
		return err
	}

	users := session.Cm.GetContactsByName(name)
	log.Printf("got %v users for name: %v\n", len(users), name)
	for _, v := range users {
		log.Printf("name: %v, text: %v\n", name, text)
		session.SendText(text, session.Bot.UserName, v.UserName)
	}
	return nil
}

// eg. sendtextquanpin("san", "hello1by-py")
func SendTextQuanPin(name, text string) error {
	err := checkBeforeSend(name, text)
	if err != nil {
		return err
	}

	if sessionReady != true {
		return fmt.Errorf("it may not logged in")
	}
	log.Printf("name: %v, text: %v\n", name, text)
	user := session.Cm.GetContactByPYQuanPin(name)
	session.SendText(text, session.Bot.UserName, user.UserName)
	return nil
}

func checkBeforeSend(name, text string) error {
	if name == "" {
		return fmt.Errorf("empty name to send")
	}
	if text == "" {
		return fmt.Errorf("empty text to send")
	}
	if sessionReady != true {
		return fmt.Errorf("it may not logged in")
	}
	return nil
}

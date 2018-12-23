package service

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/songtianyi/wechat-go/wxweb"
)

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

func SendImage(data []byte, url, name string, session *wxweb.Session) {
	log.Printf("%v bytes,url: %v, name: %v, session: %v", len(data), url, name, session)

	users := session.Cm.GetContactsByName(name)
	log.Printf("got %v users for name: %v\n", len(users), name)
	for _, v := range users {
		log.Printf("try send to name: %v\n", name)

		decoded, _ := base64.StdEncoding.DecodeString(string(data))
		// if err != nil {
		// 	log.Println("base64")
		// 	//session.SendText("robot err:\n  "+err.Error(), session.Bot.UserName, v.UserName)
		// 	return
		// }
		//todo: it's still error: BaseResponse.Ret=1
		session.SendImgFromBytes(decoded, url, session.Bot.UserName, v.UserName)
		log.Println("sended")
	}
}

// eg. sendtextquanpin("san", "hello1by-py")
func SendTextQuanPin(name, text string) error {
	err := checkBeforeSend(name, text)
	if err != nil {
		return err
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
	return nil
}

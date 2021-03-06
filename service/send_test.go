package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

// func TestMain(m *testing.M) {
// 	var err error
// 	session, err = wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
// 	checkerr(err)

// 	go func() {
// 		m.Run()
// 	}()
// 	err = session.LoginAndServe(false)

// 	log.Println("start run test...")
// 	os.Exit(0)
// }

//still somewhat error, it's hard to test
func init() {
	//wait login

}

// error: [E] BaseResponse.Ret=1
func TestSendImage(t *testing.T) {
	wait()

	//url := "http://wx2.sinaimg.cn/mw1024/9d52c073gy1foxoszeu10j20sg0zkk4y.jpg"
	url := "https://honeysanime.com/wp-content/uploads/2016/05/Chi%E2%80%99s-Sweet-Home-Wallpaper.jpg"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	name := "休比"

	aa := base64.StdEncoding.EncodeToString(data)

	SendImage([]byte(aa), url, name, session)
}

func TestSend(t *testing.T) {
	for session.Cm == nil {
		time.Sleep(3 * time.Second)
		log.Println("waiting session...")
	}
	//spew.Dump("user", session.Cm.GetContactByUserName("xiubi-dola"))
	//users := session.Cm.GetAll()
	//for i, v := range users {
	//	fmt.Println(i, v.UserName, v.NickName, v.PYInitial, v.PYQuanPin)
	//}

	users := session.Cm.GetGroupContacts()
	for i, v := range users {
		fmt.Println(i, v.UserName, v.NickName, v.PYInitial, v.PYQuanPin,
			v.RemarkName, v.RemarkPYInitial, v.RemarkPYQuanPin, v.DisplayName)
	}

	text := "hi there, robot here\n"
	sendtext("xiubi", text)        //it's works
	sendtext("三", "hello1by-name") //ok

	//sendtext("filehelper", text) //not ok
	//sendtext("Chinglin", text) //this does not work
	//sendtextquanpin("san", "hello1by-py")  //it's works
}

func checkerr(err error) {
	if err != nil {
		log.Fatal("err: ", err)
	}
}

func sendtext(name, text string) {
	users := session.Cm.GetContactsByName(name)
	fmt.Printf("got %v users for name\n", len(users))
	for _, v := range users {
		session.SendText(text, session.Bot.UserName, v.UserName)
	}
}

func sendtextquanpin(name, text string) {
	user := session.Cm.GetContactByPYQuanPin(name)
	session.SendText(text, session.Bot.UserName, user.UserName)
}

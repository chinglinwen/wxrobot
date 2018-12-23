package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tidwall/gjson"
)

var (
	robot = []byte(`
{
  "IsGroup": false,
  "MsgId": "8927133500120292699",
  "Content": "Robot",
  "FromUserName": "@a99651a071b3adfe9d4fea18915cb09e",
  "ToUserName": "@fe447f00f7ef71089b35244b706fcbd22e9ed44855bfa6fc7b3dba19ff5ee8bc",
  "Who": "@a99651a071b3adfe9d4fea18915cb09e",
  "MsgType": 1,
  "SubType": 0,
  "OriginContent": "robot",
  "At": "",
  "Url": "",
  "RecommendInfo": null
}
`)

	girl = []byte(`
{
  "IsGroup": false,
  "MsgId": "8232189154394217299",
  "Content": "美女",
  "FromUserName": "@61d8b12ebd9c155adab0b64d82619ee0",
  "ToUserName": "@8738a68fea0c6cef74e985e18493d719e8061455263c6c4565366ce7af3dffab",
  "Who": "@61d8b12ebd9c155adab0b64d82619ee0",
  "MsgType": 1,
  "SubType": 0,
  "OriginContent": "美女",
  "At": "",
  "Url": "",
  "RecommendInfo": null
}
`)
)

func TestCmd(t *testing.T) {
	reply, err := request(defaultBackend.url, robot)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(reply)
}

func TestGirl(t *testing.T) {
	reply, err := request(defaultBackend.url, girl)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(reply))

	replyType := gjson.Get(reply, "type").String()
	replyData := gjson.Get(reply, "data").String()
	replyErr := gjson.Get(reply, "error").String()

	var n int
	if len(replyData) < 10 {
		n = len(replyData)
	}
	fmt.Println("got:", replyType, len(replyData), replyData[0:n], "err", replyErr)

	decoded, err := base64.StdEncoding.DecodeString(replyData)
	if err != nil {
		t.Error("decode err", err)
	}
	err = ioutil.WriteFile("testpic.png", decoded, 0644)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("see file testpic.png")
}

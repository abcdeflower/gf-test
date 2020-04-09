package chat

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gvalid"
)

type Msg struct {
	Type string      `json:"type" v:"type@required#不能为空"`
	Data interface{} `json:"data" v:""`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ReceiveMsg(msgByte []byte) error {
	msg := &Msg{}

	if err := gjson.DecodeTo(msgByte, msg); err != nil {
		fmt.Printf("decode error: %s\n", err)
		return err
	}

	if err := gvalid.CheckStruct(msg, nil); err != nil {
		fmt.Printf("check struct error: %s\n", err)
		return err
	}
	return nil
}

func SendMsg(msg string) error {
	return nil
}

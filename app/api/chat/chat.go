package chat

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"reflect"
	"time"
	"web_test/app/service/chat"
	"web_test/common"
)

type Controller struct{}

func (c *Controller) WebSocket(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}

	for {
		msgType, msgByte, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("readmessage error: %s\n", err)
			break
		}
		fmt.Printf("%s\n", reflect.TypeOf(ws))
		if err := chat.ReceiveMsg(msgByte); err != nil {
			common.Json(r, 0, fmt.Sprintf("receive error: %s", err), 0, nil)
			break
		}

		for i := 0; i <= 1; i++ {

			if err = ws.WriteMessage(msgType, msgByte); err != nil {
				break
			}
			time.Sleep(time.Second * 5)
		}

	}
	ws.Close()
}

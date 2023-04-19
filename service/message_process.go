package gold_price

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/log"
	"gold-price/util"
	"strings"
	"time"
)

func atMessage(event *dto.WSPayload, data *dto.WSATMessageData) error {

	if strings.HasSuffix(data.Content, "> hello") { // 如果@机器人并输入 hello 则回复 你好。
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "你好"})
	}
	return nil
}

func catchErrHandler(messageData *dto.WSATMessageData) {
	if r := recover(); r != nil {
		log.Error("catchErrHandler: ", r)
		content := time.Now().String() + " atMessageHandler出现了不知道是什么玩意的异常"
		util.PostMessage(api, ctx, args.ChannelId, messageData.ID, content)
	}
}

// 消息处理器
func atMessageHandler(event *dto.WSPayload, messageData *dto.WSATMessageData) error {
	// 捕捉异常 防止程序停止
	defer catchErrHandler(messageData)
	res := message.ParseCommand(messageData.Content)
	var defaultContent = "\n未知命令[ " + res.Cmd + " ], 请联系开发者添加."

	var content = ""

	content, err := price(res.Cmd)
	if err != nil {
		log.Error(err)
		if content == "err" {
			content = "\n" + err.Error()
		} else {
			content = defaultContent
		}
	}

	log.Info(content)

	// 发消息
	util.PostMessage(api, ctx, args.ChannelId, messageData.ID, content)

	return nil
}

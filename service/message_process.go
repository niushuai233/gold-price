package gold_price

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/log"
	"gold-price/util"
	"strings"
)

func atMessage(event *dto.WSPayload, data *dto.WSATMessageData) error {

	if strings.HasSuffix(data.Content, "> hello") { // 如果@机器人并输入 hello 则回复 你好。
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "你好"})
	}
	return nil
}

// 消息处理器
func atMessageHandler(event *dto.WSPayload, messageData *dto.WSATMessageData) error {
	res := message.ParseCommand(messageData.Content)
	var defaultContent = "\n未知命令[ " + res.Cmd + " ], 请联系开发者添加."

	var content = ""

	content, err := price(res.Cmd)
	if err != nil {
		content = defaultContent
	}

	log.Info(content)

	// 发消息
	util.PostMessage(api, ctx, args.ChannelId, messageData, content)

	return nil
}

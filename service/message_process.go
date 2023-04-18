package gold_price

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"gold-price/util"
)

// 消息处理器
func atMessageHandler(event *dto.WSPayload, messageData *dto.WSATMessageData) error {
	res := message.ParseCommand(messageData.Content)
	var defaultContent = "未知命令[" + res.Cmd + "], 请联系开发者添加."

	var content = ""

	content, err := util.Price(res.Cmd)
	if err != nil {
		content = defaultContent
	}

	// 发消息
	util.PostMessage(api, ctx, args.ChannelId, messageData, content)

	return nil
}

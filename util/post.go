package util

import (
	"context"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
)

func PostMessage(api openapi.OpenAPI, ctx context.Context, channelId string, data *dto.WSATMessageData, content string) {
	api.PostMessage(ctx, channelId, &dto.MessageToCreate{
		MsgID:   data.ID,
		Content: content,
	})
}

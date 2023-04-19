package gold_price

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"gold-price/model"
	"gold-price/util"
	"os"
)

var (
	userId      uint64
	UnionOpenID string
	guildId     string
	channelId   string
)

var config *model.AppConfig
var args *model.StartArgs
var botToken *token.Token
var api openapi.OpenAPI
var ctx context.Context

func init() {
	log.Info("主业务流程初始化")
}

// ServiceRun 业务入口
func ServiceRun(appConfig *model.AppConfig, startArgs *model.StartArgs) {

	config = appConfig
	args = startArgs

	printConfig()
	// 初始化api对象
	botToken = token.BotToken(config.AppId, config.Token)
	api = botgo.NewSandboxOpenAPI(botToken)
	ctx = context.Background()

	fmt.Println("service run...")

	// 开启定时任务
	timerRun()
	// 开启对话通道
	chatChannelRun()

	//content, err := price(model.TodayPrice)
	//if err != nil {
	//	log.Error(err)
	//	os.Exit(1)
	//}
	//log.Info(content)
}

// 聊天通道 @机器人消息时回复信息
func chatChannelRun() {

	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Error("开启websocket失败, err = ", err)
		os.Exit(1)
	}

	// 消息处理器
	var atMessage event.ATMessageEventHandler = atMessageHandler
	intent := websocket.RegisterHandlers(atMessage)
	botgo.NewSessionManager().Start(ws, botToken, &intent)
}

// 定时任务 每天上午10点推送
func timerRun() {

	//cron表达式由6部分组成，从左到右分别表示 秒 分 时 日 月 星期
	timer := cron.New()

	timer.AddFunc("00 10,18 * * *", func() {
		content := getAllPriceContent()
		log.Info("定时主动推送: ", content)
		util.PostMessage(api, ctx, args.ChannelId, "", content)
	})

	timer.Start()
}

func getAllPriceContent() string {
	var result string

	result = result + "\n" + getPriceContent(model.TodayPrice)
	result = result + "\n" + getPriceContent(model.LFX)
	result = result + "\n" + getPriceContent(model.ZDS)
	result = result + "\n" + getPriceContent(model.ZSS)
	result = result + "\n" + getPriceContent(model.ZDF)
	result = result + "\n" + getPriceContent(model.ZLF)
	result = result + "\n" + getPriceContent(model.LFZB)
	result = result + "\n" + getPriceContent(model.LM)

	return result
}

func getPriceContent(brand string) string {
	content, err := price(brand)
	if err != nil {
		return brand + " 获取失败: " + err.Error()
	}
	return content
}

// 配置信息输出
func printConfig() {
	data, _ := json.Marshal(args)
	log.Debug("args: ", string(data))
	data, _ = json.Marshal(config)
	log.Debug("app config: ", string(data))
}

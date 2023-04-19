package main

import (
	"flag"
	"github.com/tencent-connect/botgo/log"
	"gold-price/model"
	gold_price "gold-price/service"
	"os"
)

func init() {
	log.Info("金价监控程序启动")
}

func main() {
	args := parseStartArgs()
	if args == nil {
		os.Exit(1)
	}
	config := model.AppConfig{
		AppId:  args.AppId,
		Token:  args.Token,
		Secret: args.Secret,
	}
	gold_price.ServiceRun(&config, args)
}

func parseStartArgs() *model.StartArgs {
	var (
		appId     uint64
		token     string
		secret    string
		userId    string
		guildId   string
		channelId string
	)
	flag.Uint64Var(&appId, "appId", 0, "机器人ID, 必填")
	flag.StringVar(&token, "token", "", "机器人Token, 必填")
	flag.StringVar(&secret, "secret", "", "秘钥")
	flag.StringVar(&userId, "userId", "", "用户id")
	flag.StringVar(&guildId, "guildId", "", "频道id")
	flag.StringVar(&channelId, "channelId", "", "子频道id, 必填")
	flag.Parse()

	if appId == 0 || token == "" || channelId == "" {
		log.Error("请检查appId, token, channelId三个必填参数值是否均存在")
		return nil
	}

	startArgs := model.StartArgs{
		AppId:     appId,
		Token:     token,
		UserId:    userId,
		GuildId:   guildId,
		ChannelId: channelId,
		Secret:    secret,
	}
	return &startArgs

}

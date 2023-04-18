package model

func init() {
}

const (
	TODAY_PRICE = "今日金价"
	LFX         = "老凤祥"
	ZDS         = "周大生"
	ZSS         = "周生生"
	ZDF         = "周大福"
	ZLF         = "周六福"
	LFZB        = "六福珠宝"
	LM          = "老庙"
)

type StartArgs struct {
	// 机器人id
	AppId uint64
	// 令牌
	Token string
	// 秘钥
	Secret string
	// 用户id
	UserId string
	// 频道id
	GuildId string
	// 子频道id
	ChannelId string
}

// AppConfig 机器人认证配置
type AppConfig struct {
	// 机器人id
	AppId uint64 `yaml:"appid"`
	// 令牌
	Token string `yaml:"token"`
	// 秘钥
	Secret string `yaml:"secret"`
}

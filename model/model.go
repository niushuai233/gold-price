package model

func init() {
}

// 产品编码
var brands = [...]Brand{
	// TodayPrice = "今日金价"
	{
		product:   TpBase_JO_52683,
		brandName: TodayPrice,
	}, {
		product:   TpBase_JO_52684,
		brandName: TodayPrice,
	}, {
		product:   TpBase_JO_52685,
		brandName: TodayPrice,
	},
	// LFX        = "老凤祥"
	{
		product:   LFX_Gold,
		brandName: LFX,
	},
	{
		product:   LFX_PtGold,
		brandName: LFX,
	},
	// ZDS        = "周大生"
	{
		product:   ZDS_Gold,
		brandName: ZDS,
	},
	{
		product:   ZDS_PtGold,
		brandName: ZDS,
	},
	// ZSS        = "周生生"
	{
		product:   ZSS_Gold,
		brandName: ZSS,
	},
	{
		product:   ZSS_PtGold,
		brandName: ZSS,
	},
	// ZDF        = "周大福"
	{
		product:   ZDF_Gold,
		brandName: ZDF,
	},
	{
		product:   ZDF_PtGold,
		brandName: ZDF,
	},
	// ZLF        = "周六福"
	{
		product:   ZLF_Gold,
		brandName: ZLF,
	},
	{
		product:   ZLF_PtGold,
		brandName: ZLF,
	},
	// LFZB       = "六福珠宝"
	{
		product:   LFZB_Gold,
		brandName: LFZB,
	},
	{
		product:   LFZB_PtGold,
		brandName: LFZB,
	},
	// LM         = "老庙"
	{
		product:   LM_Gold,
		brandName: LM,
	},
	{
		product:   LM_PtGold,
		brandName: LM,
	},
}

// productCode 定义
var (
	// 中国黄金
	TpBase_JO_52683 = ProductCode{
		code: "JO_52683",
		name: "中国黄金基础金价",
	}
	TpBase_JO_52684 = ProductCode{
		code: "JO_52684",
		name: "投资金条/储值金条/元宝金：零售价",
	}
	TpBase_JO_52685 = ProductCode{
		code: "JO_52685",
		name: "投资金条/储值金条/元宝金：回购价",
	}
	// LFX        = "老凤祥"
	LFX_PtGold = ProductCode{
		code: "JO_42658",
		name: "铂金价格",
	}
	LFX_Gold = ProductCode{
		code: "JO_42657",
		name: "黄金价格",
	}
	// ZDS        = "周大生"
	ZDS_PtGold = ProductCode{
		code: "JO_52677",
		name: "铂金价格",
	}
	ZDS_Gold = ProductCode{
		code: "JO_52678",
		name: "黄金价格",
	}
	// ZSS        = "周生生"
	ZSS_PtGold = ProductCode{
		code: "JO_42626",
		name: "铂金价格",
	}
	ZSS_Gold = ProductCode{
		code: "JO_42625",
		name: "黄金价格",
	}
	// ZDF        = "周大福"
	ZDF_PtGold = ProductCode{
		code: "JO_42661",
		name: "铂金价格",
	}
	ZDF_Gold = ProductCode{
		code: "JO_42660",
		name: "黄金价格",
	}
	// ZLF        = "周六福"
	ZLF_PtGold = ProductCode{
		code: "JO_42654",
		name: "铂金价格",
	}
	ZLF_Gold = ProductCode{
		code: "JO_42653",
		name: "黄金价格",
	}
	// LFZB       = "六福珠宝"
	LFZB_PtGold = ProductCode{
		code: "JO_42647",
		name: "铂金价格",
	}
	LFZB_Gold = ProductCode{
		code: "JO_42646",
		name: "黄金价格",
	}
	// LM         = "老庙"
	LM_PtGold = ProductCode{
		code: "JO_42635",
		name: "铂金价格",
	}
	LM_Gold = ProductCode{
		code: "JO_42634",
		name: "黄金价格",
	}
)

const (
	TodayPrice = "今日金价"
	LFX        = "老凤祥"
	ZDS        = "周大生"
	ZSS        = "周生生"
	ZDF        = "周大福"
	ZLF        = "周六福"
	LFZB       = "六福珠宝"
	LM         = "老庙"
)

// ProductCode 产品代码
type ProductCode struct {
	// code
	code string
	// name
	name string
}

// Brand 品牌
type Brand struct {
	// 品牌代码
	product ProductCode
	// 品牌名称
	brandName string
}

// PriceResp 价格响应结果
type PriceResp struct {
	// 响应标识
	flag bool
	// 错误代码
	errorCode int
	// 不同产品的价格集合
	codePrices [...]CodePrice
}

// CodePrice 产品价格
type CodePrice struct {
	// 产品代码
	code string
	// 展示code
	showCode string
	// 展示name
	showName string
	// 单位
	unit string
	// 不知道干啥的
	status int
	// 不知道干啥的
	digits int
	q1     int
	q2     int
	q3     int
	q4     int
	q60    int
	q63    int
	q70    int
	q80    int
	q128   int
	q129   int
	q193   int
}

// StartArgs 启动参数
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

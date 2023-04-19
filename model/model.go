package model

func init() {
}

// 产品编码
var BrandArray = [...]Brand{
	// TodayPrice = "今日金价"
	{
		Product:   TpBase_JO_52683,
		BrandName: TodayPrice,
	}, {
		Product:   TpBase_JO_52684,
		BrandName: TodayPrice,
	}, {
		Product:   TpBase_JO_52685,
		BrandName: TodayPrice,
	},
	// LFX        = "老凤祥"
	{
		Product:   LFX_Gold,
		BrandName: LFX,
	},
	{
		Product:   LFX_PtGold,
		BrandName: LFX,
	},
	// ZDS        = "周大生"
	{
		Product:   ZDS_Gold,
		BrandName: ZDS,
	},
	{
		Product:   ZDS_PtGold,
		BrandName: ZDS,
	},
	// ZSS        = "周生生"
	{
		Product:   ZSS_Gold,
		BrandName: ZSS,
	},
	{
		Product:   ZSS_PtGold,
		BrandName: ZSS,
	},
	// ZDF        = "周大福"
	{
		Product:   ZDF_Gold,
		BrandName: ZDF,
	},
	{
		Product:   ZDF_PtGold,
		BrandName: ZDF,
	},
	// ZLF        = "周六福"
	{
		Product:   ZLF_Gold,
		BrandName: ZLF,
	},
	{
		Product:   ZLF_PtGold,
		BrandName: ZLF,
	},
	// LFZB       = "六福珠宝"
	{
		Product:   LFZB_Gold,
		BrandName: LFZB,
	},
	{
		Product:   LFZB_PtGold,
		BrandName: LFZB,
	},
	// LM         = "老庙"
	{
		Product:   LM_Gold,
		BrandName: LM,
	},
	{
		Product:   LM_PtGold,
		BrandName: LM,
	},
}

// productCode 定义
var (
	// 中国黄金
	TpBase_JO_52683 = ProductCode{
		Code: "JO_52683",
		Name: "中国黄金基础金价",
	}
	TpBase_JO_52684 = ProductCode{
		Code: "JO_52684",
		Name: "投资金条/储值金条/元宝金：零售价",
	}
	TpBase_JO_52685 = ProductCode{
		Code: "JO_52685",
		Name: "投资金条/储值金条/元宝金：回购价",
	}
	// LFX        = "老凤祥"
	LFX_PtGold = ProductCode{
		Code: "JO_42658",
		Name: "铂金价格",
	}
	LFX_Gold = ProductCode{
		Code: "JO_42657",
		Name: "黄金价格",
	}
	// ZDS        = "周大生"
	ZDS_PtGold = ProductCode{
		Code: "JO_52677",
		Name: "铂金价格",
	}
	ZDS_Gold = ProductCode{
		Code: "JO_52678",
		Name: "黄金价格",
	}
	// ZSS        = "周生生"
	ZSS_PtGold = ProductCode{
		Code: "JO_42626",
		Name: "铂金价格",
	}
	ZSS_Gold = ProductCode{
		Code: "JO_42625",
		Name: "黄金价格",
	}
	// ZDF        = "周大福"
	ZDF_PtGold = ProductCode{
		Code: "JO_42661",
		Name: "铂金价格",
	}
	ZDF_Gold = ProductCode{
		Code: "JO_42660",
		Name: "黄金价格",
	}
	// ZLF        = "周六福"
	ZLF_PtGold = ProductCode{
		Code: "JO_42654",
		Name: "铂金价格",
	}
	ZLF_Gold = ProductCode{
		Code: "JO_42653",
		Name: "黄金价格",
	}
	// LFZB       = "六福珠宝"
	LFZB_PtGold = ProductCode{
		Code: "JO_42647",
		Name: "铂金价格",
	}
	LFZB_Gold = ProductCode{
		Code: "JO_42646",
		Name: "黄金价格",
	}
	// LM         = "老庙"
	LM_PtGold = ProductCode{
		Code: "JO_42635",
		Name: "铂金价格",
	}
	LM_Gold = ProductCode{
		Code: "JO_42634",
		Name: "黄金价格",
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
	// Code
	Code string
	// Name
	Name string
}

// Brand 品牌
type Brand struct {
	// 品牌代码
	Product ProductCode
	// 品牌名称
	BrandName string
}

// PriceResp 价格响应结果
type PriceResp struct {
	// 响应标识
	Flag bool
	// 错误代码
	ErrorCode int
	// 不同产品的价格集合
	CodePrices []CodePrice
}

// CodePrice 产品价格
type CodePrice struct {
	// 产品代码
	Code string `json:"code"`
	// 展示code
	ShowCode string `json:"showCode"`
	// 展示name
	ShowName string `json:"showName"`
	// 单位
	Unit string `json:"unit"`
	// 不知道干啥的
	Status int `json:"status"`
	// 不知道干啥的
	Digits int     `json:"digits"`
	Q1     float64 `json:"q1"`
	Q2     float64 `json:"q2"`
	Q3     float64 `json:"q3"`
	Q4     float64 `json:"q4"`
	Q60    float64 `json:"q60"`
	Q63    float64 `json:"q63"`
	Q70    float64 `json:"q70"`
	Q80    float64 `json:"q80"`
	Q128   float64 `json:"q128"`
	Q129   float64 `json:"q129"`
	Q193   float64 `json:"q193"`
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

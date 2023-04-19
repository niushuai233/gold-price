package gold_price

import (
	"errors"
	"github.com/tencent-connect/botgo/log"
	"gold-price/model"
)

const (
	BASE_MESSAGE_TEMPLATE  = "\n%s\n\t黄金基础价格：%f\n\t投资金价格：%f\n\t回收价格: %f"
	BRAND_MESSAGE_TEMPLATE = "\n%s\n\t黄金价格：%f\n\t铂金价格：%f"
)

const (
	CODE_TODAY_PRICEX = "JO_52683,JO_52684,JO_52685"
)

const (
	URL_PRICE = "https://api.jijinhao.com/quoteCenter/realTime.htm?codes=CODES&_=CURR_TIMES"

	URL_HISTORY_PRICE = "https://api.jijinhao.com/quoteCenter/history.htm?code=CODES&style=3&pageSize=10&needField=128,129,70&currentPage=1&_=1681867495109"
)

func price(brand string) (string, error) {
	log.Info("brand: ", brand)
	switch brand {
	case model.TodayPrice:

		break
	case model.LFX:
		break
	case model.ZDS:
		break
	case model.ZSS:
		break
	case model.ZDF:
		break
	case model.ZLF:
		break
	case model.LFZB:
		break
	case model.LM:
		break
	default:
		return "", errors.New("未知品牌: " + brand)
	}

	return "", nil
}

package gold_price

import (
	"errors"
	"gold-price/model"
)

const BASE_MESSAGE_TEMPLATE = "\n%s\n\t黄金基础价格：%f\n\t投资金价格：%f\n\t回收价格: %f"
const BRAND_MESSAGE_TEMPLATE = "\n%s\n\t黄金价格：%f\n\t铂金价格：%f"

func price(brand string) (string, error) {

	switch brand {

	case model.TODAY_PRICE:
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

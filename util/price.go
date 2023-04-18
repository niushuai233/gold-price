package util

import (
	"errors"
	"gold-price/model"
)

func Price(brand string) (string, error) {

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

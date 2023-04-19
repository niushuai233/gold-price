package gold_price

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tencent-connect/botgo/log"
	"gold-price/model"
	"gold-price/util"
	"strconv"
	"strings"
	"time"
)

const (
	FLAG                   = "flag"
	CODES                  = "CODES"
	CURR_TIMES             = "CURR_TIMES"
	QUOTE_JSON             = "var quote_json ="
	BASE_MESSAGE_TEMPLATE  = "\n%s\n  > %s：%s\n  > %s：%s\n  > %s：%s"
	BRAND_MESSAGE_TEMPLATE = "\n%s\n\t%s：%f\n\t%s：%f"
)

const (
	URL_PRICE = "https://api.jijinhao.com/quoteCenter/realTime.htm?codes=CODES&_=CURR_TIMES"

	URL_HISTORY_PRICE = "https://api.jijinhao.com/quoteCenter/history.htm?code=CODES&style=3&pageSize=10&needField=128,129,70&currentPage=1&_=1681867495109"
)

func price(brand string) (string, error) {
	log.Info("brand: ", brand)

	switch brand {
	case model.TodayPrice:
		return getTodayPrice()
	case model.LFX:
		return getLFXPrice()
	case model.ZDS:
		return getZDSPrice()
	case model.ZSS:
		return getZSSPrice()
	case model.ZDF:
		return getZDFPrice()
	case model.ZLF:
		return getZLFPrice()
	case model.LFZB:
		return getLFZBPrice()
	case model.LM:
		return getLMPrice()
	default:
		return "err", errors.New("未知品牌: " + brand)
	}
}

func formatUrl(url string, productCodes []model.ProductCode) string {

	target := ""
	for index, productCode := range productCodes {
		target = target + productCode.Code
		if index != len(productCodes)-1 {
			target = target + ","
		}
	}

	return strings.ReplaceAll(strings.ReplaceAll(url, CODES, target), CURR_TIMES, strconv.FormatInt(time.Now().UnixMilli(), 10))
}

func formatFloatPrice(price float64) string {

	result := strconv.FormatFloat(price, 'f', 6, 64)
	for strings.HasSuffix(result, "0") {
		result = strings.TrimSuffix(result, "0")
	}
	return result
}

func getRespMap(productCodes []model.ProductCode) (map[string]interface{}, error) {
	url := formatUrl(URL_PRICE, productCodes)
	respBody, err := util.Get(url)
	if err != nil {
		return nil, err
	}
	respBody = strings.ReplaceAll(respBody, QUOTE_JSON, "")
	log.Info(respBody)
	// 重组resp

	//var codePrice model.CodePrice
	json2Map, err := util.Json2Map(respBody)

	flag := json2Map[FLAG]
	tmp := fmt.Sprint(flag)
	if tmp != "true" {
		return nil, errors.New("请求接口失败: " + url)
	}
	return json2Map, nil
}

func getTodayPrice() (string, error) {
	json2Map, err := getRespMap([]model.ProductCode{model.TpBase_JO_52683, model.TpBase_JO_52684, model.TpBase_JO_52685})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	baseGold := json2Map[model.TpBase_JO_52683.Code]
	baseGold_TZ := json2Map[model.TpBase_JO_52684.Code]
	baseGold_TZ_HS := json2Map[model.TpBase_JO_52685.Code]

	goldJson, _ := util.Map2Json(baseGold.(map[string]interface{}))
	tzGoldJson, _ := util.Map2Json(baseGold_TZ.(map[string]interface{}))
	hsGoldJson, _ := util.Map2Json(baseGold_TZ_HS.(map[string]interface{}))

	log.Info(goldJson, tzGoldJson, hsGoldJson)

	var goldCp model.CodePrice
	var tzGoldCp model.CodePrice
	var hsGoldCp model.CodePrice
	json.Unmarshal([]byte(goldJson), &goldCp)
	json.Unmarshal([]byte(tzGoldJson), &tzGoldCp)
	json.Unmarshal([]byte(hsGoldJson), &hsGoldCp)

	content := fmt.Sprintf(BASE_MESSAGE_TEMPLATE, model.TodayPrice,
		goldCp.ShowName, formatFloatPrice(goldCp.Q1),
		tzGoldCp.ShowName, formatFloatPrice(tzGoldCp.Q1),
		hsGoldCp.ShowName, formatFloatPrice(hsGoldCp.Q1),
	)

	return content, nil
}

func getLFXPrice() (string, error) {
	return "", nil
}

func getZDSPrice() (string, error) {
	return "", nil
}

func getZSSPrice() (string, error) {
	return "", nil
}

func getZDFPrice() (string, error) {
	return "", nil
}

func getZLFPrice() (string, error) {
	return "", nil
}

func getLFZBPrice() (string, error) {
	return "", nil
}

func getLMPrice() (string, error) {
	return "", nil
}

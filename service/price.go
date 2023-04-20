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
	BRAND_MESSAGE_TEMPLATE = "\n%s\n  > %s：%s\n  > %s：%s"
)

const (
	URL_PRICE = "https://api.jijinhao.com/quoteCenter/realTime.htm?codes=CODES&_=CURR_TIMES"

	URL_HISTORY_PRICE = "https://api.jijinhao.com/quoteCenter/history.htm?code=CODES&style=3&pageSize=10&needField=128,129,70&currentPage=1&_=1681867495109"
)

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
	for strings.HasSuffix(result, ".") {
		result = strings.TrimSuffix(result, ".")
	}
	return result
}

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

func goldPrice(brand string, tmpMap map[string]interface{}, goldProduct model.ProductCode, ptGoldProduct model.ProductCode) (model.CodePrice, model.CodePrice) {
	goldJson, _ := util.Map2Json(tmpMap[goldProduct.Code].(map[string]interface{}))
	ptGoldJson, _ := util.Map2Json(tmpMap[ptGoldProduct.Code].(map[string]interface{}))

	log.Info(brand, goldJson, ptGoldJson)

	var goldCp model.CodePrice
	var ptGoldCp model.CodePrice
	json.Unmarshal([]byte(goldJson), &goldCp)
	json.Unmarshal([]byte(ptGoldJson), &ptGoldCp)
	return goldCp, ptGoldCp
}

func getTodayPrice() (string, error) {
	baseMap, err := getRespMap([]model.ProductCode{model.TpBase_JO_52683, model.TpBase_JO_52684, model.TpBase_JO_52685})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldJson, _ := util.Map2Json(baseMap[model.TpBase_JO_52683.Code].(map[string]interface{}))
	tzGoldJson, _ := util.Map2Json(baseMap[model.TpBase_JO_52684.Code].(map[string]interface{}))
	hsGoldJson, _ := util.Map2Json(baseMap[model.TpBase_JO_52685.Code].(map[string]interface{}))

	log.Info(goldJson, tzGoldJson, hsGoldJson)

	var goldCp model.CodePrice
	var tzGoldCp model.CodePrice
	var hsGoldCp model.CodePrice
	json.Unmarshal([]byte(goldJson), &goldCp)
	json.Unmarshal([]byte(tzGoldJson), &tzGoldCp)
	json.Unmarshal([]byte(hsGoldJson), &hsGoldCp)

	content := fmt.Sprintf(BASE_MESSAGE_TEMPLATE, model.TodayPrice,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		tzGoldCp.ShowName, formatFloatPrice(tzGoldCp.Q63),
		hsGoldCp.ShowName, formatFloatPrice(hsGoldCp.Q63),
	)

	return content, nil
}

func getLFXPrice() (string, error) {
	lfxMap, err := getRespMap([]model.ProductCode{model.LFX_Gold, model.LFX_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.LFX, lfxMap, model.LFX_Gold, model.LFX_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.LFX,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)

	return content, nil
}

func getZDSPrice() (string, error) {
	zdsMap, err := getRespMap([]model.ProductCode{model.ZDS_Gold, model.ZDS_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.ZDS, zdsMap, model.ZDS_Gold, model.ZDS_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.ZDS,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)

	return content, nil
}

func getZSSPrice() (string, error) {
	zssMap, err := getRespMap([]model.ProductCode{model.ZSS_Gold, model.ZSS_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.ZSS, zssMap, model.ZSS_Gold, model.ZSS_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.ZSS,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)
	return content, nil
}

func getZDFPrice() (string, error) {
	zdfMap, err := getRespMap([]model.ProductCode{model.ZDF_Gold, model.ZDF_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.ZDF, zdfMap, model.ZDF_Gold, model.ZDF_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.ZDF,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)
	return content, nil
}

func getZLFPrice() (string, error) {
	zlfMap, err := getRespMap([]model.ProductCode{model.ZLF_Gold, model.ZLF_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.ZLF, zlfMap, model.ZLF_Gold, model.ZLF_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.ZLF,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)
	return content, nil
}

func getLFZBPrice() (string, error) {
	lfzbMap, err := getRespMap([]model.ProductCode{model.LFZB_Gold, model.LFZB_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.LFZB, lfzbMap, model.LFZB_Gold, model.LFZB_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.LFZB,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)
	return content, nil
}

func getLMPrice() (string, error) {
	lmMap, err := getRespMap([]model.ProductCode{model.LM_Gold, model.LM_PtGold})
	if err != nil {
		// 获取响应结果失败
		return "err", err
	}

	goldCp, ptGoldCp := goldPrice(model.LM, lmMap, model.LM_Gold, model.LM_PtGold)

	content := fmt.Sprintf(BRAND_MESSAGE_TEMPLATE, model.LM,
		goldCp.ShowName, formatFloatPrice(goldCp.Q63),
		ptGoldCp.ShowName, formatFloatPrice(ptGoldCp.Q63),
	)
	return content, nil
}

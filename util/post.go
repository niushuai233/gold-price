package util

import (
	"context"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/openapi"
	"io/ioutil"
	"net/http"
)

func PostMessage(api openapi.OpenAPI, ctx context.Context, channelId string, msgId string, content string) {
	api.PostMessage(ctx, channelId, &dto.MessageToCreate{
		MsgID:   msgId,
		Content: content,
	})
}

func Get(url string) (string, error) {
	log.Info("url: ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("请求失败: %s, %s, %s", resp.Status, resp.StatusCode, resp)
		log.Error("请求失败err: ", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("读取body失败: ", err)
		return "", err
	}

	log.Info("resp: ", string(body))
	return string(body), nil
	//return "{\"flag\":true,\"JO_52683\":{\"code\":\"JO_52683\",\"time\":1681873834000,\"q128\":0.0,\"q80\":0.27057773,\"q129\":99.9,\"q1\":444.7,\"q193\":1.0,\"q2\":443.5,\"q3\":444.7,\"q4\":444.7,\"q70\":1.2000122,\"q60\":1.0,\"q63\":444.7,\"unit\":\"元/克\",\"showName\":\"中国黄金基础金价\",\"showCode\":\"zhongguohuangjinjichujinjia\",\"digits\":2,\"status\":100},\"JO_52684\":{\"code\":\"JO_52684\",\"time\":1681874524000,\"q128\":0.0,\"q80\":0.26229775,\"q129\":99.9,\"q1\":458.7,\"q193\":1.0,\"q2\":457.5,\"q3\":458.7,\"q4\":458.7,\"q70\":1.2000122,\"q60\":2.0,\"q63\":458.7,\"unit\":\"元/克\",\"showName\":\"投资金条/储值金条/元宝金：零售价\",\"showCode\":\"touzijintiao/chuzhijintiao/yuanbaojin：lingshoujia\",\"digits\":2,\"status\":100},\"JO_52685\":{\"code\":\"JO_52685\",\"time\":1681873823000,\"q128\":0.0,\"q80\":0.27180344,\"q129\":99.9,\"q1\":442.7,\"q193\":1.0,\"q2\":441.5,\"q3\":442.7,\"q4\":442.7,\"q70\":1.2000122,\"q60\":1.0,\"q63\":442.7,\"unit\":\"元/克\",\"showName\":\"投资金条/储值金条/元宝金：回购价\",\"showCode\":\"touzijintiao/chuzhijintiao/yuanbaojin：huigoujia\",\"digits\":2,\"status\":100},\"errorCode\":[]}", nil
}

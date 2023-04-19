package util

import (
	"encoding/json"
	"github.com/tencent-connect/botgo/log"
)

func Json2Map(source string) (map[string]interface{}, error) {

	tmpMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(source), &tmpMap)
	if err != nil {
		log.Error("json2map err = ", err)
		return nil, err
	}

	for k, v := range tmpMap {
		log.Debugf("key = %s, value = %s", k, v)
	}
	return tmpMap, nil
}

func Map2Json(sourceMap map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(sourceMap)
	if err != nil {
		log.Error("map2json err = ", err)
		return "", err
	}

	return string(bytes), nil
}

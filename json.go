package helpers

import "encoding/json"

/**
 * @Description: Json转换助手
 * @Date: 2022-08-26 15:25:36
 */
func JsonMapToString(data map[string]interface{}) string {
	if len(data) <= 0 {
		return ""
	}
	jsonObj, _ := json.Marshal(data)
	return string(jsonObj)
}

func JsonSliceToString(data []string) string {
	if len(data) <= 0 {
		return ""
	}
	jsonObj, _ := json.Marshal(data)
	return string(jsonObj)
}

func JsonToString(data interface{}) string {
	jsonObj, _ := json.Marshal(data)
	return string(jsonObj)
}

func JsonDecodeToMap(data string) map[string]interface{} {
	var res map[string]interface{}
	json.Unmarshal([]byte(data), &res)
	return res
}

func JsonDecodeToSlice(data string) []string {
	var res []string
	json.Unmarshal([]byte(data), &res)
	return res
}

func JsonDecodeToSliceMap(data string) []map[string]interface{} {
	var res []map[string]interface{}
	json.Unmarshal([]byte(data), &res)
	return res
}

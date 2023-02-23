package helpers

import "time"

/**
 * @Description: Redis操作助手
 * @Date: 2022-08-26 15:25:36
 */

/**
 * @name: GetRedisData
 * @description: 获取redis数据(有数据则直接返回，无数据则通过回调函数)
 * @param {string} key redis-key值
 * @param {time.Duration} expiration 超时时间
 * @param {func(map[string]interface{}) (map[string]interface{}, error)} callback 自定义回调函数
 * @param {map[string]interface{}} callback_param 自定义回调函数传递参数
 * @return {map[string]interface{}, error}
 * @author: gx
 *
 * //-------使用示例----------
 *
 * //自定义函数和参数
 * param := map[string]interface{}{"name": name}
 * callback := func(param map[string]interface{}) (map[string]interface{}, error) {
 *  corp := GetCorpInfoByParam(param, nil) //查询数据
 *  return map[string]interface{}{"name":corp.Name,"age":corp.Age}, nil
 * }
 * //获取redis数据
 * corp_data, err := GetRedisData("corp|"+name, time.Minute*60*3, callback, param)
 *
 */
func GetRedisData(key string, expiration time.Duration,
	callback func(map[string]interface{}) (map[string]interface{}, error),
	callback_param map[string]interface{}) (map[string]interface{}, error) {

	json_data_str, _ := redis.Get(key).Result()

	var value map[string]interface{}
	if json_data_str == "" {
		value, err := callback(callback_param)
		if err != nil {
			return nil, err
		}

		//map json 转换成string存入redis
		redis.Set(key, JsonMapToString(value), expiration)
	} else {
		//redis string数据转换成map
		value = JsonDecodeToMap(json_data_str)
	}

	return value, nil
}

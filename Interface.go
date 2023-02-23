package helpers

/**
 * @Description: Interface接口转换助手
 * @Date: 2022-08-26 15:25:36
 */

/**
 * @name: InterfaceToSliceInt
 * @description: 接口转字符串切片
 * @param {interface{}} data
 * @return {*}
 * @author: gx
 */
func InterfaceToSliceString(data interface{}) []string {
	var result []string
	for _, v := range data.([]interface{}) {
		result = append(result, v.(string))
	}

	return result
}

/**
 * @name: InterfaceToSliceInt
 * @description: 接口转整数切片
 * @param {interface{}} data
 * @return {*}
 * @author: gx
 */
type Num interface {
	~int | ~float64
}
type ToNum interface {
	~int | ~float64
}

func InterfaceToSliceInt[N Num, T ToNum](data interface{}) []T {
	var result []T
	for _, v := range data.([]interface{}) {
		num := T(v.(N))
		result = append(result, num)
	}

	return result
}

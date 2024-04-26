package utils

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"time"
)

// 使用模板函数渲染map中的字符串值
func RenderStringValues(v any, vars map[string]interface{}) error {
	dict := make(map[string]interface{}, 0)
	FromJsonString(ToJsonString(v), &dict)
	// 递归遍历data并更新值
	result := RenderStringValuesRecursive(dict, vars)
	FromJsonString(ToJsonString(dict), &v)
	return result
}

// renderStringValuesRecursive 递归遍历map，并更新字符串值
func RenderStringValuesRecursive(data map[string]interface{}, vars map[string]interface{}) error {
	vars["uuid"] = GenUUID32()
	vars["unixTimestamp"] = fmt.Sprint(time.Now().Unix())
	vars["Nanosecond"] = fmt.Sprint(time.Now().Nanosecond())

	for key, value := range data {
		switch v := value.(type) {
		case string:
			// 如果值是字符串，使用模板函数进行渲染
			renderedValue, err := RenderStringValue(v, vars)
			if err != nil {
				return err
			}
			data[key] = renderedValue
		case map[string]interface{}:
			// 如果值是map，递归遍历
			err := RenderStringValuesRecursive(v, vars)
			if err != nil {
				return err
			}
		case []interface{}:
			// 如果值是数组，遍历数组中的每个元素
			for i, item := range v {
				if "string" == reflect.TypeOf(item).String() {
					strVal := item.(string)
					renderedValue, err := RenderStringValue(strVal, vars)
					if err != nil {
						return err
					}
					v[i] = renderedValue
					continue
				}
				if itemMap, ok := item.(map[string]interface{}); ok {
					err := RenderStringValuesRecursive(itemMap, vars)
					if err != nil {
						return err
					}
					v[i] = itemMap // 更新数组中的map
				}
			}
			data[key] = v // 更新data中的数组
		default:
			// 其他类型不处理
			continue
		}
	}
	return nil
}

// renderStringValue 使用模板函数渲染字符串值
func RenderStringValue(value string, vars map[string]interface{}) (string, error) {
	// 创建一个模板实例
	tmpl := template.New("mytemplate").Funcs(template.FuncMap{
		"getFirstDataElement": func(data []interface{}) interface{} {
			if len(data) > 0 {
				return data[0]
			}
			return nil
		},
		"add": func(a int, b int) int {
			return a + b
		},
	})
	//tmpl, err := template.New("tmpl").Parse(value)
	//if err != nil {
	//	return "", err
	//}
	tmpl, err := tmpl.Parse(value)
	if err != nil {
		panic(err)
	}
	// 创建一个字符串缓冲区来存储渲染结果
	var rendered strings.Builder

	// 执行模板渲染
	err = tmpl.Execute(&rendered, vars)
	if err != nil {
		panic(err)
	}

	// 返回渲染后的字符串
	return rendered.String(), nil
}

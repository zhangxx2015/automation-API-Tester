package main

import (
	"fmt"
	"log"
	"myApiTester/models"
	"myApiTester/utils"
	"strings"
)

var vars = make(map[string]interface{}, 0)

func main() {
	utils.JsonFromFile(utils.CurrDir("vars.json"), &vars)
	files := make([]string, 0)
	utils.JsonFromFile(utils.CurrDir("tasks.json"), &files)
	// 使用解析后的数据
	for _, item := range files {
		ri := models.RequestInfo{}
		if strings.HasSuffix(item, ".go.tpl") {
			text := utils.TextFromFile(utils.CurrDir(item))
			text, err := utils.RenderStringValue(text, vars)
			if err != nil {
				panic(err)
			}
			utils.FromJsonString(text, &ri)
		} else {
			utils.JsonFromFile(utils.CurrDir(item), &ri)
		}

		// 渲染前
		//fmt.Println("\t", utils.ToJsonString(ri))
		err := utils.RenderStringValues(&ri, vars)
		if err != nil {
			panic(err)
		}
		// 渲染后
		log.Println(item)
		log.Println("\t>>>>")
		log.Println("\t", ri.Method, ri.Url)
		if len(ri.Querys) > 0 {
			log.Println("\t", utils.ToJsonString(ri.Querys))
		}
		if len(ri.Headers) > 0 {
			log.Println("\t", utils.ToJsonString(ri.Headers))
		}
		log.Println("\t", utils.ToJsonString(ri.Body))
		result := ""
		switch ri.Method {
		case "POST":
			result = utils.Post(vars["baseUrl"].(string)+ri.Url, ri.Querys, ri.Body, ri.Headers)
			break
		case "PUT":
			result = utils.Put(vars["baseUrl"].(string)+ri.Url, ri.Querys, ri.Body, ri.Headers)
			break
		case "GET":
			result = utils.Get(vars["baseUrl"].(string)+ri.Url, ri.Querys, ri.Headers)
			break
		case "DELETE":
			result = utils.Delete(vars["baseUrl"].(string)+ri.Url, ri.Querys, ri.Body, ri.Headers)
			break
		}
		log.Print("response text:", result)
		response := utils.DictFromString(result)
		log.Println("\t<<<<")
		log.Println("\t", utils.ToJsonStringIndent(response, "\t\t\t"))
		vars["response"] = response
		log.Println("\tassert")
		log.Println("\t", utils.ToJsonString(ri.Assert))
		for k, expectedV := range ri.Assert {
			kRet, err := utils.RenderStringValue(k, vars)
			if err != nil {
				panic(err)
			}
			log.Println("\t\t express:", k, "key:", kRet, "val:", expectedV, "pass:", kRet == expectedV)
			if kRet != expectedV {
				panic(fmt.Errorf("The result does not meet the expectations [%v != %v]", expectedV, kRet))
			}
		}
		fmt.Println()
	}
	log.Println("testing is passed")
}

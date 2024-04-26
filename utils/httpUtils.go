package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Post(targetUrl string, queryParams map[string]interface{}, bodyParams map[string]interface{}, headerParams map[string]interface{}) string {
	// 构造query string
	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	urlWithQuery := targetUrl + "?" + values.Encode()

	// 将body参数编码为JSON
	body, err := json.Marshal(bodyParams)
	if err != nil {
		panic(fmt.Errorf("Error marshaling body:", err))
	}

	// 创建HTTP客户端和请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", urlWithQuery, bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Errorf("Error creating request:", err))
	}

	// 设置header
	for k, v := range headerParams {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("Error sending request:", err))
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Error reading response body:", err))
	}

	//// 打印响应状态和body
	//fmt.Println("Response Status:", resp.Status)
	response := string(responseBody)
	return response
}
func Put(targetUrl string, queryParams map[string]interface{}, bodyParams map[string]interface{}, headerParams map[string]interface{}) string {
	// 构造query string
	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	urlWithQuery := targetUrl + "?" + values.Encode()

	// 将body参数编码为JSON
	body, err := json.Marshal(bodyParams)
	if err != nil {
		panic(fmt.Errorf("Error marshaling body:", err))
	}

	// 创建HTTP客户端和请求
	client := &http.Client{}
	req, err := http.NewRequest("PUT", urlWithQuery, bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Errorf("Error creating request:", err))
	}

	// 设置header
	for k, v := range headerParams {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("Error sending request:", err))
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Error reading response body:", err))
	}

	// 打印响应状态和body
	//fmt.Println("Response Status:", resp.Status)

	response := string(responseBody)
	return response
}

func Get(targetUrl string, queryParams map[string]interface{}, headerParams map[string]interface{}) string {
	// 构造query string
	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	urlWithQuery := targetUrl + "?" + values.Encode()

	// 创建HTTP客户端和请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlWithQuery, nil)
	if err != nil {
		panic(fmt.Errorf("Error creating request:", err))
	}

	// 设置header
	for k, v := range headerParams {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("Error sending request:", err))
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Error reading response body:", err))
	}

	//// 打印响应状态和body
	//fmt.Println("Response Status:", resp.Status)

	response := string(responseBody)
	return response
}

func Delete(targetUrl string, queryParams map[string]interface{}, bodyParams map[string]interface{}, headerParams map[string]interface{}) string {
	// 构造query string
	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	urlWithQuery := targetUrl + "?" + values.Encode()

	// 将body参数编码为JSON
	body, err := json.Marshal(bodyParams)
	if err != nil {
		panic(fmt.Errorf("Error marshaling body:", err))
	}

	// 创建HTTP客户端和请求
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", urlWithQuery, bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Errorf("Error creating request:", err))
	}

	// 设置header
	for k, v := range headerParams {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("Error sending request:", err))
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Error reading response body:", err))
	}

	//// 打印响应状态和body
	//fmt.Println("Response Status:", resp.Status)

	response := string(responseBody)
	return response
}

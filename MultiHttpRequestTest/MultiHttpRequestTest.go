/*
 * @Author: Wade Zhong wzhong@hso.com
 * @Date: 2024-03-21 09:58:08
 * @LastEditTime: 2024-03-21 10:00:42
 * @LastEditors: Wade Zhong wzhong@hso.com
 * @Description: 用go读取json格式的数组，然后通过http请求接口，拿到数组的每个对象，直接放入接口参数，批量发起并发请求测试。
 * @FilePath: \GoLiteAppSnippets\MultiHttpRequestTest\MultiHttpRequestTest.go
 * Copyright (c) 2024 by Wade Zhong wzhong@hso.com, All Rights Reserved.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	// 从文件读取JSON数据
	jsonData, err := ioutil.ReadFile("jsonData.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 将JSON解析到一个字符串切片中
	var params []string
	json.Unmarshal(jsonData, &params)

	// 使用WaitGroup来等待所有的HTTP请求完成
	var wg sync.WaitGroup
	wg.Add(len(params))

	for _, param := range params {
		go func(param string) {
			defer wg.Done()

			// 发起HTTP请求
			resp, err := http.Get("http://example.com/api?param=" + param)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			// 读取并打印响应
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("Response:", string(body))
		}(param)
	}

	// 等待所有的HTTP请求完成
	wg.Wait()
}

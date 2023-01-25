package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 向网络地址发送请求，并获取一段json格式额数据
	urlpath := "https://www.cnblogs.com/"
	req, _ := http.NewRequest("GET", urlpath, nil)
	// req.Header.Add()
	client := &http.Client{}
	res, _ := client.Do(req)
	fmt.Println(res)

	//var responseObj map[string]interface{}
	//json.Unmarshal(body, &responseObj)
}

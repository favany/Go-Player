package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 遇事不决 写注释

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./main.tmpl") // 请勿刻舟求剑
	if err != nil {
		fmt.Printf("Parse template failed,err: %v", err)
		return
	}
	// 3. 渲染模板
	name := "bingo"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v", err)

	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v", err)
		return
	}
}

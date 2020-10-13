package spider

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Resp 响应
type Resp struct {
	status string
	result *Result
	msg    string
}

//GetResult 爬取结果控制器
func GetResult(w http.ResponseWriter, r *http.Request) {
	var resp = &Resp{}
	proxy := r.URL.Query().Get("proxy")
	ProxyAddr := fmt.Sprintf("http://%s", proxy)
	word := r.URL.Query().Get("word")
	if word == "" {
		resp.status = "400"
		resp.msg = "Missing KeyWord! Please Check Your URL!"
		bytes, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("序列化失败", err)
			return
		}
		w.Write(bytes)
	}
	addr := fmt.Sprintf("http://www.baidu.com/s?word=%s", word)
	var result = &Result{}
	result = Spider(addr, ProxyAddr)
	resp.status = "200"
	resp.result = result
	resp.msg = "OK!"
	bytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("序列化Json失败!", err)
		return
	}
	w.Write(bytes)
}

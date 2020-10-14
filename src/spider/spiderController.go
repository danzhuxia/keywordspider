package spider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Resp 响应
type Resp struct {
	Status string   `json:"status", omitempty`
	Result []*Result `json:"result", omitempty`
	Msg    string   `json:"msg", omitempty`
}

//GetResult 爬取结果控制器
func GetResult(w http.ResponseWriter, r *http.Request) {
	proxy := r.URL.Query().Get("proxy")
	word := r.URL.Query().Get("word")

	if word != "" && proxy != "" {
		ProxyAddr := fmt.Sprintf("http://%s", proxy)
		//www.baidu.com/s?ie=UTF-8&word=
		addr := fmt.Sprintf("http://www.baidu.com/s?ie=UTF-8&word=%s", word)
		//fmt.Println(ProxyAddr, addr)
		results := Spider(addr, ProxyAddr)
		fmt.Println(results)
		resp := &Resp{
			Status: "403",
			Result: []*Result{},
			Msg:    "代理出错，获取数据失败!",
		}
		for _, res := range results {
			if res.Title == "" || res.URL == "" || res.Desc == "" {
				resp.Status = "401"
				//resp.Result = []*Result{}
				resp.Msg = "没有取得数据"
				json.NewEncoder(w).Encode(resp)
			}else {
				resp.Status = "200"
				resp.Msg = "OK!"
				resp.Result = results
			}
			//json.NewEncoder(w).Encode(resp)
		}
		json.NewEncoder(w).Encode(resp)
	} else {
		w.Header().Set("Content-Type", "application/json")
		var resp = &Resp{
			Status: "400",
			Result: []*Result{},
			Msg:    "Please Check Your URL!",
		}
		bytes, _ := json.Marshal(resp)
		io.WriteString(w, string(bytes))
	}

}

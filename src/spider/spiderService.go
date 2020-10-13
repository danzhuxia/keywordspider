package spider

import (
	"encoding/base64"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/zhshch2002/goreq"
	"github.com/zhshch2002/gospider"
)

//Spider 爬虫
func Spider(url, proxy string) *Result {
	var result = &Result{}
	s := gospider.NewSpider(
		goreq.WithRandomUA(),
		goreq.WithProxy(proxy),
	)

	var h gospider.Handler
	h = func(ctx *gospider.Context) {
		dom, _ := ctx.Resp.HTML()
		if dom != nil {
			dom.Find("div[cmatchid]").Each(func(i int, sel *goquery.Selection) {
				sel.Find("h3 a").Each(func(i int, se *goquery.Selection) {
					title := se.Text()
					result.Title = title
					//ctx.Println("title", title)
					link := se.AttrOr("href", "")
					//ctx.Println("link",link)
					if link != "" {
						r := goreq.Get(link).AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
						html := r.Do().Text
						ctx.Println(html)
						dom2, _ := r.Do().HTML()
						if dom2 != nil {
							dom2.Find("noscript META").Each(func(i int, selection *goquery.Selection) {
								url := selection.AttrOr("http-equiv", "")
								result.URL = url
								//ctx.Println(url)
							})
						}
					}
				})
				sel.Find("a[hidefocus]").Each(func(i int, se *goquery.Selection) {
					desc := se.Text()
					result.Desc = desc
					//ctx.Println("desc", desc)
				})
			})
		}
	}
	//return result
	var auth = fmt.Sprintln("18856261718@163.com:WDwandou123")
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req := goreq.Get(url).AddHeaders(map[string]string{
		"Accept":              "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Cookie":              "BIDUPSID=E86709EA9F71B83CC9A8EEAA6EFA277E; PSTM=1601260170; BAIDUID=E86709EA9F71B83C5A09134EAA391BA7:FG=1; BD_UPN=12314753; delPer=0; BD_CK_SAM=1; BD_HOME=1; shifen[186242601188_65260]=1602326348; shifen[196633781527_58314]=1602326446; BCLID=10538863840309184978; BDSFRCVID=McCOJeC62w7cQhbrA8eiU7JQ3gKah9bTH6aoS9wFiqIM0uF9CfuoEG0PVM8g0Ku-0mkwogKKWmOTH7kF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF=tb4foCt2tDL3qJ-9q46W2t_H2MreaPLXKKOLVKnXfp7keq8CD4u-QCuq0q5CbRJxQRvZbCOjWbjsbJo2y5jHhTJWLlrnaMutWR6AVl72BprpsIJM3xFWbT8U5fK8b-3NaKviaM3EBMb1MnvDBT5h2M4qMxtOLR3pWDTm_q5TtUJMeCnTD-Dhe6JQjNLjtTDqf56bVTr_MPTHD4nYKP7qq4tehHRLLpO9WDTm_DoVJqLVjxj4LJQMjlQWyRLeKMnK2n5a-pPKKRA5Mpow3-RYjfcLyHr9-U7w3mkjbUQzfn02OP5PeJbf044syPRU2xRnWnvWKfA-b4ncjRcTehoM3xI8LNj405OTbIFO0KJDJCFahKKmejtajTPW5ptX2bj05K5XQ5rJabC_HnurQjohXUI8LNDH246R0b6qQqDyLqOlMpchDpOUMxFhKRO7ttKeLGOC2R7NLpTUMqnTBU8byxL1Db37KjvMtg3tslr7BJToepvo3Pcc3MvWMPjdJJQOBKQB0KnGbUQkeq8CQft20b0EeMtjW6LEtJuq_D-btK83fP36qRrW2RkJbmT22-usMmTW2hcH0KLKJ-F9yjA5KxobMG3PtMTItaRM5f36aMb1MRjvb4T0X-uZ0MOOhnvf-DbNWl5TtUJ6JKnTDMRh-l_-XfJyKMniQKj9-pP52hQrh459XP68bTkA5bjZKxtq3mkjbPbDfn028DKujjLWj5jWDGRf-b-X2Ic8BtI8Kb7Vbp6nLUnkbfJBDltf5ho3KjrCVn5VQC3kVDOTyPQsetD7yajK2ht8b2c0QlTztP5PbqOI5tnpQT8r34FOK5Oib4DjWRv5ab3vOU3zXpO1jPPzBN5thURB2DkO-4bCWJ5TMl5jDh3Mb6ksDMDtqtJHKbDJ_IKKJUK; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; PSINO=7; H_PS_PSSID=32814_32617_1442_31660_32705_32231_7516_32115_32718; COOKIE_SESSION=1081_1_8_9_1_2_0_0_8_1_0_0_0_0_0_0_1602465391_1602326350_1602470141%7C9%2340_36_1602326300%7C9; H_PS_645EC=92b6fgPwP1%2BIUDsz9ylohIo13q%2B1QuSROIL%2FHV0N6QQ%2BvPu5cGzjiORSFt8",
		"Proxy-Authorization": basicAuth,
	})
	s.SeedTask(req, h)
	s.Wait()
	return result
}

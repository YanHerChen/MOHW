package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

var Map = make(map[string]int)

/*
 * 中央健保署- 全民健康保險醫療服務給付項目及支付標準 檔案下載
 */
func GetJokes() {
	doc, err := goquery.NewDocument("https://www.nhi.gov.tw/Content_List.aspx?n=58ED9C8D8417D00B&topn=D39E2B72B0BDFA15")
	if err != nil {
		log.Fatal(err)
	}
	var file string
	var count int
	count = 0
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		file = fmt.Sprint(s.Find("a").Attr("href"))
		if strings.Contains(file, "https://ws.nhi.gov.tw/Download.ashx?") {
			if _, exists := Map[file]; exists {
			} else {
				count = count + 1
				Map[file] = count
				fmt.Println(count, " ", strings.Replace(file, "true", "", -1))
			}
		}
	})
}

func main() {
	GetJokes()
}

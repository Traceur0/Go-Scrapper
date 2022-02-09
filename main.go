package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


var mainURL = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	totalPageNum := getPages(mainURL)
	fmt.Println(totalPageNum)
}

func getPages(mainURL string) int {
	pages := 0
	res, err := http.Get(mainURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// HTML parsing Point
	// # Priority
	// 출력된 결과 분석
	// 		현재 코드 - 원하는 결과값의 마지막 부분, 동일 결과값이 두번 출력됨
	// 결과와 코드를 수정하여 올바른 결과값 얻기 ==> 9 or 10
	
	doc.Find(".tplPagination").Each(func(idx int, sel *goquery.Selection) {
		pages = sel.Find("a").Length() // must convert string to int
	})
	
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
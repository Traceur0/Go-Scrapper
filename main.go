package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var mainURL = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	// totalPageNum := getPages(mainURL)
	// fmt.Println(totalPageNum)
	getPages(mainURL)
}

func getPages(URL string) {
	// pages := 0

	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// last_pages := doc.Find("span.pgTotal")
	last_pages := doc.Find(".tplPagination > ul > li > span")
	// doc.Find(".tplPagination").Each(func(idx int, sel *goquery.Selection) {
	// 	pages := sel.Find("span.pgTotal").Parent()
	// })
	fmt.Println(last_pages)
	return
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
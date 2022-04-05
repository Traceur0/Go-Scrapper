package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var mainURL = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	not_trimed := getPages(mainURL)

	// totalPageNum := getPages(mainURL)
	// fmt.Println(totalPageNum)
	fmt.Println(not_trimed)
}

func getPages(URL string) int {
	pages := ""

	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	pages = doc.Find("span.pgTotal").Text()
	last_page := pages[0:3]
	int_page, _ := strconv.Atoi(last_page)
	// list := doc.Find(".tplPagination > ul.clear > li")
	// list.Each(func(idx int, sel *goquery.Selection) {
	// 	last_pages = sel.Find("span.pgTotdal").Text()
	// })
	return int_page
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
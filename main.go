package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var mainURL = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	// totalPageNum := getPages(mainURL)
	// fmt.Println(totalPageNum)
	getPages(mainURL)
}

func getPages(URL string) {
	// page_list = []int{}

	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// last_pages := doc.Find("span.pgTotal")
	list := doc.Find(".tplPagination > ul.clear > li")
	list.Each(func(idx int, sel *goquery.Selection) {
		last_pages := sel.Find("span.pgTotal").Text()
		trimmed_page := strings.ReplaceAll(last_pages, "\n", "")
		fmt.Println(trimmed_page)
	})
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
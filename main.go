package main

import (
	"fmt"
	"log"
	"net/http"

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
	last_pages := 0
	
	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// last_pages := doc.Find("span.pgTotal")
	list := doc.Find(".tplPagination > ul.clear > li")
	list.Each(func(idx int, sel *goquery.Selection) {
		last_pages = sel.Find("span.pgTotdal").Length()
	})
	return last_pages
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
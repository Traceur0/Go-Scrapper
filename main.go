package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type pageNumber struct {
	p_num string
}

var mainURL = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	totalPageNum := getPages(mainURL)
	fmt.Println(totalPageNum)
}

func getPages(mainURL string) int {
	// pages := 0

	res, err := http.Get(mainURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	
	pagination := doc.Find(".tplPagination")
	pagination.Each(func(idx int, sel *goquery.Selection) {
		p_num := sel.Find(".pgTotal").Text()
		fmt.Println(p_num)
	})
	return p_num
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
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id			int		// Attr.data-gno
	company		string	// div.post-list-corp > a.title
	title		string	// div.post-list-info > a.title
	career		string	// span.exp
	location	string	// span.long
	deadline	string	// span.date
}

var mainURL string = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	// lastP := getLastPages(mainURL)

	// for i := 1; i <= lastP; i++{
	// 	extractPage(i)
	// }
	extractPage(1)
}

func extractPage(page int) {
	pageURL := mainURL + "&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	postList := doc.Find("li.list-post") // 20EA

	postList.Each(func(i int, card *goquery.Selection) {
		// id, _ := card.Attr("data-gno")
		// company := strClnr(card.Find("div.post-list-corp > a.title").Text())
		// title := strClnr(card.Find("div.post-list-info > a.title").Text())
		// career := strClnr(card.Find("span.exp").Text())
		// location := card.Find("span.long").Text()
		// deadline := strClnr(card.Find("span.date").Text())
		deadline := strClnr(card.Find("span.date").Text())
		fmt.Println(deadline)
	})
}

// func getLastPages(URL string) int {
// 	pages := ""

// 	res, err := http.Get(URL)
// 	checkErr(err)
// 	checkCode(res)

// 	defer res.Body.Close()
	
// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	checkErr(err)

// 	pages = doc.Find("span.pgTotal").Text()
// 	lastPage := pages[0:3]
// 	intPage, _ := strconv.Atoi(lastPage)
	
// 	return intPage
// }

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

func strClnr(str string) []string { // stringCleaner
	var slice []string
	append(slice, str)
	return strings.TrimSpace(str)
}

// 두가지 선택지
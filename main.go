package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// define extractedJobSet
type extractedJob struct {
	id			string		// Attr.data-gno
	company		string	// div.post-list-corp > a.title
	title		string	// div.post-list-info > a.title
	career		string	// span.exp
	location	string	// span.long
	deadline	string	// span.date
}

var mainURL string = "https://www.jobkorea.co.kr/Search/?stext=개발자&tabType=recruit"

func main() {
	var jobs []extractedJob
	// lastP := getLastPages(mainURL)

	// for i := 1; i <= lastP; i++{
	// 	extractPage(i)
	// }
	extractedjobs := getPage(1)
	jobs = append(jobs, extractedjobs...) // combine two array

	writeJobs(jobs)
	fmt.Println("Extraction Done", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	utf8Applier := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8Applier)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "COMPANY", "LOCATION", "TITLE", "CAREER", "DEADLINE"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.jobkorea.co.kr/Recruit/GI_Read/" + job.id, job.company, job.location, job.title, job.career, job.deadline}
		jobwErr := w.Write(jobSlice)
		checkErr(jobwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob

	pageURL := mainURL + "&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	postList := doc.Find("li.list-post") // 20EA -> 20(what i want)+20(dummy data)

	postList.EachWithBreak(func(i int, card *goquery.Selection) bool {
		company := strClnr(card.Find("div.post-list-corp > a").Text())
		
		job := extractPage(card)
		jobs = append(jobs, job)

		if company != "" {
			return true
		}
		return false // End of EachWithBreak
	})
	return jobs
}

func extractPage(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-gno")
	company := strClnr(card.Find("div.post-list-corp > a").Text())
	location := card.Find("span.long").Text()
	title := strClnr(card.Find("div.post-list-info > a.title").Text())
	career := strClnr(card.Find("span.exp").Text())
	deadline := strClnr(card.Find("p.option > span.date").Text())
	return extractedJob{
		id:			id,
		company:	company,
		location:	location,
		title:		title,
		career:		career,
		deadline:	deadline}
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

func strClnr(str string) string { // stringCleaner
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")// strings.fields() Deleted. // strings.join() Needed.
}

// 두가지 선택지
package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string

// Scrape Indeed by a term
func Scrape(term string) {
	file, err := os.Create("jobs.csv")
	baseURL = fmt.Sprintf("https://kr.indeed.com/jobs?q=%s", term)
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	c := make(chan []extractedJob)
	d := make(chan bool)

	totalPages := 30

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}
	for i := 0; i < totalPages; i++ {
		go writeJobs(<-c, w, d)
	}
	for i := 0; i < totalPages; i++ {
		<-d
	}
	fmt.Println("End")
}

func writeJobs(jobs []extractedJob, w *csv.Writer, d chan bool) {
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
	d <- true
}

func getPage(page int, mainC chan<- []extractedJob) {
	// strconv.Itoa(page * 50)

	c := make(chan extractedJob)

	var jobs []extractedJob
	pageURL := fmt.Sprintf("%s&start=%d", baseURL, page*10)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		jobs = append(jobs, <-c)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := card.Find(".title>a").Text()
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
	/*
		Location 찾기
			card.Find(".sjcl>div").Each(func(i int, s *goquery.Selection) {
				loc, _ := s.Attr("data-rc-loc")
				fmt.Println(loc)
			})
	*/
}

func getPages() int {
	var pages int
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
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

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

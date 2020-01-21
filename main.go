package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://www.tff.org/default.aspx?pageID=198")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTO response body. ", err)
	}

	document.Find("#ctl00_MPane_m_198_1890_ctnr_m_198_1890_Panel1").Each(func(i int, s *goquery.Selection) {

		s.Find("td").Each(func(indextr int, tr *goquery.Selection) {
			tr.Find("a").Each(func(index int, s *goquery.Selection) {
				fmt.Printf("%s\n", s.Text())
			})
		})
	})
}

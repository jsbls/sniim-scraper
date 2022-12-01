package utils

import (
	"fmt"
	"strings"

	"github.com/everitosan/snimm-scrapper/internal"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type requester struct {
	baseUrl string
}

func NewRequester(url string) *requester {
	return &requester{
		baseUrl: url,
	}
}

func (r *requester) SyncR(url string) (e *colly.HTMLElement, err error) {
	currentAttempts := 0
	finalUrl := r.fixUrl(url)
	successChan := make(chan *colly.HTMLElement)
	errorChan := make(chan error)

	for currentAttempts < 5 {

		go request(finalUrl, successChan, errorChan)

		select {
		case err = <-errorChan:
			currentAttempts += 1
			logrus.Warnf("Attempt %d failed for %s", currentAttempts, url)
		case html := <-successChan:
			return html, err
		}
	}

	return e, err
}

// Help fix the url when iframe detected
func (r *requester) fixUrl(url string) string {

	finalUrl := url

	if url[0:1] == "/" {
		finalUrl = r.baseUrl + url
	}

	if strings.Contains(url, "opcion=") {
		endPart := strings.Split(finalUrl, "opcion=")[1]
		basePart := strings.Split(finalUrl, "?")[0]
		end := strings.LastIndex(basePart, "/")

		finalUrl = basePart[:end+1] + endPart
	}

	return finalUrl
}

// Do the request based in colly
func request(url string, successChan chan *colly.HTMLElement, errorChan chan error) {
	c := colly.NewCollector()

	c.OnHTML("html", func(html *colly.HTMLElement) {
		successChan <- html
	})

	c.OnError(func(r *colly.Response, err error) {
		msg := fmt.Errorf("%w: Request URL: %s Error: %v", internal.ErrRequest, r.Request.URL, err)
		errorChan <- msg
	})

	logrus.Debugf("🌏 Requesting %s", url)
	c.Visit(url)

}

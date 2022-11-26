package scrapper

type Scrapper interface {
	RequestFromSource(url string) (SelectOptionsAsMap, error)
}

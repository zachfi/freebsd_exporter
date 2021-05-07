package exporter

type Exporter interface {
	Scrape() error
}

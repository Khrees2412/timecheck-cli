package cmd

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Info struct {
	Time       string
	Date       string
	Population string
}

func FetchInfo(location string) (*Info, error) {
	c := colly.NewCollector()
	var t, d, p string
	var e error

	c.OnError(func(_ *colly.Response, err error) {
		// log.Println("Something went wrong:", err)
		log.Println("Oops! Something went wrong, please wait for a minute and try again.")
		e = err
	})
	if e != nil {
		return nil, e
	}

	c.OnHTML("time", func(e *colly.HTMLElement) {
		t = e.Text
	})

	c.OnHTML(".clockdate", func(e *colly.HTMLElement) {
		d = e.Text
	})
	c.OnHTML("#maptext > ul", func(e *colly.HTMLElement) {
		pop := make([]string, 0)
		e.ForEach("li", func(count int, h *colly.HTMLElement) {
			pop = append(pop, h.Text)

		})
		population := pop[len(pop)-1]
		i := []string{}
		for _, v := range population {
			i = append(i, string(v))
		}
		k := i[11:]
		for _, v := range k {
			p += v
		}

	})
	c.Visit(fmt.Sprintf("https://time.is/%s", location))
	return &Info{
		Time:       t,
		Date:       d,
		Population: p,
	}, nil
}

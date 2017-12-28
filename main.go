package main

import (
    "log"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly"
)

const tokenListURL string = "https://erc20-tokens.eidoo.io"


func main() {
    // Instantiate default collector
    c := colly.NewCollector()

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        log.Println("visiting", r.URL.String())
    })

    // Extract token
    c.OnHTML("#tokensTable tbody", func(e *colly.HTMLElement) {
        e.DOM.Find("tr#coinRow").Each(func(_ int, sel *goquery.Selection) {
            name := sel.Find(".coin h4").Text()
            price := sel.Find(".price h4").Text()

            fmt.Printf("%v %v \n", name, price)
        })        
    })

    c.Visit(tokenListURL)
}


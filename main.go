package main

import (
    "log"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly"
    // For coloring print
    "github.com/fatih/color"
)

const tokenListURL string = "https://erc20-tokens.eidoo.io"

func colorPercentageChanges(amountStr string) string {
    isNegative := string(amountStr[0]) == "-"

    if isNegative {

        red := color.New(color.FgRed).SprintFunc()
        return red(amountStr)

    } else {

        green := color.New(color.FgGreen).SprintFunc()
        return green(amountStr)

    }
}

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
            rawChange := sel.Find(".change h4").Text()

            change := colorPercentageChanges(rawChange)

            fmt.Printf("%v %v %v \n", name, price, change)
        })        
    })

    c.Visit(tokenListURL)
}


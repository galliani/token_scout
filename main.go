package main

import (
    "log"
    "fmt"
    "unicode"    
    "strings"
    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly"
    // Local Dependency
    "./models"
)

const tokenListURL string = "https://erc20-tokens.eidoo.io"
const tokenListStorage string = "https://api.airtable.com/v0/appx0zwzwHaUyBg0g/Utility%20Tokens"


func main() {
    var tokens []models.Token

    // Instantiate default collector
    c := colly.NewCollector()

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        log.Println("visiting", r.URL.String())
    })

    // Extract token
    c.OnHTML("#tokensTable tbody", func(e *colly.HTMLElement) {
        e.DOM.Find("tr#coinRow").Each(func(_ int, sel *goquery.Selection) {
            fullname := sel.Find(".coin h4").Text()
            description := sel.Find(".coin p").Text()

            nameInfos := strings.Split(fullname, " ")
            name := nameInfos[0]
            rawTicker := nameInfos[1]

            ticker :=  strings.TrimFunc(rawTicker, func(c rune) bool {
                return !unicode.IsLetter(c)
            })

            if description != "" {
                tokens = append(tokens, models.Token{Name: name, Ticker: ticker, Description: description})
            }
        })

        for _, elem := range tokens {

            fmt.Printf("%v %v %v \n", elem.Name, elem.Ticker, elem.Description)            

        }        
    })

    c.Visit(tokenListURL)
}


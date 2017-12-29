package main

import (
    "log"
    // "fmt"
    "unicode"    
    "strings"
    
    "bytes"
    "net/http"
    "encoding/json"
    
    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly"
    // Local Dependency
    "./models"
)

const tokenListURL string = "https://erc20-tokens.eidoo.io"
const tokenListStorage string = "https://api.airtable.com/v0/appx0zwzwHaUyBg0g/Utility%20Tokens"


func main() {
    tokens := []models.Token

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

        client := &http.Client{}

        for _, elem := range tokens {

            tokenBody := map[string]interface{}{
                "fields": map[string]string{
                    "Name": elem.Name,
                    "Ticker":  elem.Ticker,
                    "Description": elem.Description,
                },                
            }

            bytesRepresentation, marshallErr := json.Marshal(tokenBody)
            if marshallErr != nil {
                log.Fatal("Error: ", marshallErr)
            }

            req, _ := http.NewRequest("POST", tokenListStorage, bytes.NewBuffer(bytesRepresentation))
            req.Header.Set("Authorization", "Bearer keydQbBOM9VSsBdJT")
            req.Header.Set("Content-type", "application/json")
            
            resp, respErr := client.Do(req)
            if respErr != nil {
                log.Fatalln(respErr)
            }
            // Do not forget to close the response body
            defer resp.Body.Close()

            var result map[string]interface{}

            json.NewDecoder(resp.Body).Decode(&result)
            log.Println(result)
            // fmt.Printf("%v %v %v \n", elem.Name, elem.Ticker, elem.Description)            

        }        
    })

    c.Visit(tokenListURL)
}


package main

import (
    "github.com/gocolly/colly"
)

const tokenListURL string = "https://tokenmarket.net/blockchain/ethereum/assets/"

func main() {
    // Instantiate default collector
    c := colly.NewCollector()


    c.Visit(tokenListURL)
}


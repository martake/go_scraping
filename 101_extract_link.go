package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
)

func main() {
    doc, err := goquery.NewDocument("https://github.com/PuerkitoBio/goquery")
    if err != nil {
        fmt.Print("url scarapping failed")
    }
    doc.Find("a").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.Attr("href")
          fmt.Println(url)
    })
}

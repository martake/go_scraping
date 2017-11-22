package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    fileInfos, _ := ioutil.ReadFile("./html/goquery.html")
    stringReader := strings.NewReader(string(fileInfos))
    doc, err := goquery.NewDocumentFromReader(stringReader)
    if err != nil {
        fmt.Print("url scarapping failed")
    }
    doc.Find("table > tbody > tr > td.content > span > a").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.Attr("href")
          fmt.Println(url)
    })
}

package main

import (
    "github.com/gocolly/colly"
    "fmt"
    "encoding/csv"
    "log"
    "os"
)

func main() {
    fName := "stocks.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
    }

    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    	writer.Write([]string{
            "Symbol", 
            "Name", 
            "Price (Intraday)", 
            "Change", 
            "% Change", 
            "Volume", 
            "Avg Vol (3 month)", 
            "Market Cap", 
            "PE Ratio (TTM)"})

    //fName1 := "losers.csv"
    //fName2 := "most-active.csv"    
    c := colly.NewCollector(
        colly.CacheDir("./yahoo_finance_cache"),
    )

    c.OnHTML(".simpTblRow", func(e *colly.HTMLElement) {
        //link := e.Attr("href")
        writer.Write([]string{
            e.ChildText("td[aria-label='Symbol']"),
            e.ChildText("td[aria-label='Name']"),
            e.ChildText("td[aria-label='Price (Intraday)']"),
            e.ChildText("td[aria-label='Change']"),
            e.ChildText("td[aria-label='% Change']"),
            e.ChildText("td[aria-label='Volume']"),
            e.ChildText("td[aria-label='Avg Vol (3 month)']"),
            e.ChildText("td[aria-label='Market Cap']"),
            e.ChildText("td[aria-label='PE Ratio (TTM)']"),
        })        


        fmt.Println("Symbol:           ", e.ChildText("td[aria-label=Symbol]"))
        fmt.Println("Name:             ", e.ChildText("td[aria-label=Name]"))
        fmt.Println("Price (Intraday): ", e.ChildText("td[aria-label='Price (Intraday)']"))
        fmt.Println("Change:           ", e.ChildText("td[aria-label=Change]"))
        fmt.Println("% Change:         ", e.ChildText("td[aria-label='% Change']"))
        fmt.Println("Volume:           ", e.ChildText("td[aria-label=Volume]"))
        fmt.Println("Avg Vol(3 month): ", e.ChildText("td[aria-label='Avg Vol (3 month)']"))
        fmt.Println("Market Cap:       ", e.ChildText("td[aria-label='Market Cap']"))
        fmt.Println("PE Ratio(TTM):    ", e.ChildText("td[aria-label='PE Ratio (TTM)']"))
        fmt.Println()
        //fmt.Println(e.Text)
        //c.Visit(e.Request.AbsoluteURL(link))
    })    

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })
    
    writer.Write([]string{"Gainers"})
    c.Visit("https://finance.yahoo.com/gainers/")   
    writer.Write([]string{"Losers"})
    c.Visit("https://finance.yahoo.com/losers/")
    writer.Write([]string{"Most-Active"})
    c.Visit("https://finance.yahoo.com/most-active/")
}

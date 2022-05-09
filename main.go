package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)


func check(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func WriteFile(data, filename string){
	file, err := os.Create(filename)
	defer file.Close()
	check(err)

	file.WriteString(data)
}

func main(){
	url := "https://techcrunch.com/"

	response, err := http.Get(url)
	defer response.Body.Close()
	check(err)

	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	check(err)

	file, err := os.Create("posts.csv")
	check(err)

	writer := csv.NewWriter(file)
	doc.Find("div.river").Find("div.post-block ").Each(func (index int, item *goquery.Selection)  {
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url, _ := h2.Find("a").Attr("href")

		excerpt := strings.TrimSpace(item.Find("div.post-bloc__content").Text())
		posts := []string{title, url, excerpt}

		writer.Write(posts)
	})
	writer.Flush()

}
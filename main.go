package main

import (
	"fmt"
	"net/htpp"

	"github.com/PuerkitoBio/goquery"
)


func check(err error){
	if err != nil{
		fmt.Println(err)
	}
}



func main(){
	url := "https://techcrunch.com/"

	response, err := htpp.Get(url)
	defer response.Body.Close()
	check(err)

	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	doc, err := goquery.NewDocument(response.Body)
	check(err)
	river, err := doc.Find("div.river").Html()
	check(err)

	fmt.Println(river)

}
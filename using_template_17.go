package main

import (
	"fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/xml"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
	keyword string
	Location string
}

type NewsAggPage struct{
	Title string
	News map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	var s Sitemapindex
	var n News
	resp, _:= http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	news_map := make(map[string]NewsMap)


	for _, Location := range s.Locations {
		fmt.Println(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.keywords{
			news_map[n.Titles[idx]] = NewsMap{n.keywords[idx], n.Locations[idx]}
		}
		
	}
	p := NewsAggPage{Title: "Amazing news Aggregator", News: news_map}
	t, _ := template.ParseFiles("basic.html")
	t.Execute(w, p)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1> Whoa, GO is neat!</h1>")
}

func main(){
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
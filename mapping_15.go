package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml: "sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewMap struct{
	keyword string
	Location string
}

func main(){
	var s Sitemapindex
	var n News
	nesw_map := make(map[string] NewMap)

	resp, _ := http.get("https://www.bbc.co.uk/sitemaps/https-index-uk-archive.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes,_ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles{
			nesw_map[n.Titles[idx]] = 	NewMap{
				n.keywords[idx],
				n.Locations[idx]
			}
		}
	}
	for idx, data := range nesw_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.keyword)
		fmt.Println("\n", data.Location)
	}
}
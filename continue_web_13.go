package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}


func main(){
	resp, _:= http.Get("https://www.bbc.co.uk/sitemaps/https-index-uk-archive.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}
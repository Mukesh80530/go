package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}

func main(){
	resp, _:= http.Get("https://www.bbc.co.uk/sitemaps/https-index-uk-archive.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
		
	}
}
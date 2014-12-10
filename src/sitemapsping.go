package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// The flag package provides a default help printer via -h switch

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var siteFlag = flag.String("site", "", "any valid domain ")

func main() {
	flag.Parse() // Scan the arguments list

	locale := *localeFlag
	themes := *themesFlag
	site := *siteFlag

	var Url *url.URL

	Url, err := url.Parse("http://www.bing.com")
	if err != nil {
		panic("boom")
	}

	Url.Path += "/ping"

	parameters := url.Values{}
	parameters.Add("sitemap", "http://"+site+"/sitemap.xml")
	Url.RawQuery = parameters.Encode()

	urlstr := Url.String()
	fmt.Println(urlstr)

	client := &http.Client{}

	response, err := client.Get(urlstr)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", urlstr)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}

		fmt.Println(string(contents))

	}

	makeRobotstxt(locale, themes, site)

}

func makeRobotstxt(locale string, themes string, site string) {

	rootdirectory := "/home/juno/git/hugostatic/hugostatic/www/" + locale + "/" + themes + "/" + site
	var buffer bytes.Buffer

	filerobots, err := os.OpenFile(rootdirectory+"/public/robots.txt", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer filerobots.Close()

	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	if _, err = filerobots.Write(buffer.Bytes()); err != nil {
		panic(err)
	}

}

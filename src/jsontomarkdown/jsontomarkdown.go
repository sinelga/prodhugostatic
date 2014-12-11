package main

import (
	//	"bytes"
	"domains"
	//	"encoding/json"
	"findfreeparagraph"
	"flag"
	//	"fmt"
	"io/ioutil"
	"log/syslog"
	//	"os"
	"pager"
	"strings"
	"time"
)

var startparameters []string

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

// The flag package provides a default help printer via -h switch

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var siteFlag = flag.String("site", "", "any valid domain ")

var paragraph domains.Paragraph

func main() {
	flag.Parse() // Scan the arguments list

	//	var d time.Time
	var pubdate string
	var tags []string
	var categories []string

	locale := *localeFlag
	themes := *themesFlag
	site := *siteFlag

	rootdirectory := "/home/juno/git/prodhugostatic/www/" + locale + "/" + themes + "/" + site

	content, err := ioutil.ReadFile("/home/juno/git/prodhugostatic/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
		
	}

	parameters := strings.Split(string(content), ",")
	startparameters = []string{strings.TrimSpace(parameters[0]), strings.TrimSpace(parameters[1]), strings.TrimSpace(parameters[2])}

	paragraph = findfreeparagraph.FindFromQ(*golog, locale, themes, "google", startparameters)

	pubdate = time.Now().Local().Format(time.RFC3339)

	tags = []string{
		paragraph.Phost,
		strings.ToLower(strings.Split(paragraph.Ptitle, " ")[0]),
		strings.ToLower(strings.Split(paragraph.Ptitle, " ")[1]),
	}

	categories = []string{
		paragraph.Phost,
		strings.ToLower(strings.Split(paragraph.Ptitle, " ")[0]),
		strings.ToLower(strings.Split(paragraph.Ptitle, " ")[1]),
	}

	frontmatter := domains.Frontmatter{

		Title:       paragraph.Ptitle + ".",
		Description: paragraph.Pphrase,
		Date:        pubdate,
		Tags:        tags,
		Categories:  categories,
		Slug:        tags[1] + "-" + tags[2],
		Sentences:   paragraph.Sentences,
	}

	pager.CreatePage(*golog, rootdirectory+"/content/post/"+strings.Replace(paragraph.Ptitle, " ", "-", 1)+".md", frontmatter)



}

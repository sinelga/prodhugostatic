package main

import (
	"flag"
	//    "fmt"
	"domains"
	"findfreeparagraph"
	"io/ioutil"
	"log/syslog"
	"strings"
	"createsite"
)


var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var siteFlag = flag.String("site", "", "any valid domain ")
var cssthemesFlag = flag.String("cssthemes", "", "css themes directory  ")

var startparameters []string

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

var paragraphs []domains.Paragraph

func main() {
	flag.Parse() // Scan the arguments list

	locale := *localeFlag
	themes := *themesFlag
	site := *siteFlag
	cssthemes :=*cssthemesFlag 

	content, err := ioutil.ReadFile("/home/juno/git/prodhugostatic/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
	}
	parameters := strings.Split(string(content), ",")
	startparameters = []string{strings.TrimSpace(parameters[0]), strings.TrimSpace(parameters[1]), strings.TrimSpace(parameters[2])}

	for i := 0; i < 31; i++ {

		paragraph := findfreeparagraph.FindFromQ(*golog, locale, themes, "google", startparameters)
		paragraphs = append(paragraphs, paragraph)

	}
	
	createsite.CreateNewSite(*golog,locale,themes,site,cssthemes,paragraphs)
			
}

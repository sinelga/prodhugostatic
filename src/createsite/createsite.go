package createsite

import (
	"bytes"
	"createpages"
	"domains"
	"fmt"
	"github.com/BurntSushi/toml"
	"log/syslog"
	"os"
	"os/exec"
	"strings"
)

func CreateNewSite(golog syslog.Writer, locale string, themes string, site string, cssthemes string, paragraphs []domains.Paragraph) {

	rootdirectory := "/home/juno/git/prodhugostatic/www/" + locale + "/" + themes + "/" + site

	if _, err := os.Stat(rootdirectory); os.IsNotExist(err) {

		if os.MkdirAll(rootdirectory, 0777) != nil {
			golog.Err(err.Error())

		}

		cmd := exec.Command("hugo", "new", "site", rootdirectory)

		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {

			golog.Err(err.Error())

		}
		fmt.Printf(" %q\n", out.String())

		var description []string
		var keywords []string

		for _, paragraph := range paragraphs {

			description = append(description, paragraph.Pphrase)
			keyword0 := strings.ToLower(strings.Split(paragraph.Ptitle, " ")[0])
			keyword1 := strings.ToLower(strings.Split(paragraph.Ptitle, " ")[1])

			keywords = append(keywords, keyword0)
			keywords = append(keywords, keyword1)

		}

		paramshugo := domains.Paramshugo{

			Locale:      locale,
			Themes:      themes,
			Site:        site,
			Cssthemes:   cssthemes,
			Description: description,
			Keywords:    keywords,
		}

		indexeshugo := domains.Indexeshugo{
			Category:    "categories",
			Tag:         "tags",
			Description: "descriptions",
		}

		permalinkshugo := domains.Permalinkshugo{
			Post: "/:slug/",
			
		}

		menuarr := []domains.Menuparshugo{

			domains.Menuparshugo{Name: "About", Url: "/about/", Weight: -100, Identifier: "about"},
			domains.Menuparshugo{Name: "Posts", Url: "/post/", Weight: -105, Identifier: "posts"},
			domains.Menuparshugo{Name: "Etusivu", Url: "/", Weight: -106, Identifier: "homepage"},
			//			domains.Menuparshugo {Name: "Rss",Url: "/rss.xml",Weight: -101,Identifier:"rss"},
			domains.Menuparshugo{Name: "All categories", Url: "/allcategories/", Weight: -102, Identifier: "allcategories"},
			domains.Menuparshugo{Name: "All descriptions", Url: "/alldescriptions/", Weight: -103, Identifier: "alldescriptions"},
			domains.Menuparshugo{Name: "All tags", Url: "/alltags/", Weight: -104, Identifier: "alltags"},
		}

		confmenu := domains.Menuhugo{

			Menu: menuarr,
		}

		configtoml := domains.Confighugo{
			LanguageCode: locale,
			Baseurl:      "http://" + site,
			Canonifyurls: true,
			Title:        paragraphs[0].Ptitle,
			Theme:        cssthemes,
			Indexes:      indexeshugo,
			Permalinks:   permalinkshugo,
			Params:       paramshugo,
			Menu:         confmenu,
		}

		var buffer bytes.Buffer
		encoder := toml.NewEncoder(&buffer)
		err = encoder.Encode(configtoml)
		if err != nil {
			panic(err)
		}

		file, err := os.OpenFile(rootdirectory+"/config.toml", os.O_RDWR|os.O_CREATE, 0660)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		if _, err = file.Write(buffer.Bytes()); err != nil {
			panic(err)
		}
		
		buffer.Reset()
		
		filerobots, err := os.OpenFile(rootdirectory+"/robots.txt", os.O_RDWR|os.O_CREATE, 0660)
				if err != nil {
			panic(err)
		}
		defer filerobots.Close()
		
		buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://"+site+"/sitemap.xml")
		
		if _, err = filerobots.Write(buffer.Bytes()); err != nil {
			panic(err)
		}
		
		createpages.CreateSomePages(golog, rootdirectory, paragraphs)

		if _, err := os.Stat(rootdirectory + "/themes"); os.IsNotExist(err) {

			if os.MkdirAll(rootdirectory+"/themes", 0777) != nil {
				golog.Err(err.Error())

			}
		}

		srcFolder := "/home/juno/git/prodhugothemes/" + cssthemes
		destFolder := rootdirectory + "/themes"

		cmd = exec.Command("ln", "-s", srcFolder, destFolder)
		err = cmd.Run()
		if err != nil {

			golog.Err(err.Error())

		}

		cmd = exec.Command("hugo", "-s", rootdirectory, "-t", cssthemes)

		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {

			golog.Err(err.Error())

		}
		fmt.Printf(" %q\n", out.String())
		
				

	} else {

		fmt.Println(" directory exist???", rootdirectory)

	}

}

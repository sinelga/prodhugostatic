package createpages

import (
	"domains"
	"log/syslog"
	"pager"
	"strings"
	"time"
)

func CreateSomePages(golog syslog.Writer, rootdirectory string, paragraphs []domains.Paragraph) {

	var d time.Time
	var pubdate string
	var tags []string
	var categories []string

	d = time.Now()

	//	about
	pubdate = d.Local().Format(time.RFC3339)

	frontmatter := domains.Frontmatter{

		Title:       "About",
		Description: paragraphs[0].Pphrase,
		Date:        pubdate,
		Tags:        []string{"about", "Noin"},
		Categories:  []string{"about", "Noin"},
		Descriptions: []string{paragraphs[0].Pphrase},
		Slug:        "about",
		Sentences:   paragraphs[0].Sentences,
	}

	pager.CreatePage(golog, rootdirectory+"/content/about.md", frontmatter)

	// alltags

	frontmatter = domains.Frontmatter{

		Title:       "All tags",
		Description: "All tags Index",
		Date:        pubdate,
		Tags:        []string{"tag", "tags", "index"},
		Categories:  []string{"tag", "tags", "index"},
		Descriptions: []string{"All tags Index"},
		Slug:        "alltags",
		Sentences:   []string{"All tags Index"},
	}

	pager.CreatePage(golog, rootdirectory+"/content/alltags.md", frontmatter)

	//	allcategories

	frontmatter = domains.Frontmatter{

		Title:       "All categories",
		Description: "All categories Index",
		Date:        pubdate,
		Tags:        []string{"categories", "index"},
		Categories:  []string{"categories", "index"},
		Descriptions: []string{"All categories Index"},
		Slug:        "allcategories",
		Sentences:   []string{"All categories Index"},
	}

	pager.CreatePage(golog, rootdirectory+"/content/allcategories.md", frontmatter)
	
	
	//	allcategories

	frontmatter = domains.Frontmatter{

		Title:       "All descriptions",
		Description: "All descriptions Index",
		Date:        pubdate,
		Tags:        []string{"descriptions", "index"},
		Categories:  []string{"descriptions", "index"},
		Descriptions: []string{"All descriptions Index"},
		Slug:        "alldescriptions",
		Sentences:   []string{"All descriptions Index"},
	}

	pager.CreatePage(golog, rootdirectory+"/content/alldescriptions.md", frontmatter)	
		


	for i, paragraph := range paragraphs {

		pubdate = time.Date(d.Year(), d.Month(), d.Day()-i, 0, 0, 0, 0, time.UTC).Local().Format(time.RFC3339)

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
			Descriptions: []string{ strings.Replace(paragraph.Pphrase,".","",1)},
			Slug:        tags[1] + "-" + tags[2],
			Sentences:   paragraph.Sentences,
		}

		pager.CreatePage(golog, rootdirectory+"/content/post/"+strings.Replace(paragraph.Ptitle, " ", "-", 1)+".md", frontmatter)

	}

}

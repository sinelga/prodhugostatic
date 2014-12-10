package domains

import ()

type Paragraph struct {
	Ptitle     string
	Pphrase    string
	Plocallink string
	Phost      string
	Sentences  []string
	Pushsite   string
}


type Menuparshugo struct {
		
	Name string `toml:"name"`
	Url string `toml:"url"`
	Weight int `toml:"weight"`
	Identifier string `toml:"identifier"`
			
}

type Menuhugo struct {
	
	Menu []Menuparshugo `toml:"main"`
	
}


type Paramshugo struct {
	Locale      string   `toml:"locale"`
	Themes      string   `toml:"themes"`
	Site        string   `toml:"site"`
	Cssthemes   string   `toml:"cssthemes"`
	Description []string `toml:"description"`
	Keywords    []string `toml:"keywords"`
}

type Permalinkshugo struct {
	Post string `toml:"post"`
	
}

type Indexeshugo struct {
	Category    string `toml:"category"`
	Tag         string `toml:"tag"`
	Description string `toml:"description"`
}

type Confighugo struct {
	LanguageCode string         `toml:"languageCode"`
	Baseurl      string         `toml:"baseurl"`
	Canonifyurls bool           `toml:"canonifyurls"`
	Title        string         `toml:"title"`
	Theme        string         `toml:"theme"`
	Indexes      Indexeshugo    `toml:"indexes"`
	Permalinks   Permalinkshugo `toml:"permalinks"`
	Params       Paramshugo     `toml:"params"`
	Menu	Menuhugo `toml:"menu"`
}

type Frontmatter struct {
	Title        string   `toml:"title"`
	Description  string   `toml:"description"`
	Date         string   `toml:"date"`
	Tags         []string `toml:"tags"`
	Categories   []string `toml:"categories"`
	Descriptions []string `toml:"descriptions"`
	Slug         string   `toml:"slug"`
	Sentences    []string `toml:"sentences"`
}

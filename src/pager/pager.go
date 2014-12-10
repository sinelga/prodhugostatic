package pager

import (
	"bytes"
	"domains"
	"github.com/BurntSushi/toml"
	"log/syslog"
	"os"
//	"strings"
	"path"
)

func CreatePage(golog syslog.Writer, filestr string, frontmatter domains.Frontmatter) {

	var buffer bytes.Buffer
	
	buffer.WriteString("+++\n")
	
			
		encoder := toml.NewEncoder(&buffer)
		err := encoder.Encode(frontmatter)
		if err != nil {
			
			golog.Err(err.Error())
		}
	buffer.WriteString("\n+++\n")


	for _, line := range frontmatter.Sentences {

		buffer.WriteString(line + " ")

	}
	
	dir := path.Dir(filestr)
	
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		
		if os.MkdirAll(dir, 0777) != nil {
			golog.Err(err.Error())

		}
		
	}	

	file, err := os.OpenFile(filestr, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(buffer.String()); err != nil {
		panic(err)
	}

}

package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"zhihu-hot-questions/model"
)

func CreateReadMe(data []model.Question) {
	fileName := "README.md"
	if file, err := ioutil.ReadFile(fileName); err == nil {
		reg,_ := regexp.Compile(`<!-- BEGIN -->[\W\w]*<!-- END -->`)
		allString := reg.ReplaceAllString(string(file),CreateList(data))
		if writeFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm); err == nil {
			defer writeFile.Close()
			// 先清空再替换
			writeFile.WriteString("")
			log.Fatalln(allString)
			writeFile.WriteString(allString)
		}
	}
}

func CreateArchives(data []model.Question, fileName string) {
	filePath := fmt.Sprintf("./archives/%v.md", fileName)
	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		content := fmt.Sprintf("# %v\n 共%v条\n %v", fileName, len(data), CreateList(data))
		file.WriteString(content)
	}
}

func CreateList(data []model.Question) string {
	var word string
	for _, v := range data {
		word += fmt.Sprintf("1. [%v](%v)\n", v.Title, v.Url)
	}
	template := fmt.Sprintf("<!-- BEGIN -->\n\n%v\n<!-- END -->", word)
	return template
}

package main

import (
	"encoding/json"
	"fmt"
	"go-weather-email/model"
	"go-weather-email/utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var fileName = time.Now().Format("2006-1-2")

func main() {
	result := getZhiHuHot()
	question := getQuestion(result)
	createRaw(question, fileName)
	utils.CreateReadMe(question)
	utils.CreateArchives(question,fileName)
}

func createRaw(data []model.Question, fileName string) {
	filePath := fmt.Sprintf("./raw/%v.json", fileName)
	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		bytes, _ := json.Marshal(data)
		file.Write(bytes)
	}
}

func getQuestion(result *model.HotList) []model.Question {
	var questionList []model.Question
	for _, v := range result.Data {
		question := model.Question{
			Title: v.Target.Title,
			Url:   fmt.Sprintf("https://www.zhihu.com/question/%v", v.Target.Id),
		}
		questionList = append(questionList, question)
	}
	return questionList
}

func getZhiHuHot() *model.HotList {
	resp, _ := http.Get("https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=100")
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var list model.HotList
		if err := json.Unmarshal(body, &list); err == nil {
			return &list
		}
	}
	return nil
}

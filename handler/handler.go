package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data []CovidData `json:"Data"`
}

type CovidData struct {
	ConfirmDate    string `json:"ConfirmDate"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
}

type ResponseData struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

func CovidSummary(c *gin.Context) {
	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var rawData Response

	json.Unmarshal(body, &rawData)

	datas := rawData.Data

	province := map[string]int{
		"N/A": 0,
	}
	ageGroup := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	for _, data := range datas {
		if data.ProvinceEn == "" {
			province["N/A"]++
		} else {
			if _, ok := province[data.ProvinceEn]; !ok {
				province[data.ProvinceEn] = 1
			}
			province[data.Province]++
		}

		if data.Age == 0 {
			ageGroup["N/A"]++
		} else if data.Age <= 30 {
			ageGroup["0-30"]++
		} else if data.Age <= 60 {
			ageGroup["31-60"]++
		} else {
			ageGroup["61+"]++
		}
	}

	responseData := ResponseData{
		Province: province,
		AgeGroup: ageGroup,
	}

	c.JSON(200, responseData)
}

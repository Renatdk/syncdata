package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	s "scsyncs/synccommoditymarkets/services"
	l "scsyncs/synccommoditymarkets/services/logging"
	t "scsyncs/synccommoditymarkets/types"
)

func main() {
	l.InitLogging()
	l.Trace.Println("Запуск приложения")

	file, e := ioutil.ReadFile("./configs/app.json")
	checkErr(e)
	var appConfig t.AppConfigType
	json.Unmarshal(file, &appConfig)

	for {
		updateData()
		time.Sleep(time.Duration(appConfig.TimeOutHour) * time.Hour)
	}
}

func updateData() {
	l.Trace.Println("Обновление")

	file, e := ioutil.ReadFile("./configs/config.json")
	checkErr(e)
	var jsontype t.MainType
	json.Unmarshal(file, &jsontype)

	for _, element := range jsontype.Dataset.Data {
		url := jsontype.Dataset.URL + element[2] + ".json?api_key=" + jsontype.Dataset.APIKey + "&start_date=" + jsontype.Dataset.StartDate + "&end_date=" + jsontype.Dataset.EndDate

		resp, err := http.Get(url)
		checkErr(err)
		defer resp.Body.Close()

		jsonElementType := t.RemoteMainType{}

		body, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(body, &jsonElementType)
		i, _ := strconv.Atoi(element[3])

		// fmt.Println(element[1])
		// fmt.Println(jsonElementType.Dataset.Data[0][0])
		// fmt.Println(jsonElementType.Dataset.Data[0][i])

		if !s.IsHaveData(element[1], string(jsonElementType.Dataset.Data[0][0])) {
			s.AddData(element[1], jsonElementType.Dataset.Data[0][0], jsonElementType.Dataset.Data[0][i])
		}

	}
}

func checkErr(err error) {
	if err != nil {
		l.Error.Println(err)
	}
}

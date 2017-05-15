package types

import "encoding/json"

type MainType struct {
	Dataset DatasetType `json:"dataset"`
}

type DatasetType struct {
	URL         string   `json:"url"`
	APIKey      string   `json:"api_key"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	ColumnNames []string `json:"column_names"`
	Data        [][]string
}

type RemoteMainType struct {
	Dataset RemoteDatasetType `json:"dataset"`
}

type RemoteDatasetType struct {
	Data [][]json.Number
}

type AppConfigType struct {
	TimeOutHour int `json:"time-out-hour"`
}

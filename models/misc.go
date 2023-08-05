package models

import "time"

type PHONE struct {
	ID        string `json:"ID"`
	ValueType string `json:"VALUE_TYPE"`
	Value     string `json:"VALUE"`
	TypeID    string `json:"TYPE_ID"`
}

type EMAIL struct {
	ID        string `json:"ID"`
	ValueType string `json:"VALUE_TYPE"`
	Value     string `json:"VALUE"`
	TypeID    string `json:"TYPE_ID"`
}

type WEB struct {
	ID        string `json:"ID"`
	ValueType string `json:"VALUE_TYPE"`
	Value     string `json:"VALUE"`
	TypeID    string `json:"TYPE_ID"`
}

type IM struct {
	ID        string `json:"ID"`
	ValueType string `json:"VALUE_TYPE"`
	Value     string `json:"VALUE"`
	TypeID    string `json:"TYPE_ID"`
}

type Time struct {
	Start            float64   `json:"start"`
	Finish           float64   `json:"finish"`
	Duration         float64   `json:"duration"`
	Processing       float64   `json:"processing"`
	DateStart        time.Time `json:"date_start"`
	DateFinish       time.Time `json:"date_finish"`
	OperatingResetAt int       `json:"operating_reset_at"`
	Operating        float64   `json:"operating"`
}

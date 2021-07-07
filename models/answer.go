package models

import "time"

type Answer struct {
	Id            uint64    `json:"id"`
	Question_id   uint64    `json:"question_id"`
	Answerer_id   uint64    `json:"answerer_id"`
	Answerer_name string    `json:"answerer_name"`
	Content       string    `json:"content"`
	Date_time     time.Time `json:"date_time"`
}

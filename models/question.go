package models

import "time"

type Question struct {
	Id              uint64    `json:"id"`
	Questioner_id   string    `json:"questioner_id"`
	Questioner_name string    `json:"questioner_name"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Upvotes         uint64    `json:"upvotes"`
	Downvotes       uint64    `json:"downvotes"`
	Date_time       time.Time `json:"date_time"`
}

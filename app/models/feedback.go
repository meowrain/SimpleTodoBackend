package models

import "gorm.io/gorm"

type FeedBack struct {
	gorm.Model
	Helpful  int `json:"helpful"`
	Helpless int `json:"helpless"`
}

type Comment struct {
	gorm.Model
	Text string `json:"text"`
}

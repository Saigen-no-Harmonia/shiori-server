package models

import "time"

type PresenterProfile struct {
	kaedeFlag     int8      `json:"kaedeFlag"`
	ieId          string    `json:"ieId"`
	photoUrl      string    `json:"photoUrl"`
	lastName      string    `json:"lastName"`
	firstName     string    `json:"firstName"`
	lastNameKana  string    `json:"lastNameKana"`
	firstNameKana string    `json:"firstNameKana`
	birthDate     time.Time `json:"birthDate"`
	birthPlace    string    `json:"birthPlace"`
	job           string    `json:"job"`
	hobby         string    `json:"hobby"`
	ramen         string    `json:"ramen"`
	nickName      string    `json:"nickName"`
	likeBy        string    `json:"likeBy"`
}

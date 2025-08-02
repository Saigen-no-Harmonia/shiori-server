package models

import "time"

type AccessInfo struct {
	VenueId            string    `json:"venueId" example:"LC001"`
	VenueName          string    `json:"venueName" example:"広い会場"`
	VenueAddress       string    `json:"venueAddress" example:"千葉県松戸市上本郷1-2-999"`
	VenueAccessPageUrl string    `json:"venueAccessPageUrl" example:"https://kaijou/access"`
	Latitude           string    `json:"latitude" example:"40.999999"`
	Longitude          string    `json:"longitude" example:"140.999999"`
	GatheringSpot      string    `json:"gatheringSpot" example:"松戸新田駅の近くのいなげや"`
	GatheringDate      time.Time `json:"gatheringDate" example:"2000-01-31T00:00:00+09:00"`
	StartingDate       time.Time `json:"startingDate" example:"2000-01-31T00:15:00+09:00"`
	RestaurantName     string    `json:"restaurantName" example:"oisiiRestaurant"`
	RestaurantUrl      string    `json:"restaurantUrl" example:"https://oisiiRestaurant.com"`
}

func NewAccessInfo(
	venueId string,
	venueName string,
	venueAddress string,
	venueAccessPageUrl string,
	latitude string,
	longitude string,
	gatheringSpot string,
	gatheringDate time.Time,
	startingDate time.Time,
	restaurantName string,
	restaurantUrl string,
) *AccessInfo {
	a := new(AccessInfo)
	a.VenueId = venueId
	a.VenueName = venueName
	a.VenueAddress = venueAddress
	a.VenueAccessPageUrl = venueAccessPageUrl
	a.Latitude = latitude
	a.Longitude = longitude
	a.GatheringSpot = gatheringSpot
	a.GatheringDate = gatheringDate
	a.StartingDate = startingDate
	a.RestaurantName = restaurantName
	a.RestaurantUrl = restaurantUrl

	return a
}

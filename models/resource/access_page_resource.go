package resource

import (
	"shiori-server/models"
	"shiori-server/models/util"
	"time"
)

type AccessPageResource struct {
	VenueAddress       string       `json:"venueAddress" example:"千葉県松戸市上本郷1-2-999"`
	VenueAccessPageUrl string       `json:"venueAccessPageUrl" example:"https://kaijou/access"`
	Latitude           string       `json:"latitude" example:"40.999999"`
	Longitude          string       `json:"longitude" example:"140.999999"`
	GatheringSpot      string       `json:"gatheringSpot" example:"松戸新田駅の近くのいなげや"`
	GatheringDate      util.JSTTime `json:"gatheringDate" example:"2000-01-31T00:00:00+09:00"`
	StartingDate       util.JSTTime `json:"startingDate" example:"2000-01-31T00:15:00+09:00"`
	// GatheringDate time.Time `json:"gatheringDate" example:"2000-01-31T00:00:00+09:00"`
	// StartingDate   time.Time `json:"startingDate" example:"2000-01-31T00:15:00+09:00"`
	RestaurantName string `json:"restaurantName" example:"oisiiRestaurant"`
	RestaurantUrl  string `json:"restaurantUrl" example:"https://oisiiRestaurant.com"`
}

// コンストラクタ
func NewAccessPageResource(
	venueAddress string,
	venueAccessPageUrl string,
	latitude string,
	longitude string,
	gatheringSpot string,
	gatheringDate time.Time,
	startingDate time.Time,
	restrauntName string,
	restaurantUrl string,
) *AccessPageResource {
	a := new(AccessPageResource)
	a.VenueAddress = venueAddress
	a.VenueAccessPageUrl = venueAccessPageUrl
	a.Latitude = latitude
	a.Longitude = longitude
	a.GatheringSpot = gatheringSpot
	a.GatheringDate = util.JSTTime{Time: gatheringDate}
	a.StartingDate = util.JSTTime{Time: startingDate}
	a.RestaurantName = restrauntName
	a.RestaurantUrl = restaurantUrl
	return a
}

// ModelをResourceにマップ
func MapToAccessPageResource(
	p models.AccessInfo,
) *AccessPageResource {
	return NewAccessPageResource(
		p.VenueAddress,
		p.VenueAccessPageUrl,
		p.Latitude,
		p.Longitude,
		p.GatheringSpot,
		p.GatheringDate,
		p.StartingDate,
		p.RestaurantName,
		p.RestaurantUrl,
	)
}

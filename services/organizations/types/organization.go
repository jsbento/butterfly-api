package types

import (
	"time"
)

type Organization struct {
	Id                       string       `bson:"_id" json:"id"`
	Name                     string       `bson:"name" json:"name"`
	OrgUrl                   string       `bson:"orgUrl" json:"orgUrl"`
	Username                 string       `bson:"username" json:"username"`
	Location                 *Location    `bson:"location" json:"location"`
	WebSettings              *WebSettings `bson:"webSettings" json:"webSettings"`
	Deleted                  bool         `bson:"deleted" json:"deleted"`
	Suspended                bool         `bson:"suspended" json:"suspended"`
	FeaturedButterflyDate    time.Time    `bson:"featuredButterflyDate" json:"featuredButterflyDate"`
	FeaturedButterflyId      string       `bson:"featuredButterflyId" json:"featuredButterflyId"`
	DisplayFeaturedButterfly bool         `bson:"displayFeaturedButterfly" json:"displayFeaturedButterfly"`
	DisplayHomeStats         bool         `bson:"displayHomeStats" json:"displayHomeStats"`
	CreatedAt                time.Time    `bson:"createdAt" json:"createdAt"`
	UpdatedAt                time.Time    `bson:"updatedAt" json:"updatedAt"`
}

type Location struct {
	Address   string  `bson:"address" json:"address"`
	City      string  `bson:"city" json:"city"`
	State     string  `bson:"state" json:"state"`
	Zip       string  `bson:"zip" json:"zip"`
	Country   string  `bson:"country" json:"country"`
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}

type WebSettings struct {
	HeaderColor             string `bson:"headerColor" json:"headerColor"`
	SectionHeaderColor      string `bson:"sectionHeaderColor" json:"sectionHeaderColor"`
	MenuColor               string `bson:"menuColor" json:"menuColor"`
	LinkFontColor           string `bson:"linkFontColor" json:"linkFontColor"`
	AdminIconColor          string `bson:"adminIconColor" json:"adminIconColor"`
	HomepageBackgroundImage string `bson:"homepageBackgroundImage" json:"homepageBackgroundImage"`
	Font                    string `bson:"font" json:"font"`
	LogoURL                 string `bson:"logoURL" json:"logoURL"`
	CoverMedia              string `bson:"coverMedia" json:"coverMedia"`
}

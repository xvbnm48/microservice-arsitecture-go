package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	URL         string        `bson:"url" json:"url"`
	CoverImage  string        `bson:"coverImage" json:"coverImage"`
	Description string        `bson:"description" json:"description"`
}

// add movie information

type AddMovie struct {
	Name        string `json:"name" example:"Movie Name"`
	URL         string `json:"url" example:"Movie URL"`
	CoverImage  string `json:"coverImage" example:"Movie Cover Image"`
	Description string `json:"description" example:"Movie Description"`
}

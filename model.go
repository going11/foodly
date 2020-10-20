package main

type SubmitFoodRequest struct {
	name           string `json:"_id,omitempty" bson:"_id,omitempty"`
	restaurant     string `json:"_id,omitempty" bson:"_id,omitempty"`
	desc           string `json:"_id,omitempty" bson:"_id,omitempty"`
	price          string `json:"_id,omitempty" bson:"_id,omitempty"`
	image_filename string `json:"_id,omitempty" bson:"_id,omitempty"`
	username       string `json:"_id,omitempty" bson:"_id,omitempty"`
	token          string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type RegisterRequest struct {
	username string `json:"_id,omitempty" bson:"_id,omitempty"`
	password string `json:"_id,omitempty" bson:"_id,omitempty"`
	email    string `json:"_id,omitempty" bson:"_id,omitempty"`
	role     string `json:"_id,omitempty" bson:"_id,omitempty"`
	city     string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type LoginRequest struct {
	username string `json:"_id,omitempty" bson:"_id,omitempty"`
	password string `json:"_id,omitempty" bson:"_id,omitempty"`
}

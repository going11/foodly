package main

type UserModel struct {
  username string `bson:"username"`
  password string `bson:"password"`
  email string `bson:"email"`
  role string `bson:"role"`
  city string `bson:"city"`
  token string `bson:"token"`
}

type RegisterRequest struct {
  username string `json:"_id,omitempty" bson:"_id,omitempty"`
  password string `json:"_id,omitempty" bson:"_id,omitempty"`
  email string `json:"_id,omitempty" bson:"_id,omitempty"`
  role string `json:"_id,omitempty" bson:"_id,omitempty"`
  city string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type LoginRequest struct {
  username string `json:"_id,omitempty" bson:"_id,omitempty"`
  password string `json:"_id,omitempty" bson:"_id,omitempty"`
}

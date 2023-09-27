package models

type Buyersmodel struct {
	Name        string `json:"name" bson:"name"`
	Mail        string `json:"mail" bson:"mail"`
	Age         int    `json:"age" bson:"age"`
	Phonenumber int    `json:"phonenumber" bson:"phonenumber"`
	Password    string `json:"password" bson:"password"`
}


type DBrespone struct {
	Name string `json:"name" bson:"name"`
}

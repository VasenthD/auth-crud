package models

type Modelsinfo struct {
	Name        string `json:"name", bson:"name"`
	Mail        string `json:"mail", bson:"mail"`
	age         int    `json:"age", bson:"age"`
	Phonenumber int    `json:"phonenumber", bson:"phonenumber"`
	password    string `json:"password", bson:"password"`
}

type DBrespone struct {
	Name string `json:"name", bson:"name"`
}

package cronsun

import (
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Coll_Business = "business"
)

//从 db 获取所有businessType
func GetBusinessTypeFromMongo() (businessType []int8, err error) {
	err = mgoDB.FindOne(Coll_Business, bson.M{"status":0}, businessType)
	return 
}
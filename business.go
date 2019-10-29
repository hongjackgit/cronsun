package cronsun

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Coll_Business = "business"
)

type Business struct {
	Id int64
	Name string
	State int8
}


func CreateBusiness() *Business{
	return &Business{}
}

//从 db 获取所有businessType
func GetBusinessTypeFromMongo() (businessType []string, err error) {
	// err = mgoDB.FindOne(Coll_Business, bson.M{"status":0}, businessType)
	businessType = []string{"task","cms","platform"}
	return 
}

//从 db 获取所有businessType
func (business *Business) AddBusinessType(Name string) (businessType string, err error) {
	// err = mgoDB.FindOne(Coll_Business, bson.M{"status":0}, businessType)
	business.Id = bson.NewObjectId()
	business.Name = Name
	business.State = 1
	err = mgoDB.Insert(Coll_Business,business);
	return 
}
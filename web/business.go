package web

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Coll_Business = "business"
)

type Business struct {
	Id bson.ObjectId
	Name string
	State int8
}

type Key struct {
	Id bson.ObjectId
	Bid bson.ObjectId
	Content string
	State int8
}



//从 db 获取所有business
func GetAllBusiness() (business []string, err error) {
	err = mgoDB.FindOne(Coll_Business, bson.M{"status":0}, business)
	return 
}

//从 db 获取name的business
func GetAllBusiness(name string) (business []string, err error) {
	err = mgoDB.FindOne(Coll_Business, bson.M{"name":name}, business)
	return 
}


//删除业务TSK
func (business *Business) delBusiness(name string) error {

	return
}

func CreateKey(business *Business) *Key{
	return &Key{"Bid":business.Id,"State":0}
}

//配置KEY，及内容
func (key *Key) addKey(key string, content string, business *Business) error {
	key.Id = bson.NewObjectId()
	key.Bid = business.Id
	key.Content = content
	key.State = 1
	err := mgoDB.Insert(Coll_Business,key);
	return err 
}


//从 db 获取所有businessType
func (business *Business) AddBusiness(Name string) (businessType string, err error) {
	// err = mgoDB.FindOne(Coll_Business, bson.M{"status":0}, businessType)
	business.Id = bson.NewObjectId()
	business.Name = Name
	business.State = 1
	err = mgoDB.Insert(Coll_Business,business);
	return 
}

//上线操作
func (business *Business) StartKey(key string, content string) error {
	if err := mgoDB.Upsert(Coll_Business, bson.M{"node": jl.Node, "hostname": jl.Hostname, "ip": jl.IP, "jobId": jl.JobId, "jobGroup": jl.JobGroup}, latestLog); err != nil {
		log.Errorf(err.Error())
	}
	return
}

//回滚操作
func (business *Business) RollbackKey(key string, content string) error {
	return
}


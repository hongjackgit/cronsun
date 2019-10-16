package cronsun

import (
	"os"
	"gopkg.in/mgo.v2/bson"
)

const (
	Coll_Etc = "etc"
)


type Etc struct {
	ID       string `bson:"_id" json:"id"`
	Content  string `bson:"content" json:"content"`
	CreateDate       string `bson:"create_date" json:"create_date"`
	User       string `bson:"user" json:"user"`
	State int8 `bson:"state" json:"state"`
}

func CreateNewEtc() *Etc {
	return &Etc{}
}

//从 db 获取version对应的配置文件
func GetEtcInfoFromMongo(id int64,businessType int8) (etclistInfo []*Etc, err error) {
	err = mgoDB.FindOne(Coll_Etc, bson.M{"id": id,"type":businessType}, etclistInfo)
	return 
}

//从 etcd 获取正在version号，启动时候
func GetEtcVidFromEetd(businessType int8) (etcInfo *Etc) {
	// client.G
	return
}

//写入etcinfo到本地文件
func (e *Etc) WriteEtcInfo(vid int64, businessType int8, etcInfo *Etc) (err error) {
	cVidPath := "/cronsun/ab/type/current_version.id"
	etcPath := "/cronsun/ab/type/version.json"
	var cVid int64 = 0
    if cVid, err := os.Stat(cVidPath); err !=nil {
        panic(err)
    }
    if cVid, err := os.Stat(etcPath); err !=nil {
        panic(err)
    }
    fmt.Println(cVid)
    if(cVid != vid){
    	etcCvidFile , err := os.Open(cVidPath)
    	defer etcCvidFile.Close()
    	if (err != nil) {
    		panic(err)
    	}
    	_, err = etcCvidFile.WriteString("{cvid:1}")
    	if (err != nil){
    		panic(err)
    	}
    	etcFile , err := os.Open(etcPath)
    	defer etcFile.Close()
    	if (err != nil) {
    		panic(err)
    	}
    	_, err = etcFile.WriteString("AAAAAA")
    	if (err != nil){
    		panic(err)
    	}
    }
	return 
}
//修改etcinfo文件
func (e *Etc) UpdateEtcInfo(vid int64,businessType int8) (err error) {
	return
}
//回滚到指定版本
func (e *Etc) RollBackEtcInfo(vid int64,businessType int8) (err error) {
	return
}
//删除配置文件
func (e *Etc) DeleteEtcInfo(vid int64,businessType int8) (err error) {
	return
}

func WatchEtc(){
	
}
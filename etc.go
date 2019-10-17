package cronsun

import (
	"os"
	"log"
	"fmt"
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


func WatchEtc(){
    
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
func (e *Etc) WriteEtcInfo(vid int64, businessType int8) (err error) {
	cVidPath := fmt.Sprintf("/cronsun/ab/%d/current_version.id",businessType)
	etcPath := fmt.Sprintf("/cronsun/ab/%d/version.json",businessType)
	_ , err = os.Stat(cVidPath)
	if( err != nil) {
		log.Printf("%s is not exist",cVidPath)
	}
	_ , err = os.Stat(etcPath)
	if( err != nil) {
		log.Printf("%s is not exist",etcPath)
	}
    cVidFile,err:=os.Open(cVidPath)
    defer cVidFile.Close()
    str := fmt.Sprintf("%s",vid);
    // content := []byte(str)
    //将指定内容写入到文件中
    _,err = cVidFile.WriteString(str)
    if err != nil {
        log.Println("WriteString error: ", err)
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

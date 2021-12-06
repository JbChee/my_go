package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

/**
定义连接实例
*/
func MyConn(user, password,host, db, port string) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",  user,password, host, port,db )
	db_instance, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	//查询单表， 禁止复数表存在
	db_instance.SingularTable(true)

	return db_instance
}

func orm_op(d *gorm.DB)  {
	//var rows []XzAutoServerConf
	//
	////新增
	//add_data := XzAutoServerConf{GroupZone: "20", ServerId: 811}
	//db.Create(&add_data)
	//
	////删除操作
	//err := db.Model(&rows).Where("server_id=?", 81).Delete(&XzAutoServerConf{}).Error
	//
	////查询操作
	//db.Table("xz_auto_server_conf").Where("status=?", 0).Select([]string{"group_zone", "server_id", "open_time", "server_name", "username"}).Find(&rows)
	//fmt.Println(rows)
	//
	////排序order
	//db.Table("xz_auto_server_conf").Where("status=?", 0).Select([]string{"group_zone", "server_id", "open_time", "server_name", "username"}).Order("server_id ASC").Find(&rows)
	//
	////count操作
	//var count int
	//db.Table("xz_auto_server_conf").Where("status=?", 0).Count(&count)
	//fmt.Println(count)
	//
	//
	////group操作
	//type Result struct {
	//	ServerId  string
	//	Num int64
	//}
	//var new_results []Result
	//db.Table("xz_auto_server_conf").Where("status=?", 0).Select("server_id, count(*) as num").Group("server_id").Scan(&new_results)
	//fmt.Println("%+v", new_results)
	//
	////更新操作
	//err = db.Model(&rows).Where("server_id=?", 80).Update("status", 0).Error
	//if err !=nil {
	//	fmt.Println(err)
	//}


	//USER_ACCOUNT_USER_THIRD_LOGIN := "user_openid"
	//type user struct {
	//	Uid int64 `json:"uid"`
	//}
	//res := []user{}
	//err := d.Table(USER_ACCOUNT_USER_THIRD_LOGIN).Select("uid").Where("platform = ? and create_time >= ?", "facebook2", "2021-03-29").Find(&res).Error
	//if err !=nil{
	//	return
	//}
	//fmt.Println(res)


	//type data struct {
	//	Id int64 `json:"id"`
	//}
	//res := []data{}
	//start := 0
	//offset := 20
	//err := d.Table(USER_ACCOUNT_USER_THIRD_LOGIN).Select("id").Order("id DESC").Offset(start).Limit(offset).Find(&res).Error
	//if err !=nil{
	//	return
	//}
	//fmt.Println(res)
	//type ActivityPlamTreeRank struct {
	//	Id         int        `json:"id" gorm:"column:id"`
	//	Uid        int64      `json:"uid" gorm:"column:uid"`
	//	//Period     int        `json:"period" gorm:"column:period"`
	//	Fruit      int        `json:"fruit" gorm:"column:fruit"`
	//	//ActivityId int        `json:"activity_id" gorm:"column:activity_id"`
	//	CreateTime xtime.Time `json:"create_time" gorm:"column:create_time"`
	//	UpdateTime xtime.Time `json:"update_time" gorm:"column:update_time" `
	//}
	//score := 100
	//activity_plamtree_rank := "activity_plamtree_fruit"
	//dublicateSQL := fmt.Sprintf("ON DUPLICATE KEY UPDATE fruit=fruit+(%d)", score)
	//err := d.Table(activity_plamtree_rank).Set("gorm:insert_option", dublicateSQL).Create(&ActivityPlamTreeRank{
	//	Uid:        100256,
	//	//Period:     0, //月榜
	//	Fruit:      score,
	//	//ActivityId: 20210410,
	//	CreateTime: xtime.Time(time.Now().Unix()),
	//	UpdateTime: xtime.Time(time.Now().Unix()),
	//}).Error
	//fmt.Println("err=",err)

	//
	//type TaskInfo struct {
	//	Task       int   `json:"task"`
	//	WaterDrop  int64 `json:"water_num"`
	//	//CreateTime xtime.Time `json:"create_time"`
	//}
	////
	//tem := []*TaskInfo{}
	//tabl := "activity_plamtree_get_water"
	//
	//_ = d.Table(tabl).Select("task, water_num").Where("uid = ?", 100307).
	//	Order("create_time asc").Offset(0).Limit(20).Find(&tem).Error
	//fmt.Println(tem)


	//var c int
	//err := d.Table("activity_plamtree_fruit").Where(
	//	" fruit > ? or (fruit = ? and id <= ?)", 16000, 16000, 41).Count(&c).Error
	//fmt.Println(err)
	//fmt.Println(c)
	type data struct {
		Uid int64 `json:"uid"`
	}
	//var rows []data
	row := data{}
	err := d.Table("user_openid").Select("uid").Where("uid= ? and platform = ?", 100599, "facebook2").Find(&row).Error
	fmt.Println(err==nil)
	fmt.Println(row)
}

func main() {
	user := "testuser"
	password := "123456"
	host := "10.100.130.145"
	port := "3306"
	db_name := "user_account"
	//建立连接
	db := MyConn(user,password,host, db_name, port)
	//打开调试模式，在日志中，能看到对应的执行sql
	db.LogMode(true)
	// 关闭数据库链接，defer会在函数结束时关闭数据库连接
	defer db.Close()


	//orm查询方式
	orm_op(db)
}
package main

import (
	"time"

	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gamedb struct {
	GID       int    `gorm:"primary_key"`
	Name      string `gorm:"unique"`
	StartTime time.Time
	EndTime   time.Time
	TotalRun  time.Duration
}

var db *gorm.DB
var dberr error

//DbInit 连接数据库,表迁移
func DbInit() {
	exist, err := IsExists("./Gamedb.db")
	if !exist && err == nil {
		//连接数据库
		db, dberr = gorm.Open(sqlite.Open("./Gamedb.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if dberr != nil {
			panic(dberr)
		}
		//迁移 schema
		db.AutoMigrate(&Gamedb{})
		u1 := Gamedb{}
		db.Create(&u1)
		//log.Println("Db init successful!")
	}
	//连接数据库
	db, dberr = gorm.Open(sqlite.Open("./Gamedb.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if dberr != nil {
		panic(dberr)
	}
	//log.Println("Db exist!")
}

func AddNewGame(gname string) (time.Time, time.Time, time.Duration) {
	var FindGame []Gamedb
	db.Where("Name=?", gname).Find(&FindGame)
	if len(FindGame) == 0 {
		var gamedb Gamedb
		db.Last(&gamedb)
		currentTime := time.Now()
		newgname := Gamedb{gamedb.GID + 1, gname, currentTime, currentTime, 0}
		db.Create(&newgname)
		log.Println("Successful add new game:", gname)
		return currentTime, currentTime, 0
	} else {
		//log.Println("Found exist game:", gname)
		//log.Println(FindGame)
		StartTime, EndTime, Totaltime := AddEndTime(gname)
		return StartTime, EndTime, Totaltime
	}
}

func AddEndTime(gname string) (time.Time, time.Time, time.Duration) {
	var gamedb Gamedb
	currentTime := time.Now()
	db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"EndTime": currentTime})
	db.Where("Name=?", gname).Find(&gamedb)
	//log.Printf(gname)
	Totaltime := AddTotalTime(gname)
	return gamedb.StartTime, gamedb.EndTime, Totaltime
}

func AddTotalTime(gname string) time.Duration {
	var gamedb Gamedb
	db.Where("Name=?", gname).Find(&gamedb)
	Totaltime := gamedb.TotalRun + gamedb.EndTime.Sub(gamedb.StartTime)
	log.Printf(gamedb.EndTime.String() + gamedb.StartTime.String())
	db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"TotalRun": Totaltime})
	return Totaltime
}

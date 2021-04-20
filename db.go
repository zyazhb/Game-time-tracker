package main

import (
	"time"

	// "log"

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
}

func AddNewGame(gname string) {
	var FindGame Gamedb
	db.Where("Name=?", gname).Find(&FindGame)

	if FindGame.Name == "" {
		var gamedb Gamedb
		db.Last(&gamedb)
		currentTime := time.Now()
		newgname := Gamedb{gamedb.GID + 1, gname, currentTime, currentTime, 0}
		db.Create(&newgname)
		// log.Println("Successful add new game:", gname)
	} else {
		StartTime, _, _ := ShowTime(gname)
		if StartTime.String() == "0001-01-01 00:00:00 +0000 UTC" {
			AddStartTime(gname)
		}

	}
}

func AddStartTime(gname string) {
	var gamedb Gamedb
	StartTime := time.Now()
	db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"StartTime": StartTime})
}

func AddEndTime(gname string) {
	var gamedb Gamedb
	StartTime, _, _ := ShowTime(gname)
	if StartTime.String() != "0001-01-01 00:00:00 +0000 UTC" {
		currentTime := time.Now()
		db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"EndTime": currentTime})
		AddTotalTime(gname)
	}
	ClearTime(gname)
}

func AddTotalTime(gname string) {
	var gamedb Gamedb
	db.Where("Name=?", gname).Find(&gamedb)
	Totaltime := gamedb.TotalRun + gamedb.EndTime.Sub(gamedb.StartTime)
	// log.Printf(gamedb.EndTime.String() + gamedb.StartTime.String())
	db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"TotalRun": Totaltime})
	ClearTime(gname)
}

func ShowTime(gname string) (time.Time, time.Time, time.Duration) {
	var gamedb Gamedb
	db.Where("Name=?", gname).Find(&gamedb)
	return gamedb.StartTime, gamedb.EndTime, gamedb.TotalRun
}

func ClearTime(gname string) {
	var gamedb Gamedb
	db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"StartTime": "", "EndTime": ""})
}

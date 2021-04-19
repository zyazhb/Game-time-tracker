package main

import (
	"time"

	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gamedb struct {
	GID       int    `gorm:"primary_key"`
	Name      string `gorm:"unique"`
	StartTime string
	EndTime   string
}

//DbInit 连接数据库,表迁移
func DbInit() {
	exist, err := IsExists("./Gamedb.db")
	if !exist && err == nil {
		//连接数据库
		db, err := gorm.Open(sqlite.Open("./Gamedb.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		//迁移 schema
		db.AutoMigrate(&Gamedb{})
		u1 := Gamedb{}
		db.Create(&u1)
		log.Println("Db init successful!")
	}
	log.Println("Db exist!")
}

func AddNewGame(gname string) (string, string) {
	db, err := gorm.Open(sqlite.Open("./Gamedb.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var FindGame []Gamedb
	db.Where("Name=?", gname).Find(&FindGame)
	if len(FindGame) == 0 {
		var gamedb Gamedb
		db.Last(&gamedb)
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		newgname := Gamedb{gamedb.GID + 1, gname, currentTime, currentTime}
		db.Create(&newgname)
		log.Println("Successful add new game:", gname)
		return currentTime, currentTime
	} else {
		log.Println("Found exist game:", gname)
		log.Println(FindGame)
		StartTime, EndTime := AddEndTime(gname)
		return StartTime, EndTime
	}
}

func AddEndTime(gname string) (string, string) {
	db, err := gorm.Open(sqlite.Open("./Gamedb.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var gamedb Gamedb
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	db.Model(&gamedb).Where("Name=?", gname).Updates(map[string]interface{}{"EndTime": currentTime})
	db.Where("Name=?", gname).Find(&gamedb)
	log.Printf(gname)
	return gamedb.StartTime, gamedb.EndTime
}

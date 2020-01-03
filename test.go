package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testDB/model"
	//"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
)

const (
	USERNAME = "root"
	PASSWORD = "chen"
	NETWORK  = "tcp"
	HOST     = "localhost"
	PORT     = "3306"
	DATABASE = "test"
)

//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
//"root:123@tcp(127.0.0.1:3306)/dbname?charset=utf8"
func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USERNAME,
		PASSWORD,
		NETWORK,
		HOST,
		PORT,
		DATABASE,
	)

	db, err := gorm.Open("mysql", dsn)

	checkErr(err)
	defer func() {
		err = db.Close()
		checkErr(err)
	}()

	db.DB().SetMaxOpenConns(100)

	db.LogMode(true)
	db.AutoMigrate(&model.User{})
	//db.CreateTable(&model.User{})

	user := model.User{
		Name: "zs",
	}
	fmt.Println(user)
	db.Create(&user)

	//time.Sleep(5*time.Second)
	//res,err := db.Select()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

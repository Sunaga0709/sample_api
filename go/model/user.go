package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
}

func InitDB() {
	db := connectDB()
	defer db.Close()

	// ユーザテーブルがなければ作成
	db.AutoMigrate(&User{})
}

func connectDB() *gorm.DB {
	// DB接続情報
	dbms := "mysql"
	user := "dbuser"
	pass := "dbpass"
	protocol := "tcp(localhost:3306)"
	dbname := "sample_api"
	connect := user + ":" + pass + "@" + protocol + "/" + dbname + "?parseTime=true&loc=Asia%2FTokyo"

	// DB接続
	db, err := gorm.Open(dbms, connect)
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func GetAllUser() ([]User, error) {
	var users []User  // 取得用変数
	db := connectDB() // DB接続
	defer db.Close()
	err := db.Find(&users).Error // DBから全ユーザ取得

	return users, err
}

func GetOneUser(id int) (User, error) {
	var user User     // 取得用変数
	db := connectDB() // DB接続
	defer db.Close()
	err := db.First(&user, id).Error // 指定したidのユーザを取得

	return user, err
}

func CreateUser(name string) {
	var user User     // ユーザ作成用変数
	user.Name = name  // Nameを設定
	db := connectDB() // DB接続
	defer db.Close()
	db.Create(&user) // DBにレコード作成
}

func UpdateUser(id int, name string) error {
	var user User     // 取得用変数
	db := connectDB() // DB接続
	defer db.Close()
	err := db.First(&user, id).Error // 指定したidのユーザを取得
	if err != nil {
		return err
	} else {
		user.Name = name // 取得したユーザのNameを変更
		db.Save(&user)   // 変更を保存
		return nil
	}
}

func DeleteUser(id int) error {
	var user User     // 取得用変数
	db := connectDB() // DB接続
	defer db.Close()
	err := db.First(&user, id).Error // 削除するユーザを取得
	if err != nil {
		return err
	} else {
		db.Delete(&user) // ユーザを削除
		return nil
	}
}

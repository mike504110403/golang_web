package gormdemo

import (
	"crypto/md5"
	"fmt"
	"os"

	golog "log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mike504110403/goutils/log"
)

// GormConnect : 連線 gorm
func gormConnect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:850606@tcp(127.0.0.1:3306)/testDatabase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, nil
}

func GormInserDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	user := GormUser{
		Name:     "L1212",
		Phone:    "0912345678",
		Password: md5Password("123456"),
	}
	db.Save(&user)
	//db.Create(&user)
}

// GormDelDemo : 刪除資料
func GormDelDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	user := GormUser{}
	db.Where("name = ?", "L1212").Delete(&user)
}

// GormQueryRowDemo : 查詢單筆資料
func GormQueryRowDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	user := GormUser{}
	db.First(&user, "name = ?", "L1212")
	log.Info(fmt.Sprintf("Query result: %v", user))
}

// GormQueryRowsDemo : 查詢多筆資料
func GormQueryRowsDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	users := []GormUser{}
	db.Find(&users)
	log.Info(fmt.Sprintf("Query result: %v", users))
}

// GormUpdateDemo : 更新資料
func GormUpdateDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	user := GormUser{}
	err = db.Model(&user).Where("name = ?", "L1212").Update("phone", "0987654321").Error
	if err != nil {
		log.Error(fmt.Sprintf("Update failed: %s", db.Error.Error()))
		return
	}
}

// GormTransectionDemo : 交易
func GormTransectionDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	tx := db.Begin()
	if tx.Error != nil {
		log.Error(fmt.Sprintf("Begin failed: %s", tx.Error.Error()))
		return
	}

	user := GormUser{
		Name:     "L1212",
		Phone:    "0912345678",
		Password: md5Password("123456"),
	}
	if tx.Save(&user); tx.Error != nil {
		tx.Rollback()
		log.Error(fmt.Sprintf("Save failed: %s", err.Error()))
		return
	}

	user = GormUser{}
	tx.First(&user, "name = ?", "L1212")
	log.Info(fmt.Sprintf("Query result: %v", user))

	tx.Commit()
}

func GormLogDemo() {
	db, err := gormConnect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()

	db.LogMode(true)
	db.SetLogger(golog.New(os.Stdout, "\r\n", 0))
}

// md5Password : md5 加密
func md5Password(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

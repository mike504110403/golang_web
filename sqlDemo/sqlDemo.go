package sqldemo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mike504110403/goutils/dbconn"
	"github.com/mike504110403/goutils/log"
)

func Init(cfg *dbconn.Config) {
	testDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		"user_name",
		"user_password",
		"server_ip",
		"server_port",
		"db_name",
		"dsn_option",
	)

	config := dbconn.DBConfig{}
	config.DBDriver = dbconn.DBDriverMySQL
	config.DSNSource = testDsn
	if cfg != nil {
		config.ConnConfig = cfg
	}

	cfgList := map[dbconn.DBName]dbconn.DBConfig{
		dbconn.DBName("test"): config,
	}

	dbconn.New(cfgList)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:850606@tcp(127.0.0.1:3306)/testDatabase")
	if err != nil {
		log.Error(fmt.Sprintf("Open database failed: %s", err.Error()))
		return nil, err
	}
	return db, nil
}

// 連線 demo
func ConnectDemo() (*sql.DB, error) {
	db, err := Connect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Error(fmt.Sprintf("Ping database failed: %s", err.Error()))
		return nil, err
	}
	log.Info("Connect to database success")
	return db, nil
}

// 連線並開啟 transaction
func ConnectTxDemo() (*sql.Tx, error) {
	db, err := Connect()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		log.Error(fmt.Sprintf("Begin transaction failed: %s", err.Error()))
		return nil, err
	}
	log.Info("Connect to database success")
	return tx, nil
}

// 單行查詢 demo
func QueryRowDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	// scan 後 query 連線才會釋放
	userData := User{}
	err = db.QueryRow("SELECT uid, name, phone FROM `user` WHERE uid=?", 1).
		Scan(&userData.Uid, &userData.Name, &userData.Phone)
	if err != nil {
		log.Error(fmt.Sprintf("Query failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("Query success: %v", userData))
}

// 多行查詢 demo
func QueryRowsDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	rows, err := db.Query("SELECT uid, name, phone FROM `user` WHERE uid > ?", 0)
	if err != nil {
		log.Error(fmt.Sprintf("Query failed: %s", err.Error()))
		return
	}
	// 關閉 rows 連線才會釋放
	defer rows.Close()
	userList := []User{}
	for rows.Next() {
		userData := User{}
		err := rows.Scan(&userData.Uid, &userData.Name, &userData.Phone)
		if err != nil {
			log.Error(fmt.Sprintf("Scan failed: %s", err.Error()))
			return
		}
		log.Info(fmt.Sprintf("Query success: %v", userData))
		userList = append(userList, userData)
	}

	for _, user := range userList {
		log.Info(fmt.Sprintf("User: %v", user))
	}
}

// 新增單筆 demo
func InsertDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()
	ret, err := db.Exec("INSERT INTO `user` (name, phone) VALUES (?, ?)", "testName", "testPhone")
	if err != nil {
		log.Error(fmt.Sprintf("Insert failed: %s", err.Error()))
		return
	}
	uid, err := ret.LastInsertId()
	if err != nil {
		log.Error(fmt.Sprintf("Get last insert id failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("Insert success, uid: %d", uid))
}

// 更新單筆 demo
func UpdateDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()
	ret, err := db.Exec("UPDATE `user` SET name=? WHERE uid=?", "updateName", 1)
	if err != nil {
		log.Error(fmt.Sprintf("Update failed: %s", err.Error()))
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		log.Error(fmt.Sprintf("Get affected rows failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("Update success, affected rows: %d", affected))
}

// 刪除單筆 demo
func DeleteDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	defer db.Close()
	ret, err := db.Exec("DELETE FROM `user` WHERE uid=?", 1)
	if err != nil {
		log.Error(fmt.Sprintf("Delete failed: %s", err.Error()))
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		log.Error(fmt.Sprintf("Get affected rows failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("Delete success, affected rows: %d", affected))
}

// 預處理查詢 demo
func PrepareQueryDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	stmt, err := db.Prepare("SELECT uid, name, phone FROM `user` WHERE uid > ?")
	if err != nil {
		log.Error(fmt.Sprintf("Prepare failed: %s", err.Error()))
		return
	}
	// 關閉 stmt 連線才會釋放
	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		log.Error(fmt.Sprintf("Query failed: %s", err.Error()))
		return
	}
	// 關閉 rows 連線才會釋放
	defer rows.Close()
	userList := []User{}
	for rows.Next() {
		userData := User{}
		err := rows.Scan(&userData.Uid, &userData.Name, &userData.Phone)
		if err != nil {
			log.Error(fmt.Sprintf("Scan failed: %s", err.Error()))
			return
		}
		log.Info(fmt.Sprintf("Query success: %v", userData))
		userList = append(userList, userData)
	}

	for _, user := range userList {
		log.Info(fmt.Sprintf("User: %v", user))
	}
}

// 預處理新增 demo
func PrepareInsertDemo() {
	db, err := ConnectDemo()
	if err != nil {
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	stmt, err := db.Prepare("INSERT INTO `user` (name, phone) VALUES (?, ?)")
	if err != nil {
		log.Error(fmt.Sprintf("Prepare failed: %s", err.Error()))
		return
	}
	// 關閉 stmt 連線才會釋放
	defer stmt.Close()

	ret, err := stmt.Exec("prepareName", "preparePhone")
	if err != nil {
		log.Error(fmt.Sprintf("Insert failed: %s", err.Error()))
		return
	}
	uid, err := ret.LastInsertId()
	if err != nil {
		log.Error(fmt.Sprintf("Get last insert id failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("Insert success, uid: %d", uid))
}

// 事務 demo
func TransectionDemo() {
	tx, err := ConnectTxDemo()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	_, err = tx.Exec("UPDATE `user` SET name=? WHERE uid=?", "transectionName", "test", 1)
	if err != nil {
		tx.Rollback()
		log.Error(fmt.Sprintf("Update failed: %s", err.Error()))
		return
	}
	_, err = tx.Exec("UPDATE `user` SET name=? WHERE uid=?", "transectionName", "test", 2)
	if err != nil {
		tx.Rollback()
		log.Error(fmt.Sprintf("Update failed: %s", err.Error()))
		return
	}
	err = tx.Commit()
	if err != nil {
		log.Error(fmt.Sprintf("Commit failed: %s", err.Error()))
		return
	}
}

func TransectionPrepareDemo() {
	tx, err := ConnectTxDemo()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		log.Error(fmt.Sprintf("Connect to database failed: %s", err.Error()))
		return
	}
	stmt, err := tx.Prepare("UPDATE `user` SET name=? WHERE uid=?")
	if err != nil {
		tx.Rollback()
		log.Error(fmt.Sprintf("Prepare failed: %s", err.Error()))
		return
	}
	// 關閉 stmt 連線才會釋放
	defer stmt.Close()

	// 批量更新数据
	users := map[string]int{
		"Alice": 1,
		"Bob":   2,
		"Eve":   3,
	}
	for name, id := range users {
		_, err = stmt.Exec(name, id)
		if err != nil {
			tx.Rollback()
			log.Error(fmt.Sprintf("Update failed: %s", err.Error()))
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Error(fmt.Sprintf("Commit failed: %s", err.Error()))
		return
	}
}

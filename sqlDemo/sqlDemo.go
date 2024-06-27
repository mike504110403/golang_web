package sqldemo

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mike504110403/goutils/dbconn"
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

package registry

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/legato/infrastructure/config"
)

var LegatoDB *sql.DB = nil

func init() {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Env.MySQLUserName,
		config.Env.MySQLPassword,
		config.Env.MySQLHostName,
		config.Env.MySQLPort,
		config.Env.MySQLDatabase,
	)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	LegatoDB = db
}

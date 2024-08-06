package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4",
        viper.GetString("mysql.user"),
        viper.GetString("mysql.password"),
        viper.GetString("mysql.host"),
        viper.GetInt("mysql.port"),
        viper.GetString("mysql.db"),
    )
    db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        zap.L().Error("failed to connect to database", zap.Error(err))
        return
    }
    db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
    db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
    return
}

func Close() {
    _ = db.Close()
}
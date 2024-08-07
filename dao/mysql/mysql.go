package mysql

import (
	"fmt"

	"github.com/SoonDubu923/go-forum/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *config.MySQLConfig) (err error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.DB,
    )
    db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        zap.L().Error("failed to connect to database", zap.Error(err))
        return
    }
    db.SetMaxOpenConns(cfg.MaxOpenConns)
    db.SetMaxIdleConns(cfg.MaxIdleConns)
    return
}

func Close() {
    _ = db.Close()
}
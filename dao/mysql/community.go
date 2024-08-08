package mysql

import (
	"database/sql"

	"github.com/SoonDubu923/go-forum/model"
	"go.uber.org/zap"
)

func GetCommunities() (data []*model.Community, err error) {
    if err = db.Select(&data, "SELECT community_id, name FROM community"); err != nil {
        if err == sql.ErrNoRows {
            zap.L().Warn("No communities found")
            err = nil
        }
    }
    return
}
package mysql

import (
	"database/sql"
	"errors"

	errmsg "github.com/SoonDubu923/go-forum/errors"
	"github.com/SoonDubu923/go-forum/model"

	"go.uber.org/zap"
)

// GetCommunities returns a list of communities.
func GetCommunities() (data []*model.Community, err error) {
    if err = db.Select(&data, "SELECT community_id, name FROM community"); err != nil {
        if err == sql.ErrNoRows {
            zap.L().Warn("No communities found")
            err = nil
        }
    }
    return
}

// GetCommunityDetailByID returns the details of a community.
func GetCommunityDetailByID(id int64) (data *model.CommunityDetail, err error) {
    data = new(model.CommunityDetail)
    if err = db.Get(data, "SELECT community_id, name, description, create_time FROM community WHERE community_id = ?", id); err != nil {
        if err == sql.ErrNoRows {
            err = errors.New(errmsg.ErrNotFound)
        }
    }
    return
}
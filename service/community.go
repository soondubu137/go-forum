package service

import (
	"github.com/SoonDubu923/go-forum/dao/mysql"
	"github.com/SoonDubu923/go-forum/model"
)

// GetCommunities returns a list of communities.
func GetCommunities() ([]*model.Community, error) {
    return mysql.GetCommunities()
}

// GetCommunityDetail returns the details of a community.
func GetCommunityDetail(id int64) (*model.CommunityDetail, error) {
    return mysql.GetCommunityDetailByID(id)
}
package service

import (
	"github.com/SoonDubu923/go-forum/dao/mysql"
	"github.com/SoonDubu923/go-forum/model"
)

func GetCommunities() ([]*model.Community, error) {
    return mysql.GetCommunities()
}
package service

import (
	"github.com/SoonDubu923/go-forum/dao/redis"
	"github.com/SoonDubu923/go-forum/model"
)

func VoteForPost(userID int64, p *model.ParamVoteData) error {
    return redis.VoteForPost(userID, p)
}
package redis

import (
	"strconv"
	"time"

	"github.com/SoonDubu923/go-forum/model"
	"github.com/go-redis/redis"
)

// CreatePost creates a post in Redis.
func CreatePost(p *model.Post) (err error) {
    // the following two operations are atomic
    pipe := rds.Pipeline()

    // 1. create post in time zset
    pipe.ZAdd(KEY_POST_TIME_ZSET, redis.Z{
        Score: float64(time.Now().Unix()),
        Member: strconv.FormatInt(p.ID, 10),
    })
    
    // 2. create post in score zset
    // note that the score is initialized to the current time,
    // so that if two posts have the same score, the more recent one will be ranked higher
    pipe.ZAdd(KEY_POST_SCORE_ZSET, redis.Z{
        Score: float64(time.Now().Unix()),
        Member: strconv.FormatInt(p.ID, 10),
    })

    _, err = pipe.Exec()
    return
}

func GetPostsInOrder(p *model.ParamPostList) ([]string, error) {
    // get the key based on the order parameter
    var key string
    switch p.Order {
    case "time":
        key = KEY_POST_TIME_ZSET
    case "score":
        key = KEY_POST_SCORE_ZSET
    }

    start := (p.Page - 1) * p.Size
    end := start + p.Size - 1

    return rds.ZRevRange(key, start, end).Result()
}
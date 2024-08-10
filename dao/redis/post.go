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

    // 3. create post in community set
    pipe.SAdd(KEY_COMMUNITY_SET_PREFIX + strconv.FormatInt(p.CommunityID, 10), strconv.FormatInt(p.ID, 10))

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

// GetPostVoteData gets the number of votes for each post.
func GetPostVoteData(ids []string) ([]int64, error) {
    data := make([]int64, 0, len(ids))
    
    pipe := rds.Pipeline()
    for _, id := range ids {
        pipe.ZCount(KEY_POST_VOTED_ZSET_PREFIX + id, "1", "1")
    }
    cmds, err := pipe.Exec()
    if err != nil {
        return nil, err
    }
    for _, cmd := range cmds {
        count, err := cmd.(*redis.IntCmd).Result()
        if err != nil {
            return nil, err
        }
        data = append(data, count)
    }
    return data, nil
}

// GetCommunityPostIDsInOrder gets the post IDs of a community in order.
func GetCommunityPostIDsInOrder(p *model.ParamCommunityPostList) ([]string, error) {
    communityKey := KEY_COMMUNITY_SET_PREFIX + strconv.FormatInt(p.CommunityID, 10)
    key := communityKey + "::" + p.Order

    // if the key does not exist, create it
    if rds.Exists(key).Val() == 0 {
        var zset string
        switch p.Order {
        case "time":
            zset = KEY_POST_TIME_ZSET
        case "score":
            zset = KEY_POST_SCORE_ZSET
        }
        pipe := rds.Pipeline()
        pipe.ZInterStore(key, redis.ZStore{
            Aggregate: "MAX",
        }, communityKey, zset)
        pipe.Expire(key, 60 * time.Second)
        _, err := pipe.Exec()
        if err != nil {
            return nil, err
        }
    }
    start := (p.Page - 1) * p.Size
    end := start + p.Size - 1

    return rds.ZRevRange(key, start, end).Result()
}
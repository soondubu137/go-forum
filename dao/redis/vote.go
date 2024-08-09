package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	errmsg "github.com/SoonDubu923/go-forum/errors"
	"github.com/go-redis/redis"

	"github.com/SoonDubu923/go-forum/model"
)

const (
    _WEEK_IN_SECONDS = 604800
    _SCORE_PER_VOTE = 432
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
    pipe.ZAdd(KEY_POST_TIME_ZSET, redis.Z{
        Score: float64(time.Now().Unix()),
        Member: strconv.FormatInt(p.ID, 10),
    })

    _, err = pipe.Exec()
    return
}

// VoteForPost votes for a post in Redis.
func VoteForPost(userID int64, p *model.ParamVoteData) error {
    // check if the post is within time limit
    postTime := rds.ZScore(KEY_POST_TIME_ZSET, strconv.FormatInt(p.PostID, 10)).Val()
    if float64(time.Now().Unix()) - postTime > _WEEK_IN_SECONDS {
        return errors.New(errmsg.ErrVoteTimeExpired)
    }

    // get user's current vote status
    key := KEY_POST_VOTED_ZSET_PREFIX + strconv.FormatInt(p.PostID, 10)
    curr := rds.ZScore(key, strconv.FormatInt(userID, 10)).Val()
    var coeff float64
    // coeff is 1 if the user is upvoting, -1 if downvoting
    if curr > float64(p.Direction) {
        coeff = -1
    } else {
        coeff = 1
    }
    // diff is 1 if the user hasn't voted before, 2 if the user is changing their vote
    diff := math.Abs(curr - float64(p.Direction))

    pipe := rds.Pipeline()
    pipe.ZIncrBy(KEY_POST_SCORE_ZSET, _SCORE_PER_VOTE * coeff * diff, strconv.FormatInt(p.PostID, 10))
    if p.Direction == 0 {
        pipe.ZRem(key, strconv.FormatInt(userID, 10))
    } else {
        pipe.ZAdd(key, redis.Z{
            Score: float64(p.Direction),
            Member: strconv.FormatInt(userID, 10),
        })
    }

    _, err := pipe.Exec()
    return err
}
